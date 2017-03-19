package config_test

import (
	"config"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AppConfig struct {
	config.BaseConfig
}

func TestShouldSetDefaultForPort(t *testing.T) {
	config := &AppConfig{}
	config.Load()
	assert.Equal(t, "3000", config.GetValue("port"))
}

func TestShouldSetDefaultForLogLevel(t *testing.T) {
	config := &AppConfig{}
	config.Load()
	assert.Equal(t, "warn", config.GetValue("log_level"))
}

func TestShouldSetDefaultForRedisPassword(t *testing.T) {
	config := &AppConfig{}
	config.Load()
	assert.Equal(t, "", config.GetValue("redis_password"))
}

func TestShouldSetNewRelicBasedOnApplicationConfig(t *testing.T) {
	config := &AppConfig{}
	config.Load()
	assert.Equal(t, "foo", config.Newrelic().AppName)
	assert.Equal(t, "bar", config.Newrelic().License)
}

func TestShouldGetValueBasedOnApplicationConfig(t *testing.T) {
	config := &AppConfig{}
	config.Load()
	assert.Equal(t, "bar", config.GetValue("foo"))
}

func TestShouldGetOptionalValueBasedOnApplicationConfig(t *testing.T) {
	config := &AppConfig{}
	config.Load()
	assert.Equal(t, "bar", config.GetOptionalValue("foo", "baz"))
	assert.Equal(t, "baz", config.GetOptionalValue("bar", "baz"))
}

func TestShouldGetIntValueBasedOnApplicationConfig(t *testing.T) {
	config := &AppConfig{}
	config.Load()
	assert.Equal(t, 1, config.GetIntValue("someInt"))
}

func TestShouldGetOptionalIntValueBasedOnApplicationConfig(t *testing.T) {
	config := &AppConfig{}
	config.Load()
	assert.Equal(t, 1, config.GetOptionalIntValue("someInt", 10))
	assert.Equal(t, 10, config.GetOptionalIntValue("bar", 10))
}

func TestShouldGetFeature(t *testing.T) {
	config := &AppConfig{}
	config.Load()
	assert.True(t, config.GetFeature("someFeature"))
	assert.False(t, config.GetFeature("someOtherFeature"))
	assert.False(t, config.GetFeature("someUnknownFeature"))
}
