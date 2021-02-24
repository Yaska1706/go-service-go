package http

import (
	"github.com/gin-gonic/gin"
	"go-mbv-go/pkg/api"
)

type Router interface {
	Run(addr string)
	Routes() *gin.Engine
}

type router struct {
	api.API
	routerEngine *gin.Engine
}

func NewRouter(r *gin.Engine, api api.API) Router {
	return &router{api, r}
}

func (r router) Run(addr string) {
	router := r.Routes()
	_ = router.Run(addr)
}
