package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wldnist/majootestcase/configs"
	"github.com/wldnist/majootestcase/handlers"
	middleware "github.com/wldnist/majootestcase/middlewares"
	"github.com/wldnist/majootestcase/repositories"
	"github.com/wldnist/majootestcase/services"
)

var (
	db                    = configs.SetupDBConnection()
	userRepository        = repositories.NewUserRepository(db)
	merchantRepository    = repositories.NewMerchantRepository(db)
	outletRepository      = repositories.NewOutletRepository(db)
	transactionRepository = repositories.NewTransactionRepository(db)
	authService           = services.NewAuthService(userRepository)
	jwtService            = services.NewJWTService()
	userService           = services.NewUserService(userRepository)
	merchantService       = services.NewMerchantService(merchantRepository)
	outletService         = services.NewOutletService(outletRepository)
	transactionService    = services.NewTransactionService(transactionRepository)
	authHandler           = handlers.NewAuthHandler(authService, jwtService, userService)
	userHandler           = handlers.NewUserHandler(userService, jwtService)
	merchantHandler       = handlers.NewMerchantHandler(merchantService, jwtService)
	outletHandler         = handlers.NewOutletHandler(outletService, jwtService)
	transactionHandler    = handlers.NewTransactionHandler(transactionService, jwtService)
)

func main() {
	defer configs.CloseDBConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	merchantRoutes := server.Group("api/merchant", middleware.AuthorizeJWT(jwtService))
	{
		merchantRoutes.GET("/", merchantHandler.FindMerchantByUserID)
	}

	outletRoutes := server.Group("api/outlet", middleware.AuthorizeJWT(jwtService))
	{
		outletRoutes.GET("/", outletHandler.All)
	}

	transactionRoutes := server.Group("api/transaction", middleware.AuthorizeJWT(jwtService))
	{
		transactionRoutes.GET("/", transactionHandler.All)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", handlers.Health)
	}

	server.Run()
}
