package middleware

import (
	"net/http"

	"example.com/m/helper"
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"message": message,
	})
}

func TokenAuthMiddleware() gin.HandlerFunc { // token확인 middleware

	return func(c *gin.Context) {

		if err := helper.TokenVaild(c); err != nil {
			respondWithError(c, http.StatusBadGateway, "vaild error")
			c.Abort()
			return
		}

		ad, err := helper.ExtractTokenMetadata(c)
		if err != nil {
			respondWithError(c, http.StatusBadGateway, "extract token error")
			return
		}

		if _, err := helper.FetchAuth(ad); err != nil {
			respondWithError(c, http.StatusBadGateway, "fetch auth error")
		}

		c.Next()
	}
}
