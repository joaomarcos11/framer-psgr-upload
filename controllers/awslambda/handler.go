package awslambda

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang-jwt/jwt/v5"

	"github.com/filipeandrade6/framer-psgr-upload/domain/service"
	"github.com/filipeandrade6/framer-psgr-upload/domain/usecases"
)

type Handler struct {
	log  *slog.Logger
	psgr usecases.Presigner
	msgr usecases.Messager
}

func New(log *slog.Logger, psgr usecases.Presigner, msgr usecases.Messager) *Handler {
	return &Handler{log: log, psgr: psgr, msgr: msgr}
}

func (hdlr *Handler) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	tokenString := request.Headers["Authorization"]

	claims := jwt.MapClaims{}

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		hdlr.log.Error("unauthorized", "parsing token", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
			Body:       fmt.Sprintf("Unauthorized: %v", err),
		}, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		hdlr.log.Error("unauthorized", "claims", "could not extract claims")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
			Body:       fmt.Sprintf("Unauthorized: %v", err),
		}, nil
	}

	email, ok := claims["email"].(string)
	if !ok {
		hdlr.log.Error("unauthorized", "claims", "could not find email")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Bad Request: Email claim not found",
		}, nil
	}

	url, err := service.Generate(email, hdlr.msgr, hdlr.psgr)
	if err != nil {
		hdlr.log.Error("processing", "generating pre-signed url for upload", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Internal Server Error: Generating pre-signed url for upload",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       url,
	}, nil
}
