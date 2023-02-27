package handler_test

import (
	// "encoding/json"
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"

	// "net/http/httptest"

	// "net/http/httptest"
	"testing"

	"app/config"
	"app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.LoadEnvVars()
}


func TestTEST(t *testing.T) {
	var db = database.MongoDBConnection()
	var h = database.DatabaseInit(db)


	app := fiber.New()
	r:= app.Group("/api")
	r.Get("/test", h.TEST)

	expectedData := struct {
		message string `json:"message"`
	} {
		message: "Test Hello World",
	}
	req, _ := http.NewRequest("GET", "/api/test", nil)
	// w:= httptest.NewRecorder()
	

	responseData, _ := app.Test(req, 1)
	res ,_ := ioutil.ReadAll(responseData.Body)
	resData := expectedData
	json.Unmarshal(res, resData)
	assert.Equal(t, http.StatusOK, responseData.StatusCode)
	assert.Equal(t, expectedData, resData)
}
