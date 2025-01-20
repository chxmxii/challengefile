package yaml

import (
	"fmt"
	"log"
	"os"

	"github.com/chxmxii/challengefile/v2/internal/core/domain"
	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	YamlConfigLoader
}

func NewYamlConfig(cfg YamlConfigLoader) *YamlConfig {
	return &YamlConfig{cfg}
}

func (y *YamlConfig) Load(challenge string) (*domain.Challenge, error) {
	challenges, err := y.LoadAll()
	if err != nil {
		return nil, err
	}

	for _, c := range challenges {
		if c.Name == challenge {
			return &c, nil
		}
	}

	return nil, fmt.Errorf("challenge %s not found", challenge)
}

func (y *YamlConfig) LoadAll() ([]domain.Challenge, error) {
	bytes, err := os.ReadFile(y.ConfigFilePath)
	if err != nil {
		log.Fatalf("error reading config file %q: %v", y.ConfigFilePath, err)
		return nil, err
	}

	if len(bytes) == 0 {
		return nil, fmt.Errorf("config file %q is empty", y.ConfigFilePath)
	}

	challCfg := make(map[string]domain.Challenge, 0)

	err = yaml.Unmarshal(bytes, &challCfg)
	if err != nil {
		log.Fatalf("error unmarshalling config file %q: %v", y.ConfigFilePath, err)
		return nil, err
	}

	challenges := make([]domain.Challenge, 0, len(challCfg))

	for name, c := range challCfg {
		c.Name = name
		challenges = append(challenges, c)
	}

	return challenges, nil
}
