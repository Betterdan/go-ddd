package config

import (
	"github.com/spf13/viper"
)

type NacosServerConfig struct {
	IpAddr string `mapstructure:"ipAddr"`
	Port   uint64 `mapstructure:"port"`
}

type NacosClientConfig struct {
	NamespaceId string `mapstructure:"namespaceId"`
	TimeoutMs   uint64 `mapstructure:"timeoutMs"`
	LogDir      string `mapstructure:"logDir"`
	CacheDir    string `mapstructure:"cacheDir"`
	LogLevel    string `mapstructure:"logLevel"`
}

type NacosConfig struct {
	ServerConfigs []NacosServerConfig `mapstructure:"serverConfigs"`
	ClientConfig  NacosClientConfig   `mapstructure:"clientConfig"`
	DataId        string              `mapstructure:"dataId"`
	Group         string              `mapstructure:"group"`
}

type AppConfig struct {
	Nacos NacosConfig `mapstructure:"nacos"`
}

func LoadLocalConfig() (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
