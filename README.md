# 🚀 Go IAM Service

Базовый микросервис Identity & Access Management на Go с REST API, JWT и RBAC.

---

## 🔍 О проекте
Этот сервис управляет пользователями, ролями и доступом:

- Регистрация и вход пользователей с bcrypt
- JWT access + refresh токены
- Контроль доступа на основе ролей (`admin`, `user`)
- CRUD операций по пользователям и ролям
- Автодокументация через Swagger/OpenAPI
- Контейнеризация (Docker) и базовая CI-конфигурация (GitHub Actions)

---
## 📦 Технологии

- Go + фреймворк Gin (или Echo)
- PostgreSQL + GORM
- bcrypt (`golang.org/x/crypto/bcrypt`)
- JWT (`github.com/golang-jwt/jwt`)
- Swagger-аннотации
- Тестирование (httptest + testify)
- Docker, GitHub Actions
---
## 🗂 Структура проекта
```
cmd/iam/main.go
internal/
├── models/ # GORM-модели: User, Role, UserRole
├── storage/ # CRUD по пользователям и ролям
├── auth/ # JWT логика (создание/проверка/refresh)
├── handlers/ # HTTP-эндпоинты: signup, login, users
├── middleware/ # JWT + RBAC middleware
└── log/ # Инициализация логгера
docs/
└── swagger.yaml # Сгенерированная OpenAPI документация
Dockerfile
.github/workflows/ci.yml
```
---
## 🌐 API и RBAC

- `POST /signup` – регистрация
- `POST /login` – вход
- `POST /refresh` – обновление access токена
- `GET /users` – список всех пользователей (только `admin`)
- `GET /users/:id` – просмотр пользователя (`user` может только своего)
- `PATCH /users/:id` – редактирование (только `admin`)
- `DELETE /users/:id` – удаление (только `admin`)

RBAC: `admin` — полный доступ, `user` — доступ только к своему профилю.
---
## 🛠 Дальнейшее развитие
Добавить роли: editor, viewer

Поддержка MFA и SSO/LDAP

Мониторинг/метрики (Prometheus, Grafana)

Более продвинутая политика RBAC/ABAC

Градация логирования, метрик и трассировки

