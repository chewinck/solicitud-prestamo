package dto

type SolicitudPrestamoDto struct {
	UUID               string  `json:"uuid" `
	NombreCompleto     string  `json:"nombre" `
	DocumentoIdentidad string  `json:"documento" `
	MontoSolicitado    float64 `json:"monto" `
	Estado             string  `json:"estado" `
	Score              int     `json:"score,omitempty"` // Optional field for score
}
