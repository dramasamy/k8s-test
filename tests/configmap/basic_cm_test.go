package configmap

import (
	"os"
	"testing"
	"time"

	"github.com/dramasamy/k8s-test/library"
	"github.com/stretchr/testify/assert"
)

func TestCreateConfigMap(t *testing.T) {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		t.Fatalf("KUBECONFIG environment variable not set")
	}
	k8sClient, err := library.CreateKubeClient(kubeconfig)
	assert.NoError(t, err)

	namespace := "cm2-test-namespace"
	configMapName := "cm2-test-configmap"

	// Create namespace
	err = library.CreateNamespace(k8sClient, namespace)
	assert.NoError(t, err)

	// Create ConfigMap
	configMapData := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	err = library.CreateConfigMap(k8sClient, namespace, configMapName, configMapData)
	assert.NoError(t, err)

	// Get ConfigMap and check if it exists
	configMap, err := library.GetConfigMap(k8sClient, namespace, configMapName)
	assert.NoError(t, err)
	assert.NotNil(t, configMap)

	// Delete the configmap
	err = library.DeleteConfigMap(k8sClient, namespace, configMapName)
	assert.NoError(t, err)

	// Wait for the configmap to be deleted
	err = library.WaitForConfigMapDeletion(k8sClient, namespace, configMapName, 60*time.Second)
	assert.NoError(t, err)

	// Cleanup
	err = library.DeleteNamespace(k8sClient, namespace)
	assert.NoError(t, err)

}
