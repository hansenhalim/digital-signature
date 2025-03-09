package certificate

import (
	"digital-signature/entity"
)

//go:generate mockery --name CertificateRepository
type CertificateRepository interface {
	Find(id uint) (*entity.Certificate, error)
	Save(certificate *entity.Certificate) error
	Delete(certificate *entity.Certificate) error
	Update(certificate *entity.Certificate) error
}

type UseCase struct {
	certificateRepo CertificateRepository
}

func NewUseCase(certificateRepo CertificateRepository) *UseCase {
	return &UseCase{certificateRepo}
}

func (uc *UseCase) GetByID(id uint) (*entity.Certificate, error) {
	return uc.certificateRepo.Find(id)
}

func (uc *UseCase) Enroll(certificate *entity.Certificate) error {
	return uc.certificateRepo.Save(certificate)
}

func (uc *UseCase) Revoke(certificate *entity.Certificate) error {
	return uc.certificateRepo.Delete(certificate)
}

func (uc *UseCase) Renew(certificate *entity.Certificate) error {
	return uc.certificateRepo.Update(certificate)
}
