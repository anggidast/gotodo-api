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

func dateValidation(due_date string) (err error) {
	due, _ := time.Parse("2006-01-02", due_date)

	today := time.Now()
	if due.Format("2006-01-02") < today.Format("2006-01-02") {
		err = errors.New("validation_error: Due date cannot be the day before today")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func GetAllTodos(c echo.Context) (err error) {
	db := config.Database()
	todos := []models.Todo{}

	var userId string = strconv.Itoa(int(middlewares.UserId))

	if err = db.Order("status desc, due_date desc").Find(&todos, "user_id = ?", userId).Error; err != nil {
		return err
	}

	response := models.TodoResponse{
		Message: "succeed",
		Data:    todos,
	}

	return c.JSON(http.StatusOK, response)
}

func GetTodo(c echo.Context) (err error) {
	id := c.Param("id")

	todo, _, _, err := middlewares.Authorization(id, c)
	if err != nil {
		return err
	}

	response := models.TodoResponse{
		Message: "succeed",
		Data:    todo,
	}

	return c.JSON(http.StatusOK, response)
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

	if err = dateValidation(req.Due_date); err != nil {
		return err
	}

	db := config.Database()
	db.Create(&newTodo)
	response := models.TodoResponse{
		Message: "created",
		Data:    newTodo,
	}

	return c.JSON(http.StatusCreated, response)
}

func UpdateTodo(c echo.Context) (err error) {
	id := c.Param("id")

	todo, userId, db, err := middlewares.Authorization(id, c)
	if err != nil {
		return err
	}

	req := new(models.Todo)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userIdInt, _ := strconv.Atoi(userId)

	todo.Title = req.Title
	todo.Description = req.Description
	todo.Due_date = req.Due_date
	todo.UserId = userIdInt
	todo.UpdatedAt = time.Now()

	if err = dateValidation(req.Due_date); err != nil {
		return err
	}

	db.Save(&todo)

	response := models.TodoResponse{
		Message: "updated",
		Data:    todo,
	}

	return c.JSON(http.StatusOK, response)
}

func ChangeStatus(c echo.Context) (err error) {
	id := c.Param("id")

	todo, _, db, err := middlewares.Authorization(id, c)
	if err != nil {
		return err
	}

	req := new(models.Todo)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	todo.Status = req.Status

	db.Save(&todo)

	response := models.TodoResponse{
		Message: "updated",
		Data:    todo,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteTodo(c echo.Context) (err error) {
	id := c.Param("id")

	todo, _, db, err := middlewares.Authorization(id, c)
	if err != nil {
		return err
	}

	db.Delete(&todo, id)

	response := models.TodoResponse{
		Message: "deleted",
		Data:    todo,
	}

	return c.JSON(http.StatusOK, response)
}
