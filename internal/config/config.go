package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	GRPSPort   string `env:"GRPC_PORT" env-default:"8082"`
	UserDB     string `env:"MYSQL_USER" env-default:"root"`
	PasswordDB string `env:"MYSQL_PSW" env-default:"root"`
	NameDB     string `env:"MYSQL_DB" env-default:"books"`
	PortDB     string `env:"MYSQL_PORT" env-default:"3308"`
	HostDB     string `env:"MYSQL_HOST" env-default:"localhost"`
}

//New возвращает Config
func New() (Config, error) {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, err
}
