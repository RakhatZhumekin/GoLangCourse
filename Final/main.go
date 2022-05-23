package main

import (
	"os"

	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/config"
	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/controller"
	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/middleware"
	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/repository"
	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService service.JWTService = service.NewJWTService()
	userService service.UserService = service.NewUserService(userRepository)
	authService service.AuthService = service.NewAuthService(userRepository)
	emailConfirmationService service.EmailConfirmationService = service.NewEmailConfirmationService()
	authController controller.AuthController = controller.NewAuthController(authService, jwtService, userService, emailConfirmationService)
	userController controller.UserController = controller.NewUserController(userService, jwtService, emailConfirmationService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/registration-confirm", authController.RegistrationConfirm)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}
	
	r.Run(":" + os.Getenv("PORT"))
}