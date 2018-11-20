package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterAlertGroupType                       = "clusterAlertGroup"
	ClusterAlertGroupFieldAPIVersion            = "apiVersion"
	ClusterAlertGroupFieldAnnotations           = "annotations"
	ClusterAlertGroupFieldClusterID             = "clusterId"
	ClusterAlertGroupFieldClusterRules          = "clusterRules"
	ClusterAlertGroupFieldCode                  = "code"
	ClusterAlertGroupFieldCreated               = "created"
	ClusterAlertGroupFieldCreatorID             = "creatorId"
	ClusterAlertGroupFieldDescription           = "description"
	ClusterAlertGroupFieldDetails               = "details"
	ClusterAlertGroupFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	ClusterAlertGroupFieldGroupWaitSeconds      = "groupWaitSeconds"
	ClusterAlertGroupFieldKind                  = "kind"
	ClusterAlertGroupFieldLabels                = "labels"
	ClusterAlertGroupFieldListMeta              = "metadata"
	ClusterAlertGroupFieldMessage               = "message"
	ClusterAlertGroupFieldName                  = "name"
	ClusterAlertGroupFieldNamespaceId           = "namespaceId"
	ClusterAlertGroupFieldOwnerReferences       = "ownerReferences"
	ClusterAlertGroupFieldReason                = "reason"
	ClusterAlertGroupFieldRecipients            = "recipients"
	ClusterAlertGroupFieldRemoved               = "removed"
	ClusterAlertGroupFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	ClusterAlertGroupFieldState                 = "state"
	ClusterAlertGroupFieldStatus                = "status"
	ClusterAlertGroupFieldTransitioning         = "transitioning"
	ClusterAlertGroupFieldTransitioningMessage  = "transitioningMessage"
	ClusterAlertGroupFieldUUID                  = "uuid"
)

type ClusterAlertGroup struct {
	types.Resource
	APIVersion            string             `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	Annotations           map[string]string  `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClusterID             string             `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	ClusterRules          []ClusterAlertRule `json:"clusterRules,omitempty" yaml:"clusterRules,omitempty"`
	Code                  int64              `json:"code,omitempty" yaml:"code,omitempty"`
	Created               string             `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID             string             `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Description           string             `json:"description,omitempty" yaml:"description,omitempty"`
	Details               *StatusDetails     `json:"details,omitempty" yaml:"details,omitempty"`
	GroupIntervalSeconds  int64              `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64              `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	Kind                  string             `json:"kind,omitempty" yaml:"kind,omitempty"`
	Labels                map[string]string  `json:"labels,omitempty" yaml:"labels,omitempty"`
	ListMeta              *ListMeta          `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Message               string             `json:"message,omitempty" yaml:"message,omitempty"`
	Name                  string             `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId           string             `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences       []OwnerReference   `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Reason                string             `json:"reason,omitempty" yaml:"reason,omitempty"`
	Recipients            []Recipient        `json:"recipients,omitempty" yaml:"recipients,omitempty"`
	Removed               string             `json:"removed,omitempty" yaml:"removed,omitempty"`
	RepeatIntervalSeconds int64              `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	State                 string             `json:"state,omitempty" yaml:"state,omitempty"`
	Status                string             `json:"status,omitempty" yaml:"status,omitempty"`
	Transitioning         string             `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage  string             `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                  string             `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ClusterAlertGroupCollection struct {
	types.Collection
	Data   []ClusterAlertGroup `json:"data,omitempty"`
	client *ClusterAlertGroupClient
}

type ClusterAlertGroupClient struct {
	apiClient *Client
}

type ClusterAlertGroupOperations interface {
	List(opts *types.ListOpts) (*ClusterAlertGroupCollection, error)
	Create(opts *ClusterAlertGroup) (*ClusterAlertGroup, error)
	Update(existing *ClusterAlertGroup, updates interface{}) (*ClusterAlertGroup, error)
	Replace(existing *ClusterAlertGroup) (*ClusterAlertGroup, error)
	ByID(id string) (*ClusterAlertGroup, error)
	Delete(container *ClusterAlertGroup) error

	ActionActivate(resource *ClusterAlertGroup) error

	ActionDeactivate(resource *ClusterAlertGroup) error

	ActionMute(resource *ClusterAlertGroup) error

	ActionUnmute(resource *ClusterAlertGroup) error
}

func newClusterAlertGroupClient(apiClient *Client) *ClusterAlertGroupClient {
	return &ClusterAlertGroupClient{
		apiClient: apiClient,
	}
}

func (c *ClusterAlertGroupClient) Create(container *ClusterAlertGroup) (*ClusterAlertGroup, error) {
	resp := &ClusterAlertGroup{}
	err := c.apiClient.Ops.DoCreate(ClusterAlertGroupType, container, resp)
	return resp, err
}

func (c *ClusterAlertGroupClient) Update(existing *ClusterAlertGroup, updates interface{}) (*ClusterAlertGroup, error) {
	resp := &ClusterAlertGroup{}
	err := c.apiClient.Ops.DoUpdate(ClusterAlertGroupType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterAlertGroupClient) Replace(obj *ClusterAlertGroup) (*ClusterAlertGroup, error) {
	resp := &ClusterAlertGroup{}
	err := c.apiClient.Ops.DoReplace(ClusterAlertGroupType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ClusterAlertGroupClient) List(opts *types.ListOpts) (*ClusterAlertGroupCollection, error) {
	resp := &ClusterAlertGroupCollection{}
	err := c.apiClient.Ops.DoList(ClusterAlertGroupType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterAlertGroupCollection) Next() (*ClusterAlertGroupCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterAlertGroupCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterAlertGroupClient) ByID(id string) (*ClusterAlertGroup, error) {
	resp := &ClusterAlertGroup{}
	err := c.apiClient.Ops.DoByID(ClusterAlertGroupType, id, resp)
	return resp, err
}

func (c *ClusterAlertGroupClient) Delete(container *ClusterAlertGroup) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterAlertGroupType, &container.Resource)
}

func (c *ClusterAlertGroupClient) ActionActivate(resource *ClusterAlertGroup) error {
	err := c.apiClient.Ops.DoAction(ClusterAlertGroupType, "activate", &resource.Resource, nil, nil)
	return err
}

func (c *ClusterAlertGroupClient) ActionDeactivate(resource *ClusterAlertGroup) error {
	err := c.apiClient.Ops.DoAction(ClusterAlertGroupType, "deactivate", &resource.Resource, nil, nil)
	return err
}

func (c *ClusterAlertGroupClient) ActionMute(resource *ClusterAlertGroup) error {
	err := c.apiClient.Ops.DoAction(ClusterAlertGroupType, "mute", &resource.Resource, nil, nil)
	return err
}

func (c *ClusterAlertGroupClient) ActionUnmute(resource *ClusterAlertGroup) error {
	err := c.apiClient.Ops.DoAction(ClusterAlertGroupType, "unmute", &resource.Resource, nil, nil)
	return err
}
