package main

func get_nearby_incidents(target []int, dist []int) IncidentResponse {
	db := setupDB()

	printMessage("Getting user...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM incidents WHERE $1 < latitude AND latitude < $2 AND $3 < longitude and longitude < $4")

	var incidents []Incident

	// Foreach incident
	for rows.Next() {
		var id int
		var reporterID int
		var incidentDate string
		var incidentTime string
		var latitude float32
		var longitude float32
		var incidentType string

		err = rows.Scan(&id, &reporterID, &incidentDate, &incidentTime, &latitude, &longitude, &incidentType)

		// check errors
		checkErr(err)

		incidents = append(incidents, Incident{
			IncidentID:   id,
			ReporterID:   reporterID,
			IncidentDate: incidentDate,
			IncidentTime: incidentTime,
			Latitude:     latitude,
			Longitude:    longitude,
			IncidentType: incidentType,
		})
	}

	var response = IncidentResponse{Type: "success", Data: incidents}

	return response
}
