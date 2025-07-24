package common

import "github.com/spf13/viper"

type serviceConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}
type mysqlConfig struct {
	Host     string `mapstructure:"mysql_host"`
	Port     int    `mapstructure:"mysql_port"`
	User     string `mapstructure:"mysql_user"`
	Password string `mapstructure:"mysql_password"`
	Database string `mapstructure:"mysql_database"`
}
type Config struct {
	Service serviceConfig `mapstructure:"service"`
	Mysqldb mysqlConfig   `mapstructure:"mysqldb"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
