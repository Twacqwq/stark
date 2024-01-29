package metrics

import (
	"github.com/gin-contrib/pprof"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
)

func RegisterGin(engine *gin.Engine) *gin.Engine {

	pprof.Register(engine, "/debug/pprof")
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	engine.GET("/health", HealthHandler())
	return engine
}

// type GinMetrics struct {
// 	*stark.Application
// }

// func NewWithGin(app *stark.Application) *GinMetrics {
// 	return &GinMetrics{
// 		Application: app,
// 	}
// }

// func (m *GinMetrics) RegisterGin(engine *gin.Engine) *gin.Engine {

// 	pprof.Register(engine, "/debug/pprof")
// 	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
// 	engine.GET("/health", HealthHandler())
// 	return engine
// }
