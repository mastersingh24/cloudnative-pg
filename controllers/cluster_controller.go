/*
This file is part of Cloud Native PostgreSQL.

Copyright (C) 2019-2021 EnterpriseDB Corporation.
*/

// Package controllers contains the controller of the CRD
package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	apiv1 "github.com/EnterpriseDB/cloud-native-postgresql/api/v1"
	apiv1alpha1 "github.com/EnterpriseDB/cloud-native-postgresql/api/v1alpha1"
	"github.com/EnterpriseDB/cloud-native-postgresql/pkg/postgres"
	"github.com/EnterpriseDB/cloud-native-postgresql/pkg/specs"
)

const (
	podOwnerKey = ".metadata.controller"
	pvcOwnerKey = ".metadata.controller"
	jobOwnerKey = ".metadata.controller"
)

var (
	apiGVString         = apiv1.GroupVersion.String()
	apiv1alpha1GVString = apiv1alpha1.GroupVersion.String()
)

// ClusterReconciler reconciles a Cluster objects
type ClusterReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// Alphabetical order to not repeat or miss permissions
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;update;list
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;update;list
// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;update;list
// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;delete;patch;create;watch
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;create;update
// +kubebuilder:rbac:groups=policy,resources=poddisruptionbudgets,verbs=create;delete;get;list;watch;update;patch
// +kubebuilder:rbac:groups=postgresql.k8s.enterprisedb.io,resources=clusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=postgresql.k8s.enterprisedb.io,resources=clusters/finalizers,verbs=update
// +kubebuilder:rbac:groups=postgresql.k8s.enterprisedb.io,resources=clusters/status,verbs=get;watch;update;patch
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=create;patch;update;get;list;watch
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=create;patch;update;get;list;watch
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;watch;delete;patch
// +kubebuilder:rbac:groups="",resources=configmaps/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch
// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups="",resources=nodes,verbs=get;list;watch
// +kubebuilder:rbac:groups="",resources=persistentvolumeclaims,verbs=get;list;create;watch;delete;patch
// +kubebuilder:rbac:groups="",resources=pods,verbs=get;list;delete;patch;create;watch
// +kubebuilder:rbac:groups="",resources=pods/exec,verbs=get;list;delete;patch;create;watch
// +kubebuilder:rbac:groups="",resources=pods/status,verbs=get
// +kubebuilder:rbac:groups="",resources=secrets,verbs=create;list;get;watch;delete
// +kubebuilder:rbac:groups="",resources=serviceaccounts,verbs=create;patch;update;list;watch;get
// +kubebuilder:rbac:groups="",resources=services,verbs=get;create;delete;update;patch;list;watch

// Reconcile is the operator reconcile loop
func (r *ClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	clusterControllerLog := r.Log.WithValues("namespace", req.Namespace, "name", req.Name)

	var cluster apiv1.Cluster
	if err := r.Get(ctx, req.NamespacedName, &cluster); err != nil {
		// This also happens when you delete a Cluster resource in k8s. If
		// that's the case, let's just wait for the Kubernetes garbage collector
		// to remove all the Pods of the cluster.
		if apierrs.IsNotFound(err) {
			clusterControllerLog.Info("Resource has been deleted")

			return ctrl.Result{}, nil
		}

		// This is a real error, maybe the RBAC configuration is wrong?
		return ctrl.Result{}, fmt.Errorf("cannot get the managed resource: %w", err)
	}

	var namespace corev1.Namespace
	if err := r.Get(ctx, client.ObjectKey{Namespace: "", Name: req.Namespace}, &namespace); err != nil {
		return ctrl.Result{}, fmt.Errorf("cannot get the containing namespace: %w", err)
	}

	if !namespace.DeletionTimestamp.IsZero() {
		// This happens when you delete a namespace containing a Cluster resource. If that's the case,
		// let's just wait for the Kubernetes to remove all object in the namespace.
		return ctrl.Result{}, nil
	}

	// Ensure we have the required global objects
	if err := r.createPostgresClusterObjects(ctx, &cluster); err != nil {
		return ctrl.Result{}, fmt.Errorf("cannot create Cluster auxiliary objects: %w", err)
	}

	// Update the status of this resource
	resources, err := r.getManagedResources(ctx, cluster)
	if err != nil {
		clusterControllerLog.Error(err, "Cannot extract the list of managed resources")
		return ctrl.Result{}, err
	}

	// Update the status section of this Cluster resource
	if err = r.updateResourceStatus(ctx, &cluster, resources); err != nil {
		if apierrs.IsConflict(err) {
			// Requeue a new reconciliation cycle, as in this point we need
			// to quickly react the changes
			return ctrl.Result{Requeue: true}, nil
		}

		return ctrl.Result{}, fmt.Errorf("cannot update the resource status: %w", err)
	}

	if cluster.Status.CurrentPrimary != "" &&
		cluster.Status.CurrentPrimary != cluster.Status.TargetPrimary {
		clusterControllerLog.Info("There is a switchover or a failover "+
			"in progress, waiting for the operation to complete",
			"currentPrimary", cluster.Status.CurrentPrimary,
			"targetPrimary", cluster.Status.TargetPrimary)

		return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
	}

	// Get the replication status
	instancesStatus := r.getStatusFromInstances(ctx, resources.pods)

	// Update the target primary name from the Pods status.
	// This means issuing a failover or switchover when needed.
	selectedPrimary, err := r.updateTargetPrimaryFromPods(ctx, &cluster, instancesStatus, resources)
	if err != nil {
		if err == ErrWalReceiversRunning {
			clusterControllerLog.Info("Waiting for the all WAL receivers to be down to elect a new primary")
			return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
		}
		clusterControllerLog.Info("Cannot update target primary: operation cannot be fulfilled. "+
			"An immediate retry will be scheduled",
			"cluster", cluster.Name)
		return ctrl.Result{Requeue: true}, nil
	}
	if selectedPrimary != "" {
		// If we selected a new primary, stop the reconciliation loop here
		clusterControllerLog.Info("Waiting for the new primary to notice the promotion request",
			"newPrimary", selectedPrimary)
		return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
	}

	// Update the labels for the -rw service to work correctly
	if err = r.updateLabelsOnPods(ctx, &cluster, resources.pods); err != nil {
		return ctrl.Result{}, fmt.Errorf("cannot update labels on pods: %w", err)
	}

	// Act on Pods and PVCs only if there is nothing that is currently being created or deleted
	if runningJobs := resources.countRunningJobs(); runningJobs > 0 {
		clusterControllerLog.V(2).Info("A job is currently running. Waiting", "count", runningJobs)
		return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
	}

	if len(resources.pods.Items) > 0 && resources.noPodsAreAlive() {
		return ctrl.Result{RequeueAfter: 1 * time.Second}, r.RegisterPhase(ctx, &cluster, apiv1.PhaseUnrecoverable,
			"No pods are active, the cluster needs manual intervention ")
	}

	if !resources.allPodsAreActive() {
		clusterControllerLog.V(2).Info("A managed resource is currently being created or deleted. Waiting")
		return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
	}

	// Reconcile PVC resource requirements
	if err = r.ReconcilePVCs(ctx, &cluster, resources); err != nil {
		if apierrs.IsConflict(err) {
			return ctrl.Result{Requeue: true}, nil
		}
		return ctrl.Result{}, err
	}

	// Reconcile Pods
	return r.ReconcilePods(ctx, req, &cluster, resources, instancesStatus)
}

// ReconcilePVCs align the PVCs that are backing our cluster with the user specifications
func (r *ClusterReconciler) ReconcilePVCs(ctx context.Context, cluster *apiv1.Cluster,
	resources *managedResources) error {
	if !cluster.ShouldResizeInUseVolumes() {
		return nil
	}

	quantity, err := resource.ParseQuantity(cluster.Spec.StorageConfiguration.Size)
	if err != nil {
		return fmt.Errorf("while parsing PVC size %v: %w", cluster.Spec.StorageConfiguration.Size, err)
	}

	for idx := range resources.pvcs.Items {
		oldPVC := resources.pvcs.Items[idx].DeepCopy()
		oldQuantity, ok := resources.pvcs.Items[idx].Spec.Resources.Requests["storage"]

		switch {
		case !ok:
			// Missing storage requirement for PVC
			fallthrough

		case oldQuantity.AsDec().Cmp(quantity.AsDec()) == -1:
			// Increasing storage resources
			resources.pvcs.Items[idx].Spec.Resources.Requests["storage"] = quantity
			if err = r.Patch(ctx, &resources.pvcs.Items[idx], client.MergeFrom(oldPVC)); err != nil {
				// Decreasing resources is not possible
				log.Error(err, "error while changing PVC storage requirement",
					"from", oldQuantity, "to", quantity,
					"pvcName", resources.pvcs.Items[idx].Name)

				// We are reaching two errors in two different conditions:
				//
				// 1. we hit a Conflict => a successive reconciliation loop will fix it
				// 2. the StorageClass we used don't support PVC resizing => there's nothing we can do
				//    about it
			}

		case oldQuantity.AsDec().Cmp(quantity.AsDec()) == 1:
			// Decreasing resources is not possible
			log.Info("cannot decrease storage requirement",
				"from", oldQuantity, "to", quantity,
				"pvcName", resources.pvcs.Items[idx].Name)
		}
	}

	return nil
}

// ReconcilePods decides when to create, scale up/down or wait for pods
func (r *ClusterReconciler) ReconcilePods(ctx context.Context, req ctrl.Request, cluster *apiv1.Cluster,
	resources *managedResources, instancesStatus postgres.PostgresqlStatusList) (ctrl.Result, error) {
	clusterControllerLog := r.Log.WithValues("namespace", req.Namespace, "name", req.Name)

	// If we are joining a node, we should wait for the process to finish
	if resources.countRunningJobs() > 0 {
		clusterControllerLog.V(2).Info("Waiting for jobs to finish",
			"clusterName", cluster.Name,
			"namespace", cluster.Namespace,
			"jobs", len(resources.jobs.Items))
		return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
	}

	// Work on the PVCs we currently have
	pvcNeedingMaintenance := len(cluster.Status.DanglingPVC) + len(cluster.Status.InitializingPVC)
	if pvcNeedingMaintenance > 0 {
		if !cluster.IsNodeMaintenanceWindowInProgress() && cluster.Status.ReadyInstances != cluster.Status.Instances {
			// A pod is not ready, let's retry
			clusterControllerLog.V(2).Info("Waiting for node to be ready before attaching PVCs")
			return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
		}

		if err := r.reconcilePVCs(ctx, cluster); err != nil {
			return ctrl.Result{}, err
		}

		// Do another reconcile cycle after handling a dangling PVC
		return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
	}

	// We have these cases now:
	//
	// 1 - There is no existent Pod for this PostgreSQL cluster ==> we need to create the
	// first node from which we will join the others
	//
	// 2 - There is one Pod, and that one is still not ready ==> we need to wait
	// for the first node to be ready
	//
	// 3 - We have already some Pods, all they all ready ==> we can create the other
	// pods joining the node that we already have.
	if cluster.Status.Instances == 0 {
		return r.createPrimaryInstance(ctx, cluster)
	}

	// Stop acting here if there are non-ready Pods unless in maintenance reusing PVCs.
	// The user have choose to wait for the missing nodes to come up
	if !(cluster.IsNodeMaintenanceWindowInProgress() && cluster.IsReusePVCEnabled()) &&
		cluster.Status.ReadyInstances < cluster.Status.Instances {
		clusterControllerLog.V(2).Info("Waiting for Pods to be ready")
		return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
	}

	// Are there missing nodes? Let's create one
	if cluster.Status.Instances < cluster.Spec.Instances &&
		cluster.Status.ReadyInstances == cluster.Status.Instances {
		newNodeSerial, err := r.generateNodeSerial(ctx, cluster)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("cannot generate node serial: %w", err)
		}
		return r.joinReplicaInstance(ctx, newNodeSerial, cluster)
	}

	// Are there nodes to be removed? Remove one of them
	if cluster.Status.Instances > cluster.Spec.Instances {
		if err := r.scaleDownCluster(ctx, cluster, resources); err != nil {
			return ctrl.Result{}, fmt.Errorf("cannot scale down cluster: %w", err)
		}
	}

	// Stop acting here if there are non-ready Pods
	// In the rest of the function we are sure that
	// cluster.Status.Instances == cluster.Spec.Instances and
	// we don't need to modify the cluster topology
	if cluster.Status.ReadyInstances != cluster.Status.Instances ||
		cluster.Status.ReadyInstances != int32(len(instancesStatus.Items)) ||
		!instancesStatus.IsComplete() {
		clusterControllerLog.V(2).Info("Waiting for Pods to be ready")
		return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
	}

	// If we need to rollout a restart of any instance, this is the right moment
	// Do I have to rollout a new image?
	done, err := r.rolloutDueToCondition(ctx, cluster, &instancesStatus, IsPodNeedingRollout)
	if err != nil {
		return ctrl.Result{}, err
	}
	if done {
		// Rolling upgrade is in progress, let's avoid marking stuff as synchronized
		// (but recheck in one second, just to be sure)
		return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
	}

	// When everything is reconciled, update the status
	if err = r.RegisterPhase(ctx, cluster, apiv1.PhaseHealthy, ""); err != nil {
		return ctrl.Result{}, err
	}

	// Cleanup stuff
	return ctrl.Result{}, r.cleanupCluster(ctx, cluster, resources.jobs)
}

// SetupWithManager creates a ClusterReconciler
func (r *ClusterReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager) error {
	err := r.createFieldIndexes(ctx, mgr)
	if err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1.Cluster{}).
		Owns(&corev1.Pod{}).
		Owns(&batchv1.Job{}).
		Owns(&corev1.Service{}).
		Owns(&corev1.PersistentVolumeClaim{}).
		Owns(&policyv1beta1.PodDisruptionBudget{}).
		Watches(
			&source.Kind{Type: &corev1.ConfigMap{}},
			handler.EnqueueRequestsFromMapFunc(r.mapConfigMapsToClusters(ctx)),
			builder.WithPredicates(configMapsPredicate),
		).
		Watches(
			&source.Kind{Type: &corev1.Secret{}},
			handler.EnqueueRequestsFromMapFunc(r.mapSecretsToClusters(ctx)),
			builder.WithPredicates(secretsPredicate),
		).
		Watches(
			&source.Kind{Type: &corev1.Node{}},
			handler.EnqueueRequestsFromMapFunc(r.mapNodeToClusters(ctx)),
			builder.WithPredicates(nodesPredicate),
		).
		Complete(r)
}

// createFieldIndexes creates the indexes needed by this controller
func (r *ClusterReconciler) createFieldIndexes(ctx context.Context, mgr ctrl.Manager) error {
	// Create a new indexed field on Pods. This field will be used to easily
	// find all the Pods created by this controller
	if err := mgr.GetFieldIndexer().IndexField(
		ctx,
		&corev1.Pod{},
		podOwnerKey, func(rawObj client.Object) []string {
			pod := rawObj.(*corev1.Pod)

			if ownerName, ok := isOwnedByCluster(pod); ok {
				return []string{ownerName}
			}

			return nil
		}); err != nil {
		return err
	}

	// Create a new indexed field on Pods. This field will be used to easily
	// find all the Pods created by node
	if err := mgr.GetFieldIndexer().IndexField(
		ctx,
		&corev1.Pod{},
		".spec.nodeName", func(rawObj client.Object) []string {
			pod := rawObj.(*corev1.Pod)
			if pod.Spec.NodeName == "" {
				return nil
			}

			return []string{pod.Spec.NodeName}
		}); err != nil {
		return err
	}

	// Create a new indexed field on PVCs.
	if err := mgr.GetFieldIndexer().IndexField(
		ctx,
		&corev1.PersistentVolumeClaim{},
		pvcOwnerKey, func(rawObj client.Object) []string {
			persistentVolumeClaim := rawObj.(*corev1.PersistentVolumeClaim)

			if ownerName, ok := isOwnedByCluster(persistentVolumeClaim); ok {
				return []string{ownerName}
			}

			return nil
		}); err != nil {
		return err
	}

	// Create a new indexed field on Jobs.
	return mgr.GetFieldIndexer().IndexField(
		ctx,
		&batchv1.Job{},
		jobOwnerKey, func(rawObj client.Object) []string {
			job := rawObj.(*batchv1.Job)

			if ownerName, ok := isOwnedByCluster(job); ok {
				return []string{ownerName}
			}

			return nil
		})
}

// isOwnedByCluster checks that an object is owned by a Cluster and returns
// the owner name
func isOwnedByCluster(obj client.Object) (string, bool) {
	owner := metav1.GetControllerOf(obj)
	if owner == nil {
		return "", false
	}

	if owner.Kind != apiv1.ClusterKind {
		return "", false
	}

	if owner.APIVersion != apiGVString && owner.APIVersion != apiv1alpha1GVString {
		return "", false
	}
	return owner.Name, true
}

// mapSecretsToClusters returns a function mapping cluster events watched to cluster reconcile requests
func (r *ClusterReconciler) mapSecretsToClusters(ctx context.Context) handler.MapFunc {
	return func(obj client.Object) []reconcile.Request {
		secret, ok := obj.(*corev1.Secret)
		if !ok {
			return nil
		}
		var clusters apiv1.ClusterList
		// get all the clusters handled by the operator in the secret namespaces
		err := r.List(ctx, &clusters,
			client.InNamespace(secret.Namespace),
		)
		if err != nil {
			r.Log.Error(err, "while getting cluster list", "namespace", secret.Namespace)
			return nil
		}
		// build requests for cluster referring the secret
		return filterClustersUsingSecret(clusters, secret)
	}
}

// mapNodeToClusters returns a function mapping cluster events watched to cluster reconcile requests
func (r *ClusterReconciler) mapConfigMapsToClusters(ctx context.Context) handler.MapFunc {
	return func(obj client.Object) []reconcile.Request {
		config, ok := obj.(*corev1.ConfigMap)
		if !ok {
			return nil
		}
		var clusters apiv1.ClusterList
		// get all the clusters handled by the operator in the configmap namespaces
		err := r.List(ctx, &clusters,
			client.InNamespace(config.Namespace),
		)
		if err != nil {
			r.Log.Error(err, "while getting cluster list", "namespace", config.Namespace)
			return nil
		}
		// build requests for cluster referring the configmap
		return filterClustersUsingConfigMap(clusters, config)
	}
}

func filterClustersUsingSecret(
	clusters apiv1.ClusterList,
	secret *corev1.Secret,
) (requests []reconcile.Request) {
	for _, cluster := range clusters.Items {
		if cluster.UsesSecret(secret.Name) {
			requests = append(requests,
				reconcile.Request{
					NamespacedName: types.NamespacedName{
						Name:      cluster.Name,
						Namespace: cluster.Namespace,
					},
				},
			)
			continue
		}
	}
	return requests
}

func filterClustersUsingConfigMap(
	clusters apiv1.ClusterList,
	config *corev1.ConfigMap,
) (requests []reconcile.Request) {
	for _, cluster := range clusters.Items {
		if cluster.UsesConfigMap(config.Name) {
			requests = append(requests,
				reconcile.Request{
					NamespacedName: types.NamespacedName{
						Name:      cluster.Name,
						Namespace: cluster.Namespace,
					},
				},
			)
			continue
		}
	}
	return requests
}

// mapNodeToClusters returns a function mapping cluster events watched to cluster reconcile requests
func (r *ClusterReconciler) mapNodeToClusters(ctx context.Context) handler.MapFunc {
	return func(obj client.Object) []reconcile.Request {
		node := obj.(*corev1.Node)
		// exit if the node is schedulable (e.g. not cordoned)
		// could be expanded here with other conditions (e.g. pressure or issues)
		if !node.Spec.Unschedulable {
			return nil
		}
		var childPods corev1.PodList
		// get all the pods handled by the operator on that node
		err := r.List(ctx, &childPods,
			client.MatchingFields{".spec.nodeName": node.Name},
			client.MatchingLabels{specs.ClusterRoleLabelName: specs.ClusterRoleLabelPrimary},
			client.HasLabels{specs.ClusterLabelName},
		)
		if err != nil {
			r.Log.Error(err, "while getting primary instances for node")
			return nil
		}
		var requests []reconcile.Request
		// build requests for nodes the pods are running on
		for idx := range childPods.Items {
			if cluster, ok := isOwnedByCluster(&childPods.Items[idx]); ok {
				requests = append(requests,
					reconcile.Request{
						NamespacedName: types.NamespacedName{
							Name:      cluster,
							Namespace: childPods.Items[idx].Namespace,
						},
					},
				)
			}
		}
		return requests
	}
}
