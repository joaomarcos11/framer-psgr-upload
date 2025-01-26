package service

import (
	"errors"
	"fmt"

	"github.com/filipeandrade6/framer-psgr-upload/domain/usecases"
	"github.com/google/uuid"
)

func Generate(email string, msgrSvc usecases.Messager, psgrSvc usecases.Presigner) (string, error) {
	filename := uuid.New()
	url, err := psgrSvc.GeneratePreSignedUrl("fiap44-framer-videos", filename.String())
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to generate pre-signed url to upload: %s", err))
	}

	err = msgrSvc.SendMessage("framer-status.fifo", fmt.Sprintf("%s.%s.%s", filename, "solicitado", email))
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to send message: %s", err))
	}

	return url, nil
}
