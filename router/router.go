package router

import "github.com/Kelnit/Jar/controller"
import "github.com/gin-gonic/gin"

func MainRouter(route *gin.Engine){
	route.GET("/", controller.GetMain)
	route.GET("/users", controller.GetUsers)
	route.GET("/userint/:user_id", controller.GetUsersID)
	route.GET("/insertform", controller.Formula)
	route.POST("/insertuser", controller.MoreUsers)
}