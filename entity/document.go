package entity

type Document struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	Signature Signature `json:"signature"`
}
