package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"crud/database/models"
	"gorm.io/gorm"
)

func AddCounty(db * gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){

		var country models.Countries

		if err := c.BindJSON(&country); 
		err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}


		result := db.Create(&country);

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		fmt.Println(result, "result")

		c.JSON(200, gin.H{
			"success": true,
			"message":"Country Added Successfully",
		})
	}
}

func GetAllCountryCode(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context){

		var allCountries []struct{
			Country string `json:"country"`
			CountryCode string `json:"countryCode"`
		} 

		// if err:=db.Find(&allCountries).Error; err!=nil{
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()});
		// 	return
		// }

		if err := db.Table("countries").Select("country, country_code").Find(&allCountries).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"success":true,
			"allCountries":allCountries,
		});
		
	}
}

func CountryCodeById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context){

		var id string = c.Param("countryId");

		var countryCode struct {
			CountryCode *string `json:"countryCode"` // Pointer to a string to allow null // A string holds the actual string value, while a *string (pointer to a string) holds the memory address where the string is stored. Pointers add flexibility by allowing for nil values (representing "no value"), more efficient data handling, and enabling mutation across function calls.
			
			// CountryCode string `json:"countryCode"`
		}

		query := db.Table("countries")
		query = query.Select("country_code")
		query = query.Where("id = ?", id) 

		if err:= query.Find(&countryCode).Error; err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":true,
			"countryCode":countryCode,
		})
	}
}