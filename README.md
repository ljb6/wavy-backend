# Wavy Backend

Backend in **Go** for the **Wavy** newsletter app, using **PostgreSQL** and **Docker**.

## Requirements

- Go 1.22+
- Docker + Docker Compose

## Setup

### 1. Clone the project

```bash
git clone https://github.com/your-username/wavy-backend.git
cd wavy-backend
```

## 2. Setup the database

In database/database.go and docker-compose.yml, configure the database connection.

## 3. Environment Variables

Create a .env file:

```bash
JWT_SECRET=
```

## 4. Run the backend

```bash
docker compose up -d  
```

```bash
go run main.go
```