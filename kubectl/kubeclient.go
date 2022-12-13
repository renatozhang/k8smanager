package kubectl

import (
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func kubeConfig() (clientset *kubernetes.Clientset, err error) {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig := filepath.Join(home, ".kube", "config")
		// fmt.Println(kubeconfig)
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
