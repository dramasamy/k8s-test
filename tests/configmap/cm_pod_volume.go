package configmap

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dramasamy/k8s-test/library"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConfigMapPodVolume(kubeconfig string) error {
	clientset, err := library.CreateKubeClient(kubeconfig)
	if err != nil {
		return err
	}

	namespace := "test-ns"
	err = library.CreateNamespace(clientset, namespace)
	if err != nil {
		return err
	}
	defer library.DeleteNamespace(clientset, namespace)

	cmName := "test-cm"
	cmData := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	err = library.CreateConfigMap(clientset, namespace, cmName, cmData)
	if err != nil {
		return err
	}
	defer library.DeleteConfigMap(clientset, namespace, cmName)

	podName := "test-pod"
	containerName := "test-container"
	image := "nginx:latest"
	mountPath := "/mnt/test"
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: namespace,
			Labels: map[string]string{
				"app": "test-app",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  containerName,
					Image: image,
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "config-volume",
							MountPath: mountPath,
						},
					},
				},
			},
			Volumes: []corev1.Volume{
				{
					Name: "config-volume",
					VolumeSource: corev1.VolumeSource{
						ConfigMap: &corev1.ConfigMapVolumeSource{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: cmName,
							},
						},
					},
				},
			},
		},
	}

	_, err = clientset.CoreV1().Pods(namespace).Create(context.Background(), pod, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create pod: %w", err)
	}
	defer clientset.CoreV1().Pods(namespace).Delete(context.Background(), podName, metav1.DeleteOptions{})

	err = library.WaitForPodCondition(clientset, namespace, podName, corev1.PodReady, corev1.ConditionTrue, 1*time.Minute)
	if err != nil {
		return fmt.Errorf("failed to wait for pod: %w", err)
	}

	logOutput, err := library.GetPodLog(clientset, namespace, podName, containerName)
	if err != nil {
		return fmt.Errorf("failed to get pod log: %w", err)
	}

	if !strings.Contains(logOutput, "key1=value1") || !strings.Contains(logOutput, "key2=value2") {
		return fmt.Errorf("configmap data not found in pod log")
	}

	return nil
}
