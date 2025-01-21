package k8s

import (
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func NewClient(kubeconfig string) (*kubernetes.Clientset, error) {

	if kubeconfig == "" {
		if kubeconfig = os.Getenv("KUBECONFIG"); kubeconfig == "" {
			kubeconfig = filepath.Join(homedir.HomeDir(), ".kube", "config")
		} else {
			log.Fatal("Unable to find kubeconfig, please specify the path to the kubeconfig (use -k flag) or set KUBECONFIG environment variable")
		}
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
