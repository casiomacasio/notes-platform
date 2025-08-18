# 📦 Notes Platform: Микросервисное приложение на Go

Проект реализует платформу для заметок и уведомлений с использованием микросервисной архитектуры.  
Каждый сервис написан на Go и взаимодействует через **RabbitMQ**. Инфраструктура поднимается с помощью **Docker Compose**.

---

## 🏗 Архитектура

### Gateway (API Gateway):

- Маршрутизация запросов
- Авторизация и валидация JWT
- Rate limiting (Redis)
- Reverse proxy к микросервисам

### Auth Service:

- JWT + Refresh токены
- Авторизация через HTTP-only cookie
- Валидация пользователей

### User Service:

- Управление профилем пользователя
- Получение информации о пользователях

### Note Service:

- Создание и редактирование заметок
- Хранение и управление заметками

### Notification Service:

- Отправка уведомлений
- Очереди сообщений через RabbitMQ
- Retry-механизм
- Логирование уведомлений в MongoDB

---

## ⚙️ Технологии

- **Backend:** Go (gin-gonic, sqlx)
- **База данных:** PostgreSQL
- **Кэш / Rate limiting:** Redis
- **Очереди сообщений:** RabbitMQ
- **Логирование уведомлений:** MongoDB
- **Контейнеризация:** Docker + docker-compose
- **CI/CD:** GitHub Actions
- **Мониторинг:** Prometheus + Grafana

---

## 🚀 Как запустить проект

### 1. Клонировать репозиторий

```bash
git clone https://github.com/casiomacasio/notes-platform.git
cd notes-platform
```

### 2. Создать `.env` файлы

Пример `.env` (для Auth Service):

```env
DB_PASSWORD=qwerty

SIGNING_KEY="supersecretkey"

```

### 3. Запустить сервисы

```bash
make up
```

### 4. Применить миграции

```bash
make migrate
```

### 5. Остановить сервисы

```bash
make down
```

---

## 📊 Мониторинг

Prometheus и Grafana доступны после запуска:

- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000
