package entity

import "time"

type Signature struct {
	ID          uint
	Certificate Certificate
	SignerName  string
	SignedAt    time.Time
}
