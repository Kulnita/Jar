package main

import "github.com/Kelnit/Jar/router"
import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	router.MainRouter(r)
	r.Run(":8080")
}