package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *router) CreateWeightEntry() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		r.UserService.NewUser("")

		response := map[string]string{
			"status": "success",
			"data":   "new user created",
		}

		c.JSON(http.StatusOK, response)
	}
}
