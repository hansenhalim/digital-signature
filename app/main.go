package main

import (
	"database/sql"
	"digital-signature/certificate"
	"digital-signature/entity"
	pgsqlRepo "digital-signature/internal/repository/pgsql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	dbConn, err := sql.Open(`postgres`, dsn)
	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("failed to ping database ", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("got error when closing the DB connection", err)
		}
	}()

	certificateRepo := pgsqlRepo.NewCertificateRepository(dbConn)
	certificateUseCase := certificate.NewUseCase(certificateRepo)

	certificateUseCase.Enroll(&entity.Certificate{
		Name:      "IDAS CA DS G1",
		Issuer:    "Root CA Indonesia DS G1",
		ExpiresAt: time.Now().AddDate(1, 0, 0),
	})

	log.Println("Certificate Enrolled successfully!")
}
