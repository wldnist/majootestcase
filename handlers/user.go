package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wldnist/majootestcase/commons/objects"
	"github.com/wldnist/majootestcase/commons/responses"
	"github.com/wldnist/majootestcase/dto"
	"github.com/wldnist/majootestcase/services"
)

type UserHandler interface {
	Profile(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userHandler struct {
	userService services.UserService
	jwtService  services.JWTService
}

func NewUserHandler(
	userService services.UserService,
	jwtService services.JWTService,
) UserHandler {
	return &userHandler{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userHandler) getUserIDByHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)

	if token == nil {
		response := responses.BuildErrorResponse("Error", "Failed to validate token", objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return ""
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

func (c *userHandler) Update(ctx *gin.Context) {
	var updateUserRequest dto.UpdateUserRequest

	err := ctx.ShouldBind(&updateUserRequest)
	if err != nil {
		response := responses.BuildErrorResponse("Failed to process request", err.Error(), objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id := c.getUserIDByHeader(ctx)

	if id == "" {
		response := responses.BuildErrorResponse("Error", "Failed to validate token", objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_id, _ := strconv.ParseInt(id, 0, 64)
	updateUserRequest.ID = _id
	res, err := c.userService.UpdateUser(updateUserRequest)

	if err != nil {
		response := responses.BuildErrorResponse("Error", err.Error(), objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	response := responses.BuildResponse(true, "OK", res)
	ctx.JSON(http.StatusOK, response)

}

func (c *userHandler) Profile(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)

	if token == nil {
		response := responses.BuildErrorResponse("Error", "Failed to validate token", objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user, err := c.userService.FindUserByID(id)

	if err != nil {
		response := responses.BuildErrorResponse("Error", err.Error(), objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	res := responses.BuildResponse(true, "OK", user)
	ctx.JSON(http.StatusOK, res)
}
