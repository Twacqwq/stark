package stark

// var (
// 	port = flag.Int64("p", 0, "Set server port.")
// )

// func New(app *Application) (*Application, error) {
// 	// ApplicationInstance = app
// 	if app.Name == "" {
// 		return nil, errors.New("application name can't not be empty")
// 	}

// 	flag.Parse()
// 	if *port > 0 {
// 		app.Port = *port
// 	}

// 	// 初始化配置
// 	conf, err := config.InitGlobalConfig(app.Config)
// 	if err != nil {
// 		log.Println("config.InitGlobalConfig err: ", err)
// 		return nil, err
// 	}

// 	// 初始化日志
// 	err = RegisterLogger(config.GetLoggerConfig())
// 	if err != nil {
// 		log.Println("RegisterLogger err: ", err)
// 		return nil, err
// 	}

// 	// 初始化内置组件
// 	if err := RegisterModules(conf); err != nil {
// 		log.Println("Register err: ", err)
// 		return nil, err
// 	}

// 	// 加载用户配置
// 	if app.LoadConfig != nil {
// 		err := app.LoadConfig()
// 		if err != nil {
// 			return nil, err
// 		}
// 		// 执行加载配置回调操作
// 		// err = runLoadConfigCallback(POSITION_LOAD_CONFIG,app.RegisterCallback)
// 		if err := runCallback(POSITION_LOAD_CONFIG, app.RegisterCallback); err != nil {
// 			return nil, err
// 		}
// 	}

// 	// 设置常量
// 	if app.SetupVars != nil {
// 		if err := app.SetupVars(); err != nil {
// 			return nil, err
// 		}
// 		// 执行设置常量的回调操作
// 		if err := runCallback(POSITION_SETUP_VARS, app.RegisterCallback); err != nil {
// 			return nil, err
// 		}
// 	}

// 	if err := runCallback(POSITION_NEW, app.RegisterCallback); err != nil {
// 		return nil, err
// 	}

// 	return app, nil
// }

// func runCallback(position CallbackPosition, callbacks map[CallbackPosition]CallbackFunc) error {
// 	if f, ok := callbacks[position]; ok {
// 		return f()
// 	}

// 	return nil
// }
