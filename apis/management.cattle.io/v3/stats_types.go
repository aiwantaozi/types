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
	Cronjob            = "cronjob"
	Daemonset          = "daemonset"
	Deployment         = "deployment"
	Endpoint           = "endpoint"
	Pvcprotection      = "pvcprotection"
	Pvprotection       = "pvprotection"
	Replicaset         = "replicaset"
	Replicationmanager = "replicationmanager"
	Service            = "service"
	ServiceAccount     = "serviceaccount"
	Statefulset        = "statefulset"
	Volumes            = "volumes"
)

const (
	CPU         = "cpu"
	ClientQuery = "client"
	DB          = "db"
	Disk        = "disk"
	DiskIO      = "disk_io"
	Filesystem  = "fs"
	GRPC        = "grpc"
	Load1       = "load1"
	Load5       = "load5"
	Load15      = "load15"
	Memory      = "memory"
	Network     = "network"

	Controller = "controller"
	Helper     = "helper"
	System     = "system"
	Server     = "server"
	Usage      = "usage"
	UserQuery  = "user"

	Receive  = "receive"
	Received = "received"
	Read     = "read"
	Reads    = "reads"
	Sent     = "sent"
	Transmit = "transmit"
	Writes   = "writes"

	CfsThrottled = "cfs_throttled"
	PageOut      = "page_out"
	PageIn       = "page_in"
	WorkingSet   = "working_set"

	IOReads  = "io_reads"
	IOWrites = "io_writes"

	Drop           = "drop"
	Errors         = "errors"
	Errs           = "errs"
	Packets        = "packets"
	PacketsDropped = "packets_dropped"

	Adds           = "adds"
	Depth          = "depth"
	QueueRetries   = "queue_retries"
	Retries        = "retries"
	RateLimiterUse = "rate_limiter_use"

	CommitDuration   = "commit_duration"
	QueueDuration    = "queue_duration"
	SnapshotDuration = "snapshot_duration"
	WorkDuration     = "work_duration"
	WalFsyncDuration = "wal_fsync_duration"

	AddLatencies               = "add_latencies"
	BindingLatency             = "binding_latency"
	GetLatencies               = "get_latencies"
	QueueLatency               = "queue_latency"
	SchedulingLatency          = "scheduling_latency"
	SchedulingAlgorithmLatency = "scheduling_algorithm_latency"

	Request        = "request"
	RequestCount   = "request_count"
	RequestError   = "request_error"
	RequestLatency = "request_latency"
	RequestSlow    = "request_slow"
	RequestCache   = "request_cache"

	Cache = "cache"
	Entry = "entry"
	Hit   = "hit"
	Miss  = "miss"

	Leader              = "leader"
	LeaderChangesSeen   = "leader_changes_seen"
	InsufficientMembers = "insufficient_members"

	ProposalsCommitted = "proposals_committed"
	ProposalsApplied   = "proposals_applied"
	ProposalsFailed    = "proposals_failed"
	ProposalsPending   = "proposals_pending"

	AuditEvent                         = "AuditEvent"
	APIServiceRegistrationController   = "APIServiceRegistrationController"
	APIServiceRegistration             = "APIServiceRegistration"
	ClientCertificateExpiration        = "client_certificate_expiration"
	E2e                                = "e2e"
	NamespaceContainerMemoryUsageBytes = "namespace_container_memory_usage_bytes"
	NamespaceContainerSpecCPUShares    = "namespace_container_spec_cpu_shares"
	NamespaceContainerCPUUsage         = "namespace_container_cpu_usage"
	ResidentMemory                     = "residentMemory"
)

const (
	Avg      = "avg"
	Count    = "count"
	Increase = "increase"
	Percent  = "percent"
	Quantile = "quantile"
	Rate     = "rate"
	Sum      = "sum"
	Summary  = "summary"
	Total    = "total"
)

const (
	Bytes        = "bytes"
	Milliseconds = "milliseconds"
	Seconds      = "seconds"
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
