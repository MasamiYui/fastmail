package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary config file
	content := []byte(`
server:
  port: 9090
  token: "secret"
smtp:
  host: "smtp.example.com"
  port: 25
  user: "user"
  pass: "pass"
`)
	tmpfile, err := os.CreateTemp("", "config.*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test loading from file
	cfg, err := LoadConfig(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if cfg.Server.Port != 9090 {
		t.Errorf("Expected server port 9090, got %d", cfg.Server.Port)
	}
	if cfg.SMTP.Host != "smtp.example.com" {
		t.Errorf("Expected smtp host smtp.example.com, got %s", cfg.SMTP.Host)
	}

	// Test Env Override
	os.Setenv("FASTMAIL_SERVER_PORT", "10000")
	defer os.Unsetenv("FASTMAIL_SERVER_PORT")

	cfg, err = LoadConfig(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}
	if cfg.Server.Port != 10000 {
		t.Errorf("Expected port 10000 from env, got %d", cfg.Server.Port)
	}
}
