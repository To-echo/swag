package model

type MyStruct struct {
	Name string `json:"child_name"`
}

type MyPayload struct {
	My   MyStruct
	Name string `json:"name2"`
}

type ErrorsResponse struct {
	NewTime MyPayload
}
