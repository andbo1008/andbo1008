package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DBRole     string `enviroment:"DBROLE"`
	DBHost     string `enviroment:"DBHOST"`
	DBuser     string `enviroment:"DBUSER"`
	DBname     string `enviroment:"DBNAME"`
	DBpassword string `enviroment:"DBPASSWORD"`
	DBport     string `enviroment:"DBPORT"`
	DBUrl      string `enviroment:"DBURL"`
	DBsslmode  string `enviroment:"DBSSLMODE"`
}

func GetConfig() *Config {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		logrus.Fatal(err)
	}
	return &c
}
