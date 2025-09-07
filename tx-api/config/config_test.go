package config

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempConfig(t *testing.T, content string) string {
	t.Helper()
	tmpFile := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp config: %v", err)
	}
	return tmpFile
}

func TestLoad_SuccessWithEnv(t *testing.T) {
	yamlContent := `
local:
  database:
    host: "localhost"
    port: 5432
    user: "user"
    password: "pass"
    name: "dbname"
    schema: "public"
test:
  database:
    host: "testhost"
    port: 3306
    user: "testuser"
    password: "testpass"
    name: "testdb"
    schema: "testschema"
`
	tmpFile := writeTempConfig(t, yamlContent)
	os.Setenv("APP_ENV", "test")
	defer os.Unsetenv("APP_ENV")

	// point Load() to temp file
	oldWd, _ := os.Getwd()
	os.Chdir(filepath.Dir(tmpFile))
	defer os.Chdir(oldWd)

	conf = nil // reset
	if err := Load(); err != nil {
		t.Fatalf("Load() failed: %v", err)
	}
	if conf.Database.Host != "testhost" {
		t.Errorf("expected host=testhost, got %s", conf.Database.Host)
	}
}

func TestLoad_DefaultEnv(t *testing.T) {
	yamlContent := `
local:
  database:
    host: "localhost"
    port: 5432
    user: "user"
    password: "pass"
    name: "dbname"
    schema: "public"
`
	tmpFile := writeTempConfig(t, yamlContent)
	os.Unsetenv("APP_ENV")

	oldWd, _ := os.Getwd()
	os.Chdir(filepath.Dir(tmpFile))
	defer os.Chdir(oldWd)

	conf = nil
	if err := Load(); err != nil {
		t.Fatalf("Load() failed: %v", err)
	}
	if conf.Database.Host != "localhost" {
		t.Errorf("expected host=localhost, got %s", conf.Database.Host)
	}
}

func TestLoad_FileMissing(t *testing.T) {
	os.Unsetenv("APP_ENV")
	conf = nil
	err := Load()
	if err == nil || err.Error()[:21] != "failed to read config" {
		t.Errorf("expected file missing error, got %v", err)
	}
}

func TestLoad_InvalidYAML(t *testing.T) {
	tmpFile := writeTempConfig(t, `: bad yaml :`)
	oldWd, _ := os.Getwd()
	os.Chdir(filepath.Dir(tmpFile))
	defer os.Chdir(oldWd)

	conf = nil
	err := Load()
	if err == nil {
		t.Errorf("expected yaml unmarshal error")
	}
}

func TestLoad_EnvNotFound(t *testing.T) {
	yamlContent := `
local:
  database:
    host: "localhost"
`
	tmpFile := writeTempConfig(t, yamlContent)
	os.Setenv("APP_ENV", "staging")
	defer os.Unsetenv("APP_ENV")

	oldWd, _ := os.Getwd()
	os.Chdir(filepath.Dir(tmpFile))
	defer os.Chdir(oldWd)

	conf = nil
	err := Load()
	if err == nil || err.Error() != `environment "staging" not found in config.yaml` {
		t.Errorf("expected env not found error, got %v", err)
	}
}

func TestGet_UsesCache(t *testing.T) {
	conf = &Config{Database: Database{Host: "cachedhost"}}
	c, err := Get()
	if err != nil {
		t.Fatalf("Get() failed: %v", err)
	}
	if c.Database.Host != "cachedhost" {
		t.Errorf("expected cachedhost, got %s", c.Database.Host)
	}
}

func TestGet_CallsLoad(t *testing.T) {
	yamlContent := `
local:
  database:
    host: "fromget"
`
	tmpFile := writeTempConfig(t, yamlContent)
	os.Unsetenv("APP_ENV")

	oldWd, _ := os.Getwd()
	os.Chdir(filepath.Dir(tmpFile))
	defer os.Chdir(oldWd)

	conf = nil
	c, err := Get()
	if err != nil {
		t.Fatalf("Get() failed: %v", err)
	}
	if c.Database.Host != "fromget" {
		t.Errorf("expected fromget, got %s", c.Database.Host)
	}
}
