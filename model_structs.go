package main

type Movie struct {
	MovieID   string `json:"movieid"`
	MovieName string `json:"moviename"`
}

type User struct {
	UserId    int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastname"`
	IsCoach   bool   `json:"isCoach"`
}

type Team struct {
	TeamID          int    `json:"teamID"`
	TeamName        string `json:"teamName"`
	OwnerID         int    `json:"ownerID"`
	TeamDescription string `json:"teamDescription"`
}

type Incident struct {
	IncidentID   int     `json:"incident_id"`
	ReporterID   int     `json:"reporter_id"`
	IncidentDate string  `json:"incident_date"`
	IncidentTime string  `json:"incident_time"`
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
	IncidentType string  `json:"incident_type"`
	Description  string  `json:"description"`
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
