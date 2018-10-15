package client

import "k8s.io/apimachinery/pkg/util/intstr"

const (
	EndpointType               = "endpoint"
	EndpointFieldInterval      = "interval"
	EndpointFieldParams        = "params"
	EndpointFieldPath          = "path"
	EndpointFieldPort          = "port"
	EndpointFieldScheme        = "scheme"
	EndpointFieldScrapeTimeout = "scrapeTimeout"
	EndpointFieldTargetPort    = "targetPort"
)

type Endpoint struct {
	Interval      string              `json:"interval,omitempty" yaml:"interval,omitempty"`
	Params        map[string][]string `json:"params,omitempty" yaml:"params,omitempty"`
	Path          string              `json:"path,omitempty" yaml:"path,omitempty"`
	Port          string              `json:"port,omitempty" yaml:"port,omitempty"`
	Scheme        string              `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	ScrapeTimeout string              `json:"scrapeTimeout,omitempty" yaml:"scrapeTimeout,omitempty"`
	TargetPort    intstr.IntOrString  `json:"targetPort,omitempty" yaml:"targetPort,omitempty"`
}
