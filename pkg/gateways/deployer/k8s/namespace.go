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

func CreateNameSpace(client *kubernetes.Clientset, challenge *domain.Metadata) error {
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: challenge.Namespace,
			Labels: map[string]string{
				"category":   challenge.Category,
				"managed-by": "challengefile",
			},
		},
	}
	_, err := client.CoreV1().Namespaces().Create(context.Background(), namespace, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating namespace: %v", err)
		return fmt.Errorf("error creating namespace: %w", err)
	}

	return nil
}

func DestroyNameSpace(client *kubernetes.Clientset, challenge *domain.Metadata) error {
	err := client.CoreV1().Namespaces().Delete(context.Background(), challenge.Namespace, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("Error deleting namespace: %v", err)
		return fmt.Errorf("error deleting namespace: %w", err)
	}

	return nil
}
