package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterAuditLoggingType                      = "clusterAuditLogging"
	ClusterAuditLoggingFieldAnnotations          = "annotations"
	ClusterAuditLoggingFieldClusterId            = "clusterId"
	ClusterAuditLoggingFieldCreated              = "created"
	ClusterAuditLoggingFieldCreatorID            = "creatorId"
	ClusterAuditLoggingFieldGlobalLoggingName    = "globalLogging"
	ClusterAuditLoggingFieldLabels               = "labels"
	ClusterAuditLoggingFieldName                 = "name"
	ClusterAuditLoggingFieldNamespaceId          = "namespaceId"
	ClusterAuditLoggingFieldOwnerReferences      = "ownerReferences"
	ClusterAuditLoggingFieldRemoved              = "removed"
	ClusterAuditLoggingFieldState                = "state"
	ClusterAuditLoggingFieldStatus               = "status"
	ClusterAuditLoggingFieldTransitioning        = "transitioning"
	ClusterAuditLoggingFieldTransitioningMessage = "transitioningMessage"
	ClusterAuditLoggingFieldUuid                 = "uuid"
)

type ClusterAuditLogging struct {
	types.Resource
	Annotations          map[string]string          `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClusterId            string                     `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	Created              string                     `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string                     `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	GlobalLoggingName    string                     `json:"globalLogging,omitempty" yaml:"globalLogging,omitempty"`
	Labels               map[string]string          `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                 string                     `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId          string                     `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences      []OwnerReference           `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed              string                     `json:"removed,omitempty" yaml:"removed,omitempty"`
	State                string                     `json:"state,omitempty" yaml:"state,omitempty"`
	Status               *ClusterAuditLoggingStatus `json:"status,omitempty" yaml:"status,omitempty"`
	Transitioning        string                     `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string                     `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	Uuid                 string                     `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type ClusterAuditLoggingCollection struct {
	types.Collection
	Data   []ClusterAuditLogging `json:"data,omitempty"`
	client *ClusterAuditLoggingClient
}

type ClusterAuditLoggingClient struct {
	apiClient *Client
}

type ClusterAuditLoggingOperations interface {
	List(opts *types.ListOpts) (*ClusterAuditLoggingCollection, error)
	Create(opts *ClusterAuditLogging) (*ClusterAuditLogging, error)
	Update(existing *ClusterAuditLogging, updates interface{}) (*ClusterAuditLogging, error)
	ByID(id string) (*ClusterAuditLogging, error)
	Delete(container *ClusterAuditLogging) error
}

func newClusterAuditLoggingClient(apiClient *Client) *ClusterAuditLoggingClient {
	return &ClusterAuditLoggingClient{
		apiClient: apiClient,
	}
}

func (c *ClusterAuditLoggingClient) Create(container *ClusterAuditLogging) (*ClusterAuditLogging, error) {
	resp := &ClusterAuditLogging{}
	err := c.apiClient.Ops.DoCreate(ClusterAuditLoggingType, container, resp)
	return resp, err
}

func (c *ClusterAuditLoggingClient) Update(existing *ClusterAuditLogging, updates interface{}) (*ClusterAuditLogging, error) {
	resp := &ClusterAuditLogging{}
	err := c.apiClient.Ops.DoUpdate(ClusterAuditLoggingType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterAuditLoggingClient) List(opts *types.ListOpts) (*ClusterAuditLoggingCollection, error) {
	resp := &ClusterAuditLoggingCollection{}
	err := c.apiClient.Ops.DoList(ClusterAuditLoggingType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterAuditLoggingCollection) Next() (*ClusterAuditLoggingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterAuditLoggingCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterAuditLoggingClient) ByID(id string) (*ClusterAuditLogging, error) {
	resp := &ClusterAuditLogging{}
	err := c.apiClient.Ops.DoByID(ClusterAuditLoggingType, id, resp)
	return resp, err
}

func (c *ClusterAuditLoggingClient) Delete(container *ClusterAuditLogging) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterAuditLoggingType, &container.Resource)
}
