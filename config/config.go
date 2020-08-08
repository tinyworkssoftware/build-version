package config

import (
	model "build-version/model/toml"
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

func GetAppConfig(opts ...string) *model.TomlConfig {
	configPath := "config.toml"
	if len(opts) >= 1 {
		log.Printf("Using override config path: %v\n", opts[0])
		configPath = opts[0]
	} else {
		log.Printf("Using default config path: %v\n", configPath)
	}
	var conf model.TomlConfig
	if _, err := toml.DecodeFile(configPath, &conf); err != nil {
		log.Fatalln("Failed to read config toml file. Please ensure config file is present.", err)
	}
	return &conf
}

func SetConfigAsEnvironmentVariables(opts ...string) {
	for f := range opts {
		log.Println(f)
	}
	var config *model.TomlConfig
	if len(opts) > 0 {
		config = GetAppConfig(opts[0])
	} else {
		config = GetAppConfig()
	}
	log.Println("Setting environment variable [DB_USERNAME]")
	_ = os.Setenv("DB_USERNAME", config.Database.Username)
	log.Println("Setting environment variable [DB_PASSWORD]")
	_ = os.Setenv("DB_PASSWORD", config.Database.Password)
	log.Println("Setting environment variable [DB_URL]")
	_ = os.Setenv("DB_URL", config.Database.DatabaseUrl)
	log.Println("Setting environment variable [DB_NAME]")
	_ = os.Setenv("DB_NAME", config.Database.DatabaseName)
}




