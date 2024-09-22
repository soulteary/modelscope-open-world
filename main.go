package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/modelscope-open-world/internal/user"
)

func main() {
	r := gin.Default()

	r.GET("/get-userinfo", func(c *gin.Context) {
		cookies := c.Request.Header.Get("Cookie")
		if cookies == "" {
			c.String(http.StatusUnauthorized, "Cookie is empty")
			return
		}

		userInfo, err := user.GetBasicUserInfo(cookies)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.String(http.StatusOK, userInfo.Name)
	})

	r.Run(":8084")
}
