package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/modelscope-open-world/internal/auth"
	"github.com/soulteary/modelscope-open-world/internal/user"
)

func main() {
	r := gin.Default()

	protected := r.Group("/")
	protected.Use(auth.AuthMiddleware)
	{
		protected.GET("/connect-to-modelscope", func(c *gin.Context) {
			userInfo, exist := c.Get("userInfo")
			if !exist {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "User not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Welcome", "user": userInfo.(user.Info).Name})
		})
	}

	r.Run(":8084")
}
