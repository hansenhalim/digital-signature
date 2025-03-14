package request

type EnrollCertificate struct {
	CertName   string `json:"cert_name" validate:"required"`
	CertIssuer string `json:"cert_issuer" validate:"required"`
}
