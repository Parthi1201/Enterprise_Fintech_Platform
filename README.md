# Enterprise Fintech Platform (Microservices Architecture)

This project is a distributed backend system built using Go, designed to simulate a real-world fintech platform. It follows a microservices architecture with multiple independently deployable services handling core financial operations.

## Key Features

- Microservices architecture with services for account, customer, payment, transaction, card, and balance management
- Clean architecture design (handler → service → data → mapper) for scalability and maintainability
- gRPC (Protocol Buffers) for efficient inter-service communication
- REST APIs (OpenAPI) for external client interaction
- PostgreSQL for persistent storage across services
- Snowflake-based distributed ID generation
- Dockerized services for independent deployment and scalability

## Architecture Overview

Each service is structured with clear separation of concerns:
- **Handler** – Handles incoming requests (HTTP/gRPC)
- **Service** – Business logic layer
- **Data** – Database interactions and repositories
- **Mapper** – Data transformation between layers

## Tech Stack

- Language: Go
- Communication: gRPC, REST
- Database: PostgreSQL
- DevOps: Docker
- Architecture: Microservices, Clean Architecture

## Goal

The goal of this project is to demonstrate backend system design concepts such as service decomposition, inter-service communication, and scalable architecture in a fintech context.
