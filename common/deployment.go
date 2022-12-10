package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
)

type Deploy struct {
	App      string             `json:"app"`
	Tag      string             `json:"tag"`
	Replices int32              `json:"replices"`
	Dep      *appsv1.Deployment `json:"dep"`
}

func SetDeployment(deploy *Deploy) (deployment *appsv1.Deployment) {
	data, err := ioutil.ReadFile("./deploy.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &deployment)
	if err != nil {
		fmt.Printf("unmarshal failed,err:%v\n", err)
		return
	}
	// 设置副本数
	deployment.Spec.Replicas = &deploy.Replices
	// 设置deploy metadata
	deployment.ObjectMeta.Name = deploy.App
	deployment.ObjectMeta.Labels["app"] = deploy.App
	// 设置selector
	deployment.Spec.Selector.MatchLabels["app"] = deploy.App + "-" + deploy.Tag
	// 设置container相关
	deployment.Spec.Template.ObjectMeta.Labels["app"] = deploy.App + "-" + deploy.Tag
	deployment.Spec.Template.Spec.Containers = append(deployment.Spec.Template.Spec.Containers, v1.Container{
		Name:            deploy.App,                    //"tomcat"
		Image:           deploy.App + ":" + deploy.Tag, //"tomcat:8.0.18-jre8",
		ImagePullPolicy: "IfNotPresent",
		Ports: []apiv1.ContainerPort{
			{
				Name:          "http",
				Protocol:      apiv1.ProtocolSCTP,
				ContainerPort: 8080,
			},
		},
	})

	return
}
