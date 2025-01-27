package ports

import (
	"github.com/chxmxii/challengefile/v2/internal/core/domain"
)

type ConfigManager interface {
	Load(ChallengeName string) (*domain.Challenge, error) // Load a challenge by name
	LoadAll() ([]domain.Challenge, error)                 // Load all challenges
}

type InfrastructureManager interface {
	DeployChallenge(challenge *domain.Challenge) error  // Deploy a challenge
	DestroyChallenge(challenge *domain.Challenge) error // Destroy a challenge
}
