package awslambda

import (
	"log/slog"

	"github.com/filipeandrade6/framer-psgr-upload/domain/usecases"

	"github.com/aws/aws-lambda-go/lambda"
)

func Start(log *slog.Logger, psgr usecases.Presigner, msgr usecases.Messager) {
	lambdaHndlr := New(log, psgr, msgr)
	lambda.Start(lambdaHndlr.Handler)
}
