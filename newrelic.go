package goconfig

import (
	"github.com/newrelic/go-agent"
	"github.com/spf13/viper"
)

func getNewRelicConfigOrPanic() newrelic.Config {
	config := newrelic.NewConfig(viper.GetString("new_relic_app_name"), viper.GetString("new_relic_licence_key"))
	config.Enabled = getFeature("new_relic_enabled")
	return config
}
