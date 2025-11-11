package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNoConfigFound(t *testing.T) {

	dir := t.TempDir()

	_, err := Load(dir)
	if !strings.Contains(err.Error(), "read config") {
		t.Errorf("unexpected error message: %v", err)
	}
}
func TestLoad_withDefaults(t *testing.T) {

	dir := t.TempDir()
	yaml := `
server:
  port: 8080
  mode: release
`
	os.WriteFile(filepath.Join(dir, "config.yaml"), []byte(yaml), 0644)
	config, err := Load(dir)
	if err != nil {
		t.Fatalf("expected error, got %v", err)
	}

	if config.Server.Port != 8080 {
		t.Fatalf("extened port 8080, got %v", config.Server.Port)
	}
}
