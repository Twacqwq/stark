package app

import (
	"errors"
	"flag"

	"github.com/FarmerChillax/stark"
	"github.com/FarmerChillax/stark/module"
)

var (
	port = flag.Int64("p", 0, "Set server port.")
)

func New(app *stark.Application) (*stark.Application, error) {
	stark.ApplicationInstance = app
	if app.Name == "" {
		return nil, errors.New("application name can't not be empty")
	}

	flag.Parse()
	if *port > 0 {
		app.Port = *port
	}
	// 加载配置
	if app.LoadConfig != nil {
		err := app.LoadConfig()
		if err != nil {
			return nil, err
		}

		err = runLoadConfigCallback(app.RegisterCallback)
		if err != nil {
			return nil, err
		}
	}

	// 加载全局组件
	// 初始化 mysql 链接
	if app.Config.Mysql != nil {
		err := module.RegisterMysql(app)
		if err != nil {
			return nil, err
		}
	}

	// 初始化 redis 链接
	if app.Config.Redis != nil {
		err := module.RegisterRedis(app)
		if err != nil {
			return nil, err
		}
	}

	return &stark.Application{}, nil
}

func runLoadConfigCallback(callbacks map[stark.CallbackPosition]stark.CallbackFunc) error {
	if f, ok := callbacks[stark.POSITION_LOAD_CONFIG]; ok {
		return f()
	}

	return nil
}
