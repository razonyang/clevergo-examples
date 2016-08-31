// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clevergo

import "os"

const (
	serverDefaultAddr  = ":8080"
	ServerTypeDefault  = 1
	ServerTypeUNIX     = 2
	ServerTypeTLS      = 3
	ServerTypeTLSEmbed = 4
)

type Config struct {
	ServerAddr     string
	ServerType     int
	ServerMode     os.FileMode
	ServerCertFile string
	ServerKeyFile  string
	ServerCertData []byte
	ServerKeyData  []byte
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

func (c *Config) IsServeUNIX() bool {
	return c.ServerType == ServerTypeUNIX
}

func (c *Config) IsServeTLS() bool {
	return c.ServerType == ServerTypeTLS
}

func (c *Config) IsServeTLSEmbed() bool {
	return c.ServerType == ServerTypeTLSEmbed
}
