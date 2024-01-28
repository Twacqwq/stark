package stark

// func RegisterModules(conf *config.Config) (err error) {
// 	// 初始化数据库配置
// 	if conf.Database != nil && conf.Database.Driver != "" && conf.Database.Dsn != "" {
// 		Database, err = module.NewDatabase(conf.Database)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	// 初始化 redis 配置
// 	if conf.Redis != nil && conf.Redis.Addr != "" {
// 		Redis, err = module.NewRedis(conf.Redis)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	// 初始化 HTTP 客户端配置

// 	return nil
// }

// func RegisterLogger(loggerConf *config.LoggerConfig) (err error) {
// 	Logger, err = xlog.NewLogger(loggerConf)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
