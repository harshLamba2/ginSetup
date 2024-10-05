SETUP 
1. go mod init MODULE_NAME

2. go get -u gorm.io/gorm
3. go get -u gorm.io/driver/mysql
4. go get -u github.com/gin-gonic/gin


FILE STRUCTURE
your_project/
├── main.go
└── models/
    └── user.go