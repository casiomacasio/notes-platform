# 📦 Notes Platform: Go Microservices Application

This project implements a platform for notes and notifications using a **microservices architecture**.
Each service is written in Go and communicates via **RabbitMQ**. The infrastructure is deployed using **Docker Compose**.

---

## 🏗 Architecture

### Gateway (API Gateway):

- Request routing
- JWT authentication and validation
- Rate limiting (Redis)
- Reverse proxy to microservices

### Auth Service:

- JWT + Refresh tokens
- Authentication via HTTP-only cookies
- User validation

### User Service:

- User profile management
- Fetching user information

### Note Service:

- Creating and editing notes
- Storage and management of notes

### Notification Service:

- Sending notifications
- Message queues via RabbitMQ
- Retry mechanism
- Logging notifications in MongoDB

---

## ⚙️ Technologies

- **Backend:** Go (gin-gonic, sqlx)
- **Database:** PostgreSQL
- **Rate limiting:** Redis
- **Message queues:** RabbitMQ
- **Notification logging:** MongoDB
- **Containerization:** Docker + Docker Compose
- **CI/CD:** GitHub Actions
- **Monitoring:** Prometheus + Grafana

---

## 🚀 How to Run the Project

### 1. Clone the repository

```bash
git clone https://github.com/casiomacasio/notes-platform.git
cd notes-platform
```

### 2. Create `.env` files

Example `.env` (for Auth Service):

```env
DB_PASSWORD=qwerty

SIGNING_KEY="supersecretkey"

REDIS_PASSWORD=redis
```

### 3. Start the services

```bash
make up
```

### 4. Apply migrations

```bash
make migrate
```

### 5. Stop the services

```bash
make down
```

---

## 📊 Monitoring

After starting the services, Prometheus and Grafana are available at:

- Prometheus: [http://localhost:9090](http://localhost:9090)
- Grafana: [http://localhost:3000](http://localhost:3000)
