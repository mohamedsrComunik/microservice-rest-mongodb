package handler

import (
	"log"
	"net/http"

	"github.com/communik/user-srv/model"
	"github.com/communik/user-srv/repository"
	"github.com/gin-gonic/gin"
)

type handler struct {
	repo repository.RepositoryInterface
}

func New(repo repository.RepositoryInterface) handler {
	return handler{repo: repo}
}

func (h handler) Create(c *gin.Context) {
	name := c.PostForm("name")
	user := model.User{UserName: name}
	id, err := h.repo.Create(c, user)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
