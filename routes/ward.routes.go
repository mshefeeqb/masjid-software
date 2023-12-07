package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mshefeeqb/masjid-software/controllers"
	"github.com/mshefeeqb/masjid-software/middleware"
)

type WardRouteController struct {
	wardController controllers.WardController
}

func NewWardRouteController(wardController controllers.WardController) WardRouteController {
	return WardRouteController{wardController}
}

func (wc *WardRouteController) WardRoute(rs *gin.RouterGroup) {
	router := rs.Group("wards")
	router.Use(middleware.DeserializeUser())
	router.GET("/", wc.wardController.GetWards)
	router.POST("/", wc.wardController.CreateWard)
	router.GET("/:id", wc.wardController.GetWard)
	router.PUT("/:id", wc.wardController.UpdateWard)
	router.DELETE("/:id", wc.wardController.DeleteWard)

}
