package countries

import (
	"fmt"
	"net/http"
	// "reflect" // reflect package lists the keys (field names) of a struct
	"strconv"

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

func UpdateCountryAndCode(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var id string = c.Param("countryId")

		var countryName string = c.PostForm("countryName")
		var countryCode string = c.PostForm("countryCode")

		type Country struct{
			Country, CountryCode string
		}

		country:= Country{Country: countryName, CountryCode:countryCode}

		// query:= db.Table("countries").Where("id=?", id).Update('country', countryName) // is only one clumn has to be updated

		update:= db.Table("countries").Where("id=?", id).Updates(country)
		
		if update.Error !=nil {
			c.JSON(http.StatusOK, gin.H{ "error":update.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"rowsUpdated":update.RowsAffected,
		})

		// val := reflect.ValueOf(update).Elem()
		// typ := val.Type()
		// for i := 0; i < val.NumField(); i++ {
		// 	fmt.Printf("\nKey: %s" ,typ.Field(i).Name)
		// }

	}
}


func DeleteCountry(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context){

		var countryId string= c.Param("countryId")

		// deleted:= db.Delete(&models.Countries, countryId)

		countryDeletionInfo := db.Table("countries").Where("id = ?", countryId).Delete(nil)
		// The .Delete(nil) method is used to delete the rows that match the condition. In this case, rows from the countries table where id = ? will be deleted.
		// Passing nil as an argument here tells GORM that you're deleting a record without needing to specify a specific model or struct.
		fmt.Println(countryDeletionInfo)

		if countryDeletionInfo.Error!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error": countryDeletionInfo.Error.Error(),
			});
			return
		}

		// val:= reflect.ValueOf(countryDeletionInfo).Elem()
		// typ:=val.Type()
		
		// for i:=0;i< val.NumField(); i++{
		// 	fmt.Println("\nKey:%s", typ.Field(i).Name)
		// }

		c.JSON(http.StatusOK, gin.H{
			"succes": true,
			"response": countryDeletionInfo.RowsAffected,
		});

	}
}


func DeleteWithStruct(db * gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var countryId string= c.Param("countryId")

		id, err := strconv.Atoi(countryId)
		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":"Invalid Country Id",
			})
			return
		}

		countryDeleteMetaData:= db.Delete(&models.Countries{}, id)
		// ALTERNATE: result := db.Where("id = ?", id).Delete(&models.Country{})

		// Use &models.Country{} to represent the model instance, and pass the ID for direct deletion.
		// Q1 :What is a model instance why when and how to use it?



		// The issue you're encountering arises because models.Countries refers to the type (the struct definition), not an instance (the data). In Gorm, when you use Delete, you need to pass a pointer to an instance of the model (like &country), or you can directly delete by primary key without an instance. Let's clarify this with an example.

		// Pass a pointer to an empty struct &models.Country{}.
		// Use the ID directly for deletion, which is the most straightforward for this case.


		if countryDeleteMetaData.Error!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":false,
				"error":countryDeleteMetaData.Error.Error(),
			})
			return
		}


		fmt.Println(countryDeleteMetaData)

		c.JSON(http.StatusOK, gin.H{
			"success":true,
			"rowsAffected":countryDeleteMetaData.RowsAffected,

		})
		
	}
}


func UpdateWithStruct(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){

		var countryId = c.Param("countryId")


		
		// 	// Create a model instance with the updated values
			var country models.Countries

			if err:=c.ShouldBindJSON(&country); err!=nil{
				c.JSON(http.StatusInternalServerError, gin.H{"error":err,})
				return
			}

			updateInfo:= db.Model(&models.Countries{}).Where("id=?", countryId).Updates(country)
			// only Updates the fields which are there {"country":"China"} only will uodat the country // {"countryCode":"+86"} only updates country code  {"countryCode":"+86", "country":"China"} updates both

			if updateInfo.Error!=nil{
				c.JSON(http.StatusInternalServerError, gin.H{"error":updateInfo.Error.Error(),})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"success":true,
				"rowsAffected":updateInfo.RowsAffected,
			})

	}
}