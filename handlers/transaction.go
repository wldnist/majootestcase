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

type TransactionHandler interface {
	All(ctx *gin.Context)
}

type transactionHandler struct {
	transactionService services.TransactionService
	jwtService         services.JWTService
}

func NewTransactionHandler(transactionService services.TransactionService, jwtService services.JWTService) TransactionHandler {
	return &transactionHandler{
		transactionService: transactionService,
		jwtService:         jwtService,
	}
}

func (c *transactionHandler) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	transactions, err := c.transactionService.All(userID)
	if err != nil {
		response := responses.BuildErrorResponse("Failed to process request", err.Error(), objects.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := responses.BuildResponse(true, "OK!", transactions)
	ctx.JSON(http.StatusOK, response)
}
