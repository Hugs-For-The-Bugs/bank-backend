package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginCheck(c *gin.Context) {
	session := sessions.Default(c)
	socialSecurityNumber := session.Get("id")

	if socialSecurityNumber == nil {
		c.JSON(400, gin.H{
			"msg": "not login",
		})
		c.Abort()
	}

}
