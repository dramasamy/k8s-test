package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dramasamy/k8s-test/libs"
)

var (
	parallelSuites int
	parallelTests  int
	kubeconfig     string
)

func init() {
	flag.IntVar(&parallelSuites, "parallelSuites", 1, "Number of test suites to run in parallel")
	flag.IntVar(&parallelTests, "parallelTests", 1, "Number of tests within a suite to run in parallel")
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to the kubeconfig file")
}

func main() {
	flag.Parse()

	if kubeconfig == "" {
		kubeconfig = os.Getenv("KUBECONFIG")
		if kubeconfig == "" {
			kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
		}
	}

	suites := []string{"configmap", "calico"}
	err := libs.RunTests(suites, parallelSuites, parallelTests, kubeconfig)
	if err != nil {
		fmt.Printf("Error running tests: %v\n", err)
	}
}
