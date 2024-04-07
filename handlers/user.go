package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/skill-map/database"
	"github.com/muhammedarifp/skill-map/managers"
	"github.com/muhammedarifp/skill-map/models"
)

type UserHandler struct {
	groupName   string
	userManager managers.UserManager
}

func NewUserHandlerFrom(um managers.UserManager) *UserHandler {
	return &UserHandler{
		groupName:   "api/v1/users",
		userManager: um,
	}
}

func (uh *UserHandler) RegisterApis(r *gin.Engine) {
	router := r.Group(uh.groupName)
	{
		router.POST("", uh.createUser)
	}
}

func (userHandler *UserHandler) createUser(ctx *gin.Context) {
	db := database.GetDb()
	if err := db.Create(&models.UserModel{FullName: "Arif", Email: "muhammedarif0100@gmailcom"}).Error; err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Done",
	})
}
