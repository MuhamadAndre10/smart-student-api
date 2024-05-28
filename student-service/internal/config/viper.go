package config

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

	PGHost              string `mapstructure:"PG_HOST"`
	PGPort              int    `mapstructure:"PG_PORT"`
	PGUser              string `mapstructure:"PG_USER"`
	PGPassword          string `mapstructure:"PG_PASSWORD"`
	PGName              string `mapstructure:"PG_NAME"`
	PGDatabase          string `mapstructure:"PG_DATABASE"`
	PGDBPoolMax         int    `mapstructure:"PG_DB_POOL_MAX"`
	PGDBPoolIdle        int    `mapstructure:"PG_DB_POOL_IDLE"`
	PGDBPoolMaxLifetime int    `mapstructure:"PG_DB_POOL_MAX_LIFETIME"`
	PGDriver            string `mapstructure:"PG_DRIVER"`
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
