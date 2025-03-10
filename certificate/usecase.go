package certificate

import (
	"digital-signature/entity"
)

//go:generate mockery --name CertificateRepository
type CertificateRepository interface {
	Find(id uint) (certificate *entity.Certificate, err error)
	Save(certificate *entity.Certificate) (err error)
	Delete(certificate *entity.Certificate) (err error)
	Update(certificate *entity.Certificate) (err error)
}

type UseCase struct {
	certificateRepo CertificateRepository
}

func NewUseCase(certificateRepo CertificateRepository) *UseCase {
	return &UseCase{certificateRepo}
}

func (c *UseCase) GetByID(id uint) (certificate *entity.Certificate, err error) {
	// Find Certificate in DB
	certificate, err = c.certificateRepo.Find(id)
	if err != nil {
		return nil, err
	}

	return certificate, nil
}

func (c *UseCase) Enroll(certificate *entity.Certificate) (err error) {
	/* TODO: Call Enroll CA lib */

	// Persist Certificate to DB
	err = c.certificateRepo.Save(certificate)
	if err != nil {
		return err
	}

	return nil
}

func (c *UseCase) Revoke(certificate *entity.Certificate) (err error) {
	/* TODO: Call Revoke CA lib */

	// Delete Certificate from DB
	err = c.certificateRepo.Delete(certificate)
	if err != nil {
		return err
	}

	return nil
}

func (c *UseCase) Renew(certificate *entity.Certificate) (err error) {
	oldCertificate := certificate

	/* TODO: Call Renew CA lib */
	newCertificate := certificate

	// Update old Certificate in DB
	err = c.certificateRepo.Update(oldCertificate)
	if err != nil {
		return err
	}

	// Persist new Certificate to DB
	err = c.certificateRepo.Save(newCertificate)
	if err != nil {
		return err
	}

	return nil
}
