package app

import (
	"errors"
	"flag"

	"github.com/FarmerChillax/stark"
)

var (
	port = flag.Int64("p", 0, "Set server port.")
)

func New(app *stark.Application) (*stark.Application, error) {
	// stark.ApplicationInstance = app
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
		// 执行加载配置回调操作
		// err = runLoadConfigCallback(stark.POSITION_LOAD_CONFIG,app.RegisterCallback)
		if err := runCallback(stark.POSITION_LOAD_CONFIG, app.RegisterCallback); err != nil {
			return nil, err
		}
	}

	// 初始化日志

	// 设置常量
	if app.SetupVars != nil {
		if err := app.SetupVars(); err != nil {
			return nil, err
		}
		// 执行设置常量的回调操作
		if err := runCallback(stark.POSITION_SETUP_VARS, app.RegisterCallback); err != nil {
			return nil, err
		}
	}

	// 注册内置组件
	if app.RegisterModule != nil {
		err := app.RegisterModule()
		if err != nil {
			return nil, err
		}
		if err := runCallback(stark.POSITION_MODULE, app.RegisterCallback); err != nil {
			return nil, err
		}
	}

	if err := runCallback(stark.POSITION_NEW, app.RegisterCallback); err != nil {
		return nil, err
	}

	return app, nil
}

func runCallback(position stark.CallbackPosition, callbacks map[stark.CallbackPosition]stark.CallbackFunc) error {
	if f, ok := callbacks[position]; ok {
		return f()
	}

	return nil
}
