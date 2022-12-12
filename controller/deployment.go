package controller

import (
	"k8smanager/common"
	"k8smanager/kubectl"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDeployment(ctx *gin.Context) {
	var deployment common.Deploy
	err := ctx.ShouldBindJSON(&deployment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = kubectl.CreateDeployment(&deployment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "cereate deploy success"})
}

func ListDeployment(ctx *gin.Context) {

	deploymentList, err := kubectl.ListDeployment()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, deploymentList)
}

func ScaleDeployment(ctx *gin.Context) {
	var deployment common.Deploy
	err := ctx.ShouldBindJSON(&deployment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = kubectl.ScaleDeployment(&deployment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "scale deploy success"})
}

func UpgradeDeployment(ctx *gin.Context) {
	var deployment common.Deploy
	err := ctx.ShouldBindJSON(&deployment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = kubectl.UpgradeDeployment(&deployment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "upgrade deploy success"})
}

func DeleteDeployment(ctx *gin.Context) {
	var deployment common.Deploy
	err := ctx.ShouldBindJSON(&deployment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = kubectl.DeleteDeployment(&deployment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "delete deploy success"})
}

func GetPods(ctx *gin.Context) {
	app := ctx.Query("app")
	if app == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "参数错误"})
		return
	}
	podList, err := kubectl.GetPods(app)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
	}
	ctx.JSON(http.StatusOK, podList)
}
