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

type Application struct {
	Config           *config.Config
	LoadConfig       func() error
	SetupVars        func() error
	RegisterCallback map[CallbackPosition]func() error
}

func New(appConfig *Application) *Application {
	return &Application{}
}

var Mysql MysqlConn

type MysqlConn interface {
	Get(ctx context.Context) *gorm.DB
}
