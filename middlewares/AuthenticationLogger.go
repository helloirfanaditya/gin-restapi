package middlewares

import (
	"strings"
	"trawlcode/utils"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, utils.ResError(401, "Unauthorized"))
			c.Abort()
			return
		}
		parseToken := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

		verify := utils.VerifyToken(parseToken)

		if verify == 0 {
			c.JSON(401, utils.ResError(401, "Unauthorized"))
			c.Abort()
		}
		c.Set("id", verify)
		c.Next()
	}
}
