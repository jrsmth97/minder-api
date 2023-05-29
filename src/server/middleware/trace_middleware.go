package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Trace(c *gin.Context) {
	now := time.Now()
	log.Printf("Get request with method :%v Path :%v\n", c.Request.Method, c.Request.URL)
	c.Next()
	isError := c.GetString("ERROR")
	if isError != "" {
		log.Printf("get error when try to get all typicode :%v\n", isError)
	}
	log.Printf("Finised request with method :%v Path :%v\n", c.Request.Method, c.Request.URL)

	end := time.Since(now).Milliseconds()
	log.Println("response time:", end)
}
