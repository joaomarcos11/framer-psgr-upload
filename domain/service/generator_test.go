package service

import (
	"testing"

	"github.com/filipeandrade6/framer-psgr-upload/domain/errors"
	"github.com/stretchr/testify/assert"
)

type mockMessagerSuccess struct{}

func NewMockMessagerSuccess() *mockMessagerSuccess {
	return &mockMessagerSuccess{}
}

func (mockMessagerSuccess) SendMessage(queue, message string) error {
	return nil
}

type mockMessagerFail struct{}

func NewMockMessagerFail() *mockMessagerFail {
	return &mockMessagerFail{}
}

func (mockMessagerFail) SendMessage(queue, message string) error {
	return errors.ErrSendMessage
}

type mockPresignerSuccess struct{}

func NewMockPresignerSuccess() *mockPresignerSuccess {
	return &mockPresignerSuccess{}
}

func (mockPresignerSuccess) GeneratePreSignedUrl(bucket, filename string) (string, error) {
	return "http://success.com", nil
}

type mockPresignerFail struct{}

func NewMockPresignerFail() *mockPresignerFail {
	return &mockPresignerFail{}
}

func (mockPresignerFail) GeneratePreSignedUrl(bucket, filename string) (string, error) {
	return "", errors.ErrGeneratePreSignedURL
}

func TestGenerate(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		msgr := NewMockMessagerSuccess()
		psgr := NewMockPresignerSuccess()
		_, err := Generate("email@email.com", msgr, psgr)

		assert.NoError(t, err)
	})

	t.Run("fail to generate pre-signed url", func(t *testing.T) {
		msgr := NewMockMessagerSuccess()
		psgr := NewMockPresignerFail()
		_, err := Generate("email@email.com", msgr, psgr)

		assert.Error(t, err)
		assert.ErrorAs(t, err, &errors.ErrGeneratePreSignedURL)
	})

	t.Run("fail to send message", func(t *testing.T) {
		msgr := NewMockMessagerFail()
		psgr := NewMockPresignerSuccess()
		_, err := Generate("email@email.com", msgr, psgr)

		assert.Error(t, err)
		assert.ErrorAs(t, err, &errors.ErrSendMessage)
	})
}
