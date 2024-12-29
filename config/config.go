package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Username     string
	PasswordPath string
	Instance     string
	Port         int
	DbName       string
}

func GetConfig() (Config, error) {
	dat, e := os.ReadFile("/hostname")
	if e != nil {
		panic(e)
	}
	hostname := string(dat)
	if strings.Contains(hostname, "DESKTOP") {
		fmt.Println("Using desktop config")
		viper.AddConfigPath("./config")
		viper.SetConfigName("desktop")
		viper.SetConfigType("yaml")
	} else {
		fmt.Println("Using default config")
		viper.AddConfigPath("./config")
		viper.SetConfigName("default")
		viper.SetConfigType("yaml")
	}
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	port := viper.GetInt("db.port")
	return Config{Username: fmt.Sprintf("%s", viper.Get("db.username")),
		PasswordPath: "/run/secrets/db_password",
		Instance:     fmt.Sprintf("%s", viper.Get("db.instance")),
		Port:         port,
		DbName:       fmt.Sprintf("%s", viper.Get("db.dbName"))}, nil
}
