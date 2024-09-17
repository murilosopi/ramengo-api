package http

import (
	"database/sql"
	"ramengo/application/services"
	"ramengo/domain/repositories"
	"ramengo/infrastructure/controllers"
	"ramengo/infrastructure/db"
	dbRP "ramengo/infrastructure/db/repositories"
	"ramengo/infrastructure/middlewares"

	"github.com/labstack/echo/v4"
)

func Init() {

	var dbConnection *sql.DB = db.MySQLConnect()
	defer dbConnection.Close()

	// implement
	var (
		// initialize required repositories
		addressRepo repositories.AddressRepository = dbRP.NewSQLAddressRepository(dbConnection)
		userRepo    repositories.UserRepository    = dbRP.NewSQLUserRepository(dbConnection)
		orderRepo   repositories.OrderRepository   = dbRP.NewSQLOrderRepository(dbConnection)
		kitchenRepo repositories.KitchenRepository = dbRP.NewSQLKitchenRepository(dbConnection)
		authRepo    repositories.AuthRepository    = dbRP.NewSQLAuthRepository(dbConnection)

		// initialize required services
		addressService      services.AddressService      = services.NewAddressService(addressRepo)
		userService         services.UserService         = services.NewUserService(userRepo)
		notificationService services.NotificationService = services.NewLocalNotificationService()
		orderService        services.OrderService        = services.NewOrderService(orderRepo, notificationService)
		kitchenService      services.KitchenService      = services.NewKitchenService(kitchenRepo)
		authService         services.AuthService         = services.NewAuthService(authRepo)
	)

	userController := controllers.NewUserController(userService, addressService)
	orderController := controllers.NewOrderController(orderService)
	kitchenController := controllers.NewKitchenController(kitchenService, orderService)
	authController := controllers.NewAuthController(authService)

	e := echo.New()

	e.POST("/auth", authController.Login)

	e.POST("/user", userController.Save)
	e.GET("/user/orders", userController.OrderHistory, middlewares.UserJWTMiddleware)

	e.POST("/order", orderController.Save, middlewares.UserJWTMiddleware)
	e.PATCH("/order/status", orderController.ChangeStatus, middlewares.KitchenJWTMiddleware)

	e.POST("/kitchen/user", kitchenController.AddUser, middlewares.KitchenJWTMiddleware)
	e.GET("/kitchen/orders", kitchenController.GetCurrentOrders, middlewares.KitchenJWTMiddleware)
	e.DELETE("/kitchen/orders", kitchenController.CancelNotReadyOrders, middlewares.KitchenJWTMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}
