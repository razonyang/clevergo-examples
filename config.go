package clevergo

import "os"

const (
	serverDefaultAddr = ":8080"
	// ServerTypeDefault means HTTP Application.
	ServerTypeDefault = 1
	// ServerTypeUNIX means UNIX Application.
	ServerTypeUNIX = 2
	// ServerTypeTLS means TLS Application.
	ServerTypeTLS = 3
	// ServerTypeTLSEmbed means TLSEmbed Application.
	ServerTypeTLSEmbed = 4
)

// Config for Application.
type Config struct {
	ServerAddr     string      // Server address.
	ServerType     int         // Server type.
	ServerMode     os.FileMode // Server mode for UNIX application.
	ServerCertFile string      // CertFile for TLS application.
	ServerKeyFile  string      // KeyFile  for TLS application.
	ServerCertData []byte      // CertData  for TLSEmbed application.
	ServerKeyData  []byte      // KeyData  for TLSEmbed application.
}

// NewConfig returns default configuration.
func NewConfig() *Config {
	return &Config{
		ServerAddr:     serverDefaultAddr,
		ServerType:     ServerTypeDefault,
		ServerCertFile: "",
		ServerKeyFile:  "",
	}
}

// IsServeUNIX returns a boolean indicating whether is UNIX Application.
func (c *Config) IsServeUNIX() bool {
	return c.ServerType == ServerTypeUNIX
}

// IsServeTLS returns a boolean indicating whether is TLS Application.
func (c *Config) IsServeTLS() bool {
	return c.ServerType == ServerTypeTLS
}

// IsServeTLSEmbed returns a boolean indicating whether is TLSEmbed Application.
func (c *Config) IsServeTLSEmbed() bool {
	return c.ServerType == ServerTypeTLSEmbed
}
