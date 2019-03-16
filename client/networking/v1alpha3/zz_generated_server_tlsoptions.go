package client

const (
	Server_TLSOptionsType                    = "server_TLSOptions"
	Server_TLSOptionsFieldCaCertificates     = "ca_certificates"
	Server_TLSOptionsFieldCipherSuites       = "cipher_suites"
	Server_TLSOptionsFieldCredentialName     = "credential_name"
	Server_TLSOptionsFieldHttpsRedirect      = "https_redirect"
	Server_TLSOptionsFieldMaxProtocolVersion = "max_protocol_version"
	Server_TLSOptionsFieldMinProtocolVersion = "min_protocol_version"
	Server_TLSOptionsFieldMode               = "mode"
	Server_TLSOptionsFieldPrivateKey         = "private_key"
	Server_TLSOptionsFieldServerCertificate  = "server_certificate"
	Server_TLSOptionsFieldSubjectAltNames    = "subject_alt_names"
)

type Server_TLSOptions struct {
	CaCertificates     string   `json:"ca_certificates,omitempty" yaml:"ca_certificates,omitempty"`
	CipherSuites       []string `json:"cipher_suites,omitempty" yaml:"cipher_suites,omitempty"`
	CredentialName     string   `json:"credential_name,omitempty" yaml:"credential_name,omitempty"`
	HttpsRedirect      bool     `json:"https_redirect,omitempty" yaml:"https_redirect,omitempty"`
	MaxProtocolVersion int64    `json:"max_protocol_version,omitempty" yaml:"max_protocol_version,omitempty"`
	MinProtocolVersion int64    `json:"min_protocol_version,omitempty" yaml:"min_protocol_version,omitempty"`
	Mode               int64    `json:"mode,omitempty" yaml:"mode,omitempty"`
	PrivateKey         string   `json:"private_key,omitempty" yaml:"private_key,omitempty"`
	ServerCertificate  string   `json:"server_certificate,omitempty" yaml:"server_certificate,omitempty"`
	SubjectAltNames    []string `json:"subject_alt_names,omitempty" yaml:"subject_alt_names,omitempty"`
}
