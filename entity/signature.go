package entity

import "time"

type Signature struct {
	ID          uint        `json:"id"`
	Certificate Certificate `json:"certificate"`
	SignerName  string      `json:"signer_name"`
	SignedAt    time.Time   `json:"signed_at"`
}
