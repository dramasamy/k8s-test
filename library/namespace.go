package library

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateNamespace(client *kubernetes.Clientset, namespace string) error {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
	_, err := client.CoreV1().Namespaces().Create(context.Background(), ns, metav1.CreateOptions{})
	return err
}

func DeleteNamespace(client *kubernetes.Clientset, namespace string) error {
	return client.CoreV1().Namespaces().Delete(context.Background(), namespace, metav1.DeleteOptions{})
}
