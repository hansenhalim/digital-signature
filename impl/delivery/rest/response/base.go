package response

import "time"

type EnrollCertificate struct {
	CertificateName      string    `json:"certificate_name"`
	CertificateIssuer    string    `json:"certificate_issuer"`
	CertificateExpiresAt time.Time `json:"certificate_expires_at"`
}
