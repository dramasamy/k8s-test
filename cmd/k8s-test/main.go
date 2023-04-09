package main

import (
	"flag"
	"fmt"

	"github.com/dramasamy/k8s-test/library"
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

	suites := []string{"configmap", "calico"}
	err := library.RunTests(suites, parallelSuites, parallelTests, kubeconfig)
	if err != nil {
		fmt.Printf("Error running tests: %v\n", err)
	}
}
