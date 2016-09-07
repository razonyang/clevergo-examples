package clevergo

import "testing"

func TestNewConfig(t *testing.T) {
	c := NewConfig()

	if c.ServerAddr != serverDefaultAddr {
		t.Errorf("c.ServerAddr = %s, expect %s.", c.ServerAddr, serverDefaultAddr)
	}

	if c.ServerType != ServerTypeDefault {
		t.Errorf("c.ServerType = %v, expect %v.", c.ServerType, ServerTypeDefault)
	}
}

func TestConfig_IsServeTLS(t *testing.T) {
	c := Config{
		ServerType: ServerTypeTLS,
	}

	if c.IsServeTLS() == false {
		t.Errorf("c.IsServeTLS() = %v, expect true.", c.IsServeTLS())
	}
}

func TestConfig_IsServeTLSEmbed(t *testing.T) {
	c := Config{
		ServerType: ServerTypeTLSEmbed,
	}

	if c.IsServeTLSEmbed() == false {
		t.Errorf("c.IsServeTLSEmbed() = %v, expect true.", c.IsServeTLSEmbed())
	}
}

func TestConfig_IsServeUNIX(t *testing.T) {
	c := Config{
		ServerType: ServerTypeUNIX,
	}

	if c.IsServeUNIX() == false {
		t.Errorf("c.IsServeUNIX() = %v, expect true.", c.IsServeUNIX())
	}
}
