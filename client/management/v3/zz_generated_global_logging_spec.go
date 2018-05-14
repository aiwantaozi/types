package client

const (
	GlobalLoggingSpecType                     = "globalLoggingSpec"
	GlobalLoggingSpecFieldDisplayName         = "displayName"
	GlobalLoggingSpecFieldElasticsearchConfig = "elasticsearchConfig"
	GlobalLoggingSpecFieldKafkaConfig         = "kafkaConfig"
	GlobalLoggingSpecFieldMySQLConfig         = "mySqlConfig"
	GlobalLoggingSpecFieldOutputFlushInterval = "outputFlushInterval"
	GlobalLoggingSpecFieldOutputTags          = "outputTags"
	GlobalLoggingSpecFieldSplunkConfig        = "splunkConfig"
	GlobalLoggingSpecFieldSyslogConfig        = "syslogConfig"
)

type GlobalLoggingSpec struct {
	DisplayName         string               `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	ElasticsearchConfig *ElasticsearchConfig `json:"elasticsearchConfig,omitempty" yaml:"elasticsearchConfig,omitempty"`
	KafkaConfig         *KafkaConfig         `json:"kafkaConfig,omitempty" yaml:"kafkaConfig,omitempty"`
	MySQLConfig         *MySQLConfig         `json:"mySqlConfig,omitempty" yaml:"mySqlConfig,omitempty"`
	OutputFlushInterval int64                `json:"outputFlushInterval,omitempty" yaml:"outputFlushInterval,omitempty"`
	OutputTags          map[string]string    `json:"outputTags,omitempty" yaml:"outputTags,omitempty"`
	SplunkConfig        *SplunkConfig        `json:"splunkConfig,omitempty" yaml:"splunkConfig,omitempty"`
	SyslogConfig        *SyslogConfig        `json:"syslogConfig,omitempty" yaml:"syslogConfig,omitempty"`
}
