package stark

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *Application) ListenAndServe() error {

	// 生成 gin engine
	engine := gin.New()
	app.engine = engine
	// 注册公共中间件
	// todo

	// 注册路由
	if app.RegisterRouter != nil {
		if err := app.RegisterRouter(app.engine); err != nil {
			return err
		}
	}
	// 服务启动
	errChan := make(chan error)
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	addr := fmt.Sprintf("%s:%d", app.Host, app.Port)
	server := http.Server{
		Addr:           addr,
		Handler:        app.engine,
		ReadTimeout:    time.Second * time.Duration(app.Config.ReadTimeout),
		WriteTimeout:   time.Second * time.Duration(app.Config.WriteTimeout),
		MaxHeaderBytes: 1 << 20, // 1MB
	}
	go func() {
		log.Printf("stark is running at http://%s/ . Press Ctrl+C to stop.", addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln("err:", err)
			errChan <- err
			return
		}
	}()

	// 程序退出、从注册中心注销服务实例
	select {
	case err := <-errChan:
		return err
	case <-stopChan:
		// 优雅退出
		// atomic.CompareAndSwapInt32(&num)
	}
	return nil
}
