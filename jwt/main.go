package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thirasak-rol/jwt/src/controller"
	"github.com/thirasak-rol/jwt/src/middleware"
	"github.com/thirasak-rol/jwt/src/service"
)

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.Default()

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	v1 := server.Group("/v1")
	v1.Use(middleware.AuthorizeJWT())
	{
		v1.POST("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}

	port := "8080"
	server.Run(":" + port)

}
