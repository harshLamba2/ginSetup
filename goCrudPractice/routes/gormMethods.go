package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"crud/controllers/gormMethods"
)

// func GormMethod(router * gin.Engine, db * gorm.DB){
func GormMethod(router * gin.RouterGroup, db * gorm.DB){ 
	router.GET("/first-and-last_methods", gormMethods.FirstAndLastMethod(db))
	router.GET("/update-and-updates", gormMethods.UpdateAndUpdates(db))
	router.GET("/raw-and-scan-methods", gormMethods.RawAndScanMethod(db))
	router.GET("/order-method", gormMethods.OrderMethod(db))
	router.GET("/group-and-having-methods", gormMethods.GroupAndHavingMethod(db))
	router.GET("/limit-and-offset-methods", gormMethods.LimitAndOffset(db))
	router.GET("/count-method", gormMethods.CountMethod(db))
	router.GET("/exec-method", gormMethods.ExecMethod(db))
	router.GET("/distinct-method", gormMethods.DistinctMethod(db))
	router.GET("/pluck-method", gormMethods.PluckMethods(db))
	router.GET("/unscoped-method", gormMethods.UnscopedMethod(db))

	router.POST("/first-or-create-method", gormMethods.FirstOrCreateMethod(db))
	
}

