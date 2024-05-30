package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

var Env *EnvConfigs

func InitEnvConfigs(path string) {
	Env = loadEnv(path)
}

type EnvConfigs struct {
	AppName string `mapstructure:"APP_NAME"`
	AppPort string `mapstructure:"APP_PORT"`

	DSN string `mapstructure:"DSN"`

	JWTPrivateKey string `mapstructure:"PRIVATE_KEY"`
	JWTPublicKey  string `mapstructure:"PUBLIC_KEY"`
	JwtTokenExp   int    `mapstructure:"TOKEN_EXP"`
}

func loadEnv(path string) (config *EnvConfigs) {

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	if err = viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	return

}
