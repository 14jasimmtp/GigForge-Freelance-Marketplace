package res

type CommonRes struct{
	Status string `json:"status"`
	Message string `json:"message"`
	Error string `json:"error"`
	Body any `json:"body"`
}