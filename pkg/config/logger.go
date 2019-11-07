package config

import "github.com/spf13/viper"

type Logger struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
	File   string `yaml:"file"`
}

func (l *Logger) viper() error {
	l.Level = viper.GetString(LoggerLevel)
	l.Output = viper.GetString(LoggerOutput)
	l.File = viper.GetString(LoggerFile)
	return nil
}
