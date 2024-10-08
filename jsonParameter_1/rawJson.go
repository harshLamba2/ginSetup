package main;

import (
	"github.com/gin-gonic/gin";
	"fmt"
)

type submitJSONBody struct {// You need to capitalize the field names of the struct (Name and Age) for Gin to access and bind the JSON data.
	Name string `json:"name"`
	Age int `json:"age"`
}
func main(){
	router:=gin.Default();

	// type submitJSONBody struct {
	// 	name string `json:"name"`
	// 	age int `json:"age"`
	// } // The reason why your code isn't printing the values is due to Go's struct field visibility rules. In Go, struct fields need to be exported (start with an uppercase letter) to be accessible outside of the package, including being bound by Gin when processing JSON

	router.POST("/submit", func(c *gin.Context){

		var reqBody submitJSONBody;
    
		// Bind the incoming JSON to reqBody
		if err := c.BindJSON(&reqBody); 
		err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	
		fmt.Println(reqBody.Name, reqBody.Age, "HELLO WORLD")

		c.JSON(200, gin.H{
			"message": "API has been hit",
		})

	})


	router.Run(":8081");
}