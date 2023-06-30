package delivery

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rafialariq/digital-bank/config"
	"github.com/rafialariq/digital-bank/controller"
	"github.com/rafialariq/digital-bank/manager"
	"github.com/rafialariq/digital-bank/middlewares"
)

type AppServer struct {
	serviceManager manager.ServiceManager
	engine         *gin.Engine
	host           string
}

func (s *AppServer) menu() {
	routes := s.engine.Group("/")
	// routes.Use(middlewares.LogMiddleware())
	routes.Use(middlewares.LogMiddleware())
	menu := routes.Group("/")
	menu.Use(middlewares.AuthMiddleware())
	s.registerController(routes)
	s.loginController(routes)
	s.paymentController(menu)
	s.logoutController(menu)

}

func (s *AppServer) registerController(r *gin.RouterGroup) {
	controller.NewRegisterController(r, s.serviceManager.RegisterService())
}

func (s *AppServer) loginController(r *gin.RouterGroup) {
	controller.NewLoginController(r, s.serviceManager.LoginService())
}

func (s *AppServer) paymentController(r *gin.RouterGroup) {
	controller.NewPaymentController(r, s.serviceManager.PaymentService())
}

func (s *AppServer) logoutController(r *gin.RouterGroup) {
	controller.NewLogoutController(r)
}

func (s *AppServer) Run() {
	s.menu()
	err := s.engine.Run(s.host)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application failed to run", err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
}

func Server() *AppServer {
	router := gin.Default()
	config := config.NewConfig()
	infraManager := manager.NewInfraManager(config)
	repoManager := manager.NewRepoManager(infraManager)
	serviceManager := manager.NewUsecaseManager(repoManager)
	host := config.ServerPort
	return &AppServer{
		serviceManager: serviceManager,
		engine:         router,
		host:           host,
	}
}
