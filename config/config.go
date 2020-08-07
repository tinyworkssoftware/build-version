package config

import (
	model "build-version/model/toml"
	"github.com/BurntSushi/toml"
	"log"
)

func GetAppConfig() *model.TomlConfig {
	var conf model.TomlConfig
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}
	return &conf
}




