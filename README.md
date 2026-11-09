# Marketplace

Marketplace is a microservice-based e-commerce platform developed as a demo of a modern online shop with separated backend services, a storefront, and infrastructure tools for real-world deployment.

## Overview

The project combines:

- user authentication and registration
- catalog of products and categories
- shopping cart logic
- order processing
- analytics and event-driven infrastructure
- frontend storefront for a real marketplace experience

The goal is to create a scalable platform that resembles a production-grade marketplace architecture, while keeping the codebase understandable and easy to run locally.

## Technology stack

### Backend

- Go
- Chi router
- GORM
- Python
- FastAPI
- SQLAlchemy

### Data and infrastructure

- PostgreSQL
- Redis
- Kafka
- ClickHouse
- MinIO
- Docker Compose

### Frontend

- React
- Vite
- JavaScript

## Project structure

```text
marketplace/
├── auth/                  # authentication service (Go)
├── cart/                  # cart service (Go)
├── catalog/               # catalog service (Python/FastAPI)
├── orders/                # orders service (Go)
├── analytics/             # analytics-related service
├── payment/               # payment service prototype
├── frontend/              # React storefront
├── docker/                # database initialization scripts
├── docker-compose.yml     # infrastructure services
├── run_catalog.sh         # catalog startup helper
├── run_orders.sh          # orders startup helper
├── README.md              # project documentation
├── .gitignore             # excludes generated files
└── sql/                   # SQL resources
```

## Services and ports

| Service | Description | Local port |
| --- | --- | --- |
| Frontend | marketplace storefront | 5173 |
| Auth API | registration/login | 8000 |
| Catalog API | products and categories | 8001 |
| PostgreSQL | main database | 5434 |
| Redis | cache / data store | 6379 |
| MinIO | object storage | 9000 |
| Kafka | event streaming | 9092 |
| ClickHouse | analytics | 8123 |

## Starting the infrastructure

From the project root:

```bash
docker compose up -d
```

This starts the database, cache, object storage, message broker and analytics services required by the marketplace.

## Starting the backend services

### Auth service

```bash
cd auth
go run .
```

### Catalog service

```bash
cd catalog
source .venv/bin/activate
uvicorn main:app --host 0.0.0.0 --port 8001
```

Or use the helper script:

```bash
./run_catalog.sh
```

### Orders service

```bash
./run_orders.sh
```

## Starting the frontend

```bash
cd frontend
npm install
npm run dev -- --host 0.0.0.0
```

Then open:

```text
http://localhost:5173/
```

## Environment configuration

Some Go services expect environment variables such as:

```text
DB_HOST
DB_PORT
DB_USER
DB_PASSWORD
DB_NAME
JWT_SECRET
```

Set these in the appropriate environment or `.env` files before starting the backend services.

## Current status

This project is in active development and already includes the core structure for a real marketplace:

- authentication service
- catalog API
- cart and order flows
- frontend storefront
- infrastructure stack for local development

The architecture is being extended gradually with additional microservices and business logic.

## Useful commands

```bash
# start infrastructure
docker compose up -d

# start auth
cd auth && go run .

# start catalog
./run_catalog.sh

# start orders
./run_orders.sh

# start frontend
cd frontend && npm run dev -- --host 0.0.0.0
```

## License

This project is intended for learning and development purposes within the marketplace architecture demo.
