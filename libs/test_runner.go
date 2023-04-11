package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type TestEvent struct {
	Time    string `json:"Time"`
	Action  string `json:"Action"`
	Package string `json:"Package"`
	Test    string `json:"Test"`
	Output  string `json:"Output"`
}

func RunTests(suites []string, parallelSuites, parallelTests int, kubeconfig string) error {
	var wg sync.WaitGroup

	// Create a timestamped directory for log files
	logsDir := time.Now().Format("2006-01-02_15-04-05-logs")
	err := os.MkdirAll(logsDir, 0755)
	if err != nil {
		return fmt.Errorf("Error creating logs directory: %v", err)
	}

	for _, suite := range suites {
		wg.Add(1)
		go func(suite string) {
			defer wg.Done()

			var buf bytes.Buffer
			cmd := exec.Command("go", "test", "-tags="+suite, "-count=1", fmt.Sprintf("-parallel=%d", parallelTests), "-json", "-v", "./tests/"+suite)
			cmd.Stdout = &buf
			cmd.Stderr = os.Stderr
			cmd.Env = append(os.Environ(), "KUBECONFIG="+kubeconfig)

			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error running tests for suite %s: %v\n", suite, err)
			}

			// Parse JSON output
			decoder := json.NewDecoder(strings.NewReader(buf.String()))
			var testLogs = make(map[string][]TestEvent)
			var filteredJSON bytes.Buffer
			for {
				var event TestEvent
				err := decoder.Decode(&event)
				if err != nil {
					break
				}

				if event.Action != "output" {
					jsonEvent, _ := json.Marshal(event)
					filteredJSON.WriteString(string(jsonEvent) + "\n")
				}

				if event.Test != "" {
					testLogs[event.Test] = append(testLogs[event.Test], event)
				}
			}

			// Store filtered JSON output
			filteredJSONFile := filepath.Join(logsDir, fmt.Sprintf("%s_filtered.json", suite))
			err = ioutil.WriteFile(filteredJSONFile, filteredJSON.Bytes(), 0644)
			if err != nil {
				fmt.Printf("Error writing filtered JSON to file %s: %v\n", filteredJSONFile, err)
			}

			// Write logs to separate files
			for testName, events := range testLogs {
				fileName := fmt.Sprintf("%s_%s.log", suite, testName)
				filePath := filepath.Join(logsDir, fileName)
				f, err := os.Create(filePath)
				if err != nil {
					fmt.Printf("Error creating log file %s: %v\n", fileName, err)
					continue
				}
				defer f.Close()

				for _, event := range events {
					if event.Action == "output" {
						_, err = f.WriteString(event.Output)
						if err != nil {
							fmt.Printf("Error writing log to file %s: %v\n", fileName, err)
						}
					}
				}
			}
		}(suite)
	}

	wg.Wait()

	return nil
}
