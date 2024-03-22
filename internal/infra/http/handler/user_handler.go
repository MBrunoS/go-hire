package handler

import (
	"fmt"
	"net/http"

	"github.com/mbrunos/go-hire/internal/core/dto"
	"github.com/mbrunos/go-hire/internal/core/usecases"
	"github.com/mbrunos/go-hire/pkg/middleware/jwtauth"
	"github.com/mbrunos/go-hire/pkg/router"
)

type UserHandler struct {
	userUseCase *usecases.UserUseCase
	jwtSecret   string
}

func NewUserHandler(u *usecases.UserUseCase, secret string) *UserHandler {
	return &UserHandler{
		userUseCase: u,
		jwtSecret:   secret,
	}
}

func (h *UserHandler) Login(c *router.Context) {
	var input dto.LoginInputDTO
	if err := c.BindJSON(&input); err != nil {
		c.SendError(http.StatusBadRequest, err)
		return
	}

	user, err := h.userUseCase.FindUserByEmail(input.Email)

	if err != nil {
		c.SendError(http.StatusNotFound, fmt.Errorf("user with email %s not found", input.Email))
		return
	}

	if !user.ComparePassword(input.Password) {
		c.SendError(http.StatusUnauthorized, err)
		return
	}

	token, err := jwtauth.NewToken(h.jwtSecret, user.ID.String())

	if err != nil {
		c.SendError(http.StatusInternalServerError, err)
		return
	}

	c.SendJSON(http.StatusOK, &dto.UserWithTokenOutputDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	})
}

func (h *UserHandler) SignUp(c *router.Context) {
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

	id := c.PathParam("id")
	user, err := h.userUseCase.UpdateUser(id, &input)

	if err != nil {
		c.SendError(http.StatusInternalServerError, err)
		return
	}

	c.SendJSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *router.Context) {
	id := c.PathParam("id")

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
