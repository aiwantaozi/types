package client

const (
	ServiceEntryType                 = "serviceEntry"
	ServiceEntryFieldAddresses       = "addresses"
	ServiceEntryFieldEndpoints       = "endpoints"
	ServiceEntryFieldExportTo        = "export_to"
	ServiceEntryFieldHosts           = "hosts"
	ServiceEntryFieldLocation        = "location"
	ServiceEntryFieldPorts           = "ports"
	ServiceEntryFieldResolution      = "resolution"
	ServiceEntryFieldSubjectAltNames = "subject_alt_names"
)

type ServiceEntry struct {
	Addresses       []string                `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	Endpoints       []ServiceEntry_Endpoint `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
	ExportTo        []string                `json:"export_to,omitempty" yaml:"export_to,omitempty"`
	Hosts           []string                `json:"hosts,omitempty" yaml:"hosts,omitempty"`
	Location        int64                   `json:"location,omitempty" yaml:"location,omitempty"`
	Ports           []Port                  `json:"ports,omitempty" yaml:"ports,omitempty"`
	Resolution      int64                   `json:"resolution,omitempty" yaml:"resolution,omitempty"`
	SubjectAltNames []string                `json:"subject_alt_names,omitempty" yaml:"subject_alt_names,omitempty"`
}
