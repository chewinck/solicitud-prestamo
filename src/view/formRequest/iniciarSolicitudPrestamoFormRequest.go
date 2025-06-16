package formRequest

type IniciarSolicitudPrestamoFormRequest struct {
	NombreCompleto     string  `json:"nombre" validate:"required"`
	DocumentoIdentidad string  `json:"documento" validate:"required"`
	MontoSolicitado    float64 `json:"monto" validate:"required"`
}
