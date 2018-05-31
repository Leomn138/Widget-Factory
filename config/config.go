package config

type Config struct {
	UserDBConfig *DBConfig
	WidgetDBConfig *DBConfig
	Port string
	Auth *Auth
}

type DBConfig struct {
	Username string
	Password string
	Name string
	Port int
	Host string
}

type Auth struct {
	Secret string
}

func GetConfig() *Config {

	return &Config{
		UserDBConfig: &DBConfig{
			Username: "",
			Password: "",
			Name: "users",
			Port: 5984,
			Host: "127.0.0.1",
		},
		WidgetDBConfig: &DBConfig{
			Username: "",
			Password: "",
			Name: "widgets",
			Port: 5984,
			Host: "127.0.0.1",
		},
		Port: ":8000",
		Auth: &Auth {
			Secret: "secret",
		},
	}
}