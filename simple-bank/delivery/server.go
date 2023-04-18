package delivery

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/mnc/config"
	"github.com/rizkyfazri23/mnc/controller"
	"github.com/rizkyfazri23/mnc/manager"
)

type AppServer struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
	host           string
}

func (p *AppServer) v1() {
	v1Routes := p.engine.Group("/v1")
	p.userController(v1Routes)
	p.paymentController(v1Routes)
	p.depositController(v1Routes)
	p.historyController(v1Routes)
}

func (p *AppServer) userController(rg *gin.RouterGroup) {
	controller.NewUserController(rg, p.usecaseManager.UserUsecase())
}

func (p *AppServer) paymentController(rg *gin.RouterGroup) {
	controller.NewPaymentController(rg, p.usecaseManager.PaymentUsecase())
}

func (p *AppServer) depositController(rg *gin.RouterGroup) {
	controller.NewDepositController(rg, p.usecaseManager.DepositUsecase())
}

func (p *AppServer) historyController(rg *gin.RouterGroup) {
	controller.NewHistoryController(rg, p.usecaseManager.HistoryUsecase())
}

func (p *AppServer) Run() {
	p.v1()
	err := p.engine.Run(p.host)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application failed to run", err)
		}
	}()
	if err != nil {
		panic(err)
	}
}

func Server() *AppServer {
	r := gin.Default()
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	db := infraManager.DB()
	repoManager := manager.NewRepoManager(db)
	usecaseManager := manager.NewUsecaseManager(repoManager)
	host := fmt.Sprintf(":%s", c.ApiPort)
	return &AppServer{
		usecaseManager: usecaseManager,
		engine:         r,
		host:           host,
	}
}
