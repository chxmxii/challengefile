package k8s

import (
	"encoding/json"
	"fmt"

	"github.com/chxmxii/challengefile/v2/internal/core/domain"
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

func (d *Deployer) DestroyChallenge(challenge *domain.Challenge) error {
	jsonDump, _ := json.Marshal(challenge)
	fmt.Printf("Destroying task %s with config: %s\n", challenge.Name, jsonDump)
	return nil
}

// func (d *Deployer) UpdateChallenge(challenge *domain.Challenge) error {
// 	fmt.Printf("Updating task %s\n", challenge.Name)
// 	return nil
// }
