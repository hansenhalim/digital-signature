package signature

import "digital-signature/entity"

//go:generate mockery --name SignatureRepository
type SignatureRepository interface {
	Find(id uint) (*entity.Signature, error)
	Save(signature *entity.Signature) error
}

type UseCase struct {
	signatureRepo SignatureRepository
}

func NewUseCase(signatureRepo SignatureRepository) *UseCase {
	return &UseCase{signatureRepo}
}

func (s *UseCase) GetByID(id uint) (*entity.Signature, error) {
	return s.signatureRepo.Find(id)
}

func (s *UseCase) Store(signature *entity.Signature) error {
	return s.signatureRepo.Save(signature)
}
