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
	hn := os.Getenv("HOSTNAME")
	if strings.Contains(hn, "DESKTOP") {
		viper.AddConfigPath("./config")
		viper.SetConfigName("desktop")
		viper.SetConfigType("yaml")
	} else {
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
		PasswordPath: fmt.Sprintf("%s", viper.Get("db.passwordPath")),
		Instance:     fmt.Sprintf("%s", viper.Get("db.instance")),
		Port:         port,
		DbName:       fmt.Sprintf("%s", viper.Get("db.dbName"))}, nil
}
