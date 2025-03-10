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

func (d *UseCase) GetByID(id uint) (*entity.Document, error) {
	return d.documentRepo.Find(id)
}

func (d *UseCase) Store(Document *entity.Document) error {
	return d.documentRepo.Save(Document)
}
