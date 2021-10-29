//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
This file is part of Cloud Native PostgreSQL.

Copyright (C) 2019-2021 EnterpriseDB Corporation.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/EnterpriseDB/cloud-native-postgresql/pkg/utils"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AffinityConfiguration) DeepCopyInto(out *AffinityConfiguration) {
	*out = *in
	if in.EnablePodAntiAffinity != nil {
		in, out := &in.EnablePodAntiAffinity, &out.EnablePodAntiAffinity
		*out = new(bool)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalPodAntiAffinity != nil {
		in, out := &in.AdditionalPodAntiAffinity, &out.AdditionalPodAntiAffinity
		*out = new(v1.PodAntiAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalPodAffinity != nil {
		in, out := &in.AdditionalPodAffinity, &out.AdditionalPodAffinity
		*out = new(v1.PodAffinity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AffinityConfiguration.
func (in *AffinityConfiguration) DeepCopy() *AffinityConfiguration {
	if in == nil {
		return nil
	}
	out := new(AffinityConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureCredentials) DeepCopyInto(out *AzureCredentials) {
	*out = *in
	if in.ConnectionString != nil {
		in, out := &in.ConnectionString, &out.ConnectionString
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.StorageAccount != nil {
		in, out := &in.StorageAccount, &out.StorageAccount
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.StorageKey != nil {
		in, out := &in.StorageKey, &out.StorageKey
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.StorageSasToken != nil {
		in, out := &in.StorageSasToken, &out.StorageSasToken
		*out = new(SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureCredentials.
func (in *AzureCredentials) DeepCopy() *AzureCredentials {
	if in == nil {
		return nil
	}
	out := new(AzureCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupConfiguration) DeepCopyInto(out *BackupConfiguration) {
	*out = *in
	if in.BarmanObjectStore != nil {
		in, out := &in.BarmanObjectStore, &out.BarmanObjectStore
		*out = new(BarmanObjectStoreConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupConfiguration.
func (in *BackupConfiguration) DeepCopy() *BackupConfiguration {
	if in == nil {
		return nil
	}
	out := new(BackupConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupList) DeepCopyInto(out *BackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupList.
func (in *BackupList) DeepCopy() *BackupList {
	if in == nil {
		return nil
	}
	out := new(BackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	out.Cluster = in.Cluster
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	if in.S3Credentials != nil {
		in, out := &in.S3Credentials, &out.S3Credentials
		*out = new(S3Credentials)
		**out = **in
	}
	if in.AzureCredentials != nil {
		in, out := &in.AzureCredentials, &out.AzureCredentials
		*out = new(AzureCredentials)
		(*in).DeepCopyInto(*out)
	}
	if in.StartedAt != nil {
		in, out := &in.StartedAt, &out.StartedAt
		*out = (*in).DeepCopy()
	}
	if in.StoppedAt != nil {
		in, out := &in.StoppedAt, &out.StoppedAt
		*out = (*in).DeepCopy()
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(InstanceID)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarmanObjectStoreConfiguration) DeepCopyInto(out *BarmanObjectStoreConfiguration) {
	*out = *in
	if in.S3Credentials != nil {
		in, out := &in.S3Credentials, &out.S3Credentials
		*out = new(S3Credentials)
		**out = **in
	}
	if in.AzureCredentials != nil {
		in, out := &in.AzureCredentials, &out.AzureCredentials
		*out = new(AzureCredentials)
		(*in).DeepCopyInto(*out)
	}
	if in.EndpointCA != nil {
		in, out := &in.EndpointCA, &out.EndpointCA
		*out = new(SecretKeySelector)
		**out = **in
	}
	if in.Wal != nil {
		in, out := &in.Wal, &out.Wal
		*out = new(WalBackupConfiguration)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(DataBackupConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarmanObjectStoreConfiguration.
func (in *BarmanObjectStoreConfiguration) DeepCopy() *BarmanObjectStoreConfiguration {
	if in == nil {
		return nil
	}
	out := new(BarmanObjectStoreConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapConfiguration) DeepCopyInto(out *BootstrapConfiguration) {
	*out = *in
	if in.InitDB != nil {
		in, out := &in.InitDB, &out.InitDB
		*out = new(BootstrapInitDB)
		(*in).DeepCopyInto(*out)
	}
	if in.Recovery != nil {
		in, out := &in.Recovery, &out.Recovery
		*out = new(BootstrapRecovery)
		(*in).DeepCopyInto(*out)
	}
	if in.PgBaseBackup != nil {
		in, out := &in.PgBaseBackup, &out.PgBaseBackup
		*out = new(BootstrapPgBaseBackup)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapConfiguration.
func (in *BootstrapConfiguration) DeepCopy() *BootstrapConfiguration {
	if in == nil {
		return nil
	}
	out := new(BootstrapConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapInitDB) DeepCopyInto(out *BootstrapInitDB) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(LocalObjectReference)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DataChecksums != nil {
		in, out := &in.DataChecksums, &out.DataChecksums
		*out = new(bool)
		**out = **in
	}
	if in.PostInitSQL != nil {
		in, out := &in.PostInitSQL, &out.PostInitSQL
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapInitDB.
func (in *BootstrapInitDB) DeepCopy() *BootstrapInitDB {
	if in == nil {
		return nil
	}
	out := new(BootstrapInitDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapPgBaseBackup) DeepCopyInto(out *BootstrapPgBaseBackup) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapPgBaseBackup.
func (in *BootstrapPgBaseBackup) DeepCopy() *BootstrapPgBaseBackup {
	if in == nil {
		return nil
	}
	out := new(BootstrapPgBaseBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapRecovery) DeepCopyInto(out *BootstrapRecovery) {
	*out = *in
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(LocalObjectReference)
		**out = **in
	}
	if in.RecoveryTarget != nil {
		in, out := &in.RecoveryTarget, &out.RecoveryTarget
		*out = new(RecoveryTarget)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapRecovery.
func (in *BootstrapRecovery) DeepCopy() *BootstrapRecovery {
	if in == nil {
		return nil
	}
	out := new(BootstrapRecovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesConfiguration) DeepCopyInto(out *CertificatesConfiguration) {
	*out = *in
	if in.ServerAltDNSNames != nil {
		in, out := &in.ServerAltDNSNames, &out.ServerAltDNSNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesConfiguration.
func (in *CertificatesConfiguration) DeepCopy() *CertificatesConfiguration {
	if in == nil {
		return nil
	}
	out := new(CertificatesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesStatus) DeepCopyInto(out *CertificatesStatus) {
	*out = *in
	in.CertificatesConfiguration.DeepCopyInto(&out.CertificatesConfiguration)
	if in.Expirations != nil {
		in, out := &in.Expirations, &out.Expirations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesStatus.
func (in *CertificatesStatus) DeepCopy() *CertificatesStatus {
	if in == nil {
		return nil
	}
	out := new(CertificatesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.PostgresConfiguration.DeepCopyInto(&out.PostgresConfiguration)
	if in.Bootstrap != nil {
		in, out := &in.Bootstrap, &out.Bootstrap
		*out = new(BootstrapConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ReplicaCluster != nil {
		in, out := &in.ReplicaCluster, &out.ReplicaCluster
		*out = new(ReplicaClusterConfiguration)
		**out = **in
	}
	if in.SuperuserSecret != nil {
		in, out := &in.SuperuserSecret, &out.SuperuserSecret
		*out = new(LocalObjectReference)
		**out = **in
	}
	if in.EnableSuperuserAccess != nil {
		in, out := &in.EnableSuperuserAccess, &out.EnableSuperuserAccess
		*out = new(bool)
		**out = **in
	}
	if in.Certificates != nil {
		in, out := &in.Certificates, &out.Certificates
		*out = new(CertificatesConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	in.StorageConfiguration.DeepCopyInto(&out.StorageConfiguration)
	in.Affinity.DeepCopyInto(&out.Affinity)
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(BackupConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeMaintenanceWindow != nil {
		in, out := &in.NodeMaintenanceWindow, &out.NodeMaintenanceWindow
		*out = new(NodeMaintenanceWindow)
		(*in).DeepCopyInto(*out)
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(MonitoringConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalClusters != nil {
		in, out := &in.ExternalClusters, &out.ExternalClusters
		*out = make([]ExternalCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.InstancesStatus != nil {
		in, out := &in.InstancesStatus, &out.InstancesStatus
		*out = make(map[utils.PodStatus][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.DanglingPVC != nil {
		in, out := &in.DanglingPVC, &out.DanglingPVC
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InitializingPVC != nil {
		in, out := &in.InitializingPVC, &out.InitializingPVC
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HealthyPVC != nil {
		in, out := &in.HealthyPVC, &out.HealthyPVC
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.SecretsResourceVersion.DeepCopyInto(&out.SecretsResourceVersion)
	in.ConfigMapResourceVersion.DeepCopyInto(&out.ConfigMapResourceVersion)
	in.Certificates.DeepCopyInto(&out.Certificates)
	if in.PoolerIntegrations != nil {
		in, out := &in.PoolerIntegrations, &out.PoolerIntegrations
		*out = new(PoolerIntegrations)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapKeySelector) DeepCopyInto(out *ConfigMapKeySelector) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapKeySelector.
func (in *ConfigMapKeySelector) DeepCopy() *ConfigMapKeySelector {
	if in == nil {
		return nil
	}
	out := new(ConfigMapKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapResourceVersion) DeepCopyInto(out *ConfigMapResourceVersion) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapResourceVersion.
func (in *ConfigMapResourceVersion) DeepCopy() *ConfigMapResourceVersion {
	if in == nil {
		return nil
	}
	out := new(ConfigMapResourceVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataBackupConfiguration) DeepCopyInto(out *DataBackupConfiguration) {
	*out = *in
	if in.Jobs != nil {
		in, out := &in.Jobs, &out.Jobs
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataBackupConfiguration.
func (in *DataBackupConfiguration) DeepCopy() *DataBackupConfiguration {
	if in == nil {
		return nil
	}
	out := new(DataBackupConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalCluster) DeepCopyInto(out *ExternalCluster) {
	*out = *in
	if in.ConnectionParameters != nil {
		in, out := &in.ConnectionParameters, &out.ConnectionParameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SSLCert != nil {
		in, out := &in.SSLCert, &out.SSLCert
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SSLKey != nil {
		in, out := &in.SSLKey, &out.SSLKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SSLRootCert != nil {
		in, out := &in.SSLRootCert, &out.SSLRootCert
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.BarmanObjectStore != nil {
		in, out := &in.BarmanObjectStore, &out.BarmanObjectStore
		*out = new(BarmanObjectStoreConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalCluster.
func (in *ExternalCluster) DeepCopy() *ExternalCluster {
	if in == nil {
		return nil
	}
	out := new(ExternalCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceID) DeepCopyInto(out *InstanceID) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceID.
func (in *InstanceID) DeepCopy() *InstanceID {
	if in == nil {
		return nil
	}
	out := new(InstanceID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConfiguration) DeepCopyInto(out *MonitoringConfiguration) {
	*out = *in
	if in.CustomQueriesConfigMap != nil {
		in, out := &in.CustomQueriesConfigMap, &out.CustomQueriesConfigMap
		*out = make([]ConfigMapKeySelector, len(*in))
		copy(*out, *in)
	}
	if in.CustomQueriesSecret != nil {
		in, out := &in.CustomQueriesSecret, &out.CustomQueriesSecret
		*out = make([]SecretKeySelector, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConfiguration.
func (in *MonitoringConfiguration) DeepCopy() *MonitoringConfiguration {
	if in == nil {
		return nil
	}
	out := new(MonitoringConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMaintenanceWindow) DeepCopyInto(out *NodeMaintenanceWindow) {
	*out = *in
	if in.ReusePVC != nil {
		in, out := &in.ReusePVC, &out.ReusePVC
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMaintenanceWindow.
func (in *NodeMaintenanceWindow) DeepCopy() *NodeMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := new(NodeMaintenanceWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgbouncerIntegrationStatus) DeepCopyInto(out *PgbouncerIntegrationStatus) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgbouncerIntegrationStatus.
func (in *PgbouncerIntegrationStatus) DeepCopy() *PgbouncerIntegrationStatus {
	if in == nil {
		return nil
	}
	out := new(PgbouncerIntegrationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolerIntegrations) DeepCopyInto(out *PoolerIntegrations) {
	*out = *in
	in.PgBouncerIntegration.DeepCopyInto(&out.PgBouncerIntegration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolerIntegrations.
func (in *PoolerIntegrations) DeepCopy() *PoolerIntegrations {
	if in == nil {
		return nil
	}
	out := new(PoolerIntegrations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresConfiguration) DeepCopyInto(out *PostgresConfiguration) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PgHBA != nil {
		in, out := &in.PgHBA, &out.PgHBA
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalLibraries != nil {
		in, out := &in.AdditionalLibraries, &out.AdditionalLibraries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresConfiguration.
func (in *PostgresConfiguration) DeepCopy() *PostgresConfiguration {
	if in == nil {
		return nil
	}
	out := new(PostgresConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecoveryTarget) DeepCopyInto(out *RecoveryTarget) {
	*out = *in
	if in.TargetImmediate != nil {
		in, out := &in.TargetImmediate, &out.TargetImmediate
		*out = new(bool)
		**out = **in
	}
	if in.Exclusive != nil {
		in, out := &in.Exclusive, &out.Exclusive
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecoveryTarget.
func (in *RecoveryTarget) DeepCopy() *RecoveryTarget {
	if in == nil {
		return nil
	}
	out := new(RecoveryTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaClusterConfiguration) DeepCopyInto(out *ReplicaClusterConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaClusterConfiguration.
func (in *ReplicaClusterConfiguration) DeepCopy() *ReplicaClusterConfiguration {
	if in == nil {
		return nil
	}
	out := new(ReplicaClusterConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Credentials) DeepCopyInto(out *S3Credentials) {
	*out = *in
	out.AccessKeyIDReference = in.AccessKeyIDReference
	out.SecretAccessKeyReference = in.SecretAccessKeyReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Credentials.
func (in *S3Credentials) DeepCopy() *S3Credentials {
	if in == nil {
		return nil
	}
	out := new(S3Credentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledBackup) DeepCopyInto(out *ScheduledBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledBackup.
func (in *ScheduledBackup) DeepCopy() *ScheduledBackup {
	if in == nil {
		return nil
	}
	out := new(ScheduledBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledBackupList) DeepCopyInto(out *ScheduledBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScheduledBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledBackupList.
func (in *ScheduledBackupList) DeepCopy() *ScheduledBackupList {
	if in == nil {
		return nil
	}
	out := new(ScheduledBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledBackupSpec) DeepCopyInto(out *ScheduledBackupSpec) {
	*out = *in
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
	if in.Immediate != nil {
		in, out := &in.Immediate, &out.Immediate
		*out = new(bool)
		**out = **in
	}
	out.Cluster = in.Cluster
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledBackupSpec.
func (in *ScheduledBackupSpec) DeepCopy() *ScheduledBackupSpec {
	if in == nil {
		return nil
	}
	out := new(ScheduledBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledBackupStatus) DeepCopyInto(out *ScheduledBackupStatus) {
	*out = *in
	if in.LastCheckTime != nil {
		in, out := &in.LastCheckTime, &out.LastCheckTime
		*out = (*in).DeepCopy()
	}
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		*out = (*in).DeepCopy()
	}
	if in.NextScheduleTime != nil {
		in, out := &in.NextScheduleTime, &out.NextScheduleTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledBackupStatus.
func (in *ScheduledBackupStatus) DeepCopy() *ScheduledBackupStatus {
	if in == nil {
		return nil
	}
	out := new(ScheduledBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeySelector.
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(SecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsResourceVersion) DeepCopyInto(out *SecretsResourceVersion) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsResourceVersion.
func (in *SecretsResourceVersion) DeepCopy() *SecretsResourceVersion {
	if in == nil {
		return nil
	}
	out := new(SecretsResourceVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfiguration) DeepCopyInto(out *StorageConfiguration) {
	*out = *in
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.ResizeInUseVolumes != nil {
		in, out := &in.ResizeInUseVolumes, &out.ResizeInUseVolumes
		*out = new(bool)
		**out = **in
	}
	if in.PersistentVolumeClaimTemplate != nil {
		in, out := &in.PersistentVolumeClaimTemplate, &out.PersistentVolumeClaimTemplate
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfiguration.
func (in *StorageConfiguration) DeepCopy() *StorageConfiguration {
	if in == nil {
		return nil
	}
	out := new(StorageConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WalBackupConfiguration) DeepCopyInto(out *WalBackupConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WalBackupConfiguration.
func (in *WalBackupConfiguration) DeepCopy() *WalBackupConfiguration {
	if in == nil {
		return nil
	}
	out := new(WalBackupConfiguration)
	in.DeepCopyInto(out)
	return out
}
