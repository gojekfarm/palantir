package palantir

import (
	"github.com/newrelic/go-agent"
	"github.com/spf13/viper"
)

func getNewRelicConfigOrPanic() newrelic.Config {
	config := newrelic.NewConfig(viper.GetString("NEW_RELIC_APP_NAME"), viper.GetString("NEW_RELIC_LICENCE_KEY"))
	config.Enabled = getFeature("NEW_RELIC_ENABLED")
	return config
}
