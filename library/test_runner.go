package library

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func RunTests(suites []string, parallelSuites, parallelTests int, kubeconfig string) error {
	var wg sync.WaitGroup
	wg.Add(len(suites))

	for _, suite := range suites {
		go func(suite string) {
			defer wg.Done()

			cmd := exec.Command("go", "test", "-tags="+suite, fmt.Sprintf("-parallel=%d", parallelTests), "github.com/dramasamy/k8s-test/tests/...")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Env = append(os.Environ(), "KUBECONFIG="+kubeconfig)

			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error running tests for suite %s: %v\n", suite, err)
			}
		}(suite)
	}

	wg.Wait()
	return nil
}
