package k8s

import (
	"encoding/json"
	"fmt"

	"github.com/chxmxii/challengefile/v2/internal/core/domain"
)

type FakeDeployer struct{}

func NewFakeDeployer() *FakeDeployer {
	return &FakeDeployer{}
}

func (f *FakeDeployer) DeployChallenge(challenge *domain.Challenge) error {
	fmt.Printf("Deploying task %s\n", challenge.Name)
	return nil
}

func (f *FakeDeployer) DestroyChallenge(challenge *domain.Challenge) error {
	jsonDump, _ := json.Marshal(challenge)
	fmt.Printf("Destroying task %s with config: %s\n", challenge.Name, jsonDump)
	return nil
}
