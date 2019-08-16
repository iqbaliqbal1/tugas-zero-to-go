package middlewares

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "go-restful-training/models"
  "go-restful-training/config"
  "go-restful-training/controllers"
  
)

func LogMiddlewares(e *echo.Echo) {
  e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
    Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
  }))
}

// middlewares/authMiddlewares.go
func BasicAuth(username, password string, c echo.Context) (bool, error) {
  if username == "admin" && password == "admin" {
    return true, nil
  }
  return false, nil
}

// middlewares/authMiddlewares.go
func BasicAuth2(username, password string, c echo.Context) (bool, error) {
  var db = config.DB
  var user models.User
  password = controllers.ToMD5(password)
  if err := db.Where("email = ? AND password = ?", username, password).
  First(&user).Error; err != nil {
    return false, nil
  }
  return true, nil
}

