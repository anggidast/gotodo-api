package controllers

import (
	"errors"
	"go-fancy-todo/config"
	"go-fancy-todo/middlewares"
	"go-fancy-todo/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func GetAllTodos(c echo.Context) (err error) {
	db := config.NewDB()
	todos := []models.Todo{}

	var userId string = strconv.Itoa(int(middlewares.UserId))

	if err = db.Find(&todos, "user_id = ?", userId).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, todos)
}

func GetTodo(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))

	db := config.NewDB()
	todo := models.Todo{}

	if err = db.First(&todo, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "todo not found",
		})
	}

	return c.JSON(http.StatusOK, todo)
}

func AddTodo(c echo.Context) (err error) {
	req := new(models.Todo)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var userId int = int(middlewares.UserId)

	newTodo := models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Status:      "undone",
		Due_date:    req.Due_date,
		UserId:      userId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	due, _ := time.Parse("2006-01-02", req.Due_date)
	today := time.Now()
	if due.Format("2006-01-02") < today.Format("2006-01-02") {
		err = errors.New("validation_error: Due date cannot be the day before today")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "todo not found",
		})
	}

	req := new(models.Todo)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var userId int = int(middlewares.UserId)

	todo.Title = req.Title
	todo.Description = req.Description
	todo.Status = req.Status
	todo.Due_date = req.Due_date
	todo.UserId = userId
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
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "todo not found",
		})
	}

	db.Delete(&todo, id)

	response := models.TodoResponse{
		Message: "deleted",
		Data:    todo,
	}

	return c.JSON(http.StatusOK, response)
}
