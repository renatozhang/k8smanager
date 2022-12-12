package kubectl

import (
	"context"
	"fmt"
	"k8smanager/common"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
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

func ListDeployment() (deploymentList *v1.DeploymentList, err error) {
	clientset, err := kubeConfig()
	if err != nil {
		fmt.Printf("Config error: %v", err.Error())
		return
	}
	// 得到deployment的客户端
	deploymentClient := clientset.AppsV1().Deployments("default")

	deploymentList, err = deploymentClient.List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		fmt.Printf("list deployment failed, err:%v\n", err)
		return
	}

	fmt.Printf("List deployment %s \n", deploymentList)

	return
}

func GetDeployment(app string) (deploymet *v1.Deployment, err error) {
	clientset, err := kubeConfig()
	if err != nil {
		fmt.Printf("Config error: %v", err.Error())
		return
	}
	// 得到deployment的客户端
	deploymentClient := clientset.AppsV1().Deployments("default")
	deploymet, err = deploymentClient.Get(context.TODO(), app, metav1.GetOptions{})
	return
}

func ScaleDeployment(deploy *common.Deploy) (err error) {
	clientset, err := kubeConfig()
	if err != nil {
		fmt.Printf("Config error: %v", err.Error())
		return
	}
	// 得到deployment的客户端
	deploymentClient := clientset.AppsV1().Deployments("default")
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		deployment, getErr := GetDeployment(deploy.App)
		if getErr != nil {
			err = getErr
			fmt.Printf("Failed to get latest version of Deployment: %v", getErr)
			return getErr
		}

		deployment.Spec.Replicas = &deploy.Replices // reduce replica count
		// result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
		_, updateErr := deploymentClient.Update(context.TODO(), deployment, metav1.UpdateOptions{})
		err = updateErr
		return updateErr
	})
	if retryErr != nil {
		err = retryErr
		fmt.Printf("Scale failed: %v", retryErr)
		return

	}
	fmt.Printf("Scale deployment replices %s, replicas:%d \n", deploy.App, deploy.Replices)
	return
}

func UpgradeDeployment(deploy *common.Deploy) (err error) {
	clientset, err := kubeConfig()
	if err != nil {
		fmt.Printf("Config error: %v", err.Error())
		return
	}
	// 得到deployment的客户端
	deploymentClient := clientset.AppsV1().Deployments("default")
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		deployment, getErr := GetDeployment(deploy.App)
		if getErr != nil {
			err = getErr
			fmt.Printf("Failed to get latest version of Deployment: %v", getErr)
			return getErr
		}

		deployment.Spec.Template.Spec.Containers[0].Image = deploy.App + ":" + deploy.Tag // change tomcat version
		_, updateErr := deploymentClient.Update(context.TODO(), deployment, metav1.UpdateOptions{})
		err = updateErr
		return updateErr
	})
	if retryErr != nil {
		err = retryErr
		fmt.Printf("Update failed: %v", retryErr)
		return

	}
	fmt.Printf("update deployment %s, tag:%s \n", deploy.App, deploy.Tag)
	return
}

func DeleteDeployment(deploy *common.Deploy) (err error) {
	clientset, err := kubeConfig()
	if err != nil {
		fmt.Printf("Config error: %v", err.Error())
		return
	}
	// 得到deployment的客户端
	deploymentClient := clientset.AppsV1().Deployments("default")
	deletePolicy := metav1.DeletePropagationForeground
	err = deploymentClient.Delete(context.TODO(), deploy.App, metav1.DeleteOptions{PropagationPolicy: &deletePolicy})
	if err != nil {
		fmt.Printf("Delete deployment %s failed, err:%v\n", deploy.App, err)
		return
	}
	fmt.Printf("Delete deployment %s success!", deploy.App)
	return
}
