package k8s

import (
	"context"
	"fmt"
	"log"

	"github.com/chxmxii/challengefile/v2/internal/core/domain"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

func int32Ptr(i int) *int32 {
	v := int32(i)
	return &v
}

func CreateDeployment(client *kubernetes.Clientset, challenge *domain.Challenge) error {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      challenge.Name,
			Namespace: challenge.Metadata.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(challenge.Deployment.Replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": challenge.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": challenge.Name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  challenge.Name,
							Image: challenge.Deployment.Image,
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: challenge.Service.Port,
								},
							},
							LivenessProbe: &apiv1.Probe{
								ProbeHandler: apiv1.ProbeHandler{
									TCPSocket: &apiv1.TCPSocketAction{
										Port: intstr.FromInt32(challenge.Service.Port),
									},
								},
								InitialDelaySeconds: 15,
								PeriodSeconds:       10,
								TimeoutSeconds:      5,
								FailureThreshold:    3,
							},
						},
					},
				},
			},
		},
	}

	_, err := client.AppsV1().Deployments(challenge.Metadata.Namespace).Create(
		context.TODO(),
		deployment,
		metav1.CreateOptions{},
	)

	if err != nil {
		log.Printf("Failed to create deployment: %v", err)
		return fmt.Errorf("failed to create deployment: %w", err)
	}

	return nil
}

func DestroyDeployment(client *kubernetes.Clientset, challenge *domain.Challenge) error {
	deletePolicy := metav1.DeletePropagationForeground
	err := client.AppsV1().Deployments(challenge.Metadata.Namespace).Delete(
		context.TODO(),
		challenge.Name,
		metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		},
	)

	if err != nil {
		log.Printf("Failed to delete deployment: %v", err) // Changed from Fatalf to Printf
		return fmt.Errorf("failed to delete deployment: %w", err)
	}

	return nil
}
