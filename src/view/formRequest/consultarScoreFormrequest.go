package formRequest

type ConsultarScoreFormRequest struct {
	DocumentoIdentidad string `json:"documento" validate:"required"`
}

type ConsultarEstadoFormRequest struct {
	UUID string `json:"uuid" validate:"required"`
}
