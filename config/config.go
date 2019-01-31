package config

import (
	"fishecho/pkgs/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var C = struct {
	Debug      bool
	HTTPServer string
	DBAddr     string
}{
	Debug:      true,
	HTTPServer: ":9196",
	DBAddr:     "postgres://postgres:12345678@localhost:5432/goissue?sslmode=disable",
}

// Load 加载配置
func Load(cfgPath string) error {
	if !utils.PathExist(cfgPath) {
		b, err := yaml.Marshal(&C)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(cfgPath, b, os.ModePerm); err != nil {
			return err
		}
	}

	b, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, &C)
}
