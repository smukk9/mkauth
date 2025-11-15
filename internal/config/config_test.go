package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoad_EnvConfigNotSet(t *testing.T) {

	_, err := Load()
	if !errors.Is(err, ErrEnvVarNotSet) {
		t.Fatalf("expected error, got %v", err)
	}
}
func TestLoad_withDefaults(t *testing.T) {

	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	yaml := `
server:
  port: 8088
  host: 0.0.0.0
  mode: release
  service: MkAuth
  version: v0.1.0

database:
  path: ./pb_data

admin:
  email: admin@mkauth.local
  password: changeme123

`
	os.WriteFile(path, []byte(yaml), 0644)

	os.Setenv("MKAUTH_FILE", path)

	config, err := Load()
	if err != nil {
		t.Fatalf("expected error, got %v", err)
	}

	if config.Server.Port != 8088 {
		t.Fatalf("extened port 8080, got %v", config.Server.Port)
	}
}

func TestLoad_YAMLParsingError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")

	// Intentionally invalid YAML
	yaml := `
	server:
	port: 8088
	host: 0.0.0.0
	mode: release
	service: MkAuth
	version: v0.1.0

	database:
	path: ./pb_data

	admin:
	email: admin@mkauth.local
	password: changeme123

`
	os.WriteFile(path, []byte(yaml), 0644)
	os.Setenv("MKAUTH_FILE", path)

	_, err := Load()
	if err == nil {
		t.Fatalf("expected YAML parsing error, got nil")
	}

	if !strings.Contains(err.Error(), "yaml") {
		t.Errorf("unexpected error message: %v", err)
	}
}
