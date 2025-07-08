package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"crud-with-mongodb/internal/models"
	"crud-with-mongodb/internal/services"
	"crud-with-mongodb/internal/utils"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := c.service.CreateUser(ctx.Request.Context(), &req)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(ctx, http.StatusCreated, user)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.service.GetUser(ctx.Request.Context(), id)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusNotFound, "User not found")
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var req models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := c.service.UpdateUser(ctx.Request.Context(), id, &req)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.service.DeleteUser(ctx.Request.Context(), id); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(ctx, http.StatusNoContent, nil)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers(ctx.Request.Context())
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(ctx, http.StatusOK, users)
}