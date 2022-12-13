package controller

import (
	"fmt"
	"k8smanager/common"
	"k8smanager/kubectl"
	"k8smanager/util"

	"github.com/gin-gonic/gin"
)

func CreateNameSpace(ctx *gin.Context) {
	var namesapce common.Namespace
	err := ctx.ShouldBindJSON(&namesapce)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	err = kubectl.CreateNameSpace(namesapce.Name)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{"message": "cereate namespace success"})
	util.ResponseSuccess(ctx, nil)
}

func ListNameSpace(ctx *gin.Context) {
	namespaceList, err := kubectl.ListNameSpace()
	fmt.Printf("%#v", namespaceList)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "list namespace success",
	// 	"result":  namespaceList,
	// })
	util.ResponseSuccess(ctx, namespaceList)
}

func DeleteNameSpace(ctx *gin.Context) {
	var namesapce common.Namespace
	err := ctx.ShouldBindJSON(&namesapce)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	fmt.Println(namesapce.Name)
	err = kubectl.DeleteNameSpace(namesapce.Name)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{"message": "delete namespace success"})
	util.ResponseSuccess(ctx, nil)
}
