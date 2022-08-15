# GIN CRUD WEB

## About
GIN CRUD WEB written in Go (Golang) powered by popular Gin Web Framework, Go HTML template, and Gorm ORM library. The main features is about CRUD of <b>'Berat'</b> entity.

## Features

- CRUD Berat
- Go HTML Template Frontend CRUD Demo
- Config Environment
- ORM Support MySQL, PostgreSQL, SQL Server, SQLite.
- Auto Migrate
- Gracefully Shutdown
- Unit Test
- Dockerize
- Use Clean Architecture
![image](https://user-images.githubusercontent.com/67728406/184624044-e0666d16-9ff7-4224-8606-8ba8cf9993a2.png)

## Dependencies

- Gin: github.com/gin-gonic/gin,
- Gorm: github.com/gin-gonic/gin.
- Godotenv: github.com/joho/godotenv
- Testify: github.com/stretchr/testify
## How To Run

> go mod tidy <br>
> go run app/main.go

## How To Build Docker Image 

> docker build -t  gin-crud-web:1.0 . 

## Unit Test

Run unit test :

> go test -v -cover ./...

This command will run database_test.go & beratService_test.go

## Config

Just copy .env.example file to .env

- Application name, default: "My App"
  > APP_NAME="GO - REST- API"
- HTTP PORT, default: 8080
  > APP_PORT=3000
- Gin Mode, default: debug
  > GIN_MODE=release
- Debugging executed queries, default: false
  > DB_DEBUG=true
- Database name, default: mydata
  > DB_NAME=gorestdb


## Routes

Base URL :

> http://localhost:8080/

List of path :
| Method |URI |Purpose |
|----------------|-----------------------------|-----------------------------|
|GET |`'/'` |Show all collection of berat |
|GET |`'/show/id?={{ .ID }}'` |Show detail of berat by ID |
|GET |`'/new'` |Show create berat form |
|POST |`'/insert'` |Create new berat |
|GET |`'/edit'` |Show update berat form |
|POST |`'/update/id?={{ .ID }}'` |Update berat by ID |
|GET |`'/delete/id?={{ .ID }}'` |Delete berat by ID |

note: Golang HTML Template doesn't support PUT & DELETE metod

## Authors

- **Jagad Wijaya Purnomo**
