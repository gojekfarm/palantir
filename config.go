package palantir

import (
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
	self.LoadWithOptions(map[string]interface{}{})
}

func (self BaseConfig) LoadWithOptions(options map[string]interface{}) {
	viper.SetDefault("port", "3000")
	viper.SetDefault("log_level", "warn")
	viper.SetDefault("redis_password", "")
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	if options["configPath"] != nil {
		viper.AddConfigPath(options["configPath"].(string))
	} else {
		viper.AddConfigPath("./")
		viper.AddConfigPath("../")
	}
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	config = configuration{}
	if options["newrelic"] != nil && options["newrelic"].(bool) {
		config["newrelic"] = getNewRelicConfigOrPanic()
	}
	if options["db"] != nil && options["db"].(bool) {
		config["db_config"] = LoadDbConf()
	}
}

func (self BaseConfig) setTestDBUrl(dbConf *DBConfig) {
	dbConf.url = getStringOrPanic("db_url_test")
	dbConf.slaveUrl = getStringOrPanic("db_url_test")
}

func (self BaseConfig) LoadTestConfig(options map[string]interface{}) error {
	self.LoadWithOptions(options)
	if options["db"] != nil && options["db"].(bool) {
		self.setTestDBUrl(config["db_config"].(*DBConfig))
	}
	return nil
}

func (self BaseConfig) Newrelic() newrelic.Config {
	return config["newrelic"].(newrelic.Config)
}

func (self BaseConfig) DBConfig() *DBConfig {
	return config["db_config"].(*DBConfig)
}

func (self BaseConfig) GetValue(key string) string {
	if _, ok := config[key]; !ok {
		config[key] = getStringOrPanic(key)
	}
	return config[key].(string)
}

func (self BaseConfig) GetOptionalValue(key string, defaultValue string) string {
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
