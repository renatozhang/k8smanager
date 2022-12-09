package controller

import (
	"fmt"
	"k8smanager/common"
	"k8smanager/kubectl"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNameSpace(ctx *gin.Context) {
	var namesapce common.Namespace
	err := ctx.ShouldBindJSON(&namesapce)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println(namesapce.Name)
	err = kubectl.CreateNameSpace(namesapce.Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "cereate namespace success"})
}
