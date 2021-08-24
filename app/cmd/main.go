package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"swagger/app/routers"
	"syscall"
	"time"
)

// @title Swagger student_system
// @version 1.0
// @description swagger使用示例
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email 1111x@qq.com

// @host 127.0.0.1:8080
// @BasePath /student/api/v1
func main() {
	// 路由
	r := routers.InitRouters()
	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	switch <-quit {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
	}
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
