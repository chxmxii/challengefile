package services

import (
	"github.com/chxmxii/challengefile/v2/internal/core/ports"
)

// ChallengeManager is a service that manages challenges
type ChallengeManager struct {
	ports.ConfigManager         // ConfigManager is used to load challenges
	ports.InfrastructureManager // InfrastructureManager where challenges are deployed
}

type ChallengeManagerConfigurer func(*ChallengeManager)

func NewChallengeManager(cfgs ...ChallengeManagerConfigurer) *ChallengeManager {
	cm := &ChallengeManager{}
	for _, cfg := range cfgs {
		cfg(cm)
	}

	return cm
}

func (c *ChallengeManager) DeployChallenge(name string) error {
	challengeCfg, err := c.ConfigManager.Load(name)
	if err != nil {
		return err
	}

	err = c.InfrastructureManager.DeployChallenge(challengeCfg)
	if err != nil {
		return err
	}

	return nil
}

func (c *ChallengeManager) DeployAllChallenges() error {
	challenges, err := c.ConfigManager.LoadAll()
	if err != nil {
		return err
	}

	for _, challenge := range challenges {
		err = c.InfrastructureManager.DeployChallenge(&challenge)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ChallengeManager) DestroyChallenge(name string) error {
	challengeCfg, err := c.ConfigManager.Load(name)
	if err != nil {
		return err
	}

	err = c.InfrastructureManager.DestroyChallenge(challengeCfg)
	if err != nil {
		return err
	}

	return nil
}

func (c *ChallengeManager) DestroyAllChallenges() error {
	challenges, err := c.ConfigManager.LoadAll()
	if err != nil {
		return err
	}

	for _, challenge := range challenges {
		err = c.InfrastructureManager.DestroyChallenge(&challenge)
		if err != nil {
			return err
		}
	}

	return nil
}
