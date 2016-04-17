package handlers

type JsonResponse struct {
	Ok    bool        `json:"ok"`
	Data  interface{} `json:"data"`
	Error string      `json:"error,omitempty"`
}
