package library

import (
	"context"
	"io"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
)

// WaitForPodCondition waits for the specified condition type with the given status.
func WaitForPodCondition(client *kubernetes.Clientset, namespace, name string, conditionType corev1.PodConditionType, status corev1.ConditionStatus, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return wait.PollImmediate(1*time.Second, timeout, func() (bool, error) {
		pod, err := client.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		for _, condition := range pod.Status.Conditions {
			if condition.Type == conditionType && condition.Status == status {
				return true, nil
			}
		}
		return false, nil
	})
}

func GetPodLog(clientset *kubernetes.Clientset, namespace, podName, containerName string) (string, error) {
	req := clientset.CoreV1().Pods(namespace).GetLogs(podName, &corev1.PodLogOptions{
		Container: containerName,
	})
	logs, err := req.Stream(context.Background())
	if err != nil {
		return "", err
	}
	defer logs.Close()

	buf := new(strings.Builder)
	_, err = io.Copy(buf, logs)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
