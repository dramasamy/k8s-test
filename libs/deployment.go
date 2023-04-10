package libs

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
)

func WaitForDeploymentReady(clientset *kubernetes.Clientset, namespace, deploymentName string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return wait.PollImmediate(1*time.Second, timeout, func() (bool, error) {
		deploy, err := clientset.AppsV1().Deployments(namespace).Get(ctx, deploymentName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}

		if deploy.Status.ReadyReplicas == deploy.Status.Replicas {
			return true, nil
		}

		return false, nil
	})
}
