package entity

type Document struct {
	ID        uint
	Name      string
	Content   string
	Signature Signature
}
