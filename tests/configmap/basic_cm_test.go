package configmap

import (
	"os"
	"testing"
	"time"

	"github.com/dramasamy/k8s-test/libs"
	"github.com/stretchr/testify/assert"
)

func TestCreateConfigMap(t *testing.T) {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		t.Fatalf("KUBECONFIG environment variable not set")
	}
	k8sClient, err := libs.CreateKubeClient(kubeconfig)
	assert.NoError(t, err)

	namespace := libs.GenerateRandomString(8)
	configMapName := libs.GenerateRandomString(8)

	// Create namespace
	err = libs.CreateNamespace(k8sClient, namespace)
	assert.NoError(t, err)

	// Create ConfigMap
	configMapData := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	err = libs.CreateConfigMap(k8sClient, namespace, configMapName, configMapData)
	assert.NoError(t, err)

	// Get ConfigMap and check if it exists
	configMap, err := libs.GetConfigMap(k8sClient, namespace, configMapName)
	assert.NoError(t, err)
	assert.NotNil(t, configMap)

	// Delete the configmap
	err = libs.DeleteConfigMap(k8sClient, namespace, configMapName)
	assert.NoError(t, err)

	// Wait for the configmap to be deleted
	err = libs.WaitForConfigMapDeletion(k8sClient, namespace, configMapName, 60*time.Second)
	assert.NoError(t, err)

	// Cleanup
	err = libs.DeleteNamespace(k8sClient, namespace)
	assert.NoError(t, err)

}
