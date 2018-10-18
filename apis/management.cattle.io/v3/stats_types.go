package v3

const (
	ResourceContainer         = "container"
	ResourcePod               = "pod"
	ResourceWorkload          = "workload"
	ResourceCluster           = "cluster"
	ResourceNode              = "node"
	ResourceEtcd              = "etcd"
	ResourceAPIServer         = "apiserver"
	ResourceScheduler         = "scheduler"
	ResourceControllerManager = "controllermanager"
	ResourceCommon            = "common"
	ResourceComponent         = "component"
)

const (
	Certificate                            = "certificate"
	Claims                                 = "claims"
	Cronjob                                = "cronjob"
	Daemonset                              = "daemonset"
	Deployment                             = "deployment"
	Endpoint                               = "endpoint"
	Pvcprotection                          = "pvcprotection"
	Pvprotection                           = "pvprotection"
	Replicaset                             = "replicaset"
	Replication                            = "replication"
	Replicationmanager                     = "replicationmanager"
	ResourceQuotaControllerResourceChanges = "resource_quota_controller_resource_changes"
	ResourcequotaPrimary                   = "resourcequota_primary"
	ResourcequotaPriority                  = "resourcequota_priority"
	Service                                = "service"
	ServiceAccount                         = "serviceaccount"
	ServiceaccountTokensSecret             = "serviceaccount_tokens_secret"
	ServiceaccountTokensService            = "serviceaccount_tokens_service"
	Statefulset                            = "statefulset"
	Ttlcontroller                          = "ttlcontroller"
	Volumes                                = "volumes"
)

const (
	CPU         = "cpu"
	Filesystem  = "fs"
	Memory      = "memory"
	DiskIO      = "disk_io"
	Load1       = "load1"
	Load5       = "load5"
	Load15      = "load15"
	Network     = "network"
	Disk        = "disk"
	DB          = "db"
	GRPC        = "grpc"
	ClientQuery = "client"

	Transmit = "transmit"
	Receive  = "receive"
	Reads    = "reads"
	Writes   = "writes"

	Usage          = "usage"
	UserQuery      = "user"
	System         = "system"
	CfsThrottled   = "cfs_throttled"
	IOReads        = "io_reads"
	IOWrites       = "io_writes"
	PageOut        = "page_out"
	PageIn         = "page_in"
	PacketsDropped = "packets_dropped"
	Errors         = "errors"
	Packets        = "packets"
	Controller     = "controller"
	RequestCache   = "request_cache"
	Helper         = "helper"
	Request        = "request"
	Cache          = "cache"
	Leader         = "leader"

	Entry = "entry"
	Hit   = "hit"
	Miss  = "miss"

	Adds           = "adds"
	Depth          = "depth"
	Retries        = "retries"
	QueueRetries   = "queue_retries"
	RateLimiterUse = "rate_limiter_use"

	QueueDuration    = "queue_duration"
	WorkDuration     = "work_duration"
	WalFsyncDuration = "wal_fsync_duration"
	CommitDuration   = "commit_duration"
	SnapshotDuration = " snapshot_duration"

	QueueLatency               = "queue_latency"
	SchedulingLatency          = "scheduling_latency"
	AddLatencies               = "add_latencies"
	GetLatencies               = "get_latencies"
	BindingLatency             = "binding_latency"
	SchedulingAlgorithmLatency = "scheduling_algorithm_latency"

	RequestCount   = "request_count"
	RequestError   = "request_error"
	RequestLatency = "request_latency"
	RequestSlow    = "request_slow"

	InsufficientMembers = "insufficient_members"
	LeaderChangesSeen   = "leader_changes_seen"
	ProposalsCommitted  = "proposals_committed"
	ProposalsApplied    = "proposals_applied"
	ProposalsFailed     = "proposals_failed"
	ProposalsPending    = "proposals_pending"

	ClientCertificateExpiration        = "client_certificate_expiration"
	APIServiceRegistrationController   = "APIServiceRegistrationController"
	APIServiceRegistration             = "APIServiceRegistration"
	E2e                                = "e2e"
	AuditEvent                         = "AuditEvent"
	Server                             = "server"
	ResidentMemory                     = "residentMemory"
	NamespaceContainerMemoryUsageBytes = "namespace_container_memory_usage_bytes"
	NamespaceContainerSpecCPUShares    = "namespace_container_spec_cpu_shares"
	NamespaceContainerCPUUsage         = "namespace_container_cpu_usage"
)

const (
	Sum      = "sum"
	Avg      = "avg"
	Rate     = "rate"
	Percent  = "percent"
	Quantile = "quantile"
	Count    = "count"
	Summary  = "summary"
	Increase = "increase"
	Total    = "total"
)

const (
	Milliseconds = "milliseconds"
	Seconds      = "seconds"
	Bytes        = "bytes"
)

var (
	MetricEnum = []string{
	// ResourceCluster + CPULoad1,
	// ResourceCluster + CPULoad5,
	// ResourceCluster + CPULoad15,
	// ResourceCluster + CPUUsageSecondsSumRate,
	// ResourceCluster + CPUUserSecondsSumRate,
	// ResourceCluster + CPUSystemSecondsSumRate,
	// ResourceCluster + MemoryUsagePercent,
	// ResourceCluster + MemoryPageInBytesSumRate,
	// ResourceCluster + MemoryPageOutBytesSumRate,
	// ResourceCluster + NetworkReceiveBytesSumRate,
	// ResourceCluster + NetworkReceivePacketsDroppedSumRate,
	// ResourceCluster + NetworkReceiveErrorsSumRate,
	// ResourceCluster + NetworkReceivePacketsSumRate,
	// ResourceCluster + NetworkTransmitPacketsSumRate,
	// ResourceCluster + NetworkTransmitBytesSumRate,
	// ResourceCluster + NetworkTransmitPacketsDroppedSumRate,
	// ResourceCluster + NetworkTransmitErrorsSumRate,
	// ResourceCluster + FilesystemUseagePercent,
	// ResourceCluster + DiskIOReadsBytesSumRate,
	// ResourceCluster + DiskIOWritesBytesSumRate,

	// ResourceNode + CPULoad1,
	// ResourceNode + CPULoad5,
	// ResourceNode + CPULoad15,
	// ResourceNode + CPUUsageSecondsSumRate,
	// ResourceNode + CPUUserSecondsSumRate,
	// ResourceNode + CPUSystemSecondsSumRate,
	// ResourceNode + MemoryUsagePercent,
	// ResourceNode + MemoryPageInBytesSumRate,
	// ResourceNode + MemoryPageOutBytesSumRate,
	// ResourceNode + NetworkReceiveBytesSumRate,
	// ResourceNode + NetworkReceivePacketsDroppedSumRate,
	// ResourceNode + NetworkReceiveErrorsSumRate,
	// ResourceNode + NetworkReceivePacketsSumRate,
	// ResourceNode + NetworkTransmitPacketsSumRate,
	// ResourceNode + NetworkTransmitBytesSumRate,
	// ResourceNode + NetworkTransmitPacketsDroppedSumRate,
	// ResourceNode + NetworkTransmitErrorsSumRate,
	// ResourceNode + FilesystemUseagePercent,
	// ResourceNode + DiskIOReadsBytesSumRate,
	// ResourceNode + DiskIOWritesBytesSumRate,

	// ResourcePod + CPUUsageSecondsSumRate,
	// ResourcePod + CPUUserSecondsSumRate,
	// ResourcePod + CPUSystemSecondsSumRate,
	// ResourcePod + CPUCfsThrottledSecondsSumRate,
	// ResourcePod + MemoryUsageBytesSum,
	// ResourcePod + MemoryUsagePercent,
	// ResourcePod + DiskIOReadsBytesSumRate,
	// ResourcePod + DiskIOWritesBytesSumRate,
	// ResourcePod + FsByteSum,
	// ResourcePod + NetworkReceiveBytesSumRate,
	// ResourcePod + NetworkTransmitBytesSumRate,
	// ResourcePod + NetworkReceiveErrorsSumRate,
	// ResourcePod + NetworkTransmitErrorsSumRate,
	// ResourcePod + NetworkReceivePacketsSumRate,
	// ResourcePod + NetworkTransmitPacketsSumRate,
	// ResourcePod + NetworkReceivePacketsDroppedSumRate,
	// ResourcePod + NetworkTransmitPacketsDroppedSumRate,
	// ResourceContainer + CPUUsageSecondsSumRate,
	// ResourceContainer + CPUUserSecondsSumRate,
	// ResourceContainer + CPUSystemSecondsSumRate,
	// ResourceContainer + CPUCfsThrottledSecondsSumRate,
	// ResourceContainer + MemoryUsageBytesSum,
	// ResourceContainer + MemoryUsagePercent,
	// ResourceContainer + DiskIOReadsBytesSumRate,
	// ResourceContainer + DiskIOWritesBytesSumRate,
	// ResourceContainer + FsByteSum,
	}
)
