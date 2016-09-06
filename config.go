package clevergo

import "os"

const (
	serverDefaultAddr  = ":8080"
	ServerTypeDefault  = 1 // HTTP Application.
	ServerTypeUNIX     = 2 // UNIX Application.
	ServerTypeTLS      = 3 // TLS Application.
	ServerTypeTLSEmbed = 4 // TLSEmbed Application.
)

// Application configuration.
type Config struct {
	ServerAddr     string      // Server address.
	ServerType     int         // Server type.
	ServerMode     os.FileMode // Server mode for UNIX application.
	ServerCertFile string      // CertFile for TLS application.
	ServerKeyFile  string      // KeyFile  for TLS application.
	ServerCertData []byte      // CertData  for TLSEmbed application.
	ServerKeyData  []byte      // KeyData  for TLSEmbed application.
}

// Returns default configuration.
func NewConfig() *Config {
	return &Config{
		ServerAddr:     serverDefaultAddr,
		ServerType:     ServerTypeDefault,
		ServerCertFile: "",
		ServerKeyFile:  "",
	}
}

// Returns a boolean indicating whether is UNIX Application.
func (c *Config) IsServeUNIX() bool {
	return c.ServerType == ServerTypeUNIX
}

// Returns a boolean indicating whether is TLS Application.
func (c *Config) IsServeTLS() bool {
	return c.ServerType == ServerTypeTLS
}

// Returns a boolean indicating whether is TLSEmbed Application.
func (c *Config) IsServeTLSEmbed() bool {
	return c.ServerType == ServerTypeTLSEmbed
}
