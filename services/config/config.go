package config

import "github.com/spf13/viper"

// Configuration is a struct that stores basic config objects
type Configuration struct {
	PORT     string
	DIALECT  string
	URL      string
	PROTOCOL string
}

// InitConfig is what Initializes a configuration for our application
func InitConfig() (Configuration, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return Configuration{}, err
	}
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("DIALECT", "sqlite3")
	viper.SetDefault("URL", "data.db")
	viper.SetDefault("PROTOCOL", "tcp")
	var config Configuration
	err = viper.Unmarshal(&config)
	return config, err
}
