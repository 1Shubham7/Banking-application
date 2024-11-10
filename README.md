# Banking application

A project made using Postgres, Go, 

Database Schema- https://dbdiagram.io/d/665d6000b65d93387953a9f5

# Banking Application

This project is a banking application built using Go and PostgreSQL, designed to showcase essential features of a modern backend system. I have used Go as the backend language, PostgreSQL as the database, SQLC for generating type-safe SQL queries, and `go-playground/validator` for input validation. The project also includes comprehensive unit tests written with `testify/assert` assertion library and `GoMock` for mocking dependencies if needed. The project also includes CI tests written using GitHub workflows and code coverage is also tracked and uploaded to Codecov ([find it here](https://app.codecov.io/gh/1shubham7/banking-application)). Build automation and commands are managed using Makefile.

### Tech Stack

- **Backend Language:** Go (Golang)
- **Database:** PostgreSQL
- **ORM/SQL Generator:** [SQLC](https://sqlc.dev/)
- **Validation:** go-playground/validator
- **Testing:** Comprehensive unit tests with GoMock for mocking
- **Code Coverage:** Coverage reports uploaded to Codecov. ([find it here](https://app.codecov.io/gh/1shubham7/banking-application))
- **Build and Commands:** Managed using Makefile

## How to use

The only prerequisit to this application are Docker and `make`, you can simply install Docker and `make` and then run the following command to start the application:

`make postgres`

`make createdb`

```
make start_app
```

## Steps to reproduce:

### Step 1. Clone the repository:

```
git clone https://github.com/1Shubham7/Banking-application.git
```

### Step 2. Start Postgres DB using the Postgres container:

```
docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:12-alpine
```

### Step 3. Install `go-migrate` with root previliges:

```
sudo go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1
```

### Step 4. Now you can simply run the application by entering the following make command:

```
make server
```

### Hurray! your appilcation is running on PORT 8080

https://codecov.io/gh/1Shubham7/Banking-application/graphs/sunburst.svg?token=X5WO4RO683

[![codecov](https://codecov.io/gh/1Shubham7/Banking-application/graph/badge.svg?token=X5WO4RO683)](https://codecov.io/gh/1Shubham7/Banking-application)

https://dbdocs.io/1Shubham7/Smyik