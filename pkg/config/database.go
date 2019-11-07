package config

import (
	"github.com/spf13/viper"
	"time"
)

type Database struct {
	Type            string        `yaml:"type"`
	Host            string        `yaml:"host"`
	Name            string        `yaml:"name"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	MaxIdleConn     int64         `yaml:"maxIdleConn"`
	MaxOpenConn     int64         `yaml:"maxOpenConn"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
	LogQueries      bool          `yaml:"logQueries"`
	URL             string        `yaml:"url"`
}

func (d *Database) viper() error {
	d.Type = viper.GetString(DatabaseType)
	d.Host = viper.GetString(DatabaseHost)
	d.Name = viper.GetString(DatabaseName)
	d.User = viper.GetString(DatabaseUser)
	d.Password = viper.GetString(DatabasePassword)
	d.MaxIdleConn = viper.GetInt64(DatabaseMaxIdleConn)
	d.MaxOpenConn = viper.GetInt64(DatabaseMaxOpenConn)
	d.ConnMaxLifetime = viper.GetDuration(DatabaseConnMaxLifetime)
	d.LogQueries = viper.GetBool(DatabaseLogQueries)
	d.URL = viper.GetString(DatabaseURL)
	return nil
}
