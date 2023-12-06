package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mshefeeqb/masjid-software/controllers"
)

type MemberRouteController struct {
	memberController controllers.MemberController
}

func NewMemberRouteController(memberController controllers.MemberController) MemberRouteController {
	return MemberRouteController{memberController}
}

func (mc *MemberRouteController) MemberRoute(rs *gin.RouterGroup) {
	routes := rs.Group("members")

	routes.GET("/", mc.memberController.GetAllMembers)
}
