package client

import (
	"github.com/rancher/norman/types"
)

const (
	ProjectGraphType                 = "projectGraph"
	ProjectGraphFieldAnnotations     = "annotations"
	ProjectGraphFieldCreated         = "created"
	ProjectGraphFieldCreatorID       = "creatorId"
	ProjectGraphFieldLabels          = "labels"
	ProjectGraphFieldName            = "name"
	ProjectGraphFieldNamespaceId     = "namespaceId"
	ProjectGraphFieldOwnerReferences = "ownerReferences"
	ProjectGraphFieldProjectID       = "projectId"
	ProjectGraphFieldRemoved         = "removed"
	ProjectGraphFieldUUID            = "uuid"
)

type ProjectGraph struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId     string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	ProjectID       string            `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ProjectGraphCollection struct {
	types.Collection
	Data   []ProjectGraph `json:"data,omitempty"`
	client *ProjectGraphClient
}

type ProjectGraphClient struct {
	apiClient *Client
}

type ProjectGraphOperations interface {
	List(opts *types.ListOpts) (*ProjectGraphCollection, error)
	Create(opts *ProjectGraph) (*ProjectGraph, error)
	Update(existing *ProjectGraph, updates interface{}) (*ProjectGraph, error)
	Replace(existing *ProjectGraph) (*ProjectGraph, error)
	ByID(id string) (*ProjectGraph, error)
	Delete(container *ProjectGraph) error
}

func newProjectGraphClient(apiClient *Client) *ProjectGraphClient {
	return &ProjectGraphClient{
		apiClient: apiClient,
	}
}

func (c *ProjectGraphClient) Create(container *ProjectGraph) (*ProjectGraph, error) {
	resp := &ProjectGraph{}
	err := c.apiClient.Ops.DoCreate(ProjectGraphType, container, resp)
	return resp, err
}

func (c *ProjectGraphClient) Update(existing *ProjectGraph, updates interface{}) (*ProjectGraph, error) {
	resp := &ProjectGraph{}
	err := c.apiClient.Ops.DoUpdate(ProjectGraphType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ProjectGraphClient) Replace(obj *ProjectGraph) (*ProjectGraph, error) {
	resp := &ProjectGraph{}
	err := c.apiClient.Ops.DoReplace(ProjectGraphType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ProjectGraphClient) List(opts *types.ListOpts) (*ProjectGraphCollection, error) {
	resp := &ProjectGraphCollection{}
	err := c.apiClient.Ops.DoList(ProjectGraphType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ProjectGraphCollection) Next() (*ProjectGraphCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ProjectGraphCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ProjectGraphClient) ByID(id string) (*ProjectGraph, error) {
	resp := &ProjectGraph{}
	err := c.apiClient.Ops.DoByID(ProjectGraphType, id, resp)
	return resp, err
}

func (c *ProjectGraphClient) Delete(container *ProjectGraph) error {
	return c.apiClient.Ops.DoResourceDelete(ProjectGraphType, &container.Resource)
}
