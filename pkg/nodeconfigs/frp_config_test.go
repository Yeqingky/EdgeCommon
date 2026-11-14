package nodeconfigs

import "testing"

func TestFRPConfigInit(t *testing.T) {
	config := &FRPConfig{}
	if err := config.Init(); err != nil {
		t.Fatal(err)
	}
	if config.BindAddr != "0.0.0.0" {
		t.Fatalf("unexpected bind address: %s", config.BindAddr)
	}
	if config.BindPort != 7000 {
		t.Fatalf("unexpected bind port: %d", config.BindPort)
	}
}

func TestFRPConfigInvalidPort(t *testing.T) {
	config := &FRPConfig{BindPort: 70000}
	if err := config.Init(); err == nil {
		t.Fatal("expected an invalid port error")
	}
}
