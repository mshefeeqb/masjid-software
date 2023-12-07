package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mshefeeqb/masjid-software/controllers"
	"github.com/mshefeeqb/masjid-software/initializers"
	"github.com/mshefeeqb/masjid-software/routes"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	MemberController      controllers.MemberController
	MemberRouteController routes.MemberRouteController

	FeePackageController      controllers.FeePackageController
	FeePackageRouteController routes.FeePackageRouteController

	WardController      controllers.WardController
	WardRouteController routes.WardRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	MemberController = controllers.NewMemberController(initializers.DB)
	MemberRouteController = routes.NewMemberRouteController(MemberController)

	FeePackageController = controllers.NewFeePackageController(initializers.DB)
	FeePackageRouteController = routes.NewFeePackageRouteController(FeePackageController)

	WardController = controllers.NewWardController(initializers.DB)
	WardRouteController = routes.NewWardRouteController(WardController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	MemberRouteController.MemberRoute(router)
	FeePackageRouteController.FeePackageRoute(router)
	WardRouteController.WardRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
