package main

import (
	"k8smanager/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/namespace", controller.CreateNameSpace)

	router.Run()
}