package model

type MyStruct struct {
	Name string `json:"child_name"`
}

type MyPayload struct {
	My   MyStruct
	Name string `json:"name1"`
}

type ErrorsResponse struct {
	NewTime MyPayload
}
