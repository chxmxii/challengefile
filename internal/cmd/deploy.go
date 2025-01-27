package cmd

import (
	"github.com/chxmxii/challengefile/v2/internal/core/services"
	"github.com/chxmxii/challengefile/v2/pkg/gateways/config_manager/yaml"
	"github.com/chxmxii/challengefile/v2/pkg/gateways/deployer/k8s"
	_ "github.com/chxmxii/challengefile/v2/pkg/gateways/validation"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringP("challenge", "c", "", "Challenge name")
	deployCmd.Flags().StringP("file", "f", "challengefile", "Path to the challenge configuration file")
	deployCmd.Flags().StringP("kubeconfig", "k", "", "Path to the kubeconfig file")

}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a challenge",
	RunE:  deployHandler,
}

func deployHandler(cmd *cobra.Command, args []string) error {

	configFile, _ := cmd.Flags().GetString("file")
	challengeName, _ := cmd.Flags().GetString("challenge")
	kubecfg, _ := cmd.Flags().GetString("kubeconfig")

	yamlCfg := yaml.NewYamlConfig(
		yaml.YamlConfigLoader{ConfigFilePath: configFile},
	)

	clientset, err := k8s.NewClient(kubecfg)
	if err != nil {
		return err
	}

	deployer := k8s.NewDeployer(clientset)

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

	err = challenge.DeployChallenge(challengeName)
	if err != nil {
		return err
	}

	return nil
}
