package library

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateConfigMap(client *kubernetes.Clientset, namespace, name string, data map[string]string) error {
	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: data,
	}
	_, err := client.CoreV1().ConfigMaps(namespace).Create(context.Background(), cm, metav1.CreateOptions{})
	return err
}

func GetConfigMap(client *kubernetes.Clientset, namespace, name string) (*corev1.ConfigMap, error) {
	return client.CoreV1().ConfigMaps(namespace).Get(context.Background(), name, metav1.GetOptions{})
}

func DeleteConfigMap(client *kubernetes.Clientset, namespace, name string) error {
	return client.CoreV1().ConfigMaps(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}

func WaitForConfigMapDeletion(client *kubernetes.Clientset, namespace, name string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timed out waiting for configmap deletion")
		default:
			_, err := client.CoreV1().ConfigMaps(namespace).Get(context.Background(), name, metav1.GetOptions{})
			if errors.IsNotFound(err) {
				return nil
			}
			time.Sleep(1 * time.Second)
		}
	}
}
