package routes

import (
	"github.com/Tsuzat/zipit/src/config"
	"github.com/Tsuzat/zipit/src/controller"
	"github.com/Tsuzat/zipit/src/middleware"
)

func InitAuthRouter() {
	group := config.APP.Group("/api/v1/auth")

	group.Post("/signup", controller.SignUpUser)
	group.Post("/login", controller.LoginUser)
	group.Get("/me", middleware.Authenticate, controller.Me)
	group.Get("/logout", controller.LogOut)
	group.Post("/refresh-access-token", controller.RefreshToken)
	group.Get("/verify", controller.VerifyUserEmail)
}
