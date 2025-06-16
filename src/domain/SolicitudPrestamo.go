package domain


type SolicitudPrestamo struct {
	ID              string `json:"id"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Email           string `json:"email"`
	Telefono        string `json:"telefono"`
	MontoSolicitado float64 `json:"monto_solicitado"`
	FechaSolicitud  string `json:"fecha_solicitud"`
	Estado          string `json:"estado"` // Ejemplo: "pendiente", "aprobado", "rechazado"
	Comentarios     string `json:"comentarios,omitempty"` // Comentarios adicionales, si los hay
	FechaAprobacion string `json:"fecha_aprobacion,omitempty"` // Fecha de aprobación, si aplica
	FechaRechazo    string `json:"fecha_rechazo,omitempty"` // Fecha de rechazo, si aplica
	MetodoPago      string `json:"metodo_pago,omitempty"` // Método de pago, si aplica
	MontoAprobado   float64 `json:"monto_aprobado,omitempty"` // Monto aprobado, si aplica
	Plazo           int    `json:"plazo,omitempty"` // Plazo del préstamo en meses, si aplica
	TasaInteres     float64 `json:"tasa_interes,omitempty"` // Tasa de interés del préstamo, si aplica				
}