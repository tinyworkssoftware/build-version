package model

type TomlConfig struct {
	Database database `toml:"database"`
	Application application `toml:"application"`
	Git git `toml:"git"`
	Auth auth `toml:"auth"`
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
	ApiPrefix string `toml:"api_prefix"`
	ApiBypass []string `toml:"api_bypass"`
	LogLevel string `toml:"log_level"`
}

type git struct {
	Source string `toml:"source"`
	AccessToken string `toml:"access_token"`
}

type auth struct {
	UserInfoUrl string `toml:"user_info_url"`
	PublicTokenUrl string `toml:"public_token_url"`
	AuthUrl string `toml:"auth_url"`
}