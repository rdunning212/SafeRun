package main

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Movie `json:"data"`
	Message string  `json:"message"`
}

type UserResponse struct {
	Type    string `json:"type"`
	Data    []User `json:"data"`
	Message string `json:"message"`
}

type IncidentResponse struct {
	Type    string     `json:"type"`
	Data    []Incident `json:"data"`
	Message string     `json:"message"`
}

type TeamResponse struct {
	Type    string `json:"type"`
	Data    []Team `json:"data"`
	Message string `json:"message"`
}
