package cmd

import (
	"fmt"

	"github.com/chxmxii/challengefile/v2/internal/core/services"
	"github.com/chxmxii/challengefile/v2/pkg/gateways/config_manager/yaml"
	validate "github.com/chxmxii/challengefile/v2/pkg/gateways/validation"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringP("file", "f", "challengefile", "Path to the challenge configuration file")
}

// validator
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates challengefile",
	RunE:  validateHandler,
}

func validateHandler(cmd *cobra.Command, args []string) error {
	configFile, _ := cmd.Flags().GetString("file")

	yamlCfg := yaml.NewYamlConfig(
		yaml.YamlConfigLoader{ConfigFilePath: configFile},
	)

	challenge := services.NewChallengeManager(func(s *services.ChallengeManager) {
		s.ConfigManager = yamlCfg
	})

	challenges, err := challenge.ConfigManager.LoadAll()
	if err != nil {
		return err
	}

	for _, c := range challenges {
		err := validate.Validate(c)
		if err != nil {
			return err
		}
	}

	fmt.Println("your challengefile is valid (◠‿◠)")
	return nil
}
