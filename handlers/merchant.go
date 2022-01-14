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

type MerchantHandler interface {
	FindMerchantByUserID(ctx *gin.Context)
}

type merchantHandler struct {
	merchantService services.MerchantService
	jwtService      services.JWTService
}

func NewMerchantHandler(merchantService services.MerchantService, jwtService services.JWTService) MerchantHandler {
	return &merchantHandler{
		merchantService: merchantService,
		jwtService:      jwtService,
	}
}

func (c *merchantHandler) FindMerchantByUserID(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	merchants, err := c.merchantService.FindMerchantByUserID(userID)
	if err != nil {
		response := responses.BuildErrorResponse("Failed to process request", err.Error(), objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := responses.BuildResponse(true, "OK!", merchants)
	ctx.JSON(http.StatusOK, response)
}
