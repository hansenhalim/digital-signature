package main

import (
	"database/sql"
	"digital-signature/certificate"
	"digital-signature/impl/delivery/rest"
	"digital-signature/impl/delivery/rest/middleware"
	"digital-signature/impl/lib/emudhra"
	"digital-signature/impl/repository/pgsql"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

const (
	defaultAddress = ":8080"
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
	dbConn, err := sql.Open("postgres", dsn)
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

	// prepare echo
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Timeout())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Prepare Repository
	certRepo := pgsql.NewCertificateRepository(dbConn)

	// Prepare libs
	certAuth := emudhra.NewCertificateAuthority()

	// Build usecase Layer
	certificateUseCase := certificate.NewUseCase(certRepo, certAuth)
	rest.NewCertificateHandler(e, *certificateUseCase)

	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address)) //nolint
}
