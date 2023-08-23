package models

import (
	"time"

	"github.com/spf13/viper"
)

type AuthConfig struct {
	AccessTokenPrivateKey  string        `json:"ACCESS_TOKEN_PRIVATE_KEY" bson:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `json:"ACCESS_TOKEN_PUBLIC_KEY" bson:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `json:"REFRESH_TOKEN_PRIVATE_KEY" bson:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `json:"REFRESH_TOKEN_PUBLIC_KEY" bson:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiry      time.Duration `json:"ACCESS_TOKEN_EXPIRY" bson:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry     time.Duration `json:"REFRESH_TOKEN_EXPIRY" bson:"REFRESH_TOKEN_EXPIRY"`
	AccessTokenMaxAge      int           `json:"ACCESS_TOKEN_MAXAGE" bson:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `json:"REFRESH_TOKEN_MAXAGE" bson:"REFRESH_TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config AuthConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
