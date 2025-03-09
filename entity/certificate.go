package entity

import "time"

type Certificate struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Issuer    string    `json:"issuer"`
	ExpiresAt time.Time `json:"expires_at"`
}
