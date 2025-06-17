package domain

type IniciarSolicitud struct {
}

func NewIniciarSolicitud() *IniciarSolicitud {
	return &IniciarSolicitud{}
}

func (i *IniciarSolicitud) CambiarEstado(solicitudPrestamo *SolicitudPrestamo) {
	solicitudPrestamo.GetRepository().IniciarSolicitud(solicitudPrestamo.GetSolicitudPrestadoDto())
	// solicitudPrestamo.SetEstado(NewConsultarScore())
	// solicitudPrestamo.Avanzar()
}
