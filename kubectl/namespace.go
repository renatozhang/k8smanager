package kubectl

import (
	"context"
	"fmt"
	"path/filepath"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func kubeConfig() (clientset *kubernetes.Clientset, err error) {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig := filepath.Join(home, ".kube", "config")
		fmt.Println(kubeconfig)
		config, errRet := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			fmt.Println("Out cluster config error")
			err = errRet
			return
		}
		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			fmt.Println("Create clientset error")
			return
		}
	}
	return
}

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
	fmt.Printf("list namespace %#v \n", namesapceList)
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
