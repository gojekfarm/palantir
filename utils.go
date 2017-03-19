package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

func getIntOrPanic(key string) int {
	checkKey(key)
	v, err := strconv.Atoi(viper.GetString(key))
	panicIfErrorForKey(err, key)
	return v
}

func getFeature(key string) bool {
	v, err := strconv.ParseBool(viper.GetString(key))
	if err != nil {
		return false
	}
	return v
}

func checkKey(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Errorf("%s key is not set", key))
	}
}

func panicIfErrorForKey(err error, key string) {
	if err != nil {
		panic(fmt.Errorf("Could not parse key: %s. Error: %v", key, err))
	}
}

func getStringOrPanic(key string) string {
	checkKey(key)
	return viper.GetString(key)
}
