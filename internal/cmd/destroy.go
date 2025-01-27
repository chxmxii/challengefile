package cmd

import (
	"github.com/chxmxii/challengefile/v2/internal/core/services"
	"github.com/chxmxii/challengefile/v2/pkg/gateways/config_manager/yaml"
	"github.com/chxmxii/challengefile/v2/pkg/gateways/deployer/k8s"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(destroyCmd)

	destroyCmd.Flags().StringP("challenge", "c", "", "Challenge name")
	destroyCmd.Flags().StringP("file", "f", "challengefile", "Path to the challenge configuration file")
	destroyCmd.Flags().StringP("kubeconfig", "k", "", "Path to the kubeconfig file")

}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroys challenges",
	RunE:  destroyerHandler,
}

func destroyerHandler(cmd *cobra.Command, args []string) error {

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

	destoyer := k8s.NewKM(clientset)

	challenge := services.NewChallengeManager(func(s *services.ChallengeManager) {
		s.ConfigManager = yamlCfg
		s.InfrastructureManager = destoyer
	})

	if challengeName == "" {
		err := challenge.DestroyAllChallenges()
		if err != nil {
			return err
		}
		return nil
	}

	err = challenge.DestroyChallenge(challengeName)
	if err != nil {
		return err
	}

	return nil
}
