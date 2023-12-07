package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mshefeeqb/masjid-software/controllers"
	"github.com/mshefeeqb/masjid-software/middleware"
)

type FeePackageRouteController struct {
	feePackageController controllers.FeePackageController
}

func NewFeePackageRouteController(feePackageController controllers.FeePackageController) FeePackageRouteController {
	return FeePackageRouteController{feePackageController}
}

func (fpc *FeePackageRouteController) FeePackageRoute(rs *gin.RouterGroup) {
	router := rs.Group("fee-packages")
	router.Use(middleware.DeserializeUser())
	router.GET("/", fpc.feePackageController.GetFeePackages)
	router.POST("/", fpc.feePackageController.CreateFeePackage)
	router.GET("/:id", fpc.feePackageController.GetFeePackage)
	router.PUT("/:id", fpc.feePackageController.UpdateFeePackage)
	router.DELETE("/:id", fpc.feePackageController.DeleteFeePackage)
}
