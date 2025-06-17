package formRequest

type VerificarIdentidadFormRequest struct {
	UUID               string `json:"uuid" validate:"required,uuid"`
	DocumentoIdentidad string `json:"documento" validate:"required"`
	NombreCompleto     string `json:"nombre" validate:"required"`
}
