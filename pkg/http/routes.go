package http

import "github.com/gin-gonic/gin"

func (r *router) Routes() *gin.Engine {
	router := r.Engine

	return router
}
