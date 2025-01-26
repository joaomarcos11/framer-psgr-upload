package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/filipeandrade6/framer-psgr-upload/adapters/message/awssqs"
	"github.com/filipeandrade6/framer-psgr-upload/adapters/presigner/awss3"
	"github.com/filipeandrade6/framer-psgr-upload/controllers/awslambda"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	presigner, err := awss3.New()
	if err != nil {
		log.Fatalf("failed to configure pre-signer: %s", err)
	}

	messager, err := awssqs.New()
	if err != nil {
		log.Fatalf("failed to configure messager: %s", err)
	}

	awslambda.Start(logger, presigner, messager)
}
