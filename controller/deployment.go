package controller

import (
	"k8smanager/common"
	"k8smanager/kubectl"
	"k8smanager/util"

	"github.com/gin-gonic/gin"
)

func CreateDeployment(ctx *gin.Context) {
	var deployment common.Deploy
	err := ctx.ShouldBindJSON(&deployment)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}

	err = kubectl.CreateDeployment(&deployment)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{"message": "cereate deploy success"})
	util.ResponseSuccess(ctx, nil)
}

func ListDeployment(ctx *gin.Context) {

	deploymentList, err := kubectl.ListDeployment()
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	// ctx.JSON(http.StatusOK, deploymentList)
	util.ResponseSuccess(ctx, deploymentList)
}

func ScaleDeployment(ctx *gin.Context) {
	var deployment common.Deploy
	err := ctx.ShouldBindJSON(&deployment)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	err = kubectl.ScaleDeployment(&deployment)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{"message": "scale deploy success"})
	util.ResponseSuccess(ctx, nil)
}

func UpgradeDeployment(ctx *gin.Context) {
	var deployment common.Deploy
	err := ctx.ShouldBindJSON(&deployment)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	err = kubectl.UpgradeDeployment(&deployment)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{"message": "upgrade deploy success"})
	util.ResponseSuccess(ctx, nil)
}

func DeleteDeployment(ctx *gin.Context) {
	var deployment common.Deploy
	err := ctx.ShouldBindJSON(&deployment)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	err = kubectl.DeleteDeployment(&deployment)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{"message": "delete deploy success"})
	util.ResponseSuccess(ctx, nil)
}

func GetPods(ctx *gin.Context) {
	app := ctx.Query("app")
	if app == "" {
		// ctx.JSON(http.StatusBadRequest, gin.H{"err": "参数错误"})
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	podList, err := kubectl.GetPods(app)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
		util.ResponseError(ctx, util.ErrCodeServerBusy)
	}
	// ctx.JSON(http.StatusOK, podList)
	util.ResponseSuccess(ctx, podList)
}
