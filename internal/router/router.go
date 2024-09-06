package router

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/errors"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/hash"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/storage"
)

type UserHandler struct {
	logger  *slog.Logger
	storage storage.Queries
}

func NewMux(logger *slog.Logger, storage storage.Queries) http.Handler {
	mux := http.NewServeMux()
	userHandler := UserHandler{
		logger:  logger,
		storage: storage,
	}
	mux.HandleFunc("POST /api/users/register", userHandler.Register)
	return mux
}

func (u UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userRegisterReq UserRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&userRegisterReq); err != nil {
		u.logger.Error("failed to decode user register data: ", err)
		http.Error(w, errors.ErrDecodeUserRegister.Error(), http.StatusBadRequest)
		return
	}

	passwordHash, err := hash.GenerateFromPassword(userRegisterReq.Password)
	if err != nil {
		u.logger.Error("failed to hash user password", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user, err := u.storage.CreateUser(r.Context(), storage.CreateUserParams{
		Username:     userRegisterReq.Username,
		PasswordHash: passwordHash,
		Email:        userRegisterReq.Email,
	})
	if err != nil {
		u.logger.Error("failed to create user from db", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	userRegisterResp :=UserRegisterResponse{
		ID :int(user.ID),
		Username: user.Username,
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userRegisterResp)

}
