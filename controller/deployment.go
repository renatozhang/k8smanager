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
