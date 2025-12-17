package api

import (
	"Lesson07/dao"
	"Lesson07/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := dao.CheckUser(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Token": token})
}

func Register(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dao.CreateUser(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Register ok!"})
}

func InitRouterGin() {
	r := gin.Default()
	r.GET("/show", Show)
	r.GET("/login", Login)
	r.GET("/register", Register)
	r.POST("/create", Create)
	r.POST("/update", Update)
	r.POST("/delete", Delete)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
