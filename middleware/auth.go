package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthCheck(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil || user != "admin" {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "not login",
		})
		c.Abort()
	}

}
