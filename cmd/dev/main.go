package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/m/internal/common"
	"example.com/m/internal/component"
	cpnmysqldb "example.com/m/internal/component/mysqldb"
	itemgin "example.com/m/internal/module/item/transport/gin"
	itemresourcegin "example.com/m/internal/module/item_resource/transport/gin"
	itemtypegin "example.com/m/internal/module/item_type/transport/gin"
	resourcegin "example.com/m/internal/module/resource/transport/gin"
	resourcetypegin "example.com/m/internal/module/resource_type/transport/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := common.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	gorm := cpnmysqldb.SQLGORM{}
	mysqlDB, err := gorm.Connect(config)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check endpoint
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	appCtx := component.NewAppContext(config, mysqlDB)

	// Setup modules
	modules := []common.AppModule{
		itemresourcegin.NewItemResourceModule(appCtx),
		resourcegin.NewResourceModule(appCtx),
		resourcetypegin.NewResourceTypeModule(appCtx),
		itemtypegin.NewItemTypeModule(appCtx),
		itemgin.NewItemModule(appCtx),
	}

	for _, module := range modules {
		module.SetupGin(r)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Service.Port),
		Handler: r,
	}

	go func() {
		log.Printf("Starting server on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}
