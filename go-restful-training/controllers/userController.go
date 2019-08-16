package controllers

import (
  "net/http"
  "strconv"
  "crypto/md5"
  "go-restful-training/models"
  "github.com/labstack/echo"
  "fmt"
)

// get all users
func GetUsersController(c echo.Context) error {
  users, err := models.GetUsers()
  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, users)
}


// get single user by id
func GetUserController(c echo.Context) error {
  id, _ := strconv.Atoi(c.Param("id"))
  user, err := models.GetUser(id)

  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, user)
}

// create new user
func CreateUserController(c echo.Context) error {
  user := models.User{}
  c.Bind(&user)

  user.Password = ToMD5(user.Password)

  result, err := models.CreateUser(&user)
  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusCreated, result)
}

func DeleteUserController(c echo.Context) error {
  id, _ := strconv.Atoi(c.Param("id"))

  user, _ := models.DeleteUser(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "yah udah ke delete :(",
		"user": user,
		})
  }
  
  func UpdateUserController(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    newUser := models.User{}
    c.Bind(&newUser)

    user, _ := models.UpdateUser(id,newUser)
    return c.JSON(http.StatusOK, map[string]interface{}{
      "message": "keupdate broo!! :(",
      "user": user,
      })
  }
 

func ToMD5(s string) string {
  return fmt.Sprintf("%x", md5.Sum([]byte (s)))
}