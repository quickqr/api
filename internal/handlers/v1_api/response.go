package v1_api

type errorResponse struct {
	Message string `json:"message" example:"You messed up!"`
}

type httpError struct {
	Status  int
	Message string
}
