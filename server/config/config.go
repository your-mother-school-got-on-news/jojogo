package config

import (
	"fmt"

	"jojogo/server/utils/log"

	"github.com/spf13/viper"
)

// Val Val
var Val Config

// Config Config
type Config struct {
	JWTTokenLife int    `mapstructure:"JWT_TOKEN_LIFE"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
}

// Init Init
func Init() {
	// 讀config.yaml
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %v ", err))
	}
	if err := viper.Unmarshal(&Val); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	log.Info("Config loaded", log.String("Val", fmt.Sprintf("%v", Val)))
}
