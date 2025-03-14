package rest

import (
	"digital-signature/certificate"
	"digital-signature/entity"
	"digital-signature/impl/delivery/rest/request"
	"digital-signature/impl/delivery/rest/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message"`
}

type CertificateHandler struct {
	UseCase certificate.UseCase
}

func NewCertificateHandler(e *echo.Echo, uc certificate.UseCase) {
	handler := &CertificateHandler{
		UseCase: uc,
	}

	e.GET("/certificates/:id", handler.GetByID)
	e.POST("/certificates/enroll", handler.Enroll)
}

func (r *CertificateHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.ErrNotFound.Error())
	}

	id := uint(idP)

	certificate, err := r.UseCase.GetByID(id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, certificate)
}

func isRequestValid(m *request.EnrollCertificate) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *CertificateHandler) Enroll(c echo.Context) (err error) {
	var enrollRequest request.EnrollCertificate
	err = c.Bind(&enrollRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&enrollRequest); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	certificate := entity.Certificate{
		Name:   enrollRequest.CertName,
		Issuer: enrollRequest.CertIssuer,
	}

	err = r.UseCase.Enroll(&certificate)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	enrollResponse := response.EnrollCertificate{
		CertificateName:      certificate.Name,
		CertificateIssuer:    certificate.Issuer,
		CertificateExpiresAt: certificate.ExpiresAt,
	}

	return c.JSON(http.StatusCreated, enrollResponse)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case echo.ErrInternalServerError:
		return http.StatusInternalServerError
	case echo.ErrNotFound:
		return http.StatusNotFound
	case echo.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
