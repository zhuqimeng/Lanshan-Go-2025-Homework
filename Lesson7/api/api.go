package api

import (
	"Lesson07/dao"
	"Lesson07/model"
	"Lesson07/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	var req model.QueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if fg, err := utils.CheckToken(req.Token); err != nil || fg == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}

	query := dao.GetToDos(&req)

	var results []dao.ToDo
	if err := query.Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

func Create(c *gin.Context) {
	var req model.QueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if fg, err := utils.CheckToken(req.Token); err != nil || fg == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}
	if err := dao.AddToDo(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Create ok!"})
}

func Update(c *gin.Context) {
	var req model.QueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if fg, err := utils.CheckToken(req.Token); err != nil || fg == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}
	if err := dao.UpdateToDo(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update ok!"})
}

func Delete(c *gin.Context) {
	var req model.QueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if fg, err := utils.CheckToken(req.Token); err != nil || fg == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}
	if err := dao.DeleteToDo(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete ok!"})
}

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
