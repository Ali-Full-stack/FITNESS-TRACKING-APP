package requests

import (
	"database/sql"
	"time"
)

type CreateWorkoutRequest struct {
	User_ID     int32    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateWorkoutResponse struct {
	ID          int32            `json:"id"`
	User_ID     int32           `json:"user_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description,omitempty"`
	Date        time.Time      `json:"date"`
	Created_at  string      `json:"created_at"`
	Updated_at  string      `json:"updated_at"`
}

type UpdateWorkoutRequest struct{
	ID int32 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type WorkoutResponse struct{
	Message string `json:"message"`
}
