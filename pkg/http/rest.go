package http

import (
	"github.com/gin-gonic/gin"
	"go-mbv-go/pkg/api"
)

type Router interface {
	routes() *gin.Engine
	Run(addr string) error
}

type router struct {
	api.API
	routerEngine *gin.Engine
}

func NewRouter(r *gin.Engine, api api.API) Router {
	return &router{api, r}
}

func (r *router) Run(addr string) error {
	routes := r.routes()

	return routes.Run(addr)
}
