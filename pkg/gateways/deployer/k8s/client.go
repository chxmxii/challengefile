package k8s

import (
	"fmt"

	"github.com/chxmxii/challengefile/v2/internal/core/domain"
	"k8s.io/client-go/kubernetes"
)

type KubeManager struct {
	Client *kubernetes.Clientset
}

// NewKM creates a new KubeManager instance
func NewKM(client *kubernetes.Clientset) *KubeManager {
	return &KubeManager{Client: client}
}

func (k *KubeManager) DeployChallenge(challenge *domain.Challenge) error {
	clientset := k.Client
	err := CreateNameSpace(clientset, challenge.Metadata)
	if err != nil {
		return err
	}
	err = CreateDeployment(clientset, challenge)
	if err != nil {
		return err
	}
	err = CreateService(clientset, challenge)
	if err != nil {
		return err
	}

	fmt.Printf("Deployed task %s within namespace %s\n", challenge.Name, challenge.Metadata.Namespace)

	return nil
}

func (k *KubeManager) DestroyChallenge(challenge *domain.Challenge) error {
	clientset := k.Client
	err := DestroyNameSpace(clientset, challenge.Metadata)
	if err != nil {
		return err
	}

	fmt.Printf("Destroyed task %s within namespace %s\n", challenge.Name, challenge.Metadata.Namespace)

	return nil
}

// func (d *Deployer) UpdateChallenge(challenge *domain.Challenge) error {
// 	fmt.Printf("Updating task %s\n", challenge.Name)
// 	return nil
// }
