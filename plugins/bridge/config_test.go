package bridge

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_Success(t *testing.T) {
	config, err := LoadConfig([]byte(`{"name": "test-bridge", "mtu": 1500}`))
	assert.NoError(t, err)
	assert.Equal(t, "test-bridge", config.Name)
	assert.Equal(t, 1500, config.MTU)
}

func TestLoadConfig_InvalidJSON(t *testing.T) {
	_, err := LoadConfig([]byte(`invalid-json`))
	assert.Error(t, err)
}

func TestLoadConfig_MissingName(t *testing.T) {
	_, err := LoadConfig([]byte(`{"mtu": 1500}`))
	assert.Error(t, err)
}

func TestValidateConfig_Valid(t *testing.T) {
	config := &Config{Name: "test-bridge", MTU: 1500}
	err := ValidateConfig(config)
	assert.NoError(t, err)
}

func TestValidateConfig_InvalidMTU(t *testing.T) {
	config := &Config{Name: "test-bridge", MTU: -1}
	err := ValidateConfig(config)
	assert.Error(t, err)
}

func TestValidateConfig_EmptyName(t *testing.T) {
	config := &Config{Name: "", MTU: 1500}
	err := ValidateConfig(config)
	assert.Error(t, err)
}

func TestValidateConfig_MaxMTU(t *testing.T) {
	config := &Config{Name: "test-bridge", MTU: 65535}
	err := ValidateConfig(config)
	assert.NoError(t, err)
}