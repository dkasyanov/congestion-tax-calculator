package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	ApplicationPort string `yaml:"ApplicationPort"`
	MonitoringPort  string `yaml:"MonitoringPort"`
	Environment     string
	Env             *Env
	CacheTTLSeconds int `yaml:"CacheTTLSeconds"`
}

type Env struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

// LocalEnvironment should be get from environment: local/staging/production
const LocalEnvironment = "local"

func Load() *Config {
	environment := LocalEnvironment

	viper.AutomaticEnv()
	viper.SetDefault("ENVIRONMENT", LocalEnvironment)

	viper.SetConfigType("yaml")
	viper.SetConfigName(environment)
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("cannot read config file"))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("cannot unmarshal config"))
	}

	// add config from env variables
	config.Env = loadEnv()

	return &config
}

func loadEnv() *Env {
	return &Env{
		DbUser:     viper.GetString("MONGO_USER"),
		DbPassword: viper.GetString("MONGO_PASSWORD"),
		DbHost:     viper.GetString("MONGO_HOST"),
		DbPort:     viper.GetString("MONGO_PORT"),
		DbName:     viper.GetString("MONGO_DB"),
	}
}
