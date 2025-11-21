package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizqdwan/go-chats-api/internal/services"
)

type Handler struct {
	service services.UserService
}

func NewHandler(s services.UserService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.CreateUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, token, err := h.service.SignIn(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}