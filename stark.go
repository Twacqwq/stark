package stark

import (
	"context"

	"github.com/FarmerChillax/stark/config"
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

var ApplicationInstance *Application

type MysqlConn interface {
	Get(ctx context.Context) *gorm.DB
}

var Mysql MysqlConn

type RedisConn interface{}

var Redis RedisConn
