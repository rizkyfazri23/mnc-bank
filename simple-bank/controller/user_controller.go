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

type UserController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.UserUsecase
}

func NewUserController(r *gin.RouterGroup, u usecase.UserUsecase) *UserController {
	controller := UserController{
		router:  r,
		usecase: u,
	}
	usrGroup := r.Group("/user")
	usrGroup.Use(middlewares.JwtAuthMiddleware())
	usrGroup.PUT("/", controller.Edit)
	usrGroup.GET("/profile", controller.CurrentMember)
	usrGroup.POST("/logout", controller.Logout)

	r.POST("/register", controller.Add)
	r.POST("/login", controller.Login)

	return &controller
}

func (c *UserController) Add(ctx *gin.Context) {
	var user entity.User

	if err := ctx.BindJSON(&user); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		return
	}

	if user.Username == "" || user.Password == "" {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("one or more required fields are missing"))
		return
	}

	res, err := c.usecase.Add(&user)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to create user"))
		return
	}

	c.Success(ctx, http.StatusCreated, "01", "Successfully created new user", res)
}

func (c *UserController) Edit(ctx *gin.Context) {
	var user entity.User

	user_id, err := utils.ExtractTokenID(ctx)

	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError(err.Error()))
		return
	}

	if err := ctx.BindJSON(&user); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError("invalid request body"))
		return
	}

	res, err := c.usecase.Edit(&user, user_id)
	if err != nil {
		c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("user with id %d not found", user.User_Id)))
		return
	}

	c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully updated user with User_Id %d", user.User_Id), res)
}

func (c *UserController) Login(ctx *gin.Context) {
	var input entity.User

	if err := ctx.BindJSON(&input); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError("invalid request body"))
		return
	}

	token, err := c.usecase.LoginCheck(input.Username, input.Password)

	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError("Username or Password is Incorrect"))
		return
	}

	ctx.Header("Authorization", "Bearer "+token)

	c.Success(ctx, http.StatusOK, "", token, nil)
}

func (c *UserController) CurrentMember(ctx *gin.Context) {
	user_id, err := utils.ExtractTokenID(ctx)

	fmt.Println(user_id)

	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError(err.Error()))
		return
	}

	u, err := c.usecase.GetOne(int(user_id))

	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError(err.Error()))
		return
	}

	c.Success(ctx, http.StatusOK, "", "Successfully get data from User Login", u)

}

func (c *UserController) Logout(ctx *gin.Context) {
	user_id, err := utils.ExtractTokenID(ctx)
	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError(err.Error()))
		return
	}

	token, err := c.usecase.Logout(user_id)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError("failed to logout user"))
		return
	}

	c.Success(ctx, http.StatusOK, "", "Successfully logout user", token)
}
