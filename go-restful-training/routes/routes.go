package routes

import (
  c "go-restful-training/controllers"
  "github.com/labstack/echo"
  m "github.com/labstack/echo/middleware"
  auth "go-restful-training/middlewares"
)

func New() *echo.Echo {
  e := echo.New()

  // user routing
  e.GET("/api/users", c.GetUsersController, m.BasicAuth(auth.BasicAuth2))
  e.GET("/api/users/:id", c.GetUserController, m.BasicAuth(auth.BasicAuth2))
  e.POST("/api/users", c.CreateUserController, m.BasicAuth(auth.BasicAuth2))
  e.DELETE("/api/users/:id", c.DeleteUserController, m.BasicAuth(auth.BasicAuth2))
  e.PUT("/api/users/:id", c.UpdateUserController, m.BasicAuth(auth.BasicAuth2))
  
  return e
}