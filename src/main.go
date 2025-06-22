package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sohosai/go-gin-sqlx-pg-template/internal/domain"
	"github.com/sohosai/go-gin-sqlx-pg-template/internal/infrastructure"
	"github.com/sohosai/go-gin-sqlx-pg-template/internal/presentation/hadlers"
)

func main() {

	fmt.Println("Application Initializing...")

	dsn := os.Getenv("POSTGRES_DSN")
	db, err := infrastructure.Connect(dsn)
	if err != nil {
		fmt.Printf("Can NOT Connect Database %s", err)
		return
	}
	repos := domain.Repositories{PostRepo: infrastructure.PgPostRepository{DB: db}}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))

	hadlers.Handle(r, repos)

	fmt.Println("Application Starts!")
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
