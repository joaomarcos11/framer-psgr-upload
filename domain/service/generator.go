package service

import (
	"fmt"
	"log/slog"

	"github.com/filipeandrade6/framer-psgr-upload/domain/errors"
	"github.com/filipeandrade6/framer-psgr-upload/domain/usecases"
	"github.com/google/uuid"
)

func Generate(email string, msgrSvc usecases.Messager, psgrSvc usecases.Presigner) (string, error) {
	filename := uuid.New()
	url, err := psgrSvc.GeneratePreSignedUrl("fiap44-framer-videos", filename.String())
	if err != nil {
		err = fmt.Errorf("%w: %w", errors.ErrGeneratePreSignedURL, err)
		slog.Error("generate", "err", err)
		return "", err
	}

	err = msgrSvc.SendMessage("framer-status.fifo", fmt.Sprintf("%s.%s.%s", filename, "SOLICITADO", email))
	if err != nil {
		err = fmt.Errorf("%w: %w", errors.ErrSendMessage, err)
		slog.Error("generate", "err", err)
		return "", err
	}

	return url, nil
}
