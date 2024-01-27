package main

import (
	"context"
	"fmt"

	"github.com/FarmerChillax/stark/internal/otel"
	"github.com/FarmerChillax/stark/internal/xlog"
	"github.com/FarmerChillax/stark/pkg/helper"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	ctx := context.Background()
	log := helper.NewLogger(xlog.NewFormatter())
	tp, err := otel.RegisterTracer("base-demo", nil)
	if err != nil {
		log.Errorf("otel.RegisterTracer err: %v", err)
	}
	defer tp.Shutdown(ctx)

	ctx, span := tp.Tracer("base-demo").Start(ctx, "base-demo")
	log.WithContext(ctx).Infof("hello world")
	span.End()
	log.WithContext(ctx).Infof("test")

	app := gin.Default()
	app.Use(otelgin.Middleware("base-demo"))
	app.GET("/ping", func(c *gin.Context) {
		log.WithContext(c.Request.Context()).Infof("hello ping.")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Infof("listen on :6000")
	if err := app.Run(":6000"); err != nil {
		panic(err)
	}

	log.WithContext(ctx).Infof("hello world")
	fmt.Println("done.")
}

// app, err := app.New(&stark.Application{
// 	Name:   "base-demo",
// 	Host:   "127.0.0.1",
// 	Port:   6000,
// 	Config: &starkConf.Config{},
// 	LoadConfig: func() error {
// 		return nil
// 	},
// 	SetupVars: func() error {
// 		return nil
// 	},
// 	RegisterCallback: make(map[stark.CallbackPosition]stark.CallbackFunc),
// 	// RegisterRouter
// })
// if err != nil {
// 	log.Fatalf("app.New err: %v", err)
// }

// if err := app.ListenAndServe(); err != nil {
// 	log.Fatalf("app.ListenAndServe err: %v", err)
// }
