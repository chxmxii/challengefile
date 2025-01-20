package cmd

import (
	"github.com/chxmxii/challengefile/v2/internal/core/services"
	"github.com/chxmxii/challengefile/v2/pkg/gateways/config_manager/yaml"
	"github.com/chxmxii/challengefile/v2/pkg/gateways/deployer/k8s"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringP("challenge", "c", "", "Challenge name")
	deployCmd.Flags().StringP("file", "f", "challengefile", "Path to the challenge configuration file")
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a challenge",
	RunE:  deployHandler,
}

func deployHandler(cmd *cobra.Command, args []string) error {

	configFile, _ := cmd.Flags().GetString("file")
	challengeName, _ := cmd.Flags().GetString("challenge")

	yamlCfg := yaml.NewYamlConfig(
		yaml.YamlConfigLoader{ConfigFilePath: configFile},
	)

	deployer := k8s.NewFakeDeployer()

	challenge := services.NewChallengeManager(func(s *services.ChallengeManager) {
		s.ConfigManager = yamlCfg
		s.InfrastructureManager = deployer
	})

	if challengeName == "" {
		err := challenge.DeployAllChallenges()
		if err != nil {
			return err
		}
		return nil
	}

	err := challenge.DeployChallenge(challengeName)
	if err != nil {
		return err
	}

	return nil
}
