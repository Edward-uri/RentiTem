package controller

import (
	"net/http"

	"main/src/features/users/application"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	getUC    *application.GetProfileUseCase
	updateUC *application.UpdateProfileUseCase
}

func NewUserController(getUC *application.GetProfileUseCase, updateUC *application.UpdateProfileUseCase) *UserController {
	return &UserController{getUC: getUC, updateUC: updateUC}
}

// GetMe returns the current user's profile.
// @Summary Get current user profile
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /api/v1/users/me [get]
func (c *UserController) GetMe(ctx *gin.Context) {
	idVal, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := idVal.(uint)

	user, err := c.getUC.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":          user.ID,
		"full_name":   user.FullName,
		"email":       user.Email,
		"phone":       user.Phone,
		"address":     user.Address,
		"profile_pic": user.ProfilePicURL,
		"role":        user.Role,
	})
}

type updateMeRequest struct {
	FullName   *string `json:"full_name"`
	Phone      *string `json:"phone"`
	Address    *string `json:"address"`
	ProfilePic *string `json:"profile_pic"`
}

// UpdateMe updates the current user's profile.
// @Summary Update current user profile
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param payload body updateMeRequest true "Update profile payload"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /api/v1/users/me [put]
func (c *UserController) UpdateMe(ctx *gin.Context) {
	idVal, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := idVal.(uint)

	var req updateMeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := application.UpdateProfileInput{
		FullName:      req.FullName,
		Phone:         req.Phone,
		Address:       req.Address,
		ProfilePicURL: req.ProfilePic,
	}

	user, err := c.updateUC.Execute(userID, input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}
