package handler

import (
	"fmt"
	"net/http"

	"github.com/mbrunos/go-hire/internal/core/dto"
	"github.com/mbrunos/go-hire/internal/core/usecases"
	"github.com/mbrunos/go-hire/pkg/router"
)

type UserHandler struct {
	userUseCase *usecases.UserUseCase
	JwtAuth     router.Middleware
}

func NewUserHandler(usecase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (h *UserHandler) Create(c *router.Context) {
	var input dto.CreateUserInputDTO
	if err := c.BindJSON(&input); err != nil {
		c.SendError(http.StatusBadRequest, err)
		return
	}

	user, err := h.userUseCase.CreateUser(&input)

	if err != nil {
		c.SendError(http.StatusInternalServerError, err)
		return
	}

	c.SendJSON(http.StatusCreated, user)
}

func (h *UserHandler) Update(c *router.Context) {
	var input dto.UpdateUserInputDTO
	if err := c.BindJSON(&input); err != nil {
		c.SendError(http.StatusBadRequest, err)
		return
	}

	id := c.URLParam("id")
	user, err := h.userUseCase.UpdateUser(id, &input)

	if err != nil {
		c.SendError(http.StatusInternalServerError, err)
		return
	}

	c.SendJSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *router.Context) {
	id := c.URLParam("id")

	_, err := h.userUseCase.FindUserByID(id)

	if err != nil {
		c.SendError(http.StatusNotFound, fmt.Errorf("user with id %s not found", id))
		return
	}

	if err := h.userUseCase.DeleteUser(id); err != nil {
		c.SendError(http.StatusInternalServerError, err)
		return
	}

	c.SendJSON(http.StatusNoContent, nil)
}
