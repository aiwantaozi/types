package client

import (
	"github.com/rancher/norman/types"
)

const (
	LoggingTargetType                      = "loggingTarget"
	LoggingTargetFieldAnnotations          = "annotations"
	LoggingTargetFieldAppliedSpec          = "appliedSpec"
	LoggingTargetFieldClusterIDs           = "clusterIds"
	LoggingTargetFieldConditions           = "conditions"
	LoggingTargetFieldCreated              = "created"
	LoggingTargetFieldCreatorID            = "creatorId"
	LoggingTargetFieldElasticsearchConfig  = "elasticsearchConfig"
	LoggingTargetFieldKafkaConfig          = "kafkaConfig"
	LoggingTargetFieldLabels               = "labels"
	LoggingTargetFieldMySQLConfig          = "mySqlConfig"
	LoggingTargetFieldName                 = "name"
	LoggingTargetFieldOutputFlushInterval  = "outputFlushInterval"
	LoggingTargetFieldOutputTags           = "outputTags"
	LoggingTargetFieldOwnerReferences      = "ownerReferences"
	LoggingTargetFieldRemoved              = "removed"
	LoggingTargetFieldSplunkConfig         = "splunkConfig"
	LoggingTargetFieldState                = "state"
	LoggingTargetFieldSyslogConfig         = "syslogConfig"
	LoggingTargetFieldTransitioning        = "transitioning"
	LoggingTargetFieldTransitioningMessage = "transitioningMessage"
	LoggingTargetFieldUuid                 = "uuid"
)

type LoggingTarget struct {
	types.Resource
	Annotations          map[string]string    `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AppliedSpec          *LoggingTargetSpec   `json:"appliedSpec,omitempty" yaml:"appliedSpec,omitempty"`
	ClusterIDs           []string             `json:"clusterIds,omitempty" yaml:"clusterIds,omitempty"`
	Conditions           []LoggingCondition   `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Created              string               `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string               `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	ElasticsearchConfig  *ElasticsearchConfig `json:"elasticsearchConfig,omitempty" yaml:"elasticsearchConfig,omitempty"`
	KafkaConfig          *KafkaConfig         `json:"kafkaConfig,omitempty" yaml:"kafkaConfig,omitempty"`
	Labels               map[string]string    `json:"labels,omitempty" yaml:"labels,omitempty"`
	MySQLConfig          *MySQLConfig         `json:"mySqlConfig,omitempty" yaml:"mySqlConfig,omitempty"`
	Name                 string               `json:"name,omitempty" yaml:"name,omitempty"`
	OutputFlushInterval  int64                `json:"outputFlushInterval,omitempty" yaml:"outputFlushInterval,omitempty"`
	OutputTags           map[string]string    `json:"outputTags,omitempty" yaml:"outputTags,omitempty"`
	OwnerReferences      []OwnerReference     `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed              string               `json:"removed,omitempty" yaml:"removed,omitempty"`
	SplunkConfig         *SplunkConfig        `json:"splunkConfig,omitempty" yaml:"splunkConfig,omitempty"`
	State                string               `json:"state,omitempty" yaml:"state,omitempty"`
	SyslogConfig         *SyslogConfig        `json:"syslogConfig,omitempty" yaml:"syslogConfig,omitempty"`
	Transitioning        string               `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string               `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	Uuid                 string               `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type LoggingTargetCollection struct {
	types.Collection
	Data   []LoggingTarget `json:"data,omitempty"`
	client *LoggingTargetClient
}

type LoggingTargetClient struct {
	apiClient *Client
}

type LoggingTargetOperations interface {
	List(opts *types.ListOpts) (*LoggingTargetCollection, error)
	Create(opts *LoggingTarget) (*LoggingTarget, error)
	Update(existing *LoggingTarget, updates interface{}) (*LoggingTarget, error)
	ByID(id string) (*LoggingTarget, error)
	Delete(container *LoggingTarget) error
}

func newLoggingTargetClient(apiClient *Client) *LoggingTargetClient {
	return &LoggingTargetClient{
		apiClient: apiClient,
	}
}

func (c *LoggingTargetClient) Create(container *LoggingTarget) (*LoggingTarget, error) {
	resp := &LoggingTarget{}
	err := c.apiClient.Ops.DoCreate(LoggingTargetType, container, resp)
	return resp, err
}

func (c *LoggingTargetClient) Update(existing *LoggingTarget, updates interface{}) (*LoggingTarget, error) {
	resp := &LoggingTarget{}
	err := c.apiClient.Ops.DoUpdate(LoggingTargetType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LoggingTargetClient) List(opts *types.ListOpts) (*LoggingTargetCollection, error) {
	resp := &LoggingTargetCollection{}
	err := c.apiClient.Ops.DoList(LoggingTargetType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *LoggingTargetCollection) Next() (*LoggingTargetCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &LoggingTargetCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *LoggingTargetClient) ByID(id string) (*LoggingTarget, error) {
	resp := &LoggingTarget{}
	err := c.apiClient.Ops.DoByID(LoggingTargetType, id, resp)
	return resp, err
}

func (c *LoggingTargetClient) Delete(container *LoggingTarget) error {
	return c.apiClient.Ops.DoResourceDelete(LoggingTargetType, &container.Resource)
}
