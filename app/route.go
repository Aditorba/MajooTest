package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"majooTest/config"
	"majooTest/controller"
	"majooTest/middleware"
	"majooTest/repository"
	"majooTest/service"
)

var (
	Config          = config.PopulateConfigData("local.yml")
	db     *gorm.DB = config.ConnectionDB(Config)

	// Repository
	userRepository        repository.UsersRepository       = repository.NewUsersRepository(db)
	merchantRepository    repository.MerchantRepository    = repository.NewMerchantRepository(db)
	outletRepository      repository.OutletRepository      = repository.NewOutletRepository(db)
	transactionRepository repository.TransactionRepository = repository.NewTransactionRepository(db)

	// Service
	authService   service.AuthService   = service.NewAuthService(Config, userRepository)
	reportService service.ReportService = service.NewReportService(userRepository, merchantRepository, outletRepository, transactionRepository)

	//Controller
	authController   controller.AuthController   = controller.NewAuthController(authService)
	reportController controller.ReportController = controller.NewReportController(authService, reportService)
)

func Route() {
	r := gin.Default()
	service := r.Group("")
	{
		service.GET("/")
		service.POST("/login", authController.Login)
	}

	reportsService := r.Group("report", middleware.AuthMiddleware)
	{
		reportsService.GET("/")
		reportsService.POST("/merchant", reportController.GetMerchantReport)
		reportsService.POST("/outlet", reportController.GetOutletReport)
	}

	err := r.Run()
	if err != nil {
		return
	}
}
