package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mshefeeqb/masjid-software/controllers"
	"github.com/mshefeeqb/masjid-software/middleware"
)

type MemberRouteController struct {
	memberController controllers.MemberController
}

func NewMemberRouteController(memberController controllers.MemberController) MemberRouteController {
	return MemberRouteController{memberController}
}

func (mc *MemberRouteController) MemberRoute(rs *gin.RouterGroup) {
	router := rs.Group("members")
	router.Use(middleware.DeserializeUser())
	router.GET("/", mc.memberController.GetAllMembers)
	router.POST("/", mc.memberController.CreateMember)
	router.GET("/:id", mc.memberController.GetMember)
	router.PUT("/:id", mc.memberController.UpdateMember)
	router.DELETE("/:id", mc.memberController.DeleteMember)

}
