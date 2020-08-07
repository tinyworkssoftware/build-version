package model

type TomlConfig struct {
	Database database `toml:"database"`
	Application application `toml:"application"`
	Git git `toml:"git"`
}

type database struct {
	DatabaseUrl string `toml:"database_url"`
	DatabaseName string `toml:"database_name"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type application struct {
	Timeout string `toml:"timeout"`
	EnableCache bool `toml:"enable_cache"`
	Cache string `toml:"cache"`
	SelfHosted bool `toml:"self_hosted"`
}

type git struct {
	Source string `toml:"source"`
	AccessToken string `toml:"access_token"`
}