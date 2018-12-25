package main

import (
	. "../app/controllers/apis"
	. "../app/controllers/msg"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/**/*") //渲染模板
	//静态资源
	router.Static("/public","./public")
	// router.GET("/", IndexApi)
	router.POST("/user", AddUserApi)
	router.GET("/userlist/list", GetUserApi)
	router.GET("/user/:id", GetUserByIdApi)
	router.GET("/delUser/:id", DelUserApi)
	router.GET("/index", IndexList)
	return router
}
