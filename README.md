# Pulumi Zephyr
This project aims to closely follow the Pulumi blog series "[IaC Recommended Practices](https://www.pulumi.com/blog/iac-recommended-practices-code-organization-and-stacks/)".
In which the fictional, Zephyr Archaeotech Emporium, uses Pulumi to deploy and manage resources
for their online retail store.

The implementation in this repository will be done in [Go](https://golang.org/), with package/tool
management provided by [Nix](https://nixos.org/).

## Zephyr Archaeotech Emporium Online Store

This is the source repository for the online store of the Zephyr Archaeotech Emporium. It's used in Pulumi's Zephyr series of blog posts to discuss best practices when using Pulumi to manage infrastructure and applications.

This application is based on [this source repository](https://github.com/aws-containers/retail-store-sample-app); the original `README.md` for the source is now found in the `/docs` folder.

### Deploying with Pulumi

Information on using Pulumi to deploy this application and its associated infrastructure can be found in the `infra` folder.

### Deploying to Docker Compose

This application can be deployed locally via Docker Compose; use the `docker-compose.yml` file in the `deploy/docker-compose` folder.

### Deploying to Kubernetes

This application can be deployed to Kubernetes directly; use the `deploy.yaml` file in the `deploy/kubernetes` folder.
