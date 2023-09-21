package stark

import (
	"context"

	"github.com/FarmerChillax/stark/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const (
	VERSION = "v0.0.1"
)

type CallbackPosition int

const (
	// 调用 LoadConfig 方法后
	POSITION_LOAD_CONFIG CallbackPosition = iota + 1
	// 调用 SetupVars 方法后
	POSITION_SETUP_VARS
	// 调用 New 方法后
	POSITION_NEW
)

type CallbackFunc func() error

type Application struct {
	Name string
	Host string
	Port int64
	// 全局请求超时
	// RequestTimeout   int64
	// ReadTimeout      int64
	// WriteTimeout     int64
	Config           *config.Config
	LoadConfig       func() error
	SetupVars        func() error
	RegisterCallback map[CallbackPosition]CallbackFunc
}

// var ApplicationInstance *Application

type MysqlConn interface {
	Get(ctx context.Context) *gorm.DB
}

var Mysql MysqlConn

type RedisConn interface {
	Get(ctx context.Context) *redis.Client
}

var Redis RedisConn

// type Config struct {
// 	Name string
// 	Host string
// 	Port int32
// 	// 全局请求超时
// 	RequestTimeout int64
// 	ReadTimeout    int64
// 	WriteTimeout   int64
// 	Mysql          *MysqlConfig `json:"mysql,omitempty"`
// 	Redis          *RedisConfig `json:"redis,omitempty"`
// }

// type MysqlConfig struct {
// 	Dsn               string `json:"dsn,omitempty"`
// 	Username          string `json:"username,omitempty"`
// 	Password          string `json:"password,omitempty"`
// 	Host              string `json:"host,omitempty"`
// 	Port              int32  `json:"port,omitempty"`
// 	DBName            string `json:"db_name,omitempty" mapstructure:"db_name"`
// 	Charset           string `json:"charset,omitempty"`
// 	Loc               string `json:"loc,omitempty"`
// 	ParseTime         string `json:"parse_time,omitempty"`
// 	Timeout           int64  `json:"timeout,omitempty"`
// 	MaxOpen           int    `json:"max_open,omitempty"`
// 	MaxIdle           int    `json:"max_idle,omitempty"`
// 	ConnMaxLifeSecond int    `json:"conn_max_life_second,omitempty"`
// }

// type RedisConfig struct {
// 	Addr     string
// 	Password string
// 	DB       int
// 	PoolSize int
// 	MaxIdle  int
// }
