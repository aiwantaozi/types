package client

const (
	TLSSettingsType                   = "tlsSettings"
	TLSSettingsFieldCaCertificates    = "ca_certificates"
	TLSSettingsFieldClientCertificate = "client_certificate"
	TLSSettingsFieldMode              = "mode"
	TLSSettingsFieldPrivateKey        = "private_key"
	TLSSettingsFieldSni               = "sni"
	TLSSettingsFieldSubjectAltNames   = "subject_alt_names"
)

type TLSSettings struct {
	CaCertificates    string   `json:"ca_certificates,omitempty" yaml:"ca_certificates,omitempty"`
	ClientCertificate string   `json:"client_certificate,omitempty" yaml:"client_certificate,omitempty"`
	Mode              int64    `json:"mode,omitempty" yaml:"mode,omitempty"`
	PrivateKey        string   `json:"private_key,omitempty" yaml:"private_key,omitempty"`
	Sni               string   `json:"sni,omitempty" yaml:"sni,omitempty"`
	SubjectAltNames   []string `json:"subject_alt_names,omitempty" yaml:"subject_alt_names,omitempty"`
}
