package config

import "github.com/spf13/viper"

type Cache struct {
	Type   string `yaml:"type"`
	URL    string `yaml:"url"`
	Enable bool   `yaml:"enable"`
}

func (c *Cache) viper() error {
	c.Type = viper.GetString(CacheType)
	c.URL = viper.GetString(CacheURL)
	c.Enable = viper.GetBool(CacheEnable)
	return nil
}
