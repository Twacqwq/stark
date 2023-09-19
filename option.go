package stark

import (
	"github.com/FarmerChillax/stark/config"
)

type StarkOption func(*Application) *Application

func WithConfig(configPath ...string) StarkOption {
	return func(app *Application) *Application {
		config.Load(configPath...)
		app.Config = config.Get()
		return app
	}
}
