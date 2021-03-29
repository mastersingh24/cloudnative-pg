/*
This file is part of Cloud Native PostgreSQL.

Copyright (C) 2019-2021 EnterpriseDB Corporation.
*/

// Package metrics enables to expose a set of metrics and collectors on a given postgres instance
package metrics

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/EnterpriseDB/cloud-native-postgresql/pkg/management/log"
	"github.com/EnterpriseDB/cloud-native-postgresql/pkg/management/postgres"
)

const namespace = "cnp"

// Exporter exports a set of metrics and collectors on a given postgres instance
type Exporter struct {
	instance *postgres.Instance
	metrics  metrics
	queries  *QueriesCollector
}

// metrics here are related to the exporter itself, which is instrumented to
// expose them
type metrics struct {
	CollectionsTotal   prometheus.Counter
	PgCollectionErrors *prometheus.CounterVec
	Error              prometheus.Gauge
	PostgreSQLUp       prometheus.Gauge
	CollectionDuration *prometheus.GaugeVec
}

// NewExporter creates an exporter
func NewExporter(instance *postgres.Instance) *Exporter {
	return &Exporter{
		instance: instance,
		metrics:  newMetrics(),
	}
}

// newMetrics returns collector metrics
func newMetrics() metrics {
	subsystem := "collector"
	return metrics{
		CollectionsTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "collections_total",
			Help:      "Total number of times PostgreSQL was accessed for metrics.",
		}),
		PgCollectionErrors: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "collection_errors_total",
			Help:      "Total errors occurred accessing PostgreSQL for metrics.",
		}, []string{"collector"}),
		Error: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "last_collection_error",
			Help:      "1 if the last collection ended with error, 0 otherwise.",
		}),
		PostgreSQLUp: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "up",
			Help:      "1 if PostgreSQL is up, 0 otherwise.",
		}),
		CollectionDuration: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "collection_duration_seconds",
			Help:      "Collection time duration in seconds",
		}, []string{"collector"}),
	}
}

// Describe implements prometheus.Collector, defining the metrics we return.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.metrics.CollectionsTotal.Desc()
	ch <- e.metrics.Error.Desc()
	e.metrics.PgCollectionErrors.Describe(ch)
	ch <- e.metrics.PostgreSQLUp.Desc()
	e.metrics.CollectionDuration.Describe(ch)

	if e.queries != nil {
		e.queries.Describe(ch)
	}
}

// Collect implements prometheus.Collector, collecting the metrics values to
// export.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.collectPgMetrics(ch)

	ch <- e.metrics.CollectionsTotal
	ch <- e.metrics.Error
	e.metrics.PgCollectionErrors.Collect(ch)
	ch <- e.metrics.PostgreSQLUp
	e.metrics.CollectionDuration.Collect(ch)
}

func (e *Exporter) collectPgMetrics(ch chan<- prometheus.Metric) {
	e.metrics.CollectionsTotal.Inc()
	collectionStart := time.Now()
	db, err := e.instance.GetApplicationDB()
	if err != nil {
		log.Log.Error(err, "Error opening connection to PostgreSQL")
		e.metrics.Error.Set(1)
		return
	}

	// First, let's check the connection. No need to proceed if this fails.
	if err := db.Ping(); err != nil {
		log.Log.Error(err, "Error pinging PostgreSQL")
		e.metrics.PostgreSQLUp.Set(0)
		e.metrics.Error.Set(1)
		e.metrics.CollectionDuration.WithLabelValues("Collect.up").Set(time.Since(collectionStart).Seconds())
		return
	}

	e.metrics.PostgreSQLUp.Set(1)
	e.metrics.Error.Set(0)
	e.metrics.CollectionDuration.WithLabelValues("Collect.up").Set(time.Since(collectionStart).Seconds())

	// Work on predefined metrics and custom queries
	if e.queries != nil {
		label := "Collect." + e.queries.Name()
		collectionStart := time.Now()
		if err := e.queries.Collect(ch); err != nil {
			log.Log.Error(err, "Error during collection", "collector", e.queries.Name())
			e.metrics.PgCollectionErrors.WithLabelValues(label).Inc()
			e.metrics.Error.Set(1)
		}
		e.metrics.CollectionDuration.WithLabelValues(label).Set(time.Since(collectionStart).Seconds())
	}
}

// PgCollector is the interface for all the collectors that need to do queries
// on PostgreSQL to gather the results
type PgCollector interface {
	// Name is the unique name of the collector
	Name() string

	// Collect collects data and send the metrics on the channel
	Collect(ch chan<- prometheus.Metric) error

	// Describe collects metadata about the metrics we work with
	Describe(ch chan<- *prometheus.Desc)
}

// ClearCustomQueries removes every custom queries previously added to this
// collector
func (e *Exporter) ClearCustomQueries() {
	e.queries = nil
}

// AddCustomQueries read the custom queries from the passed content
// and add the relative Prometheus collector
func (e *Exporter) AddCustomQueries(queriesContent []byte) error {
	if e.queries == nil {
		e.queries = NewQueriesCollector("cnp", e.instance)
	}

	err := e.queries.ParseQueries(queriesContent)
	if err != nil {
		return fmt.Errorf("while parsing custom queries: %s", err)
	}

	return nil
}
