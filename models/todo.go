package models

import "time"

type (
	Todo struct {
		ID          int       `json:"id"`
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		Due_date    string    `json:"due_date" validate:"required"`
		UserId      int       `json:"user_id" gorm:"column:user_id"`
		CreatedAt   time.Time `json:"createdAt" gorm:"column:createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"  gorm:"column:updatedAt"`
	}

	TodoResponse struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)
