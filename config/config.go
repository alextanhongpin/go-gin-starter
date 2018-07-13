package config

import "github.com/spf13/viper"

// Config represents an alias to viper config
type Config = viper.Viper

// New returns a new pointer to the config
func New() *Config {
	v := viper.New()
	v.SetDefault("port", ":3030")
	v.SetDefault("usersvc_on", true)
	return v
}
