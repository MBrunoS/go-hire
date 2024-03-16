package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mbrunos/go-hire/internal/core/dto"
	"github.com/mbrunos/go-hire/internal/core/usecases"
)

type UserHandler struct {
	userUseCase *usecases.UserUseCase
}

func NewUserHandler(usecase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.userUseCase.CreateUser(&input)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	sendSuccess(w, http.StatusCreated, user)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")

	user, err := h.userUseCase.FindUserByEmail(email)

	if err != nil {
		sendError(w, http.StatusNotFound, err)
		return
	}

	sendSuccess(w, http.StatusOK, user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateUserInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	id := r.PathValue("id")
	user, err := h.userUseCase.UpdateUser(id, &input)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	sendSuccess(w, http.StatusOK, user)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")

	err := h.userUseCase.DeleteUser(email)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	sendSuccess(w, http.StatusNoContent, nil)
}
