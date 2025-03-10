package certificate_test

import (
	"digital-signature/certificate"
	"digital-signature/certificate/mocks"
	"digital-signature/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByID(t *testing.T) {
	oldTime, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z07:00")

	mockCertificateRepo := new(mocks.CertificateRepository)
	mockCertificate := &entity.Certificate{
		ID:        1,
		Name:      "IDAS CA DS G1",
		Issuer:    "Root CA Indonesia DS G1",
		ExpiresAt: oldTime,
	}

	mockCertificateRepo.
		On("Find", mock.AnythingOfType("uint")).
		Return(mockCertificate, nil).
		Once()

	certificateUseCase := certificate.NewUseCase(mockCertificateRepo)
	certificate, err := certificateUseCase.GetByID(mockCertificate.ID)

	assert.NoError(t, err)
	assert.NotNil(t, certificate)

	mockCertificateRepo.AssertExpectations(t)
}

func TestEnroll(t *testing.T) {

}

func TestRevoke(t *testing.T) {

}

func TestRenew(t *testing.T) {

}
