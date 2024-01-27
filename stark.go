package stark

import (
	"context"

	"github.com/FarmerChillax/stark/config"
	"github.com/gin-gonic/gin"
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
	// 调用 RegisterModule 后
	POSITION_MODULE
)

type CallbackFunc func() error

type Application struct {
	Name string
	Host string
	Port int64
	// 全局请求超时
	Config           *config.Config
	LoadConfig       func() error
	SetupVars        func() error
	RegisterModule   func() error
	RegisterCallback map[CallbackPosition]CallbackFunc
	RegisterRouter   func(*gin.Engine) error
	engine           *gin.Engine
}

// var ApplicationInstance *Application

type DatabaseIface interface {
	Get(ctx context.Context) *gorm.DB
}

var Database DatabaseIface

type RedisConn interface {
	Get(ctx context.Context) *redis.Client
}

var Redis RedisConn
