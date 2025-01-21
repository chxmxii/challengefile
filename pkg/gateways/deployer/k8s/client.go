package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/chxmxii/challengefile/v2/internal/core/domain"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Deployer struct {
	Client *kubernetes.Clientset
}

func NewDeployer(client *kubernetes.Clientset) *Deployer {
	return &Deployer{Client: client}
}

func (d *Deployer) DeployChallenge(challenge *domain.Challenge) error {
	clientset := d.Client
	for {
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		namespace := "default"
		pod := "test"
		_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
				pod, namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		}

		time.Sleep(10 * time.Second)
	}
	fmt.Printf("Deploying task %s\n", challenge.Name)
	return nil
}

func (d *Deployer) DestroyChallenge(challenge *domain.Challenge) error {
	jsonDump, _ := json.Marshal(challenge)
	fmt.Printf("Destroying task %s with config: %s\n", challenge.Name, jsonDump)
	return nil
}

// func (d *Deployer) UpdateChallenge(challenge *domain.Challenge) error {
// 	fmt.Printf("Updating task %s\n", challenge.Name)
// 	return nil
// }
