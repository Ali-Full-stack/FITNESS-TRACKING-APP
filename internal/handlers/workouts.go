package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/errors"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/requests"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/storage"
)

func (h Handler) CreateWorkouts(w http.ResponseWriter, r *http.Request) {
	var createWorkoutReq requests.CreateWorkoutRequest
	if err := json.NewDecoder(r.Body).Decode(&createWorkoutReq); err != nil {
		h.Logger.Error("failed to decode  workouts data: ", slog.Any("error", err))
		http.Error(w, errors.ErrDecodeRequestBody.Error(), http.StatusBadRequest)
		return
	}

	workout, err := h.Storage.CreateWorkout(r.Context(), storage.CreateWorkoutParams{
		UserID: int32(createWorkoutReq.User_ID),
		Name:   createWorkoutReq.Name,
		Description: sql.NullString{
			String: createWorkoutReq.Description,
			Valid:  true,
		},
	})
	if err != nil {
		h.Logger.Error("failed to create workout from db", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	workoutResp := requests.CreateWorkoutResponse{
		ID:          workout.ID,
		User_ID:     workout.UserID,
		Name:        workout.Name,
		Description: workout.Description,
		Date:        workout.Date,
		Created_at:  workout.CreatedAt.Format(time.ANSIC),
		Updated_at:  workout.UpdatedAt.Format(time.ANSIC),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workoutResp)
}

func (h Handler) GetWorkoutsByID(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, _ := strconv.Atoi(idString)

	workouts, err := h.Storage.GetWorkoutByUserID(r.Context(), int32(id))
	if err != nil {
		h.Logger.Error("failed to get  workouts from db", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var listWorkouts []requests.CreateWorkoutResponse
	for _, w :=range workouts {
		workoutResp := requests.CreateWorkoutResponse{
			ID:          w.ID,
			User_ID:     w.UserID,
			Name:        w.Name,
			Description: w.Description,
			Date:        w.Date,
			Created_at:  w.CreatedAt.Format(time.ANSIC),
			Updated_at:  w.UpdatedAt.Format(time.ANSIC),
		}
		listWorkouts = append(listWorkouts, workoutResp)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listWorkouts)
}
