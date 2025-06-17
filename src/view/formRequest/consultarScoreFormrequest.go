package formRequest

type ConsultarScoreFormRequest struct {
	UUID string `json:"uuid" validate:"required,uuid"`
}
