package main

import (
	"fmt"
	log "log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/totvslabs/couchbase-exporter/client"
	"github.com/totvslabs/couchbase-exporter/collector"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// nolint: gochecknoglobals,lll
var (
	version           = "dev"
	app               = kingpin.New("couchbase-exporter", "exports couchbase metrics in the prometheus format")
	logLevel          = app.Flag("log.level", "Couchbase URL to scrape").Default("info").String()
	listenAddress     = app.Flag("web.listen-address", "Address to listen on for web interface and telemetry").Default(":9420").String()
	metricsPath       = app.Flag("web.telemetry-path", "Path under which to expose metrics").Default("/metrics").String()
	couchbaseURL      = app.Flag("couchbase.url", "Couchbase URL to scrape").Default("http://localhost:8091").String()
	couchbaseUsername = app.Flag("couchbase.username", "Couchbase username").String()
	couchbasePassword = app.Flag("couchbase.password", "Couchbase password").OverrideDefaultFromEnvar("COUCHBASE_PASSWORD").String()

	tasks   = app.Flag("collectors.tasks", "Whether to collect tasks metrics").Default("true").Bool()
	buckets = app.Flag("collectors.buckets", "Whether to collect buckets metrics").Default("true").Bool()
	nodes   = app.Flag("collectors.nodes", "Whether to collect nodes metrics").Default("true").Bool()
	cluster = app.Flag("collectors.cluster", "Whether to collect cluster metrics").Default("true").Bool()

	tasksCollectorBucketsCacheRefreshIntervalSeconds = app.Flag("collector.tasks.buckets-cache-refresh-interval-seconds", "Buckets cache refresh interval seconds for tasks").Default("300").Int()

	bucketsCollectorBucketsCacheRefreshIntervalSeconds = app.Flag("collector.buckets.buckets-cache-refresh-interval-seconds", "Buckets cache refresh interval seconds for buckets").Default("15").Int()
	bucketsCollectorBucketsMaxConcurrent               = app.Flag("collectors.buckets.max-concurent", "Max concurent operation for buckets").Default("5").Int()
	bucketsCollectorBucketsLimitlessConcurrent         = app.Flag("collectors.buckets.limitless-concurent", "Limitless concurent operation for buckets").Default("false").Bool()
)

func main() {
	app.Version(version)
	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))

	var level log.Level
	switch *logLevel {
	case "debug":
		level = log.LevelDebug
	case "info":
		level = log.LevelInfo
	case "warn":
		level = log.LevelWarn
	case "error":
		level = log.LevelError
	default:
		level = log.LevelInfo
	}

	logger := log.New(
		log.NewJSONHandler(
			os.Stdout, &log.HandlerOptions{
				Level: level,
			},
		),
	)

	log.SetDefault(logger)

	log.Info(fmt.Sprintf("starting couchbase-exporter %s...", version))
	log.Info(fmt.Sprintf("Tasks Collector's bucket-cache-refresh-interval: '%d', Buckets Collector's bucket-cache-refresh-interval: '%d'", *tasksCollectorBucketsCacheRefreshIntervalSeconds, *bucketsCollectorBucketsCacheRefreshIntervalSeconds))

	var client = client.New(*couchbaseURL, *couchbaseUsername, *couchbasePassword)

	if *tasks {
		prometheus.MustRegister(collector.NewTasksCollector(client, *tasksCollectorBucketsCacheRefreshIntervalSeconds))
	}
	if *buckets {
		prometheus.MustRegister(
			collector.NewBucketsCollector(
				client,
				*bucketsCollectorBucketsCacheRefreshIntervalSeconds,
				*bucketsCollectorBucketsMaxConcurrent,
				*bucketsCollectorBucketsLimitlessConcurrent,
			),
		)
	}
	if *nodes {
		prometheus.MustRegister(collector.NewNodesCollector(client))
	}
	if *cluster {
		prometheus.MustRegister(collector.NewClusterCollector(client))
	}

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,
			`
			<html>
			<head><title>Couchbase Exporter</title></head>
			<body>
				<h1>Couchbase Exporter</h1>
				<p><a href="`+*metricsPath+`">Metrics</a></p>
			</body>
			</html>
			`)
	})

	log.Info(fmt.Sprintf("server listening on %s", *listenAddress))
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Error(fmt.Sprintf("failed to start server: %v", err))
		os.Exit(1)
	}
}
