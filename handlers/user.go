package handlers

import (
	"net/http"
	"w2s-backend/services"
	"w2s-backend/utils"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	uid := c.GetString("userId")
	user, err := services.GetUserProfile(uid)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.SendJSON(c, user)
}

func UpdateEmail(c *gin.Context) {
	uid := c.GetString("userId")
	var req struct {
		NewEmail string `json:"newEmail" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := services.UpdateUserEmail(uid, req.NewEmail); err != nil {
		utils.SendError(c, http.StatusConflict, err.Error())
		return
	}
	utils.SendStatusMessage(c, "email updated successfully")
}

func UpdateProfile(c *gin.Context) {
	uid := c.GetString("userId")
	var req struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := services.UpdateUserProfile(uid, req.Name); err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendStatusMessage(c, "profile updated successfully")
}
