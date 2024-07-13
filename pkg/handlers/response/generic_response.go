package response

type GenericResponse struct {
	Message *string `json:"message"`
	Error   *string `json:"error"`
}
