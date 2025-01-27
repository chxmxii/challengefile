package validation

import (
	"fmt"
	"strings"

	"github.com/chxmxii/challengefile/v2/internal/core/domain"
)

var (
	ErrRequiredParam = "Missing required parameter: %s."
	ErrInvalidValue  = "Invalid value for parameter: %s. Please provide a valid value."
)

func requiredParamErr(param string) error {
	return fmt.Errorf(ErrRequiredParam, param)
}

func Validate(config domain.Challenge) error {

	if err := blackListed(config.Name); err != nil {
		return err
	}
	if err := validateMetadata(config.Metadata); err != nil {
		return err
	}
	if err := validateDeployment(config.Deployment); err != nil {
		return err
	}
	if err := validateService(config.Service); err != nil {
		return err
	}
	return nil
}

func validateMetadata(metadata *domain.Metadata) error {
	if metadata.Namespace == "" {
		return requiredParamErr("namespace")
	}
	if metadata.Category == "" {
		return requiredParamErr("category")
	}
	return nil
}

func validateDeployment(deployment *domain.Deployment) error {
	if err := blackListed(deployment.Name); err != nil {
		return err
	}
	if err := validateImage(deployment.Image); err != nil {
		return err
	}
	if err := validateReplicas(deployment.Replicas); err != nil {
		return err
	}
	if err := validateHPA(deployment.HPA); err != nil {
		return err
	}
	if err := validateHealthCheck(deployment.HealthCheck); err != nil {
		return err
	}
	return nil
}

func validateService(service *domain.Service) error {
	if err := validatePort(service.Port); err != nil {
		return err
	}
	if err := validateProtocol(service.Protocol); err != nil {
		return err
	}
	// if err := validateDNS(service.DnsEndpoint); err != nil {
	// 	return err
	// }
	return nil
}

// Validate image param
func validateImage(image string) error {
	if image == "" {
		return requiredParamErr("image")
	}
	if err := blackListed(image); err != nil {
		return err
	}
	return nil
}

func blackListed(input string) error {
	specialCharacters := []string{".", "*", "\\", "+", "@", "`", "§", "#", "~", "!", "$", "%", "^", "&", "="}
	if input == "" {
		return requiredParamErr(input)
	}

	for _, c := range specialCharacters {
		if strings.Contains(input, c) {
			return fmt.Errorf("'%s' contains invalid characters. The following characters are not allowed: %v", input, specialCharacters)
		}
	}
	return nil
}

// Validate port param
func validatePort(port int32) error {
	if port == 0 {
		return requiredParamErr("port")
	}
	if port < 30000 || port > 35000 {
		return fmt.Errorf("invalid port: %d. Please choose a port between 30000 and 35000", port)
	}
	return nil
}

// Validate replicas param
func validateReplicas(replicas int) error {
	if replicas == 0 {
		return requiredParamErr("replicas")
	}
	if replicas < 0 {
		return fmt.Errorf("invalid number of replicas: %d.", replicas)
	}
	return nil
}

// Validate HPA param
func validateHPA(hpa bool) error {
	if hpa {
		return nil
	}
	return nil
}

// Validate healthCheck param
func validateHealthCheck(healthCheck bool) error {
	if healthCheck {
		return nil
	}
	return nil
}

// Validate protocol param
func validateProtocol(protocol domain.Protocol) error {
	validProtocols := []string{"TCP", "UDP"}
	if protocol == "" {
		return requiredParamErr("protocol")
	}

	if protocol != "TCP" && protocol != "UDP" {
		return fmt.Errorf("invalid protocol: %s. Please choose a protocol between %v", protocol, validProtocols)
	}

	return nil
}

// Validate DNS endpoint param
// func validateDNS(dns string) error {
// if dns == "" {
// return requiredParamErr("dnsEndpoint")
// }
// return nil
// }
