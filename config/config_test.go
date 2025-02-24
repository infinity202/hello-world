package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Test with environment variables
	expectedName := "Test App"
	expectedVersion := "2.0.0"
	expectedPort := "3000"
	expectedEnv := "test"

	// Set test environment variables
	os.Setenv("APP_NAME", expectedName)
	os.Setenv("APP_VERSION", expectedVersion)
	os.Setenv("APP_PORT", expectedPort)
	os.Setenv("APP_ENV", expectedEnv)

	// Load configuration
	cfg := LoadConfig()

	// Verify configuration values
	if cfg.AppName != expectedName {
		t.Errorf("Expected AppName %s, got %s", expectedName, cfg.AppName)
	}
	if cfg.AppVersion != expectedVersion {
		t.Errorf("Expected AppVersion %s, got %s", expectedVersion, cfg.AppVersion)
	}
	if cfg.AppPort != expectedPort {
		t.Errorf("Expected AppPort %s, got %s", expectedPort, cfg.AppPort)
	}
	if cfg.AppEnv != expectedEnv {
		t.Errorf("Expected AppEnv %s, got %s", expectedEnv, cfg.AppEnv)
	}
}

func TestLoadConfigDefaults(t *testing.T) {
	// Clear environment variables
	os.Unsetenv("APP_NAME")
	os.Unsetenv("APP_VERSION")
	os.Unsetenv("APP_PORT")
	os.Unsetenv("APP_ENV")

	// Load configuration
	cfg := LoadConfig()

	// Verify default values
	if cfg.AppName != "Go Hello World" {
		t.Errorf("Expected default AppName 'Go Hello World', got %s", cfg.AppName)
	}
	if cfg.AppVersion != "1.0.0" {
		t.Errorf("Expected default AppVersion '1.0.0', got %s", cfg.AppVersion)
	}
	if cfg.AppPort != "8080" {
		t.Errorf("Expected default AppPort '8080', got %s", cfg.AppPort)
	}
	if cfg.AppEnv != "development" {
		t.Errorf("Expected default AppEnv 'development', got %s", cfg.AppEnv)
	}
}

func TestGetEnv(t *testing.T) {
	// Test with existing environment variable
	expectedValue := "test-value"
	os.Setenv("TEST_KEY", expectedValue)
	value := getEnv("TEST_KEY", "default-value")
	if value != expectedValue {
		t.Errorf("Expected value %s, got %s", expectedValue, value)
	}

	// Test with non-existing environment variable
	os.Unsetenv("TEST_KEY")
	fallbackValue := "fallback-value"
	value = getEnv("TEST_KEY", fallbackValue)
	if value != fallbackValue {
		t.Errorf("Expected fallback value %s, got %s", fallbackValue, value)
	}
}
