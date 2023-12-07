package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mshefeeqb/masjid-software/models"
	"gorm.io/gorm"
)

type MemberController struct {
	DB *gorm.DB
}

func NewMemberController(DB *gorm.DB) MemberController {
	return MemberController{DB}
}

func (mc *MemberController) CreateMember(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.MemberRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	now := time.Now()
	newMember := models.Member{
		Name:          payload.Name,
		Email:         payload.Email,
		Phone:         payload.Phone,
		Address:       payload.Address,
		DepartmentID:  payload.DepartmentID,
		Department:    payload.Department,
		DateOfBirth:   payload.DateOfBirth,
		Age:           payload.Age,
		Photo:         payload.Photo,
		Gender:        payload.Gender,
		WardID:        payload.WardID,
		BloodGroup:    payload.BloodGroup,
		ParentId:      payload.ParentId,
		MarrigeStatus: payload.MarrigeStatus,
		Occupation:    payload.Occupation,
		Education:     payload.Education,
		IsActive:      payload.IsActive,
		PendingAmount: payload.PendingAmount,
		FeePackageId:  payload.FeePackageId,
		CreatedAt:     now,
		UpdatedAt:     now,
		CreatedBy:     currentUser.ID,
		UpdatedBy:     currentUser.ID,
	}

	result := mc.DB.Create(&newMember)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Member Already Exists"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Member Created Successfully"})
}

func (mc *MemberController) UpdateMember(ctx *gin.Context) {
	memberId := ctx.Param("id")
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.MemberRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var updatedMember models.Member
	result := mc.DB.First(&updatedMember, "id=?", memberId)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	now := time.Now()
	memberToUpdate := models.Member{
		Name:         payload.Name,
		Email:        payload.Email,
		Phone:        payload.Phone,
		Address:      payload.Address,
		DepartmentID: payload.DepartmentID,
		Department:   payload.Department,
		DateOfBirth:  payload.DateOfBirth,
		Age:          payload.Age,
		Photo:        payload.Photo,
		Gender:       payload.Gender,
		WardID:       payload.WardID,
		BloodGroup:   payload.BloodGroup,
		ParentId:     payload.ParentId,
		Occupation:   payload.Occupation,
		Education:    payload.Education,
		IsActive:     payload.IsActive,
		UpdatedAt:    now,
		UpdatedBy:    currentUser.ID,
	}

	result = mc.DB.Model(&updatedMember).Updates(memberToUpdate)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Member Updated Successfully"})
}

func (mc *MemberController) GetAllMembers(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var members []models.Member
	result := mc.DB.Offset(offset).Limit(intLimit).Find(&members)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(200, gin.H{"status": "success", "message": members})

}

func (mc *MemberController) GetMember(ctx *gin.Context) {
	memberId := ctx.Param("id")

	var member models.Member
	result := mc.DB.First(&member, "id=?", memberId)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(200, gin.H{"status": "success", "message": member})
}

func (mc *MemberController) DeleteMember(ctx *gin.Context) {
	memberId := ctx.Param("id")
	currentUser := ctx.MustGet("currentUser").(models.User)
	if currentUser.Role != "superadmin" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "You are not allowed to delete member"})
		return
	}

	var member models.Member
	result := mc.DB.First(&member, "id=?", memberId)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	result = mc.DB.Delete(&member)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(200, gin.H{"status": "success", "message": "Member Deleted Successfully"})
}
