package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wldnist/majootestcase/commons/objects"
	"github.com/wldnist/majootestcase/commons/responses"
	"github.com/wldnist/majootestcase/dto"
	"github.com/wldnist/majootestcase/services"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authHandler struct {
	authService services.AuthService
	jwtService  services.JWTService
	userService services.UserService
}

func NewAuthHandler(
	authService services.AuthService,
	jwtService services.JWTService,
	userService services.UserService,
) AuthHandler {
	return &authHandler{
		authService: authService,
		jwtService:  jwtService,
		userService: userService,
	}
}

func (c *authHandler) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	err := ctx.ShouldBind(&loginRequest)

	if err != nil {
		response := responses.BuildErrorResponse("Failed to process request", err.Error(), objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = c.authService.VerifyCredential(loginRequest.UserName, loginRequest.Password)
	if err != nil {
		response := responses.BuildErrorResponse("Failed to login", err.Error(), objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, _ := c.userService.FindUserByUserName(loginRequest.UserName)

	token := c.jwtService.GenerateToken(strconv.FormatInt(user.ID, 10))
	user.Token = token
	response := responses.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusOK, response)

}

func (c *authHandler) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest

	err := ctx.ShouldBind(&registerRequest)
	if err != nil {
		response := responses.BuildErrorResponse("Failed to process request", err.Error(), objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.userService.CreateUser(registerRequest)
	if err != nil {
		response := responses.BuildErrorResponse(err.Error(), err.Error(), objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	token := c.jwtService.GenerateToken(strconv.FormatInt(user.ID, 10))
	user.Token = token
	response := responses.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusCreated, response)
}
