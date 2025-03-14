package certificate

import (
	"digital-signature/entity"
	"time"
)

//go:generate mockery --name CertificateRepository
type CertificateRepository interface {
	Find(id uint) (certificate *entity.Certificate, err error)
	Save(certificate *entity.Certificate) (err error)
	Delete(certificate *entity.Certificate) (err error)
	Update(certificate *entity.Certificate) (err error)
}

//go:generate mockery --name CertificateAuthority
type CertificateAuthority interface {
	Status() (err error)
	Enroll() (err error)
	Revoke() (err error)
	Renew() (err error)
}

type UseCase struct {
	certRepo CertificateRepository
	certAuth CertificateAuthority
}

func NewUseCase(certRepo CertificateRepository, certAuth CertificateAuthority) *UseCase {
	return &UseCase{certRepo, certAuth}
}

func (c *UseCase) GetByID(id uint) (certificate *entity.Certificate, err error) {
	// Find Certificate in DB
	certificate, err = c.certRepo.Find(id)
	if err != nil {
		return nil, err
	}

	return certificate, nil
}

func (c *UseCase) Enroll(certificate *entity.Certificate) (err error) {
	c.certAuth.Enroll()

	certificate.ExpiresAt = time.Now().AddDate(1, 0, 0)

	// Persist Certificate to DB
	err = c.certRepo.Save(certificate)
	if err != nil {
		return err
	}

	return nil
}

func (c *UseCase) Revoke(certificate *entity.Certificate) (err error) {
	/* TODO: Call Revoke CA lib */

	// Delete Certificate from DB
	err = c.certRepo.Delete(certificate)
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
	err = c.certRepo.Update(oldCertificate)
	if err != nil {
		return err
	}

	// Persist new Certificate to DB
	err = c.certRepo.Save(newCertificate)
	if err != nil {
		return err
	}

	return nil
}
