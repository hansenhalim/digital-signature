package postgresql

import (
	"database/sql"
	"digital-signature/entity"
)

type CertificateRepository struct {
	Conn *sql.DB
}

func NewCertificateRepository(conn *sql.DB) *CertificateRepository {
	return &CertificateRepository{conn}
}

func (p *CertificateRepository) Find(id uint) (certificate *entity.Certificate, err error) {
	query := `SELECT id, name, issuer, expires_at FROM certificates WHERE id = $1`

	// Initialize the Certificate struct
	certificate = &entity.Certificate{}

	// Execute the query and scan the result into the Certificate struct
	err = p.Conn.QueryRow(query, id).Scan(&certificate.ID, &certificate.Name, &certificate.Issuer, &certificate.ExpiresAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case where no rows were returned
			return nil, nil
		}
		// Handle other errors
		return nil, err
	}

	return certificate, nil
}

func (p *CertificateRepository) Save(certificate *entity.Certificate) (err error) {
	query := `INSERT INTO certificates (name, issuer, expires_at) VALUES ($1, $2, $3) RETURNING id`

	// Execute the query and scan the returned ID into the certificate struct
	err = p.Conn.QueryRow(query, certificate.Name, certificate.Issuer, certificate.ExpiresAt).Scan(&certificate.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *CertificateRepository) Delete(certificate *entity.Certificate) (err error) {
	query := `DELETE FROM certificates WHERE id = $1`

	// Execute the delete query
	_, err = p.Conn.Exec(query, certificate.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *CertificateRepository) Update(certificate *entity.Certificate) (err error) {
	query := `UPDATE certificates SET name = $1, issuer = $2, expires_at = $3 WHERE id = $4`

	// Execute the update query
	_, err = p.Conn.Exec(query, certificate.Name, certificate.Issuer, certificate.ExpiresAt, certificate.ID)
	if err != nil {
		return err
	}

	return nil
}
