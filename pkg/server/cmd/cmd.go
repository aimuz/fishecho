package cmd

import (
	"os"
	"path/filepath"
	"time"

	"github.com/aimuz/fishecho/pkg/config"
	"github.com/aimuz/fishecho/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewServerDefaultCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: filepath.Base(os.Args[0]) + " is blog server",
		Long: `Support seamless upgrade from Typecho
Complete documentation is available at github.com/aimuz/fishecho`,
		Version: version.Version + " gitCommit " + version.GitCommit,
	}
	cmd.Flags().StringP(config.ServerHTTPAddr, "H", "0.0.0.0", "http server addr")
	cmd.Flags().IntP(config.ServerHTTPPort, "p", 3001, "http server port")
	cmd.Flags().BoolP(config.ServerHTTPGzip, "g", false, "http server enable gzip")

	cmd.Flags().StringP(config.LoggerLevel, "v", "info", "logger level [debug info warn error fail]")
	cmd.Flags().String(config.LoggerOutput, "stderr", "logger output [stderr stdout file]")
	cmd.Flags().String(config.LoggerFile, "", "logger output file")

	cmd.Flags().String(config.DatabaseType, "", "database type [mysql]")
	cmd.Flags().String(config.DatabaseHost, "", "database host")
	cmd.Flags().String(config.DatabaseName, "", "database name")
	cmd.Flags().String(config.DatabaseUser, "", "database user")
	cmd.Flags().String(config.DatabasePassword, "", "database password")
	cmd.Flags().String(config.DatabaseURL, "mysql://user:secret@host:port/database", "database url")
	cmd.Flags().Int(config.DatabaseMaxIdleConn, 2, "database max idle setting")
	cmd.Flags().Int(config.DatabaseMaxOpenConn, 0, "database max conn setting")
	cmd.Flags().Duration(config.DatabaseConnMaxLifetime, time.Hour*4, "database Connection Max Lifetime")
	cmd.Flags().Bool(config.DatabaseLogQueries, false, "database set to true to log the sql calls and execution times.")

	cmd.Flags().String(config.CacheType, "database", "remote cache type[database redis]")
	cmd.Flags().String(config.CacheURL, "", "remote cache url")
	cmd.Flags().Bool(config.CacheEnable, false, "enable cache")

	cmd.Flags().StringP(config.File, "c", "", "config file, support file format [YAML JSON INI]")

	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		panic(err)
	}
	return cmd
}
