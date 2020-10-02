package test

type Request struct {
	Name string `json:"name"`
}

type AddRequest struct {
	Request
}

type AddResponse struct {
	ID string `json:"id"`
}

type UpdateRequest struct {
	Request
}

type UpdateResponse struct{}
