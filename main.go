package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-biu/app/admin"
	"go-biu/app/api"
	"go-biu/library/config"
	"go-biu/library/database/cache"
	"go-biu/library/database/orm"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConf()
	db := orm.InitGorm()
	rdb := cache.InitRedisConn()
	defer rdb.Close()

	gin.SetMode(config.YamlConf.Base.Model)
	router := gin.New()
	router.Use(gin.Recovery())

	// 设置时区
	time.Local = time.FixedZone("CST", 8*3600)

	admin.NewRouter(router, db)
	api.InitRouter(router, db)
	server := &http.Server{
		Addr:           ":9090",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
