package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MemberController struct {
	DB *gorm.DB
}

func NewMemberController(DB *gorm.DB) MemberController {
	return MemberController{DB}
}

func (mc *MemberController) GetAllMembers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "success", "message": "All Members"})
}
