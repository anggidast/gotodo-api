package routes

import (
	"net/http"

	"go-fancy-todo/config"
	"go-fancy-todo/controllers"
	"go-fancy-todo/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo { // * function yang langsung berjalan, ketika run project, mereturn instance
	e := echo.New()
	e.Validator = &models.CustomValidator{
		Validator: validator.New(),
	}
	config.NewDB()

	e.GET("/", func(c echo.Context) error { // * echo.Context untuk hanlde request dan response
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/todos", controllers.GetAllTodos)
	e.GET("/todos/:id", controllers.GetTodo)
	e.POST("/todos", controllers.AddTodo)
	e.PUT("/todos/:id", controllers.UpdateTodo)
	e.DELETE("/todos/:id", controllers.DeleteTodo)

	e.GET("/users", controllers.GetAllUsers)
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)
	e.PUT("/user/:id", controllers.UpdateUser)
	e.DELETE("/user/:id", controllers.DeleteUser)

	return e
}
