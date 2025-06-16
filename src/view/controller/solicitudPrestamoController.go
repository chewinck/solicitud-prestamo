package controller

import (
	"net/http"
	"solicitudPrestamo/src/dao"
	"solicitudPrestamo/src/domain"
	"solicitudPrestamo/src/infraestructure/util"
	usecase "solicitudPrestamo/src/useCase"
	"solicitudPrestamo/src/view/dto"
	formrequest "solicitudPrestamo/src/view/formRequest" // Replace with the actual path to the formrequest package

	"github.com/gin-gonic/gin"
)

var solicitudPrestamoRepository domain.SolicitudPrestamoRepository = dao.NewSqlLite()

func IniciarSolicitudPrestamo(c *gin.Context) {

	var req formrequest.IniciarSolicitudPrestamoFormRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	uuid := util.GenerateUUID()

	iniciarSolicitudUsecase := usecase.NewIniciarSolicitudUseCase(solicitudPrestamoRepository)

	solicitudPrestamoDto := dto.SolicitudPrestamoDto{
		UUID:               uuid,
		DocumentoIdentidad: req.DocumentoIdentidad,
		NombreCompleto:     req.NombreCompleto,
		MontoSolicitado:    req.MontoSolicitado,
		Estado:             "Solicitud de prestamo Iniciada",
	}

	go iniciarSolicitudUsecase.Execute(solicitudPrestamoDto)

	c.JSON(http.StatusOK, gin.H{"uuid": uuid})
}
