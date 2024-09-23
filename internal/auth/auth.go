package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/modelscope-open-world/internal/user"
)

// Perform authorization verification
func AuthMiddleware(c *gin.Context) {
	cookies := c.Request.Header.Get("Cookie")
	if cookies == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	userInfo, err := user.GetBasicUserInfo(cookies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.Set("userInfo", userInfo)
	c.Next()
}
