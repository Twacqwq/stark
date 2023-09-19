package config

import "github.com/spf13/viper"

type Config struct {
	Name string
	Host string
	Port int32
	// 全局请求超时
	RequestTimeout int64
	ReadTimeout    int64
	WriteTimeout   int64
	Mysql          *MysqlConfig `json:"mysql,omitempty"`
	Redis          *RedisConfig `json:"redis,omitempty"`
}

var config *Config

func Get() *Config {
	return config
}

func Load(paths ...string) error {
	if len(paths) != 0 {
		for _, path := range paths {
			viper.SetConfigFile(path)
		}
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigType("json")
	viper.AddConfigPath("./local")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	return nil
}

type MysqlConfig struct {
	Dsn               string `json:"dsn,omitempty"`
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
	Host              string `json:"host,omitempty"`
	Port              int32  `json:"port,omitempty"`
	DBName            string `json:"db_name,omitempty" mapstructure:"db_name"`
	Charset           string `json:"charset,omitempty"`
	Loc               string `json:"loc,omitempty"`
	ParseTime         string `json:"parse_time,omitempty"`
	Timeout           int64  `json:"timeout,omitempty"`
	MaxOpen           int    `json:"max_open,omitempty"`
	MaxIdle           int    `json:"max_idle,omitempty"`
	ConnMaxLifeSecond int    `json:"conn_max_life_second,omitempty"`
}

func GetMysql() *MysqlConfig {
	// if config.Mysql == nil {
	// 	config.Mysql = &MysqlConfig{}
	// }
	return config.Mysql
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
	MaxIdle  int
}

func GetRedis() *RedisConfig {
	// if config.Redis == nil {
	// 	return &defaultRedisConfig
	// }
	return config.Redis
}