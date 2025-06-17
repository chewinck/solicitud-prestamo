package domain

type EstadoSolicitudPrestamo interface {
	CambiarEstado(solicitudPrestamo *SolicitudPrestamo)
}
