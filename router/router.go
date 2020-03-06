package router

import (
	"github.com/communik/user-srv/handler"
	"github.com/gin-gonic/gin"
)

type router struct {
	routerGroup *gin.RouterGroup
	handler     handler.HandlerInterface
}

// New to create new instance of router struct with deps
func New(routerGroup *gin.RouterGroup, handler handler.HandlerInterface) router {
	return router{
		routerGroup,
		handler,
	}
}

func (rt router) SetRoutes() {
	rt.routerGroup.POST("/user", rt.handler.Create)
}
