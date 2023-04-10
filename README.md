# k8s-test

<p style="color: red;">This project is under development</p>

## Introduction

`k8s-test` is a comprehensive Kubernetes testing framework developed in Golang, leveraging the Testify library. It empowers users to write and run test cases for their Kubernetes clusters with flexibility, ease of use, and high customizability. With features such as parallel test execution, xfail and xpass markers, and customizable reporting, `k8s-test` offers a robust solution to validate and optimize Kubernetes deployments.

## Key Features

1. **Parallel Test Execution:** `k8s-test` allows you to run tests concurrently, significantly reducing the total time required to execute your test cases while improving resource efficiency.
2. **Parallel Test Suites:** `k8s-test` extends the parallelism capabilities to test suites, enabling you to validate different aspects of your Kubernetes cluster simultaneously, further streamlining the testing process.
3. **xfail and xpass Markers:** `k8s-test` incorporates xfail and xpass markers that help you flag tests that are expected to fail or pass, respectively. This feature is particularly useful when dealing with known issues or tracking improvements in your Kubernetes deployments.
4. **Customizable Reporting:** `k8s-test` generates test reports in multiple formats, such as JUnit XML, JSON, and HTML, catering to different reporting needs. Additionally, it offers the option to create custom reporters to generate reports in your preferred format.

## Usage

`k8s-test` can be configured and run through a CLI interface. The following command-line options are available:

- `-parallelSuites`: Number of test suites to run in parallel (default: 1)
- `-parallelTests`: Number of tests within a suite to run in parallel (default: 1)
- `-kubeconfig`: Path to the kubeconfig file (default: empty string)

To run `k8s-test`, specify the test suites you want to run and the desired CLI options. For example:

```
k8s-test -parallelSuites=2 -kubeconfig=path/to/kubeconfig.yaml configmap calico
```

This command will run the configmap and calico test suites in parallel, using the specified kubeconfig file.

## Contributing

To contribute to `k8s-test`, please follow these guidelines:

1. Fork the repository and create a new branch for your feature or bug fix.
2. Write tests for your code.
3. Run `go test` and ensure that all tests pass.
4. Submit a pull request.

## License

`k8s-test` is licensed under the MIT License. See the LICENSE file for more information.

