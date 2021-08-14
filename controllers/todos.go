package controllers

import (
	"go-fancy-todo/config"
	"go-fancy-todo/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func GetAllTodos(c echo.Context) (err error) {
	db := config.NewDB()
	todos := []models.Todo{}

	if err = db.Find(&todos).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, todos)
}

func GetTodo(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))

	db := config.NewDB()
	todo := models.Todo{}

	if err = db.First(&todo, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, todo)
}

func AddTodo(c echo.Context) (err error) {
	req := new(models.Todo)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newTodo := models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Status:      "undone",
		Due_date:    req.Due_date,
		UserId:      req.UserId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	db := config.NewDB()
	db.Create(&newTodo)
	response := models.TodoResponse{
		Message: "created",
		Data:    newTodo,
	}

	return c.JSON(http.StatusCreated, response)
}

func UpdateTodo(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))

	db := config.NewDB()
	todo := models.Todo{}

	if err = db.First(&todo, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	req := new(models.Todo)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	todo.Title = req.Title
	todo.Description = req.Description
	todo.Status = req.Status
	todo.Due_date = req.Due_date
	todo.UserId = req.UserId
	todo.UpdatedAt = time.Now()

	db.Save(&todo)

	response := models.TodoResponse{
		Message: "updated",
		Data:    todo,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteTodo(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))

	db := config.NewDB()
	todo := models.Todo{}

	if err = db.First(&todo, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	db.Delete(&todo, id)

	response := models.TodoResponse{
		Message: "deleted",
		Data:    todo,
	}

	return c.JSON(http.StatusOK, response)
}
