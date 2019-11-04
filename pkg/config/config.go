package config

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/aimuz/fishecho/pkg/register"
	"github.com/spf13/viper"
)

// config file
const (
	File = "config.file"
)

// server config
const (
	ServerHTTPAddr = "server.http.addr"
	ServerHTTPPort = "server.http.port"
	ServerGzip     = "server.gzip"
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
	DatabaseMaxIdleConn     = "database.max.conn.idle"
	DatabaseMaxOpenConn     = "database.max.conn.open"
	DatabaseConnMaxLifetime = "database.max.conn.lifetime"
	DatabaseLogQueries      = "database.log.queries"
)

// cache config
const (
	CacheType = "cache.type"
	CacheURL  = "cache.url"
)

func init() {
	viper.SetEnvPrefix("FE")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	register.Initer("config", Load)
}

func Load(_ context.Context) error {
	if filename := viper.GetString(File); len(filename) > 0 {
		viper.SetConfigType(filepath.Ext(filename))

		f, err := os.Open(filename)
		if err != nil {
			return nil
		}
		defer f.Close()
		err = viper.MergeConfig(f)
		if err != nil {
			return err
		}
	}
	return nil
}

//func defaultConfig() {
//	// server config default
//	viper.SetDefault(ServerHTTPAddr, "0.0.0.0")
//	viper.SetDefault(ServerHTTPPort, 3002)
//	viper.SetDefault(ServerGzip, false)
//
//	// logger config default
//	viper.SetDefault(LoggerLevel, "info")
//	viper.SetDefault(LoggerOutput, "stderr")
//	viper.SetDefault(LoggerFile, "")
//
//	// database config default
//	viper.SetDefault(DatabaseType, "")
//	viper.SetDefault(DatabaseHost, "")
//	viper.SetDefault(DatabaseName, "fishecho")
//	viper.SetDefault(DatabaseUser, "")
//	viper.SetDefault(DatabasePassword, "")
//	viper.SetDefault(DatabaseURL, "mysql://user:secret@host:port/database")
//	viper.SetDefault(DatabaseMaxOpenConn, 0)
//	viper.SetDefault(DatabaseMaxIdleConn, 2)
//	viper.SetDefault(DatabaseConnMaxLifetime, time.Hour*4)
//	viper.SetDefault(DatabaseLogQueries, false)
//
//	// cache config default
//	viper.SetDefault(CacheType, "database")
//	viper.SetDefault(CacheURL, "")
//}
