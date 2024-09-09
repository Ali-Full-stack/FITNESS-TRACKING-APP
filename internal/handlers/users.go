package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/errors"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/hash"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/requests"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/storage"
)



func (u Handler) UserRegister(w http.ResponseWriter, r *http.Request) {
	var userRegisterReq requests.UserRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&userRegisterReq); err != nil {
		u.Logger.Error("failed to decode user register data: ", slog.Any("error", err))
		http.Error(w, errors.ErrDecodeRequestBody.Error(), http.StatusBadRequest)
		return
	}

	passwordHash, err := hash.GenerateFromPassword(userRegisterReq.Password)
	if err != nil {
		u.Logger.Error("failed to hash user password", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user, err := u.Storage.CreateUser(r.Context(), storage.CreateUserParams{
		Username:     userRegisterReq.Username,
		PasswordHash: passwordHash,
		Email:        userRegisterReq.Email,
	})
	if err != nil {
		u.Logger.Error("failed to create user from db", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	userRegisterResp := requests.UserRegisterResponse{
		ID:       int(user.ID),
		Username: user.Username,
		Email:    user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userRegisterResp)

}

func (u Handler) UserLogin(w http.ResponseWriter, r *http.Request) {
	var LoginReq requests.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&LoginReq); err != nil {
		u.Logger.Error("failed to decode user login data: ", slog.Any("error", err))
		http.Error(w, errors.ErrDecodeRequestBody.Error(), http.StatusBadRequest)
		return
	}
}
