package validator

import "time"

type ToDoCreateRequest struct {
	Name      string    `json:"name" validate:"required"`
	Priority  string    `json:"priority" validate:"required"`
	StartTime time.Time `json:"start_time" validate:"required"`
	EndTime   time.Time `json:"end_time" validate:"required"`
}

type ToDoUpdateRequest struct {
	ID        uint      `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"omitempty"`
	Priority  string    `json:"priority" validate:"omitempty"`
	StartTime time.Time `json:"start_time" validate:"omitempty"`
	EndTime   time.Time `json:"end_time" validate:"omitempty"`
}

type ToDoDeleteRequest struct {
	ID uint `json:"id" validate:"required"`
}
