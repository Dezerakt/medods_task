package configs

import (
	"errors"
	"github.com/spf13/viper"
)

var (
	EnvConfigObject *EnvConfig

	ErrReadingFromConfigFile = errors.New("ошибка при чтении из config файла")
	ErrUnmarshalConfig       = errors.New("ошибка при парсинге данных из config файла")
)

type EnvConfig struct {
	ServerPort string `mapstructure:"SERVER_PORT"`

	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`

	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func NewEnvConfig() *EnvConfig {
	if EnvConfigObject == nil {
		EnvConfigObject = &EnvConfig{}
	}

	return EnvConfigObject
}

func (config *EnvConfig) LoadConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	err = viper.Unmarshal(config)

	if err != nil {
		return err
	}

	return nil
}
