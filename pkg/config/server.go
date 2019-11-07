package config

import (
	"github.com/spf13/viper"
)

type Server struct {
	HTTPAddr string `yaml:"httpAddr"`
	HTTPPort string `yaml:"httpPort"`
	HTTPGzip string `yaml:"httpGzip"`
}

func (s *Server) viper() error {
	s.HTTPAddr = viper.GetString(ServerHTTPAddr)
	s.HTTPPort = viper.GetString(ServerHTTPPort)
	s.HTTPGzip = viper.GetString(ServerHTTPGzip)
	return nil
}
