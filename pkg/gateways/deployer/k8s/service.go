package k8s

import (
	"context"
	"fmt"
	"log"

	"github.com/chxmxii/challengefile/v2/internal/core/domain"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateService(client *kubernetes.Clientset, challenge *domain.Challenge) error {
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      challenge.Name,
			Namespace: challenge.Metadata.Namespace,
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"category":   challenge.Metadata.Category,
				"managed-by": "challengefile",
			},
			Ports: []v1.ServicePort{
				{
					Port: challenge.Service.Port,
				},
			},
		},
	}

	_, err := client.CoreV1().Services(challenge.Metadata.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return fmt.Errorf("error creating service: %w", err)
	}

	return nil
}

func DestroyService(client *kubernetes.Clientset, challenge *domain.Challenge) error {
	err := client.CoreV1().Services(challenge.Metadata.Namespace).Delete(context.Background(), challenge.Name, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("Error deleting service: %v", err)
		return fmt.Errorf("error deleting service: %w", err)
	}

	return nil
}
