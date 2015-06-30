package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/interval-io/backend/api"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	users := api.NewUsersResource(db)
	r.GET("/users", users.Index)
	r.POST("/users", users.Create)
	r.GET("/users/:userID", users.Show)
	r.PATCH("/users/:userID", users.Update)
	r.DELETE("/users/:userID", users.Destroy)

	r.Run(":8000")
}
