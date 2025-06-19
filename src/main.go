package main

import (
	"fmt"
	"os"

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
	hadlers.Handle(r, repos)

	fmt.Println("Application Starts!")
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
