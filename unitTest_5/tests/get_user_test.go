// run comand "go test "folder/path" // if the folder path is dame then "go test ." and in terms of this file go like "go test ./tests"
// make sure all the test files has _test.go in the end 

package tests;

import (
	"tests/controller"
	"testing"
	"net/http"
	"net/http/httptest"
	
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
);


// The function TestGetUser(t *testing.T) is a test function, and in Go, test functions are automatically discovered and executed by the Go testing framework when you run tests with the go test command.
// Test functions in Go are typically named TestXYZ
func TestGetUser(t *testing.T){

	gin.SetMode(gin.TestMode);
	router:= gin.Default();

	router.GET("/getUser", controller.GetUser);

	req, _ := http.NewRequest(http.MethodGet, "/getUser", nil) //the nil is used to specify the body of the HTTP request. //body: This is an io.Reader that represents the body of the request, which can be nil if no body is needed (such as in GET requests). 
	// SIMULATING POST REQUEST
	// jsonData := `{"name": "John Doe", "email": "john@example.com"}`
	// req, _ := http.NewRequest(http.MethodPost, "/user", strings.NewReader(jsonData))

	w:= httptest.NewRecorder(); //w is commonly used as a shorthand name for the response writer in HTTP testing scenarios. the variable which will hold the response from the request

	router.ServeHTTP(w, req);
	assert.Equal(t, http.StatusOK, w.Code); // assert.Equal is an assertion used in unit tests to verify that the actual value (w.Code) matches the expected value (http.StatusOK). 

	expected:= `{
	"success": true,
	"user": {
			"name": "Harsh",
			"age": 24
  			}
	}`

	// check the response body
	assert.JSONEq(t, expected, w.Body.String(), "Response Body Mismatch");

}

