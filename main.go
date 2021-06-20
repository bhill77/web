package main

import (
	"fmt"
	"github.com/bhill77/web/app/api"
	"github.com/bhill77/web/app/api/handler"
	"github.com/bhill77/web/module/cache"
	newsModule "github.com/bhill77/web/module/news"
	"github.com/bhill77/web/service/news"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	e := echo.New()
	db := dbConnect()
	articleRepo := newsModule.NewRepository(db)
	redisClient := redisConnect()
	redisCache := cache.NewRedisCacheRepository(redisClient)
	newsService := news.NewService(articleRepo, redisCache)
	articleHandler := handler.NewArticleHandler(newsService)
	api.RegisterPath(e, articleHandler)
	e.Logger.Fatal(e.Start(":8000"))
}

func dbConnect() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"",
		"localhost",
		3306,
		"news")

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db, _ := connection.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(10 * time.Minute)
	return connection
}

func redisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}
