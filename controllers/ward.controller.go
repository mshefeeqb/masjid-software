package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mshefeeqb/masjid-software/models"
	"gorm.io/gorm"
)

type WardController struct {
	DB *gorm.DB
}

func NewWardController(DB *gorm.DB) WardController {
	return WardController{DB}
}

func (wc *WardController) CreateWard(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.WardRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	now := time.Now()
	newWard := models.Ward{
		Name:       payload.Name,
		WardNumber: payload.WardNumber,
		CreatedAt:  now,
		UpdatedAt:  now,
		CreatedBy:  currentUser.ID,
		UpdatedBy:  currentUser.ID,
	}

	result := wc.DB.Create(&newWard)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Ward created successfully!"})
}

func (wc *WardController) GetWards(ctx *gin.Context) {
	var wards []models.Ward
	result := wc.DB.Find(&wards)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": wards})
}

func (wc *WardController) GetWard(ctx *gin.Context) {
	var ward models.Ward
	id := ctx.Param("id")
	result := wc.DB.Where("id = ?", id).First(&ward)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": ward})
}

func (wc *WardController) UpdateWard(ctx *gin.Context) {
	wardId := ctx.Param("id")
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.WardRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var updatedWard models.Ward
	result := wc.DB.First(&updatedWard, "id=?", wardId)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	updatedWard.Name = payload.Name
	updatedWard.WardNumber = payload.WardNumber
	updatedWard.UpdatedAt = time.Now()
	updatedWard.UpdatedBy = currentUser.ID

	result = wc.DB.Save(&updatedWard)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Ward updated successfully!"})
}

func (wc *WardController) DeleteWard(ctx *gin.Context) {
	wardId := ctx.Param("id")
	var ward models.Ward
	result := wc.DB.First(&ward, "id=?", wardId)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	result = wc.DB.Delete(&ward)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Ward deleted successfully!"})
}
