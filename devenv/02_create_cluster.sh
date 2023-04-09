#!/bin/bash

# Default Kubernetes version and cluster name
K8S_VERSION="1.26"
CLUSTER_NAME="kind"

# Check if a version argument was provided
if [ $# -eq 1 ]; then
    K8S_VERSION=$1
    CLUSTER_NAME="kind-$(echo $K8S_VERSION | tr . -)"
fi

# Set the node image based on the Kubernetes version
case "$K8S_VERSION" in
    "1.26")
        NODE_IMAGE="kindest/node:v1.26.3@sha256:61b92f38dff6ccc29969e7aa154d34e38b89443af1a2c14e6cfbd2df6419c66f"
        ;;
    "1.25")
        NODE_IMAGE="kindest/node:v1.25.8@sha256:00d3f5314cc35327706776e95b2f8e504198ce59ac545d0200a89e69fce10b7f"
        CLUSTER_NAME="kind-1-25"
        ;;
    "1.24")
        NODE_IMAGE="kindest/node:v1.24.12@sha256:1e12918b8bc3d4253bc08f640a231bb0d3b2c5a9b28aa3f2ca1aee93e1e8db16"
        ;;
    "1.23")
        NODE_IMAGE="kindest/node:v1.23.17@sha256:e5fd1d9cd7a9a50939f9c005684df5a6d145e8d695e78463637b79464292e66c"
        ;;
    "1.22")
        NODE_IMAGE="kindest/node:v1.22.17@sha256:c8a828709a53c25cbdc0790c8afe12f25538617c7be879083248981945c38693"
        ;;
    "1.21")
        NODE_IMAGE="kindest/node:v1.21.14@sha256:27ef72ea623ee879a25fe6f9982690a3e370c68286f4356bf643467c552a3888"
        ;;
    *)
        echo "Invalid Kubernetes version: $K8S_VERSION"
        exit 1
        ;;
esac

# Create the Kind cluster using the node image, config, and name
cat <<EOF | kind create cluster --name=$CLUSTER_NAME --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  image: $NODE_IMAGE
- role: worker
  image: $NODE_IMAGE
- role: worker
  image: $NODE_IMAGE
EOF

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml

kubectl wait --namespace metallb-system \
                --for=condition=ready pod \
                --selector=app=metallb \
                --timeout=90s