package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"crud/controllers/gormMethods"
)

// func GormMethod(router * gin.Engine, db * gorm.DB){
func GormMethod(router * gin.RouterGroup, db * gorm.DB){ 
	router.GET("/firstMethod", gormMethods.FirstMethod(db))
}

