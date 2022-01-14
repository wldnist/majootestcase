package handlers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wldnist/majootestcase/commons/objects"
	"github.com/wldnist/majootestcase/commons/responses"
	"github.com/wldnist/majootestcase/services"
)

type OutletHandler interface {
	All(ctx *gin.Context)
}

type outletHandler struct {
	outletService services.OutletService
	jwtService    services.JWTService
}

func NewOutletHandler(outletService services.OutletService, jwtService services.JWTService) OutletHandler {
	return &outletHandler{
		outletService: outletService,
		jwtService:    jwtService,
	}
}

func (c *outletHandler) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	outlets, err := c.outletService.All(userID)
	if err != nil {
		response := responses.BuildErrorResponse("Failed to process request", err.Error(), objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := responses.BuildResponse(true, "OK!", outlets)
	ctx.JSON(http.StatusOK, response)
}
