package main

import (
	"k8smanager/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// namespace
	router.POST("/namespace", controller.CreateNameSpace)
	router.GET("/namespace", controller.ListNameSpace)
	router.DELETE("/namespace", controller.DeleteNameSpace)

	// deployment
	router.POST("/deploy", controller.CreateDeployment)
	router.GET("/deploy", controller.ListDeployment)
	router.PUT("/deploy/scale", controller.ScaleDeployment)
	router.PUT("/deploy/upgrade", controller.UpgradeDeployment)
	router.DELETE("/deploy", controller.DeleteDeployment)
	router.GET("/deploy/pods", controller.GetPods)

	router.Run()
}
