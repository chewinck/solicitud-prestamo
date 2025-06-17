package dto

type SolicitudPrestamoDto struct {
	UUID               string  `json:"uuid" `
	NombreCompleto     string  `json:"nombre" `
	DocumentoIdentidad string  `json:"documento" `
	MontoSolicitado    float64 `json:"monto" `
	Estado             string  `json:"estado" `
	Score              int     `json:"score,omitempty"`
	SimulacionPrestamoDto SimulacionPrestamoDto `json:"simulacion_prestamo,omitempty"`
}


type SimulacionPrestamoDto struct {
	MontoSolicitado float64 `json:"monto_solicitado" validate:"required,gt=0"`
	PlazoMeses      int     `json:"plazo_meses" validate:"required,gt=0"`
	TasaInteres     float64 `json:"tasa_interes" validate:"required,gt=0"`
	TipoCredito     string  `json:"tipo_credito" validate:"required"`      
}
