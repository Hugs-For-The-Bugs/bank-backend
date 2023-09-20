package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func AuthCheck(c *gin.Context) {
	session := sessions.Default(c)
	socialSecurityNumber := session.Get("ssn")
	if socialSecurityNumber == nil {
		c.JSON(400, gin.H{
			"msg": "not login",
		})
		c.Abort()
	}

}

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func Middleware() gin.HandlerFunc {
	// Initialize UserRepository
	userRepo := New() // Assume New() initializes UserRepository

	return func(c *gin.Context) {
		// Logging: Capture the current time
		t := time.Now()

		// Logging: Set example variable (existing code)
		c.Set("example", "12345")

		// User Authentication: Fetch the default session
		session := sessions.Default(c)

		// User Authentication: Fetch the SSN from the session
		ssn := session.Get("ssn")

		// User Authentication: Find user by SSN
		user := userRepo.FindBySSN(ssn.(string))

		// User Authentication: If user does not exist, redirect to home
		if user.SocialSecurityNumber == "" {
			c.Redirect(http.StatusFound, "/")
			return
		}

		// Before handling the request, call c.Next() to move to the next middleware
		c.Next()

		// After request: Calculate latency and log it
		latency := time.Since(t)
		log.Print(latency)

		// After request: Fetch and log the HTTP status code
		status := c.Writer.Status()
		log.Println(status)
	}
}
