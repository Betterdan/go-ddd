package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Config struct {
	Server      ServerConfig `yaml:"server"`
	DbConfig    DbConfig     `yaml:"db_config"`
	CacheConfig CacheConfig  `yaml:"cache_config"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DbConfig struct {
	DBHost     string `yaml:"db_host"`
	DBPort     int    `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
}

type CacheConfig struct {
	Host        string        `yaml:"host"`
	Port        string        `yaml:"port"`
	Password    string        `yaml:"password"`
	Db          int           `yaml:"db"`
	PoolSize    int           `yaml:"pool_size"`
	MinIdleCon  int           `yaml:"min_idle_con"`
	PoolTimeout time.Duration `yaml:"pool_timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func LoadConfig() (*Config, error) {
	// 从本地配置文件加载配置
	appConfig, err := LoadLocalConfig()
	if err != nil {
		return nil, err
	}

	// 创建 Nacos 服务端配置
	var serverConfigs []constant.ServerConfig
	for _, sc := range appConfig.Nacos.ServerConfigs {
		serverConfig := *constant.NewServerConfig(sc.IpAddr, sc.Port)
		serverConfigs = append(serverConfigs, serverConfig)
	}

	// 创建 Nacos 客户端配置
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(appConfig.Nacos.ClientConfig.NamespaceId),
		constant.WithTimeoutMs(appConfig.Nacos.ClientConfig.TimeoutMs),
		constant.WithLogDir(appConfig.Nacos.ClientConfig.LogDir),
		constant.WithCacheDir(appConfig.Nacos.ClientConfig.CacheDir),
		constant.WithLogLevel(appConfig.Nacos.ClientConfig.LogLevel),
	)

	// 创建 Nacos 配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		return nil, err
	}

	// 从 Nacos 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: appConfig.Nacos.DataId,
		Group:  appConfig.Nacos.Group,
	})
	if err != nil {
		return nil, err
	}

	// 解析配置
	cfg := &Config{}
	err = yaml.Unmarshal([]byte(content), cfg)
	if err != nil {
		return nil, err
	}
	log.Println("content:" + content)
	return cfg, nil
}
