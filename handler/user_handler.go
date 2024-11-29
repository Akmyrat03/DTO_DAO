package handler

import (
	"dto/dto"
	service "dto/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

type errResponse struct {
	Message string
}

func NewUserHandler(UserService *service.UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

// Create a new user
// @Summary Create a new user
// @Description Create a new user by providing the username and password
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.UserRequest true "User"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /users/create [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var userRequest dto.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, errResponse{Message: "Invalid ID"})
		return
	}

	err := h.UserService.Create(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// GetAllUsers returns a list of all users
// @Summary get all users
// @Description get all users
// @Tags User
// @Success 200 {object} dto.UserResponse
// @Failure 500 {object} errResponse
// @Router /users/all [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.UserService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID returns a user by ID
// @Summary get user by ID
// @Description get user by ID
// @Tags User
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse{Message: "Invalid ID"})
		return
	}

	user, err := h.UserService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, errResponse{Message: "Failed to get user by ID"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUserBy ID delete user by ID
// @Summary delete user
// @Description delete user by ID
// @Tags User
// @Param id path int true "User ID"
// @Success 200 {object} string "User deleted successfully"
// @Failure 400 {object} errResponse
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.UserService.DeleteByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
