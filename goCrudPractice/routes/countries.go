package routes

import (
	"github.com/gin-gonic/gin"
	"crud/controllers/countries"
	"gorm.io/gorm"
)

func CountriesRoute(router *gin.Engine, db *gorm.DB){

	// CRUD
	router.POST("/add_country", countries.AddCounty(db));

	router.GET("/get_all_countries", countries.GetAllCountryCode(db));
	router.GET("/get_country_code/:countryId", countries.CountryCodeById(db));

	router.POST("/update_country_and_code/:countryId", countries.UpdateCountryAndCode(db));
	router.POST("/update_country_with_struct/:countryId", countries.UpdateWithStruct(db));
	
	router.DELETE("/delete_country/:countryId", countries.DeleteCountry(db));
	router.DELETE("/delete_country_with_struct/:countryId", countries.DeleteWithStruct(db));

	// OTHER GORM METHODS
	


}