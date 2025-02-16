package config

type Config struct {
	Cluster struct {
		Server                   string `toml:"server"`
		CertificateAuthorityData string `toml:"certificate_authority_data"`
	} `toml:"cluster"`

	Auth struct {
		ClientCertificateData string `toml:"client_certificate_data"`
		ClientKeyData         string `toml:"client_key_data"`
	} `toml:"auth"`

	Namespace struct {
		Default string `toml:"default"`
	} `toml:"namespace"`

  Display struct {
    Refresh string `toml:"refresh"`
  } `toml:"display"`
}
