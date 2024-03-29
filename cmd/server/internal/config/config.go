package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL" default:"postgres://qrincard:supersecret@0.0.0.0/qrincard?sslmode=disable"`
	StaticPath  string `envconfig:"STATIC_PATH" default:"./static"`
	Port        string `envconfig:"PORT" default:"2345"`
}

func New() Config {
	var env Config
	_ = envconfig.Process("", &env)

	return env
}
