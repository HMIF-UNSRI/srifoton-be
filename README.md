# Sriwijaya Exhibition Backend

A web-based competition registration platform at the Sriwijaya Exhibition with a golang backend.

## How to Run

1. Install the latest stable version of [Go](https://go.dev/dl/)
2. Install the latest stable version of [PostgresSQL](https://www.postgresql.org/download/)
3. Create database and table with [the following scheme](config/db/schema.sql)
4. Rename and configure [`.env.example`](.env.example) to `.env`
5. If you want to enable production-mode, add `export GIN_MODE=release`
6. Start server with the following command :

```bash
$ go run cmd/server/main.go
```

## Feature Documentation

This is a documentation of how each feature works using UML with PlantUML. You can use the PlantUML extension on VSCode,
Jetbrains, or the official website.

### Account Register

![Login](_docs/exported_uml/account_registration.png)

### Account Login

![Login](_docs/exported_uml/account_login-0.png)

### Account Forgot Password

![Login](_docs/exported_uml/account_forgot_password-0.png)

### Competition Register

![Login](_docs/exported_uml/competition_registration-0.png)
