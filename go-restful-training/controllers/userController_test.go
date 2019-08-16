package controllers



import (
	"encoding/json"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo"
	"strings"

)
// // test addition this is for testing addition
// func TestAddition (t *testing.T)  {
// 	// if A + B = C 
// 	// expect 1 + 2 equals 3
// 	testResult := Addition(1, 2)

// 	assert.Equal(t, 3, testResult)
// }


func TestGetAllReseller (t *testing.T) {
	data := `{"name" : "iqbal" , "email" : "iqbal@alterra.id", "password" : "112233"}`

	e := echo.New()
	// success case
	// setup
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext (req, rec)
	c.SetPath("/api/v1/reseller/all")
	h := CreateUserController(c)

	// Assertions
	if assert.NoError (t, h){
		assert.Equal(t, 201, rec.Code)

		var jsonBody interface{}

		json.NewDecoder(rec.Body).Decode(&jsonBody)

		jsonMap := jsonBody.(map[string]interface{})

		assert.Equal(t, "iqbal", jsonMap["name"].(string))
		assert.Equal(t, "iqbal@alterra.id", jsonMap["email"].(string))
		assert.Equal(t, "112233", jsonMap["password"].(string))
	}
}

// func TestCreateUsersController (t *testing.T) {

// 	e := echo.New()
// 	// success case
// 	// setup
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext (req, rec)
// 	c.SetPath("/api/v1/reseller/all")
// 	h := GetUsersController(c)

// 	// assertions
// 	if assert.NoError (t, h){
// 		assert.Equal(t, 200, rec.Code)
// 	}
// }