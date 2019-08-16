package main
import (
  "net/http"
  "github.com/labstack/echo"
)

type User struct {
	Id    int
	Name  string
	Email string
  }
  
  // User Controller
  func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{Id: id, Name: "Ismail", Email: "ismail@alterra.id"}
	// Render Data - JSON Response
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "user": user,
	})
  }
  
  func main() {
	e := echo.New()
	// routing
	e.GET("/users/:id", GetUserController)
  
	e.Start(":8000")
  }
  