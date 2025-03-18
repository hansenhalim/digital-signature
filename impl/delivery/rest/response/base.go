package response

import "time"

type EnrollCertificate struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Issuer    string    `json:"issuer"`
	ExpiresAt time.Time `json:"expires_at"`
}

type GetCertificate struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Issuer    string    `json:"issuer"`
	ExpiresAt time.Time `json:"expires_at"`
}
