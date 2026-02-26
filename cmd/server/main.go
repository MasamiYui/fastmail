package main

import (
	"context"
	"fastmail/internal/api"
	"fastmail/internal/config"
	"fastmail/internal/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

func main() {
	var cfgFile string
	pflag.StringVarP(&cfgFile, "config", "c", "", "config file (default is ./config.yaml)")
	pflag.Parse()

	// 1. Load config
	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Init service
	emailService := service.NewEmailService(cfg)

	// 3. Init handler
	handler := api.NewHandler(emailService)

	// 4. Setup Gin
	r := gin.Default()

	// Public routes
	r.GET("/health", handler.Health)

	// Protected routes
	apiV1 := r.Group("/api/v1")
	apiV1.Use(api.AuthMiddleware(cfg.Server.Token))
	{
		apiV1.POST("/send", handler.SendEmail) // Use handler.SendEmail, not handler.Send
	}

	// 5. Run server with graceful shutdown
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: r,
	}

	go func() {
		// service connections
		log.Printf("Server starting on port %d", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
