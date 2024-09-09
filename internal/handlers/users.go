package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/auth/token"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/errors"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/auth/hash"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/requests"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/storage"
)

func (h Handler) UserRegister(w http.ResponseWriter, r *http.Request) {
	var userRegisterReq requests.UserRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&userRegisterReq); err != nil {
		h.Logger.Error("failed to decode user register data: ", slog.Any("error", err))
		http.Error(w, errors.ErrDecodeRequestBody.Error(), http.StatusBadRequest)
		return
	}

	passwordHash, err := hash.GenerateFromPassword(userRegisterReq.Password)
	if err != nil {
		h.Logger.Error("failed to hash user password", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user, err := h.Storage.CreateUser(r.Context(), storage.CreateUserParams{
		Username:     userRegisterReq.Username,
		PasswordHash: passwordHash,
		Email:        userRegisterReq.Email,
	})
	if err != nil {
		h.Logger.Error("failed to create user from db", slog.Any("error", err))
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

func (h Handler) UserLogin(w http.ResponseWriter, r *http.Request) {
	var LoginReq requests.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&LoginReq); err != nil {
		h.Logger.Error("failed to decode user login data: ", slog.Any("error", err))
		http.Error(w, errors.ErrDecodeRequestBody.Error(), http.StatusBadRequest)
		return
	}
	hash, err := h.Storage.VerifyUserLogin(r.Context(), LoginReq.ID)
	if err != nil {
		h.Logger.Error("failed to check user login  from db", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if hash != LoginReq.Password {
		http.Error(w, "Invalid Password !!", http.StatusBadRequest)
		return
	}
	accessToken, err := token.GenerateToken(LoginReq.ID, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
}
