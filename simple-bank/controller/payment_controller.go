package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/mnc/middlewares"
	"github.com/rizkyfazri23/mnc/model/app_error"
	"github.com/rizkyfazri23/mnc/model/entity"
	"github.com/rizkyfazri23/mnc/usecase"
	"github.com/rizkyfazri23/mnc/utils"
)

type PaymentController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.PaymentUsecase
}

func NewPaymentController(r *gin.RouterGroup, u usecase.PaymentUsecase) *PaymentController {
	controller := PaymentController{
		router:  r,
		usecase: u,
	}
	payGroup := r.Group("/payment")
	payGroup.Use(middlewares.JwtAuthMiddleware())
	payGroup.POST("/", controller.Create)

	return &controller
}

func (c *PaymentController) Create(ctx *gin.Context) {
	var newPayment *entity.PaymentInfo

	sender_Id, err := utils.ExtractTokenID(ctx)
	if err := ctx.BindJSON(&newPayment); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		return
	}

	if newPayment.ReceiptUsername == "" {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("one or more required fields are missing"))
		return
	}
	if newPayment.Payment_Amount < 1 {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid amount"))
		return
	}

	res, err := c.usecase.TransferBalance(newPayment, sender_Id)

	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to transfer fund"))
		return
	}

	c.Success(ctx, http.StatusCreated, "01", "Successfully transfer fund", res)
}
