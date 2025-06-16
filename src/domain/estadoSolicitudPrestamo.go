package domain

type EstadoSolicitudPrestamo interface {
	// CrearSolicitud crea una nueva solicitud de préstamo
	CrearSolicitud(solicitud SolicitudPrestamo) (SolicitudPrestamo, error)

}