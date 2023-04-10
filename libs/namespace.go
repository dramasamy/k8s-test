package libs

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Generate a random string of the given length
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func CreateNamespace(clientset *kubernetes.Clientset, namespace string) error {
	exists, err := IsNamespaceExists(clientset, namespace)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("namespace '%s' already exists", namespace)
	}

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}

	_, err = clientset.CoreV1().Namespaces().Create(context.Background(), ns, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create namespace: %w", err)
	}

	return nil
}

func DeleteNamespace(client *kubernetes.Clientset, namespace string) error {
	return client.CoreV1().Namespaces().Delete(context.Background(), namespace, metav1.DeleteOptions{})
}

func IsNamespaceExists(clientset *kubernetes.Clientset, namespace string) (bool, error) {
	_, err := clientset.CoreV1().Namespaces().Get(context.Background(), namespace, metav1.GetOptions{})
	if err == nil {
		return true, nil
	} else if errors.IsNotFound(err) {
		return false, nil
	} else {
		return false, err
	}
}
