package routes

import (
	"net/http"

	"go-fancy-todo/config"
	"go-fancy-todo/controllers"
	"go-fancy-todo/middlewares"
	"go-fancy-todo/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(c echo.Context) *echo.Echo {
	c.Request().Header.Set(echo.HeaderAccessControlAllowOrigin, "*")

	e := echo.New()
	e.Validator = &models.CustomValidator{
		Validator: validator.New(),
	}
	e.Use(middlewares.ACAOHeaderOverwriteMiddleware, middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowMethods, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, "access_token"},
	}))
	config.NewDB()

	// * define group, with group level middleware
	todos := e.Group("/todos", middlewares.Authentication)
	todos.Use()

	e.GET("/", func(c echo.Context) error { // * echo.Context for handle request and response
		return c.String(http.StatusOK, "Go ToDo API Connected")
	})

	// * using group
	todos.GET("", controllers.GetAllTodos)
	todos.GET("/:id", controllers.GetTodo)
	todos.POST("", controllers.AddTodo)
	todos.PUT("/:id", controllers.UpdateTodo)
	todos.PATCH("/:id", controllers.ChangeStatus)
	todos.DELETE("/:id", controllers.DeleteTodo)

	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users/login", controllers.Login)
	e.POST("/users/register", controllers.Register)
	e.PUT("/user/:id", controllers.UpdateUser)
	e.DELETE("/user/:id", controllers.DeleteUser)

	e.OPTIONS("/users/login", controllers.Login)

	return e
}
