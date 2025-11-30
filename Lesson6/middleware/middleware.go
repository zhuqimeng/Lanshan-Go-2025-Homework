package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Example1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mid1 before")
		//先执行next,后执行after
		c.Next()
		fmt.Println("mid1 after")
	}
}

func Example2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mid2 before")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
		//先执行next,后执行after
		c.Next()
		fmt.Println("mid2 after")
	}
}
