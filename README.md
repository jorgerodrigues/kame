# UpKame - the simple monitoring tool.

UpKame is the simplest uptime monitoring tool around. It does all that you need and nothing that you dont'.  
This repo hosts the backend api server..

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests

```bash
make all build
```

build the application

```bash
make build
```

run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB container

```bash
make docker-down
```

live reload the application

```bash
make watch
```

run the test suite

```bash
make test
```

clean up binary from the last build

```bash
make clean
```

## Database migrations

### how to generate a new migration

`migrate create -dir ./migrations -seq -ext .sql name-of-the-migration`

### How to apply a migration

`migrate -path=./migrations -database="postgres://kame:password1234@localhost:5432/upkame?sslmode=disable" up`
