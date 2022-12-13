package kubectl

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateNameSpace(name string) (err error) {
	clientset, err := kubeConfig()
	if err != nil {
		fmt.Printf("Config error: %v", err.Error())
		return
	}
	namespaceClient := clientset.CoreV1().Namespaces()
	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	result, err := namespaceClient.Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Create namesapce error: %v", err.Error())
		return
	}
	fmt.Printf("ceate namespace %s \n", result.GetName())
	return
}

func ListNameSpace() (namesapceList *apiv1.NamespaceList, err error) {
	clientset, err := kubeConfig()
	if err != nil {
		fmt.Printf("Config error: %v", err.Error())
		return
	}
	namespaceClient := clientset.CoreV1().Namespaces()
	namesapceList, err = namespaceClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("list namesapce error: %v", err.Error())
		return
	}
	// fmt.Printf("list namespace %#v \n", namesapceList)
	return
}

func DeleteNameSpace(name string) (err error) {
	clientset, err := kubeConfig()
	if err != nil {
		fmt.Printf("Config error: %v", err.Error())
		return
	}
	namespaceClient := clientset.CoreV1().Namespaces()
	return namespaceClient.Delete(context.TODO(), name, metav1.DeleteOptions{})
}
