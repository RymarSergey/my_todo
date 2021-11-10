package handler

import (
	todo "github.com/RymarSergey/my_todo"
	"github.com/gin-gonic/gin"
	"net/http"
)

//registration
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type signinInput struct {
	UserName string `json:"username" binding:"required"`
	Pasword  string `json:"password" binding:"required"`
}

//authentification
func (h *Handler) signIn(c *gin.Context) {
	var input signinInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	token, err := h.services.Authorization.GenerateToken(input.UserName, input.Pasword)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
