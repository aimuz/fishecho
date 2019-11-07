package config

import (
	"context"
	"github.com/spf13/viper"
)

// config file
const (
	File = "config.file"
)

// server config
const (
	ServerHTTPAddr = "server.http-addr"
	ServerHTTPPort = "server.http-port"
	ServerHTTPGzip = "server.http-gzip"
)

// logger config
const (
	LoggerLevel  = "logger.level"
	LoggerOutput = "logger.output"
	LoggerFile   = "logger.file"
)

// database config
const (
	DatabaseType            = "database.type"
	DatabaseHost            = "database.host"
	DatabaseName            = "database.name"
	DatabaseUser            = "database.user"
	DatabasePassword        = "database.password"
	DatabaseURL             = "database.url"
	DatabaseMaxIdleConn     = "database.max-conn-idle"
	DatabaseMaxOpenConn     = "database.max-conn-open"
	DatabaseConnMaxLifetime = "database.max-conn-lifetime"
	DatabaseLogQueries      = "database.log-queries"
)

// cache config
const (
	CacheType   = "cache.type"
	CacheURL    = "cache.url"
	CacheEnable = "cache.enable"
)

type Config struct {
	Files    []string `yaml:"-"`
	Server   Server   `yaml:"server"`
	Logger   Logger   `yaml:"logger"`
	Database Database `yaml:"database"`
	Cache    Cache    `yaml:"cache"`
}

type viperLoad interface {
	viper() error
}

func (c *Config) Load(ctx context.Context) error {
	err := c.LoadWithViper(ctx)
	if err != nil {
		return err
	}
	if filename := viper.GetString(File); len(filename) > 0 {
		err = c.LoadWithFiles(ctx, filename)
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadWithViper load config with env or cli args
func (c *Config) LoadWithViper(_ context.Context) error {
	for _, v := range []viperLoad{&c.Server, &c.Logger, &c.Database, &c.Cache} {
		if err := v.viper(); err != nil {
			return err
		}
	}
	return nil
}

// LoadWithFiles load config with files
func (c *Config) LoadWithFiles(_ context.Context, file string) error {
	c.Files = append(c.Files, file)
	return nil
}

// Reload
func (c *Config) Reload(ctx context.Context) error {
	return nil
}

func init() {
	viper.SetEnvPrefix("FE")
	viper.AutomaticEnv()
}
