package main

import (
	"fmt"

	"github.com/communik/user-srv/database"
	"github.com/communik/user-srv/handler"
	"github.com/communik/user-srv/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("hello")
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}
	userRepo := repository.NewUserRepo(db)
	handler := handler.New(userRepo)
	router := gin.Default()
	router.POST("/user", handler.Create)
	router.Run()
}
