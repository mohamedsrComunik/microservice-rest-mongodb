package main

import (
	"fmt"

	"github.com/communik/user-srv/database"
	"github.com/communik/user-srv/handler"
	"github.com/communik/user-srv/repository"
	"github.com/communik/user-srv/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}
	userRepo := repository.NewUserRepo(db)
	handler := handler.New(userRepo)
	r := gin.Default()
	version1 := r.Group("/user-srv/api/v1/")
	routes := router.New(version1, handler)
	routes.SetRoutes()
	r.Run()
}
