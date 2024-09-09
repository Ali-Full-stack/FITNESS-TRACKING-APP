package router

import (
	"log/slog"
	"net/http"

	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/handlers"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/storage"
)

func NewMux(logger *slog.Logger, storage storage.Queries) http.Handler {
	mux := http.NewServeMux()
	
	handler :=handlers.NewHandler(logger, storage)

	// USERS
	mux.HandleFunc("POST /api/users/register", handler.UserRegister)
	mux.HandleFunc("POST /api/users/login", handler.UserLogin)

	//WORKOUTS
	mux.HandleFunc("POST /api/workouts",  handler.CreateWorkouts)
	mux.HandleFunc("GET /api/workouts/{id}", handler.GetWorkoutsByUserID)
	mux.HandleFunc("GET /api/workouts", handler.GetWorkoutsByID)
	mux.HandleFunc("PUT /api/workouts/{id}", handler.UpdateWorkoutsByUserID)
	mux.HandleFunc("DELETE /api/workouts", handler.DeleteWorkoutsByID)

	return mux
}



