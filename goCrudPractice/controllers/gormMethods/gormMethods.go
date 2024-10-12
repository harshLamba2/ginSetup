// Model(&models.Country{}): This specifies the table (countries) by referring to the model Country. The empty {} just initializes the struct so GORM knows its structure.

// Find(&groupedCountries): This collects the results from the database query and fills them into the provided slice. GORM expects to write the query results into this slice, so you just pass a reference to it without needing to instantiate the slice elements.

package gormMethods

import(
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"fmt"
	// "reflect"

	"crud/database/models"
)

func FirstAndLastMethod(db * gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){

		var data *models.Countries

		// if err:= db.First(&data).Error; err!=nil{ // fetch all columns
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"succes":false, "error":err.Error(),
		// 	})
		// return
		// }

		// if err:=db.Select("country_code").First(&data).Error; err!=nil{ //need a custom struct to fetch specific data else data will be like { "id": 0, "country": "", "countryCode": "+91" }
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"success":false,
		// 		"error":err.Error(),
		// 	})
		// return
		// }

		// if the result is 0 it comes as an error as no rows selected
		if err:= db.Table("countries").Where("country_code LIKE '+%'").First(&data).Error; err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"succes":false, "error":err.Error(),
			})
			return
		}

		// Last(&data) does the same thing except it ORDER the field in DESC order LIMIT 1 


		// val:= reflect.ValueOf(firstCountry).Elem()
		// typ:= val.Type()
		// for i:=0; i< val.NumField(); i++{
		// 	fmt.Println("\n Key: %s", typ.Field(i).Name)
		// }

		fmt.Println(data)
		c.JSON(http.StatusOK, gin.H{
			"success":true,
			"data": data,
		})


	}

}


func SaveMethod(){

	// GORM's Save() function operates specifically based on the primary key (ID) to determine whether to insert a new record or update an existing one
	
	// var user = models.User{
	// 	Name: "John",
	// 	Email: "john@example.com",
	// 	Age:  30,
	// }
	// db.Save(&user) // This will insert a new user record into the database




// var user models.User
// // First, find an existing user by ID
// db.First(&user, 1) // Finds the user with ID = 1

// // Update user fields
// user.Name = "Jane"
// user.Age = 32

// // Save the changes, this will perform an UPDATE in the database
// db.Save(&user)



// db.Model(&user).Update("name", "Jane Updated") //UPDATE "users" SET "name"='Jane Updated' WHERE "id"=1;
// db.Model(&user).Updates(models.User{Name: "Jane Updated", Age: 32}) || db.Model(&user).Updates(map[string]interface{}{"name": "Jane Updated", "age": 32}) //UPDATE "users" SET "name"='Jane Updated', "age"=32 WHERE "id"=1;



// EXPLICITLY DEFINING ID
// user := User{
// 	ID:   1, // Specify the ID in the struct to save the exact struct
// 	Name: "John Doe",
// }

// // Save the user (this will insert it if it doesn't exist, or update it if it does)
// db.Save(&user)

// Set the ID field in the struct before calling Save().
// If an entry with that ID already exists, Save() will update it; if not, it will create a new record with that ID.
// Ensure that the ID is unique to avoid conflicts.


}


func UpdateAndUpdates(db * gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){

		// Use Update when you want to modify ONE COLUMNS only.
		// Use Updates when you want to modify MULTIPLE COLUMNS.

		updateInfo:= db.Model(&models.Countries{}).Where("id=?", 1).Limit(1).Update("country", "Saudi Arabia")
		if updateInfo.Error!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":updateInfo.Error.Error()})
			return
		}

		updatesInfo:= db.Model(&models.Countries{}).Where("country_code LIKE '+%' ").Updates(map[string]interface{}{"colonizer":"United Kingdom", "government":"Kingship"})
		if updatesInfo.Error!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":updatesInfo.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"UpdateRowsAffected":updateInfo.RowsAffected,
			"UpdatesRowsAffected":updatesInfo.RowsAffected,
		})

	}
}


// Scan() is used to map the results of raw SQL queries or specific fields to a struct.
// This is useful when you want to select custom fields or join data from multiple tables and map it to a struct, but you don't need to load the full GORM model.
func RawAndScanMethod(db * gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){

		var selectedFields [] struct{
			CountryCode string `json:"countryCode"`
			Country string `json:"country`
		}

		// result:= db.Table("countries").Select("country, country_code").Where("country IS NOT NULL").Scan(&selectedFields)
		
		result:= db.Raw(`SELECT country, country_code FROM countries WHERE country LIKE 'N%'`).Scan(&selectedFields)

		if result.Error!=nil{
			c.JSON(http.StatusInternalServerError, map[string]interface{}{ "error": result.Error.Error() })
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"success":true,
			"result":selectedFields,
		})

	}
}


func OrderMethod(db * gorm.DB) gin.HandlerFunc{
	return func(c * gin.Context){

		var data []models.Countries

		// result:= db.Order("country_code, id").Find(&data)
		result:= db.Order("country_code desc, id").Find(&data)
		if result.Error!=nil{
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error":result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{"success":true, "data":data})
		return
	}
}

func GroupAndHavingMethod(db *gorm.DB) gin.HandlerFunc{
	return func (c * gin.Context){

		var groupedCountries []struct{
			Country string `json:"country"`
			Count int `json:"count"`
		}
		
		// db.Model(&models.Countries).Select("country, count").group("country").Find(&groupedCountries) // what i made
		// if err:=db.Model(&models.Countries{}).Select("country, count(*) as count").Group("country").Find(&groupedCountries).Error;err!=nil{
		if err:=db.Model(&models.Countries{}).Select("country, count(*) as count").Group("country").Having("count(*) > 2").Find(&groupedCountries).Error;err!=nil{
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error":err.Error()})
			return
		}
		
		c.JSON(http.StatusOK, map[string]interface{}{"success":true, "result":groupedCountries})
		return

	}
}


func LimitAndOffset(db * gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		
		var ids []struct{
			CountryCode string `json:"country_code"` // preffer int64
		}

		// if err:=db.Model(&models.Countries{}).Select("country_code").Order("country_code").Find(&ids).Error; err!=nil{
		// if err:=db.Model(&models.Countries{}).Select("country_code").Order("country_code").Limit(1).Find(&ids).Error; err!=nil{
		if err:=db.Model(&models.Countries{}).Select("country_code").Order("country_code").Limit(1).Offset(1).Find(&ids).Error; err!=nil{
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error":err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":true,
			"result":ids,
		})

	}
}



func CountMethod(db * gorm.DB)gin.HandlerFunc {
	return func(c *gin.Context){

		var countryCount int64

		if err:= db.Model(&models.Countries{}).Where("country LIKE ?", "I%").Count(&countryCount).Error; err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success":true, "countryCount":countryCount})

	}
}

func ExecMethod(db *gorm.DB) gin.HandlerFunc{ //Exec is used for raw queries for non Select Queries
	return func(c *gin.Context){

	insertResult:= db.Exec("INSERT INTO countries (country, country_code) VALUES (?, ?)", "Mayanmar", "+95")
	updateResult:= db.Exec("UPDATE countries SET government=? WHERE country_code=?", "Military Rule", "+95")

	if insertResult.Error!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":insertResult.Error.Error()})
	}

	if updateResult.Error!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":updateResult.Error.Error()})
	}


	c.JSON(http.StatusOK, gin.H{
		"success":true,
		"insertedRows":insertResult.RowsAffected,
		"updatedRows":updateResult.RowsAffected,
	})
	// db.Exec("DELETE FROM users WHERE country_code =?", "+95")
	}
}

func DistinctMethod(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){

		var distinctGovernments []struct{
			Government *string `json:"government"`
		}

		if err:= db.Model(&models.Countries{}).Distinct("government").Find(&distinctGovernments).Error; err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success":true, "distinctGovernments": distinctGovernments})

	}
}

func PluckMethods(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){

		var countryCodes []string // dosent have to similar to the model key

		// db.Model(&User{}).Joins("JOIN orders ON orders.user_id = users.id").Where("orders.status = ?", "completed").Pluck("users.username", &usernames)// use empty Distinct() to fetch distinct pluck values
		if err:= db.Model(&models.Countries{}).Pluck("country_code", &countryCodes).Where("country_code IS NOT NULL").Error; err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success":true, "countryCodes":countryCodes})
		
	}
}


// EXPLORE THIS METHOD MORE
// var user User
// db.Select("username", "email").First(&user)
// fmt.Println(user.Username, user.Email)

func UnscopedMethod(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		// USELESS DID NOT WOKR
		// lets you bypass default behaviour and includes soft-deleted records in your queries

		if err:= db.Model(&models.Countries{}).Unscoped().Where("country", "Mayanmar").Error; err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success":true})

	}
}



func FirstOrCreateMethod(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){

		
		var country models.Countries

		if err := c.BindJSON(&country); 
		err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var countryCode *string = country.CountryCode



	// // Search for the user by email
	// // If found, only update the Phone field
	// // If not found, insert the new record with all the fields
	// db.Where(models.Countries{CountryCode: countryCode}).Attrs(&country).FirstOrUpdate(&user, User{Phone: user.Phone})
	db.Where("country_code= ?", countryCode).FirstOrUpdate(&country)






		// result := db.Create(&country);

		// if result.Error != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		// 	return
		// }

		// fmt.Println(result, "result")

		c.JSON(200, gin.H{
			"success": true,
			"message":countryCode,
		})
		// db.Where(User{Name: "John"}).FirstOrUpdate(&user, User{Email: "john@example.com"})
		// looks for user with name jhon //if found Email updates to "john@example.com" else record insersion (name, email) VALUES('John', 'john@example.com')
	}
}