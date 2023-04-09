# k8s-test

<p style="color: red;">This project is under development</p>

## Introduction
k8s-test is a comprehensive Kubernetes testing framework developed in Golang, leveraging the Testify library. It empowers users to write and run test cases for their Kubernetes clusters with flexibility, ease of use, and high customizability. With features such as parallel test execution, xfail and xpass markers, and customizable reporting, k8s-test offers a robust solution to validate and optimize Kubernetes deployments.

## Key Features
1. **Parallel Test Execution:** k8s-test allows you to run tests concurrently, significantly reducing the total time required to execute your test cases while improving resource efficiency.
2. **Parallel Test Suites:** k8s-test extends the parallelism capabilities to test suites, enabling you to validate different aspects of your Kubernetes cluster simultaneously, further streamlining the testing process.
3. **xfail and xpass Markers:** k8s-test incorporates xfail and xpass markers that help you flag tests that are expected to fail or pass, respectively. This feature is particularly useful when dealing with known issues or tracking improvements in your Kubernetes deployments.
4. **Customizable Reporting:** k8s-test generates test reports in multiple formats, such as JUnit XML, JSON, and HTML, catering to different reporting needs. Additionally, it offers the option to create custom reporters to generate reports in your preferred format.
