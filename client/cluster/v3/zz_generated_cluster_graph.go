package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterGraphType                      = "clusterGraph"
	ClusterGraphFieldAnnotations          = "annotations"
	ClusterGraphFieldCreated              = "created"
	ClusterGraphFieldCreatorID            = "creatorId"
	ClusterGraphFieldLabels               = "labels"
	ClusterGraphFieldMonitorGraphSelector = "monitorGraphSelector"
	ClusterGraphFieldName                 = "name"
	ClusterGraphFieldNamespaceId          = "namespaceId"
	ClusterGraphFieldOwnerReferences      = "ownerReferences"
	ClusterGraphFieldRemoved              = "removed"
	ClusterGraphFieldState                = "state"
	ClusterGraphFieldStatus               = "status"
	ClusterGraphFieldTransitioning        = "transitioning"
	ClusterGraphFieldTransitioningMessage = "transitioningMessage"
	ClusterGraphFieldUUID                 = "uuid"
)

type ClusterGraph struct {
	types.Resource
	Annotations          map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created              string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels               map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	MonitorGraphSelector map[string]string `json:"monitorGraphSelector,omitempty" yaml:"monitorGraphSelector,omitempty"`
	Name                 string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId          string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed              string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	State                string            `json:"state,omitempty" yaml:"state,omitempty"`
	Status               *StatsStatus      `json:"status,omitempty" yaml:"status,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                 string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ClusterGraphCollection struct {
	types.Collection
	Data   []ClusterGraph `json:"data,omitempty"`
	client *ClusterGraphClient
}

type ClusterGraphClient struct {
	apiClient *Client
}

type ClusterGraphOperations interface {
	List(opts *types.ListOpts) (*ClusterGraphCollection, error)
	Create(opts *ClusterGraph) (*ClusterGraph, error)
	Update(existing *ClusterGraph, updates interface{}) (*ClusterGraph, error)
	Replace(existing *ClusterGraph) (*ClusterGraph, error)
	ByID(id string) (*ClusterGraph, error)
	Delete(container *ClusterGraph) error
}

func newClusterGraphClient(apiClient *Client) *ClusterGraphClient {
	return &ClusterGraphClient{
		apiClient: apiClient,
	}
}

func (c *ClusterGraphClient) Create(container *ClusterGraph) (*ClusterGraph, error) {
	resp := &ClusterGraph{}
	err := c.apiClient.Ops.DoCreate(ClusterGraphType, container, resp)
	return resp, err
}

func (c *ClusterGraphClient) Update(existing *ClusterGraph, updates interface{}) (*ClusterGraph, error) {
	resp := &ClusterGraph{}
	err := c.apiClient.Ops.DoUpdate(ClusterGraphType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterGraphClient) Replace(obj *ClusterGraph) (*ClusterGraph, error) {
	resp := &ClusterGraph{}
	err := c.apiClient.Ops.DoReplace(ClusterGraphType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ClusterGraphClient) List(opts *types.ListOpts) (*ClusterGraphCollection, error) {
	resp := &ClusterGraphCollection{}
	err := c.apiClient.Ops.DoList(ClusterGraphType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterGraphCollection) Next() (*ClusterGraphCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterGraphCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterGraphClient) ByID(id string) (*ClusterGraph, error) {
	resp := &ClusterGraph{}
	err := c.apiClient.Ops.DoByID(ClusterGraphType, id, resp)
	return resp, err
}

func (c *ClusterGraphClient) Delete(container *ClusterGraph) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterGraphType, &container.Resource)
}
