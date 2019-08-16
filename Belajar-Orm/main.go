package main

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"net/http"
	"fmt"
)
 
var (
  db *gorm.DB
)

type User struct {
  gorm.Model
  Name     string `json:"name" form:"name"`
  Email    string `json:"email" form:"email"`
  Password string `json:"password" form:"password"`
}

func InitialMigration() {
  db.AutoMigrate(&User{})
}

func InitDB(connectionString string) {
	var err error
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
	  log.Panic(err)
	}
	if err = db.DB().Ping(); err != nil {
	  panic(err)
	}
	fmt.Println(err)
  }

//   get user controller
  func GetUsersController(c echo.Context) error {
	var users []User
  
	if err := db.Find(&users).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all users",
	  "users":   users,
	})
  }

// create user controller
  func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)
  
	if err := db.Save(&user).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new user",
	  "user":    user,
	})
  }
  func getusercontrollerid(c echo.Context) error {
	id := c.Param("id")
	user := User{}
	
	// err := db.First(&user,id).Error
	idku := db.First(&user,id)

	// fmt.Println(err)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "Not Found",
	// 		"user":"kosong cukk",
	// 		})
	// }
	return c.JSON(http.StatusOK, map[string]interface{}{
	"message": "success create new user",
	"user":	idku,
	})
}

func Deleteusercontroller(c echo.Context) error {
	id := c.Param("id")
	user := User{}

	hapus_id := db.Delete(&user,id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "yah udah ke delete :(",
		"user":	hapus_id,
		})
	}

	
	func Updateusercontroller(c echo.Context) error {
		id := c.Param("id")
		user := User{}
	
		data := db.First(&user,id)
		c.Bind(&user)
		db.Save(user)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Udah ke update brooo!!",
			"user":	data,
			})
		}


  func main() {
	// create a new echo instance
	e := echo.New()
	InitDB("root:Iqbaliqbal1@/go_db?charset=utf8&parseTime=True&loc=Local")
	InitialMigration()
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", getusercontrollerid)
	e.DELETE("/users/:id", Deleteusercontroller)
	e.PUT("/users/:id", Updateusercontroller)
	// start the server
	e.Start(":8000")
  }
  
  