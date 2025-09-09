package handlers

import (
	"net/http"
	"w2s-backend/services"
	"w2s-backend/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := services.RegisterUser(req.Name, req.Email); err != nil {
		utils.SendError(c, http.StatusConflict, err.Error())
		return
	}

	utils.SendStatusMessage(c, "OTP sent to email. Verify to complete registration")
}

func VerifyOTP(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		OTP      string `json:"otp" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := services.VerifyUserOTP(req.Email, req.OTP, req.Password); err != nil {
		utils.SendError(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SendStatusMessage(c, "user verified & registered successfully")
}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := services.LoginUser(req.Email, req.Password)
	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SendJSON(c, gin.H{"token": token})
}

func ForgotPassword(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	otp := utils.GenerateOTP()

	// if err := utils.SendOTPEmail(req.Email, otp); err != nil {
	// 	utils.SendError(c, http.StatusInternalServerError, "failed to send OTP")
	// 	return
	// }

	if err := utils.SaveOTP(req.Email, otp); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "failed to save OTP")
		return
	}

	utils.SendStatusMessage(c, "OTP sent to email for password reset")
}

func ResetPassword(c *gin.Context) {
	var req struct {
		Email       string `json:"email" binding:"required,email"`
		OTP         string `json:"otp" binding:"required"`
		NewPassword string `json:"newPassword" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if !utils.VerifyOTP(req.Email, req.OTP) {
		utils.SendError(c, http.StatusUnauthorized, "invalid or expired OTP")
		return
	}

	if err := services.ResetPassword(req.Email, req.NewPassword); err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendStatusMessage(c, "password reset successfully")
}
