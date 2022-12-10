package kubectl

import (
	"context"
	"fmt"
	"k8smanager/common"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateDeployment(deploy *common.Deploy) (err error) {
	clientset, err := kubeConfig()
	if err != nil {
		fmt.Printf("Config error: %v", err.Error())
		return
	}
	// 得到deployment的客户端
	deploymentClient := clientset.AppsV1().Deployments("default")
	deploy.Dep = common.SetDeployment(deploy)
	// fmt.Printf("debug %#v\n", deploy)
	// 实例化一个数据结构
	// deployment := deploy.Dep
	// deployment := &appsv1.Deployment{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name: "test-deployment",
	// 	},
	// 	Spec: appsv1.DeploymentSpec{
	// 		Replicas: pointer.Int32Ptr(2),
	// 		Selector: &metav1.LabelSelector{
	// 			MatchLabels: map[string]string{
	// 				"app": "tomcat",
	// 			},
	// 		},

	// 		Template: apiv1.PodTemplateSpec{
	// 			ObjectMeta: metav1.ObjectMeta{
	// 				Labels: map[string]string{
	// 					"app": "tomcat",
	// 				},
	// 			},
	// 			Spec: apiv1.PodSpec{
	// 				Containers: []apiv1.Container{
	// 					{
	// 						Name:            "tomcat",
	// 						Image:           "tomcat:8.0.18-jre8",
	// 						ImagePullPolicy: "IfNotPresent",
	// 						Ports: []apiv1.ContainerPort{
	// 							{
	// 								Name:          "http",
	// 								Protocol:      apiv1.ProtocolSCTP,
	// 								ContainerPort: 8080,
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	
	result, err := deploymentClient.Create(context.TODO(), deploy.Dep, metav1.CreateOptions{})
	// result, err := deploymentClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("create deploy failed, err:%v\n", err)
		return
	}

	fmt.Printf("Create deployment %s \n", result.GetName())

	return
}
