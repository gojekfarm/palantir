package config

import (
	"fmt"

	"github.com/newrelic/go-agent"
	"github.com/spf13/viper"
)

type Config interface {
	GetValue(string) string
	GetIntValue(string) int
}

type configuration map[string]interface{}

var config configuration

type BaseConfig struct {
}

func (self BaseConfig) Load() {
	viper.SetDefault("port", "3000")
	viper.SetDefault("log_level", "warn")
	viper.SetDefault("redis_password", "")
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	config = configuration{}
	config["newrelic"] = getNewRelicConfigOrPanic()
}

func (self BaseConfig) Newrelic() newrelic.Config {
	return config["newrelic"].(newrelic.Config)
}

func (self BaseConfig) GetValue(key string) string {
	if _, ok := config[key]; !ok {
		config[key] = getStringOrPanic(key)
	}
	return config[key].(string)
}

func (self BaseConfig) GetOptionalValue(key string, defaultValue string) string {
	fmt.Println(config)
	if _, ok := config[key]; !ok {
		var value string
		if value = viper.GetString(key); !viper.IsSet(key) {
			value = defaultValue
		}
		config[key] = value
	}
	return config[key].(string)
}

func (self BaseConfig) GetIntValue(key string) int {
	if _, ok := config[key]; !ok {
		config[key] = getIntOrPanic(key)
	}
	return config[key].(int)
}

func (self BaseConfig) GetOptionalIntValue(key string, defaultValue int) int {
	if _, ok := config[key]; !ok {
		var value int
		if value = viper.GetInt(key); !viper.IsSet(key) {
			value = defaultValue
		}
		config[key] = value
	}
	return config[key].(int)
}

func (self BaseConfig) GetFeature(key string) bool {
	if _, ok := config[key]; !ok {
		config[key] = getFeature(key)
	}
	return config[key].(bool)
}
