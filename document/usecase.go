package document

import "digital-signature/entity"

//go:generate mockery --name DocumentRepository
type DocumentRepository interface {
	Find(id uint) (*entity.Document, error)
	Save(Document *entity.Document) error
}

type UseCase struct {
	documentRepo DocumentRepository
}

func NewUseCase(documentRepo DocumentRepository) *UseCase {
	return &UseCase{documentRepo}
}

func (uc *UseCase) GetByID(id uint) (*entity.Document, error) {
	return uc.documentRepo.Find(id)
}

func (uc *UseCase) Store(Document *entity.Document) error {
	return uc.documentRepo.Save(Document)
}
