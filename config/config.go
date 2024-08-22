package config

import (
	"flag"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	Debug       bool `yaml:"debug" env:"APP_DEBUG"`
	Development bool `yaml:"development" env:"APP_DEVELOPMENT"`
}

type JwtConfig struct {
	Key string `env-required:"true" yaml:"key" env:"JWT_KEY"`
}

type DbConfig struct {
	Url     string `env-required:"true" yaml:"url" env:"DB_URL"`
	Name    string `env-required:"true" yaml:"name" env:"DB_NAME"`
	Sync    bool   `yaml:"sync" env:"DB_SYNC"`
	Seed    bool   `yaml:"seed" env:"DB_SEED"`
	Clear   bool   `yaml:"clear" env:"DB_CLEAR"`
	Migrate bool   `yaml:"migrate" env:"DB_MIGRATE"`
}

type ServerConfig struct {
	Port int `env-required:"true" yaml:"port" env:"SERVER_PORT"`
}

/*
type MinIoConfig struct {
	Endpoint        string `env-required:"true" yaml:"endpoint" env:"MINIO_ENDPOINT"`
	AccessKeyID     string `env-required:"true" yaml:"accessKeyID" env:"MINIO_ACCESS_KEY_ID"`
	SecretAccessKey string `env-required:"true" yaml:"secretAccessKey" env:"MINIO_SECRET_ACCESS_KEY"`
}

type ZimbraConfig struct {
	Port uint   `env-required:"true" yaml:"port" env:"ZIMBRA_PORT"`
	Url  string `env-required:"true" yaml:"url" env:"ZIMBRA_URL"`
}
*/

type Config struct {
	Db     DbConfig     `yaml:"db"`
	App    AppConfig    `yaml:"app"`
	Jwt    JwtConfig    `yaml:"jwt"`
	Server ServerConfig `yaml:"server"`
	//MinIo  MinIoConfig  `yaml:"minio"`
	//Zimbra ZimbraConfig `yaml:"zimbra"`
}

func Parse() (*Config, error) {
	var cfgFilePath = flag.String("config-file", "/etc/config.yml", "A filepath to the yml file containing the microservice configuration")
	flag.Parse()

	cfg := &Config{}

	if err := cleanenv.ReadConfig(*cfgFilePath, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
