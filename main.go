package main

import (
	"github.com/gin-gonic/gin"
	"go-judge/component"
	"go-judge/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	secretKey := os.Getenv("SYSTEM_SECRET")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	rand.Seed(time.Now().UnixNano())
	if err != nil {
		log.Fatalln(err)
	}
	//db = db.Debug()
	if err := runService(db, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, secretKey string) error {
	appCtx := component.NewAppContext(db, secretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.Use(middleware.AllowCORS())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//r.NoRoute(func(c *gin.Context) {
	//	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	//})
	return mainRoute(appCtx, r)
}

func mainRoute(appCtx *component.AppCtx, r *gin.Engine) error {

	//v1 := r.Group("/v1")

	return r.Run()
}
