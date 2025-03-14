package emudhra

type CertificateAuthority struct {
}

func NewCertificateAuthority() *CertificateAuthority {
	return &CertificateAuthority{}
}

func (e *CertificateAuthority) Status() (err error) {
	/* TODO: curl to emudhra server */

	return nil
}

func (e *CertificateAuthority) Enroll() (err error) {
	/* TODO: curl to emudhra server */

	return nil
}

func (e *CertificateAuthority) Revoke() (err error) {
	/* TODO: curl to emudhra server */

	return nil
}

func (e *CertificateAuthority) Renew() (err error) {
	/* TODO: curl to emudhra server */

	return nil
}
