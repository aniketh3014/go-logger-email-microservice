# Go Microservices Project

![Go Microservices Logo](https://via.placeholder.com/200x100?text=Go+Microservices)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/aniketh3014/go-logger-email-microservice)](https://goreportcard.com/report/github.com/aniketh3014/go-logger-email-microservice)
[![Go Version](https://img.shields.io/github/go-mod/go-version/aniketh3014/go-logger-email-microservice)](https://golang.org/)

A robust, scalable microservices architecture built with Golang, showcasing modern development patterns and deployment options including Docker, Kubernetes, and Docker Swarm.

## 🚀 Overview

This project demonstrates a complete microservices ecosystem with several independent services that work together:

- **Broker Service**: API gateway that routes requests to appropriate microservices
- **Authentication Service**: Handles user authentication and password verification
- **Logger Service**: Centralized logging with MongoDB storage and multiple access protocols (HTTP, RPC, gRPC)
- **Mail Service**: Email handling service for sending notifications
- **Listener Service**: Event-driven service that listens for RabbitMQ messages
- **Front-End**: Simple web interface to interact with the microservices

## 📋 Features

- **Polyglot Persistence**: Uses PostgreSQL and MongoDB for different data needs
- **Multiple Communication Protocols**:
  - RESTful HTTP API
  - RPC
  - gRPC
  - Message Broker (RabbitMQ)
- **Deployment Options**:
  - Docker Compose for development
  - Kubernetes for container orchestration
  - Docker Swarm for cluster management
- **Observability**: Centralized logging system
- **Security**: Authentication and secure service-to-service communication
- **Scalability**: Each service can be scaled independently

## 🛠️ Technology Stack

- **Backend**: Go (Golang 1.24+)
- **Web Framework**: Chi Router
- **Databases**:
  - PostgreSQL for user data
  - MongoDB for logs
- **Message Broker**: RabbitMQ
- **Containerization**: Docker
- **Orchestration**: Kubernetes, Docker Swarm
- **Reverse Proxy**: Caddy
- **Mail Testing**: MailHog

## 🔧 Getting Started

### Prerequisites

- Go 1.24 or higher
- Docker and Docker Compose
- kubectl (for Kubernetes deployment)
- Make

### Running with Docker Compose

```bash
# Clone the repository
git clone https://github.com/yourusername/go-micro.git
cd go-micro/proj

# Start all services
make up_build

# To stop all services
make down
```

### Running with Kubernetes

```bash
# Start Minikube or connect to your K8s cluster
minikube start

# Deploy services
kubectl apply -f k8s/

# Set up ingress (if needed)
kubectl apply -f ingress.yml
```

### Running with Docker Swarm

```bash
# Initialize swarm
docker swarm init

# Deploy stack
docker stack deploy -c swarm.yml microservices
```

## 📁 Project Structure

```
go-micro/
├── authentication-service/  # User authentication
├── broker-service/          # API Gateway
├── front-end/               # Web interface
├── listener-service/        # RabbitMQ event consumer
├── logger-service/          # Logging with MongoDB
├── mail-service/            # Email handling
└── proj/                    # Deployment configurations
    ├── k8s/                 # Kubernetes configurations
    ├── docker-compose.yml   # Docker Compose configuration
    └── swarm.yml            # Docker Swarm configuration
```

## Resources

- [Chi Router](https://github.com/go-chi/chi)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)
- [PostgreSQL Driver](https://github.com/jackc/pgx)
- [RabbitMQ Client](https://github.com/rabbitmq/amqp091-go)
