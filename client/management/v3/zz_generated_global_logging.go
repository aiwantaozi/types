package client

import (
	"github.com/rancher/norman/types"
)

const (
	GlobalLoggingType                      = "globalLogging"
	GlobalLoggingFieldAnnotations          = "annotations"
	GlobalLoggingFieldAppliedSpec          = "appliedSpec"
	GlobalLoggingFieldClusterIDs           = "clusterIds"
	GlobalLoggingFieldConditions           = "conditions"
	GlobalLoggingFieldCreated              = "created"
	GlobalLoggingFieldCreatorID            = "creatorId"
	GlobalLoggingFieldElasticsearchConfig  = "elasticsearchConfig"
	GlobalLoggingFieldKafkaConfig          = "kafkaConfig"
	GlobalLoggingFieldLabels               = "labels"
	GlobalLoggingFieldMySQLConfig          = "mySqlConfig"
	GlobalLoggingFieldName                 = "name"
	GlobalLoggingFieldOutputFlushInterval  = "outputFlushInterval"
	GlobalLoggingFieldOutputTags           = "outputTags"
	GlobalLoggingFieldOwnerReferences      = "ownerReferences"
	GlobalLoggingFieldRemoved              = "removed"
	GlobalLoggingFieldSplunkConfig         = "splunkConfig"
	GlobalLoggingFieldState                = "state"
	GlobalLoggingFieldSyslogConfig         = "syslogConfig"
	GlobalLoggingFieldTransitioning        = "transitioning"
	GlobalLoggingFieldTransitioningMessage = "transitioningMessage"
	GlobalLoggingFieldUuid                 = "uuid"
)

type GlobalLogging struct {
	types.Resource
	Annotations          map[string]string    `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AppliedSpec          *GlobalLoggingSpec   `json:"appliedSpec,omitempty" yaml:"appliedSpec,omitempty"`
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
type GlobalLoggingCollection struct {
	types.Collection
	Data   []GlobalLogging `json:"data,omitempty"`
	client *GlobalLoggingClient
}

type GlobalLoggingClient struct {
	apiClient *Client
}

type GlobalLoggingOperations interface {
	List(opts *types.ListOpts) (*GlobalLoggingCollection, error)
	Create(opts *GlobalLogging) (*GlobalLogging, error)
	Update(existing *GlobalLogging, updates interface{}) (*GlobalLogging, error)
	ByID(id string) (*GlobalLogging, error)
	Delete(container *GlobalLogging) error
}

func newGlobalLoggingClient(apiClient *Client) *GlobalLoggingClient {
	return &GlobalLoggingClient{
		apiClient: apiClient,
	}
}

func (c *GlobalLoggingClient) Create(container *GlobalLogging) (*GlobalLogging, error) {
	resp := &GlobalLogging{}
	err := c.apiClient.Ops.DoCreate(GlobalLoggingType, container, resp)
	return resp, err
}

func (c *GlobalLoggingClient) Update(existing *GlobalLogging, updates interface{}) (*GlobalLogging, error) {
	resp := &GlobalLogging{}
	err := c.apiClient.Ops.DoUpdate(GlobalLoggingType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *GlobalLoggingClient) List(opts *types.ListOpts) (*GlobalLoggingCollection, error) {
	resp := &GlobalLoggingCollection{}
	err := c.apiClient.Ops.DoList(GlobalLoggingType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *GlobalLoggingCollection) Next() (*GlobalLoggingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &GlobalLoggingCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *GlobalLoggingClient) ByID(id string) (*GlobalLogging, error) {
	resp := &GlobalLogging{}
	err := c.apiClient.Ops.DoByID(GlobalLoggingType, id, resp)
	return resp, err
}

func (c *GlobalLoggingClient) Delete(container *GlobalLogging) error {
	return c.apiClient.Ops.DoResourceDelete(GlobalLoggingType, &container.Resource)
}
