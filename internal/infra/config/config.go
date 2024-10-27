package config

import "github.com/spf13/viper"

type Config struct {
	source *viper.Viper
}

func NewConfig(source *viper.Viper) *Config {
	return &Config{source: source}
}

func (c *Config) Variation() float64 {
	return c.source.GetFloat64("VARIATION")
}

func (c *Config) Port() string {
	return c.source.GetString("PORT")
}

func (c *Config) LogLevel() string {
	return c.source.GetString("LOG_LEVEL")
}
