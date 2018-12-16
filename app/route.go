package main

import (
	. "../app/controllers/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexApi)
	router.POST("/user", AddUserApi)
	router.GET("/userlist/list", GetUserApi)
	router.GET("/user/:id", GetUserByIdApi)
	router.GET("/delUser/:id", DelUserApi)
	return router
}
