package client

import (
	"fmt"
	log "log/slog"
	"time"

	"github.com/pkg/errors"
)

// Buckets returns the results of /pools/default/buckets
func (c Client) Buckets() ([]Bucket, error) {
	var buckets []Bucket
	err := c.get("/pools/default/buckets", &buckets)
	return buckets, errors.Wrap(err, "failed to get buckets")
}

// BucketStats returns the results of /pools/default/buckets/<bucket_name>/stats
func (c Client) BucketStats(name string) (BucketStats, error) {
	unixMilli := time.Now().UnixMilli() - 500
	url := fmt.Sprintf("/pools/default/buckets/%s/stats?haveTStamp=%d", name, unixMilli)
	log.Debug(fmt.Sprintf("[bucket-stats]: Url: %s\n", url))

	var stats BucketStats
	err := c.get(url, &stats)
	log.Debug(fmt.Sprintf("[bucket-stats]: Url: %s, CouchTotalDiskSize:Len: %d\n", url, len(stats.Op.Samples.CouchTotalDiskSize)))
	return stats, errors.Wrap(err, "failed to get bucket stats")
}

// BucketStats (/pools/default/buckets/<bucket_name>/stats)
type BucketStats struct {
	Op struct {
		Samples struct {
			CouchTotalDiskSize          []float64 `json:"couch_total_disk_size"`
			CouchDocsFragmentation      []float64 `json:"couch_docs_fragmentation"`
			CouchViewsFragmentation     []float64 `json:"couch_views_fragmentation"`
			HitRatio                    []float64 `json:"hit_ratio"`
			EpCacheMissRate             []float64 `json:"ep_cache_miss_rate"`
			EpResidentItemsRate         []float64 `json:"ep_resident_items_rate"`
			VbAvgActiveQueueAge         []float64 `json:"vb_avg_active_queue_age"`
			VbAvgReplicaQueueAge        []float64 `json:"vb_avg_replica_queue_age"`
			VbAvgPendingQueueAge        []float64 `json:"vb_avg_pending_queue_age"`
			VbAvgTotalQueueAge          []float64 `json:"vb_avg_total_queue_age"`
			VbActiveResidentItemsRatio  []float64 `json:"vb_active_resident_items_ratio"`
			VbReplicaResidentItemsRatio []float64 `json:"vb_replica_resident_items_ratio"`
			VbPendingResidentItemsRatio []float64 `json:"vb_pending_resident_items_ratio"`
			AvgDiskUpdateTime           []float64 `json:"avg_disk_update_time"`
			AvgDiskCommitTime           []float64 `json:"avg_disk_commit_time"`
			AvgBgWaitTime               []float64 `json:"avg_bg_wait_time"`
			AvgActiveTimestampDrift     []float64 `json:"avg_active_timestamp_drift"`
			AvgReplicaTimestampDrift    []float64 `json:"avg_replica_timestamp_drift"`
			// EpDcpViewsIndexesCount            []float64 `json:"ep_dcp_views+indexes_count"`
			// EpDcpViewsIndexesItemsRemaining   []float64 `json:"ep_dcp_views+indexes_items_remaining"`
			// EpDcpViewsIndexesProducerCount    []float64 `json:"ep_dcp_views+indexes_producer_count"`
			// EpDcpViewsIndexesTotalBacklogSize []float64 `json:"ep_dcp_views+indexes_total_backlog_size"`
			// EpDcpViewsIndexesItemsSent       []float64 `json:"ep_dcp_views+indexes_items_sent"`
			// EpDcpViewsIndexesTotalBytes      []float64 `json:"ep_dcp_views+indexes_total_bytes"`
			// EpDcpViewsIndexesBackoff         []float64 `json:"ep_dcp_views+indexes_backoff"`
			// BgWaitCount                      []float64 `json:"bg_wait_count"`
			// BgWaitTotal                      []float64 `json:"bg_wait_total"`
			BytesRead               []float64 `json:"bytes_read"`
			BytesWritten            []float64 `json:"bytes_written"`
			CasBadval               []float64 `json:"cas_badval"`
			CasHits                 []float64 `json:"cas_hits"`
			CasMisses               []float64 `json:"cas_misses"`
			CmdGet                  []float64 `json:"cmd_get"`
			CmdSet                  []float64 `json:"cmd_set"`
			CouchDocsActualDiskSize []float64 `json:"couch_docs_actual_disk_size"`
			CouchDocsDataSize       []float64 `json:"couch_docs_data_size"`
			CouchDocsDiskSize       []float64 `json:"couch_docs_disk_size"`
			// CouchSpatialDataSize             []float64 `json:"couch_spatial_data_size"`
			// CouchSpatialDiskSize             []float64 `json:"couch_spatial_disk_size"`
			// CouchSpatialOps                  []float64 `json:"couch_spatial_ops"`
			CouchViewsActualDiskSize []float64 `json:"couch_views_actual_disk_size"`
			CouchViewsDataSize       []float64 `json:"couch_views_data_size"`
			// CouchViewsDiskSize               []float64 `json:"couch_views_disk_size"`
			CouchViewsOps   []float64 `json:"couch_views_ops"`
			CurrConnections []float64 `json:"curr_connections"`
			CurrItems       []float64 `json:"curr_items"`
			CurrItemsTot    []float64 `json:"curr_items_tot"`
			DecrHits        []float64 `json:"decr_hits"`
			DecrMisses      []float64 `json:"decr_misses"`
			DeleteHits      []float64 `json:"delete_hits"`
			DeleteMisses    []float64 `json:"delete_misses"`
			DiskCommitCount []float64 `json:"disk_commit_count"`
			// DiskCommitTotal                  []float64 `json:"disk_commit_total"`
			DiskUpdateCount []float64 `json:"disk_update_count"`
			// DiskUpdateTotal                  []float64 `json:"disk_update_total"`
			DiskWriteQueue          []float64 `json:"disk_write_queue"`
			EpActiveAheadExceptions []float64 `json:"ep_active_ahead_exceptions"`
			EpActiveHlcDrift        []float64 `json:"ep_active_hlc_drift"`
			// EpActiveHlcDriftCount            []float64 `json:"ep_active_hlc_drift_count"`
			EpBgFetched                      []float64 `json:"ep_bg_fetched"`
			EpClockCasDriftThresholdExceeded []float64 `json:"ep_clock_cas_drift_threshold_exceeded"`
			EpDcp2IBackoff                   []float64 `json:"ep_dcp_2i_backoff"`
			EpDcp2ICount                     []float64 `json:"ep_dcp_2i_count"`
			EpDcp2IItemsRemaining            []float64 `json:"ep_dcp_2i_items_remaining"`
			EpDcp2IItemsSent                 []float64 `json:"ep_dcp_2i_items_sent"`
			EpDcp2IProducerCount             []float64 `json:"ep_dcp_2i_producer_count"`
			EpDcp2ITotalBacklogSize          []float64 `json:"ep_dcp_2i_total_backlog_size"`
			EpDcp2ITotalBytes                []float64 `json:"ep_dcp_2i_total_bytes"`
			// EpDcpFtsBackoff                  []float64 `json:"ep_dcp_fts_backoff"`
			// EpDcpFtsCount                []float64 `json:"ep_dcp_fts_count"`
			// EpDcpFtsItemsRemaining       []float64 `json:"ep_dcp_fts_items_remaining"`
			// EpDcpFtsItemsSent            []float64 `json:"ep_dcp_fts_items_sent"`
			// EpDcpFtsProducerCount        []float64 `json:"ep_dcp_fts_producer_count"`
			// EpDcpFtsTotalBacklogSize     []float64 `json:"ep_dcp_fts_total_backlog_size"`
			// EpDcpFtsTotalBytes           []float64 `json:"ep_dcp_fts_total_bytes"`
			EpDcpOtherBackoff            []float64 `json:"ep_dcp_other_backoff"`
			EpDcpOtherCount              []float64 `json:"ep_dcp_other_count"`
			EpDcpOtherItemsRemaining     []float64 `json:"ep_dcp_other_items_remaining"`
			EpDcpOtherItemsSent          []float64 `json:"ep_dcp_other_items_sent"`
			EpDcpOtherProducerCount      []float64 `json:"ep_dcp_other_producer_count"`
			EpDcpOtherTotalBacklogSize   []float64 `json:"ep_dcp_other_total_backlog_size"`
			EpDcpOtherTotalBytes         []float64 `json:"ep_dcp_other_total_bytes"`
			EpDcpReplicaBackoff          []float64 `json:"ep_dcp_replica_backoff"`
			EpDcpReplicaCount            []float64 `json:"ep_dcp_replica_count"`
			EpDcpReplicaItemsRemaining   []float64 `json:"ep_dcp_replica_items_remaining"`
			EpDcpReplicaItemsSent        []float64 `json:"ep_dcp_replica_items_sent"`
			EpDcpReplicaProducerCount    []float64 `json:"ep_dcp_replica_producer_count"`
			EpDcpReplicaTotalBacklogSize []float64 `json:"ep_dcp_replica_total_backlog_size"`
			EpDcpReplicaTotalBytes       []float64 `json:"ep_dcp_replica_total_bytes"`
			EpDcpViewsBackoff            []float64 `json:"ep_dcp_views_backoff"`
			EpDcpViewsCount              []float64 `json:"ep_dcp_views_count"`
			EpDcpViewsItemsRemaining     []float64 `json:"ep_dcp_views_items_remaining"`
			EpDcpViewsItemsSent          []float64 `json:"ep_dcp_views_items_sent"`
			EpDcpViewsProducerCount      []float64 `json:"ep_dcp_views_producer_count"`
			EpDcpViewsTotalBacklogSize   []float64 `json:"ep_dcp_views_total_backlog_size"`
			EpDcpViewsTotalBytes         []float64 `json:"ep_dcp_views_total_bytes"`
			EpDcpXdcrBackoff             []float64 `json:"ep_dcp_xdcr_backoff"`
			EpDcpXdcrCount               []float64 `json:"ep_dcp_xdcr_count"`
			EpDcpXdcrItemsRemaining      []float64 `json:"ep_dcp_xdcr_items_remaining"`
			EpDcpXdcrItemsSent           []float64 `json:"ep_dcp_xdcr_items_sent"`
			EpDcpXdcrProducerCount       []float64 `json:"ep_dcp_xdcr_producer_count"`
			EpDcpXdcrTotalBacklogSize    []float64 `json:"ep_dcp_xdcr_total_backlog_size"`
			EpDcpXdcrTotalBytes          []float64 `json:"ep_dcp_xdcr_total_bytes"`
			EpDiskqueueDrain             []float64 `json:"ep_diskqueue_drain"`
			EpDiskqueueFill              []float64 `json:"ep_diskqueue_fill"`
			EpDiskqueueItems             []float64 `json:"ep_diskqueue_items"`
			EpFlusherTodo                []float64 `json:"ep_flusher_todo"`
			EpItemCommitFailed           []float64 `json:"ep_item_commit_failed"`
			EpKvSize                     []float64 `json:"ep_kv_size"`
			EpMaxSize                    []float64 `json:"ep_max_size"`
			EpMemHighWat                 []float64 `json:"ep_mem_high_wat"`
			EpMemLowWat                  []float64 `json:"ep_mem_low_wat"`
			EpMetaDataMemory             []float64 `json:"ep_meta_data_memory"`
			EpNumNonResident             []float64 `json:"ep_num_non_resident"`
			EpNumOpsDelMeta              []float64 `json:"ep_num_ops_del_meta"`
			EpNumOpsDelRetMeta           []float64 `json:"ep_num_ops_del_ret_meta"`
			EpNumOpsGetMeta              []float64 `json:"ep_num_ops_get_meta"`
			EpNumOpsSetMeta              []float64 `json:"ep_num_ops_set_meta"`
			EpNumOpsSetRetMeta           []float64 `json:"ep_num_ops_set_ret_meta"`
			EpNumValueEjects             []float64 `json:"ep_num_value_ejects"`
			EpOomErrors                  []float64 `json:"ep_oom_errors"`
			EpOpsCreate                  []float64 `json:"ep_ops_create"`
			EpOpsUpdate                  []float64 `json:"ep_ops_update"`
			EpOverhead                   []float64 `json:"ep_overhead"`
			EpQueueSize                  []float64 `json:"ep_queue_size"`
			EpReplicaAheadExceptions     []float64 `json:"ep_replica_ahead_exceptions"`
			EpReplicaHlcDrift            []float64 `json:"ep_replica_hlc_drift"`
			// EpReplicaHlcDriftCount       []float64 `json:"ep_replica_hlc_drift_count"`
			EpTmpOomErrors []float64 `json:"ep_tmp_oom_errors"`
			EpVbTotal      []float64 `json:"ep_vb_total"`
			Evictions      []float64 `json:"evictions"`
			GetHits        []float64 `json:"get_hits"`
			GetMisses      []float64 `json:"get_misses"`
			IncrHits       []float64 `json:"incr_hits"`
			IncrMisses     []float64 `json:"incr_misses"`
			MemUsed        []float64 `json:"mem_used"`
			Misses         []float64 `json:"misses"`
			Ops            []float64 `json:"ops"`
			// Timestamp               []float64 `json:"timestamp"`
			VbActiveEject           []float64 `json:"vb_active_eject"`
			VbActiveItmMemory       []float64 `json:"vb_active_itm_memory"`
			VbActiveMetaDataMemory  []float64 `json:"vb_active_meta_data_memory"`
			VbActiveNum             []float64 `json:"vb_active_num"`
			VbActiveNumNonResident  []float64 `json:"vb_active_num_non_resident"`
			VbActiveOpsCreate       []float64 `json:"vb_active_ops_create"`
			VbActiveOpsUpdate       []float64 `json:"vb_active_ops_update"`
			VbActiveQueueAge        []float64 `json:"vb_active_queue_age"`
			VbActiveQueueDrain      []float64 `json:"vb_active_queue_drain"`
			VbActiveQueueFill       []float64 `json:"vb_active_queue_fill"`
			VbActiveQueueSize       []float64 `json:"vb_active_queue_size"`
			VbPendingCurrItems      []float64 `json:"vb_pending_curr_items"`
			VbPendingEject          []float64 `json:"vb_pending_eject"`
			VbPendingItmMemory      []float64 `json:"vb_pending_itm_memory"`
			VbPendingMetaDataMemory []float64 `json:"vb_pending_meta_data_memory"`
			VbPendingNum            []float64 `json:"vb_pending_num"`
			VbPendingNumNonResident []float64 `json:"vb_pending_num_non_resident"`
			VbPendingOpsCreate      []float64 `json:"vb_pending_ops_create"`
			VbPendingOpsUpdate      []float64 `json:"vb_pending_ops_update"`
			VbPendingQueueAge       []float64 `json:"vb_pending_queue_age"`
			VbPendingQueueDrain     []float64 `json:"vb_pending_queue_drain"`
			VbPendingQueueFill      []float64 `json:"vb_pending_queue_fill"`
			VbPendingQueueSize      []float64 `json:"vb_pending_queue_size"`
			VbReplicaCurrItems      []float64 `json:"vb_replica_curr_items"`
			VbReplicaEject          []float64 `json:"vb_replica_eject"`
			VbReplicaItmMemory      []float64 `json:"vb_replica_itm_memory"`
			VbReplicaMetaDataMemory []float64 `json:"vb_replica_meta_data_memory"`
			VbReplicaNum            []float64 `json:"vb_replica_num"`
			VbReplicaNumNonResident []float64 `json:"vb_replica_num_non_resident"`
			VbReplicaOpsCreate      []float64 `json:"vb_replica_ops_create"`
			VbReplicaOpsUpdate      []float64 `json:"vb_replica_ops_update"`
			VbReplicaQueueAge       []float64 `json:"vb_replica_queue_age"`
			VbReplicaQueueDrain     []float64 `json:"vb_replica_queue_drain"`
			VbReplicaQueueFill      []float64 `json:"vb_replica_queue_fill"`
			VbReplicaQueueSize      []float64 `json:"vb_replica_queue_size"`
			VbTotalQueueAge         []float64 `json:"vb_total_queue_age"`
			XdcOps                  []float64 `json:"xdc_ops"`
			CPUIdleMs               []float64 `json:"cpu_idle_ms"`
			CPULocalMs              []float64 `json:"cpu_local_ms"`
			CPUUtilizationRate      []float64 `json:"cpu_utilization_rate"`
			HibernatedRequests      []float64 `json:"hibernated_requests"`
			HibernatedWaked         []float64 `json:"hibernated_waked"`
			MemActualFree           []float64 `json:"mem_actual_free"`
			MemActualUsed           []float64 `json:"mem_actual_used"`
			MemFree                 []float64 `json:"mem_free"`
			MemTotal                []float64 `json:"mem_total"`
			MemUsedSys              []float64 `json:"mem_used_sys"`
			RestRequests            []float64 `json:"rest_requests"`
			SwapTotal               []float64 `json:"swap_total"`
			SwapUsed                []float64 `json:"swap_used"`
		} `json:"samples"`
		SamplesCount float64 `json:"samplesCount"`
		IsPersistent bool    `json:"isPersistent"`
		LastTStamp   float64 `json:"lastTStamp"`
		Interval     float64 `json:"interval"`
	} `json:"op"`
}

// Bucket (/pools/default/buckets)
type Bucket struct {
	Name string `json:"name"`
	// UUID          string  `json:"uuid"`
	// ReplicaNumber float64 `json:"replicaNumber"`
	// ThreadsNumber float64 `json:"threadsNumber"`
	// Quota         struct {
	// 	RAM    float64 `json:"ram"`
	// 	RawRAM float64 `json:"rawRAM"`
	// } `json:"quota"`
	BasicStats struct {
		QuotaPercentUsed       float64 `json:"quotaPercentUsed"`
		OpsPerSec              float64 `json:"opsPerSec"`
		DiskFetches            float64 `json:"diskFetches"`
		ItemCount              float64 `json:"itemCount"`
		DiskUsed               float64 `json:"diskUsed"`
		DataUsed               float64 `json:"dataUsed"`
		MemUsed                float64 `json:"memUsed"`
		VbActiveNumNonResident float64 `json:"vbActiveNumNonResident"`
	} `json:"basicStats"`
}
