package middlewares

import (
	"context"
	"time"

	"github.com/FarmerChillax/stark"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type GinMiddleware struct {
	*stark.GinApplication
}

func NewWithGin(ginApp *stark.GinApplication) *GinMiddleware {
	return &GinMiddleware{
		GinApplication: ginApp,
	}
}

func (m *GinMiddleware) Register(engine *gin.Engine) {
	// 注册公共中间件
	engine.Use(gin.Recovery())
	engine.Use(otelgin.Middleware(m.Application.Name))
	engine.Use(m.AccessLog())

	if m.Application.Config != nil {
		engine.Use(ContextTimeout(time.Second * time.Duration(m.Application.Config.Timeout)))
	}
}

func (m *GinMiddleware) AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		m.TracerLogger.WithContext(c.Request.Context()).Infof("Request log")
		c.Next()
		m.TracerLogger.WithContext(c.Request.Context()).Infof("response log")
	}
}

func ContextTimeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
