package controller

import (
	"net/http"

	"main/src/features/auth/application"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	uc *application.AuthUseCase
}

func NewAuthController(uc *application.AuthUseCase) *AuthController {
	return &AuthController{uc: uc}
}

type registerRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type authResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

// Register handles user registration.
// @Summary Register user
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body registerRequest true "Register payload"
// @Success 201 {object} authResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/v1/auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := c.uc.Register(application.RegisterInput{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		Address:  req.Address,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, authResponse{
		Token: token,
		User: gin.H{
			"id":        user.ID,
			"full_name": user.FullName,
			"email":     user.Email,
			"phone":     user.Phone,
			"address":   user.Address,
		},
	})
}

// Login handles user login.
// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body loginRequest true "Login payload"
// @Success 200 {object} authResponse
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Router /api/v1/auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := c.uc.Login(application.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "invalid credentials" {
			status = http.StatusUnauthorized
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, authResponse{
		Token: token,
		User: gin.H{
			"id":        user.ID,
			"full_name": user.FullName,
			"email":     user.Email,
			"phone":     user.Phone,
			"address":   user.Address,
		},
	})
}
