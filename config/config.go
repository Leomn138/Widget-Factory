package config

type Config struct {
	UserDBConfig *DBConfig
	WidgetDBConfig *DBConfig
	Port string
	Auth *Auth
}

type DBConfig struct {
	Name string
	Port int
	Host string
	Protocol string
}

type Auth struct {
	Secret string
}

func GetConfig() *Config {

	return &Config{
		UserDBConfig: &DBConfig{
			Name: "users",
			Port: 5984,
			Host: "127.0.0.1",
			Protocol: "http",
		},
		WidgetDBConfig: &DBConfig{
			Name: "widgets",
			Port: 5984,
			Host: "127.0.0.1",
			Protocol: "http",
		},
		Port: ":8000",
		Auth: &Auth {
			Secret: "secret",
		},
	}
}