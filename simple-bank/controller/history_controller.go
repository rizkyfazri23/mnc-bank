package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/mnc/middlewares"
	"github.com/rizkyfazri23/mnc/model/app_error"
	"github.com/rizkyfazri23/mnc/usecase"
	"github.com/rizkyfazri23/mnc/utils"
)

type HistoryController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.HistoryUsecase
}

func NewHistoryController(r *gin.RouterGroup, u usecase.HistoryUsecase) *HistoryController {
	controller := HistoryController{
		router:  r,
		usecase: u,
	}

	hGroup := r.Group("/history")
	hGroup.Use(middlewares.JwtAuthMiddleware())

	hGroup.GET("/transaction", controller.GetAllTransaction)
	hGroup.GET("/auth", controller.GetAllAuth)

	return &controller
}

func (c *HistoryController) GetAllTransaction(ctx *gin.Context) {
	id, err := utils.ExtractTokenID(ctx)
	res, err := c.usecase.GetAllTransaction(id)

	fmt.Println(res)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(""))
		return
	}

	c.Success(ctx, http.StatusOK, "", "Successfully retrieved all history data", res)
}

func (c *HistoryController) GetAllAuth(ctx *gin.Context) {
	id, err := utils.ExtractTokenID(ctx)
	res, err := c.usecase.GetAllAuth(id)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(""))
		return
	}

	c.Success(ctx, http.StatusOK, "", "Successfully retrieved all payment data", res)
}
