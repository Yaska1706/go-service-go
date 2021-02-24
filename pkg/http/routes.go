package http

import "github.com/gin-gonic/gin"

// routes holds all the routes you need for your service
func (r *router) routes() *gin.Engine {
	v1 := r.routerEngine.Group("/v1/api")
	{
		v1.GET("/hello-world", func(c *gin.Context) {
			c.Status(200)
		})
	}

	return r.routerEngine
}
