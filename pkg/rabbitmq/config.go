package rabbitmq

import "github.com/spf13/viper"

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
}

func GetRabbitMQConfig() *Config {
	return &Config{
		Host:     viper.GetString("rabbitmq.host"),
		Port:     viper.GetInt("rabbitmq.port"),
		Username: viper.GetString("rabbitmq.username"),
		Password: viper.GetString("rabbitmq.password"),
	}
}
