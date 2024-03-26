package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mbrunos/go-hire/internal/core/dto"
	"github.com/mbrunos/go-hire/internal/core/usecases"
	"github.com/mbrunos/go-hire/pkg/middleware/jwtauth"
	"github.com/mbrunos/go-hire/pkg/router"
)

type UserHandler struct {
	userUseCase *usecases.UserUseCase
	jwtSecret   string
	jwtExp      time.Duration
}

func NewUserHandler(u *usecases.UserUseCase, secret string, exp time.Duration) *UserHandler {
	return &UserHandler{
		userUseCase: u,
		jwtSecret:   secret,
		jwtExp:      exp,
	}
}

// Login User 	godoc
// @Summary 		Login User
// @Description Login User
// @Tags 				users
// @Accept 			json
// @Produce 		json
// @Param 			request body dto.LoginInputDTO true "User credentials"
// @Success 		200 {object} dto.UserWithTokenOutputDTO
// @Failure 		401 {object} dto.ErrorOutputDTO
// @Failure 		404 {object} dto.ErrorOutputDTO
// @Failure 		500 {object} dto.ErrorOutputDTO
// @Router 			/login [post]
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

	token, err := jwtauth.NewToken(h.jwtSecret, user.ID.String(), h.jwtExp)

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

// Create User 	godoc
// @Summary 		Creates a new user
// @Description Creates a new user
// @Tags 				users
// @Accept 			json
// @Produce 		json
// @Param 			request body dto.CreateUserInputDTO true "User data"
// @Success 		201 {object} dto.UserOutputDTO
// @Failure 		400 {object} dto.ErrorOutputDTO
// @Failure 		500 {object} dto.ErrorOutputDTO
// @Router 			/signup [post]
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

// Update User 	godoc
// @Summary 		Updates a user
// @Description Updates a user
// @Tags 				users
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "User ID"
// @Param 			request body dto.UpdateUserInputDTO true "User data"
// @Success 		200 {object} dto.UserOutputDTO
// @Failure 		400 {object} dto.ErrorOutputDTO
// @Failure 		403 {object} dto.ErrorOutputDTO
// @Failure 		500 {object} dto.ErrorOutputDTO
// @Router 			/users/{id} [put]
// @Security 		ApiKeyAuth
func (h *UserHandler) Update(c *router.Context) {
	var input dto.UpdateUserInputDTO
	if err := c.BindJSON(&input); err != nil {
		c.SendError(http.StatusBadRequest, err)
		return
	}

	id := c.PathParam("id")
	user_id := c.Get("user_id").(string)

	if id != user_id {
		c.SendError(http.StatusForbidden, errors.New("you can only update your own user"))
		return
	}

	user, err := h.userUseCase.UpdateUser(id, &input)

	if err != nil {
		c.SendError(http.StatusInternalServerError, err)
		return
	}

	c.SendJSON(http.StatusOK, user)
}

// Delete User 	godoc
// @Summary 		Deletes a user
// @Description Deletes a user
// @Tags 				users
// @Produce 		json
// @Param 			id path string true "User ID"
// @Success 		204
// @Failure 		403 {object} dto.ErrorOutputDTO
// @Failure 		404 {object} dto.ErrorOutputDTO
// @Failure 		500 {object} dto.ErrorOutputDTO
// @Router 			/users/{id} [delete]
// @Security 		ApiKeyAuth
func (h *UserHandler) Delete(c *router.Context) {
	id := c.PathParam("id")
	user_id := c.Get("user_id").(string)

	if id != user_id {
		c.SendError(http.StatusForbidden, errors.New("you can only delete your own user"))
		return
	}

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
