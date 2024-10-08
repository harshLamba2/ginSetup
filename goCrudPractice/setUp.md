SETUP 
1. go mod init MODULE_NAME

2. go get -u gorm.io/gorm
3. go get -u gorm.io/driver/mysql
4. go get -u github.com/gin-gonic/gin


FILE STRUCTURE
C:.
│   go.mod
│   go.sum
│   main.go
│   README.md
│
├───config // includes configurations which might change accoring to development and production variables
|    ├── config.dev.env
|    ├── config.prod.env
|    └── config.go
│
controllers //The controllers folder should only handle HTTP requests and delegate business logic to services or other layers.
|    ├── countries/
|    │   └── countries_controller.go
|    └ users/
|       └── users_controller.go
│
├───database
│   ├───models
│   │       countries.go
│   └───connection.go
│
├───middlewares
│   └───auth.go
│   └───logging.go
│
├───routes
│       countries.go
│
├───services //IMPORTANT: The services layer handles business logic. This layer is responsible for orchestrating operations that involve different components, like calling repositories (database operations), making API calls, or performing complex data transformations.
│   ├── countries_service.go
|   └── users_service.go
│
└───test
        countries_test.go
