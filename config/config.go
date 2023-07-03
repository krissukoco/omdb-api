package config

import (
	"runtime"

	"github.com/spf13/viper"
)

type dbConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Dbname    string `mapstructure:"dbname"`
	EnableSsl bool   `mapstructure:"enable_ssl"`
}

type config struct {
	Database *dbConfig
	Port     int
}

func Load(file string) (*config, error) {
	_, filename, _, _ := runtime.Caller(0)
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filename + "/../../config")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
