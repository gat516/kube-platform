# kube-platform

A personal Kubernetes platform built to demonstrate production-style DevOps practices — automated deployments, Helm packaging, and live observability through Prometheus and Grafana.

## Stack

- Kubernetes (k3s local, EKS production)
- Terraform
- GitHub Actions
- Prometheus + Grafana
- Postgres (RDS)
- JavaScript (Netlify/Vercel)
- AWS

## Architecture

Frontend is hosted on Netlify/Vercel. Backend services and infrastructure run on AWS via EKS, with GitHub Actions handling CI/CD and Helm managing deployments. Prometheus and Grafana provide observability.

## Getting Started

### Prerequisites

- Docker
- kubectl
- Helm
- k3s (for local development)

### Local Setup

coming soon

### Deploying to AWS

coming soon

## Status

Work in progress.

## Roadmap

- [ ] Local k3s cluster setup
- [ ] Demo service with `/health` and `/metrics` endpoints
- [ ] Dockerfile + GHCR image publishing
- [ ] Helm chart
- [ ] GitHub Actions pipeline
- [ ] Prometheus + Grafana stack
- [ ] Live public dashboard
