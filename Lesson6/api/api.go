package api

import (
	"lesson06/middleware"
	"net/http"
	"time"

	"lesson06/dao"
	"lesson06/model"
	"lesson06/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}
	// 如果用户存在，这里这种是用户名可以一致的，即只要密码不一致就视为不同用户
	if dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user already exists",
		})
		return
	}
	dao.AddUser(req.Username, req.Password)
	c.JSON(http.StatusOK, gin.H{
		"message": "register success",
	})
	dao.Save()
}

func Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	// 检查用户是否存在且密码是否正确
	if !dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}
	token, err := utils.MakeToken(req.Username, time.Now().Add(10*time.Minute))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
	})
	dao.Save()
}

func UpdateUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	if !dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}
	ok, _ := utils.CheckToken(req.Token)
	if ok == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}
	token, err := utils.MakeToken(req.Username, time.Now().Add(10*time.Minute))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	dao.UpdateUser(req.Username, req.Password2)
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"token":   token,
	})
	dao.Save()
}

func refresh(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	ok, _ := utils.CheckToken(req.Token)
	if ok == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}
	token, err := utils.RefreshToken(req.Token, []byte("my_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "refresh success",
		"token":   token,
	})
}

func InitRouterGin() {
	dao.Init()
	r := gin.Default()
	r.POST("/ping", middleware.Example1(), middleware.Example2())
	r.POST("/login", Login)
	r.POST("/user", UpdateUser)
	r.POST("/register", Register)
	r.POST("/refresh", refresh)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
