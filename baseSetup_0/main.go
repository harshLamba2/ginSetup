package main;

import (
	"fmt"
    "github.com/gin-gonic/gin"
)


func main() {

    router := gin.Default() //creates a new Gin router instance with default middleware (like logging and recovery). This router will handle incoming HTTP requests. // const app = express()

    // Define a simple route
    router.GET("/", func(c *gin.Context) { //The function takes a gin.Context object (c *gin.Context), which is used to handle request and response operations.
        c.JSON(200, gin.H{ // c.JSON sends a JSON response back to the client.  //gin.H{} is a shortcut provided by the Gin framework to create a map in Go, where the keys are strings and the values are of any type (like an object in JavaScript).
            "message": "Hello, Gin!",
        })
    })

    router.POST("/submit", func(c *gin.Context){

        name:= c.PostForm("name"); //When you're sending raw JSON data in a POST request, the function c.PostForm("name") won't work because that method is specifically for extracting form-encoded data (like application/x-www-form-urlencoded).
        fmt.Println(name);

        c.JSON(200, gin.H{
            "message":"Successfully Recived",
            "name":name,
        })

    });

    // PARAMS
    router.GET("/user/:id", func(c *gin.Context) {
        id := c.Param("id")
        c.JSON(200, gin.H{"user_id": id})
    })

    // QUERY PARAMS
    router.GET("/filter", func(c *gin.Context){
        name:=c.Query("name");
        fmt.Println(name);

        c.JSON(200, gin.H{
            "QueryParameterWas":name,
        })

    })

    // Start the server on port 8080
    router.Run(":8080") // similar to app.listen(8080)
}

