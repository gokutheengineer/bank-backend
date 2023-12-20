package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment         string        `mapstructure:"ENVIRONMENT"`
	DB_Source           string        `mapstructure:"DB_SOURCE"`
	Migration_Url       string        `mapstructure:"MIGRATION_URL"`
	HttpServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GrpcServerAddress   string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// TODO check env to differantiate between env to apply different configs for dev, test, prod
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
