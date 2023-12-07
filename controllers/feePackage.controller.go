package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mshefeeqb/masjid-software/models"
	"gorm.io/gorm"
)

type FeePackageController struct {
	DB *gorm.DB
}

func NewFeePackageController(DB *gorm.DB) FeePackageController {
	return FeePackageController{DB}
}

func (fpc *FeePackageController) CreateFeePackage(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.FeePackageRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	now := time.Now()
	newFeePackage := models.FeePackage{
		Description: payload.Description,
		Amount:      payload.Amount,
		IsActive:    payload.IsActive,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   currentUser.ID,
		UpdatedBy:   currentUser.ID,
	}

	result := fpc.DB.Create(&newFeePackage)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "FeePackage created successfully!"})
}

func (fpc *FeePackageController) GetFeePackages(ctx *gin.Context) {
	var feePackages []models.FeePackage
	result := fpc.DB.Where("is_active = ?", true).Find(&feePackages)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": feePackages})
}

func (fpc *FeePackageController) GetFeePackage(ctx *gin.Context) {
	var feePackage models.FeePackage
	feePackageId := ctx.Param("id")
	result := fpc.DB.Where("id = ?", feePackageId).First(&feePackage)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": feePackage})
}

func (fpc *FeePackageController) UpdateFeePackage(ctx *gin.Context) {
	feePackageId := ctx.Param("id")
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.FeePackageRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var updatedFeePackage models.FeePackage
	result := fpc.DB.First(&updatedFeePackage, "id=?", feePackageId)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	updatedFeePackage.Description = payload.Description
	updatedFeePackage.Amount = payload.Amount
	updatedFeePackage.IsActive = payload.IsActive
	updatedFeePackage.UpdatedAt = time.Now()
	updatedFeePackage.UpdatedBy = currentUser.ID

	result = fpc.DB.Save(&updatedFeePackage)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "FeePackage updated successfully!"})
}

func (fpc *FeePackageController) DeleteFeePackage(ctx *gin.Context) {
	feePackageId := ctx.Param("id")
	var feePackage models.FeePackage
	result := fpc.DB.First(&feePackage, "id=?", feePackageId)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	result = fpc.DB.Delete(&feePackage)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "FeePackage deleted successfully!"})
}
