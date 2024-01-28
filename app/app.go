package app

import (
	"errors"
	"flag"
	"log"

	"github.com/FarmerChillax/stark"
	"github.com/FarmerChillax/stark/config"
	"github.com/FarmerChillax/stark/internal/module"
	"github.com/FarmerChillax/stark/internal/xlog"
)

var (
	host = flag.String("h", "127.0.0.1", "Set server listen host.")
	port = flag.Int64("p", 5500, "Set server port.")
)

func validateAndMergeAppConfig(app *stark.Application) error {
	if app.Name == "" {
		return errors.New("application name can't not be empty")
	}

	flag.Parse()
	if *port > 0 {
		app.Port = *port
	}

	if app.Port > 65535 || app.Port < 0 {
		return errors.New("invalid port")
	}

	if app.Host == "" {
		app.Host = *host
	}

	return nil
}

func New(app *stark.Application) (*stark.Application, error) {
	// stark.ApplicationInstance = app
	if err := validateAndMergeAppConfig(app); err != nil {
		return nil, err
	}

	// 初始化配置
	conf, err := config.InitGlobalConfig(app.Config)
	if err != nil {
		log.Println("InitGlobalConfig err: ", err)
		return nil, err
	}
	err = runCallback(stark.POSITION_GLOBAL_CONFIG, app.RegisterCallback)
	if err != nil {
		log.Println("runCallback POSITION_GLOBAL_CONFIG err: ", err)
		return nil, err
	}

	// 初始化日志
	err = xlog.Register(config.GetLoggerConfig())
	if err != nil {
		log.Println("Register Logger err: ", err)
		return nil, err
	}
	err = runCallback(stark.POSITION_INIT_LOGGER, app.RegisterCallback)
	if err != nil {
		log.Println("runCallback POSITION_INIT_LOGGER err: ", err)
		return nil, err
	}

	// 初始化内置组件
	if err := module.Register(conf); err != nil {
		log.Println("Register Modules err: ", err)
		return nil, err
	}
	err = runCallback(stark.POSITION_MODULE_REGISTER, app.RegisterCallback)
	if err != nil {
		log.Println("runCallback POSITION_MODULE_REGISTER err: ", err)
		return nil, err
	}

	// 加载用户配置
	if app.LoadConfig != nil {
		err := app.LoadConfig()
		if err != nil {
			return nil, err
		}
		// 执行加载配置回调操作
		if err := runCallback(stark.POSITION_LOAD_CONFIG, app.RegisterCallback); err != nil {
			return nil, err
		}
	}

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
