package routes

import (
	"github.com/gin-gonic/gin"
	"crud/controllers"
	"gorm.io/gorm"
)

func CountriesRoute(router *gin.Engine, db *gorm.DB){

	router.POST("/add_country", controllers.AddCounty(db));
	router.GET("/get_all_countries", controllers.GetAllCountryCode(db));
	router.GET("/get_country_code/:countryId", controllers.CountryCodeById(db));
	router.POST("/update_country_and_code/:countryId", controllers.UpdateCountryAndCode(db));

}