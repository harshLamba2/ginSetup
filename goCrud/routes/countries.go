package routes

import (
	"github.com/gin-gonic/gin"
	"crud/controllers"
)

func CountriesRoute(router *gin.Engine){

	router.POST("/add_country", controllers.AddCounty)

}