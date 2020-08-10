package config

import (
	model "build-version/model/toml"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"sync"
)

var (
	loadedConfig *model.TomlConfig
	once sync.Once
)
func GetAppConfig(opts ...string) *model.TomlConfig {
	once.Do(func() {
		configPath := "config.toml"
		if len(opts) >= 1 {
			log.Printf("Using override config path: %v\n", opts[0])
			configPath = opts[0]
		} else {
			log.Printf("Using default config path: %v\n", configPath)
		}
		if _, err := toml.DecodeFile(configPath, &loadedConfig); err != nil {
			log.Fatalln("Failed to read config toml file. Please ensure config file is present.", err)
		}
	})
	return loadedConfig

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
	log.Println("Setting environment variable [API_VERSION]")
	_ = os.Setenv("API_VERSION", config.Application.ApiPrefix)
}




