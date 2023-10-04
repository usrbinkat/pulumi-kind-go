# Zot Developer IaC

## 🌟 Overview

This Infrastructure as Code (IaC) project leverages the power of Pulumi and KinD (Kubernetes in Docker) to provision a Zot service and accompanying resources for Zot development and testing.

### 📊 What It Solves

- **Rapid Prototyping**: Ideal for developing Zot on Kubernetes.
- **Local Development**: Perfect for testing locally on Kind Kubernetes.
- **Automation**: Easily integrate into CI/CD pipelines.

## 🚀 Getting Started

These instructions will help you set up a Zot on KinD cluster.

    ```bash
    usrbinkat@mordor:~/uor-zot/iac$ time pulumi up -y
    Updating (dev)

         Type                              Name                           Status              Info
     +   pulumi:pulumi:Stack               zot-dev                        created (0.36s)     14 messages
     +   ├─ my:kind:KindCluster            default-cluster                created (42s)
     +   │  ├─ command:local:Command       default-cluster-deleteCluster  created (0.14s)
     +   │  └─ command:local:Command       default-cluster-createCluster  created (42s)
     +   ├─ command:local:Command          default-cluster-volumeCheck    created (0.36s)
     +   ├─ kubernetes:apps/v1:Deployment  zot-deployment                 created (27s)
     +   └─ kubernetes:core/v1:Service     zot-service                    created (10s)

    Resources:
        + 7 created

    Duration: 1m29s

    real    1m38.564s
    user    0m1.734s
    sys     0m1.302s
    ```

### 📋 Prerequisites

- [Go](https://golang.org/dl/) - v1.16+
- [Pulumi](https://www.pulumi.com/docs/get-started/install/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

### 🖥️ Installation

1. **Clone the Repository**

    ```bash
    git clone https://github.com/usrbinkat/pulumi-kind-go.git
    cd uor-zot/iac
    ```

2. **Download Go Deps**

    ```bash
    go get -u ./...
    go mod download
    go mod tidy
    ```

3. **Initialize Pulumi**

    ```bash
    pulumi stack init dev
    ```

4. **Run Pulumi**

    ```bash
    pulumi up
    ```

5. **Check Your Cluster**

    ```bash
    kubectl get nodes
    ```

## 🛠️ How It Works

### 📚 Components

1. **Pulumi**: Orchestrates the provisioning.
2. **KinD**: Creates a local Kubernetes cluster using Docker containers.
3. **Go**: The IaC logic is written in Go, leveraging Pulumi's Go SDK.
4. **Zot**: The Zot service is deployed to the cluster.

### 🔨 Code Structure

- `kind/kind.go`: KinD cluster creation.
- `kind/volumes.go`: Accelerate kind cluster builds with stateful containerd storage volume.
- `helper/helper.go`: Utility functions.
- `zot/zot.go`: Zot service deployment.

### 📈 Workflow

1. **Initialize Configuration**: Sets up Pulumi context and configuration.
2. **Load Cluster Config**: Reads the desired Kind cluster setup.
3. **Manage Docker Volumes**: Checks, creates, or deletes Docker volumes as needed.
4. **Create/Delete KinD Cluster**: Uses Pulumi's `local.Command` to run the KinD CLI commands.
5. **Create/Delete Zot Service**: Deploy Zot to k8s with Pulumi's `kubernetes.apps.v1.{Service,Deployment}` resources.

## 📜 Customizing the Cluster

For now, the code is set up to deploy a basic zot configuration. Customizations will be added in future releases.

## 🤝 Contributing

Please read [CONTRIBUTING.md](../CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## 📄 License

This project is licensed under the Apache 2.0 License - see the [LICENSE.md](../LICENSE) file for details.

## 🙏 Acknowledgments

- The Pulumi Team for their fantastic IaC tool.
- Kubernetes for revolutionizing the way we think about container orchestration.
- [The ReadME Project](https://www.readme.com/the-readme-project) for README best practices.

-------------------------------------------------------

## Local Development

```bash
export GOFLAGS='-replace=github.com/usrbinkat/pulumi-kind-go=.'
```
