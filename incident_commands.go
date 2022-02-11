package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetIncidents(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting incidents...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM incidents")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var incidents []Incident

	// Foreach movie
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

	json.NewEncoder(w).Encode(response)
}

func GetIncident(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	params := mux.Vars(r)

	targetID := params["incidentid"]

	printMessage("Getting incident...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM incidents WHERE id IS" + string(targetID))

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var incidents []Incident

	// Foreach movie
	for rows.Next() {
		var id int
		var reporterID int
		var incidentDate string
		var incidentTime string
		var latitude float32
		var longitude float32
		var incidentType string
		var description string

		err = rows.Scan(
			&id,
			&reporterID,
			&incidentDate,
			&incidentTime,
			&latitude,
			&longitude,
			&incidentType,
			&description)

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
			Description:  description,
		})
	}

	var response = IncidentResponse{Type: "success", Data: incidents}

	json.NewEncoder(w).Encode(response)
}

func GetIncidentsByUser(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	params := mux.Vars(r)

	reporterID := params["reporterID"]

	printMessage("Getting incident...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM incidents WHERE reporterID IS" + string(reporterID))

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var incidents []Incident

	// Foreach movie
	for rows.Next() {
		var id int
		var reporterID int
		var incidentDate string
		var incidentTime string
		var latitude float32
		var longitude float32
		var incidentType string
		var description string

		err = rows.Scan(
			&id,
			&reporterID,
			&incidentDate,
			&incidentTime,
			&latitude,
			&longitude,
			&incidentType,
			&description)

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
			Description:  description,
		})
	}

	var response = IncidentResponse{Type: "success", Data: incidents}

	json.NewEncoder(w).Encode(response)
}

func CreateIncident(w http.ResponseWriter, r *http.Request) {
	reporterID := r.FormValue("reporterID")
	incidentDate := r.FormValue("incidentDate")
	incidentTime := r.FormValue("incidentTime")
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")
	incidentType := r.FormValue("incidentType")
	description := r.FormValue("description")

	var response = JsonResponse{}

	if reporterID == "" || incidentDate == "" || incidentTime == "" || latitude == "" || longitude == "" || incidentType == "" || description == "" {
		response = JsonResponse{Type: "error", Message: "You are missing a creation parameter."}
	} else {
		db := setupDB()

		printMessage("Inserting user into DB")

		fmt.Println("Inserting new incident")

		var lastInsertID int
		err := db.QueryRow("INSERT INTO incidents(reporterID, incidentDate, incidentTime, latitude, longitude"+
			"incidentType, description) VALUES($1, $2, $3, $4, $5, $6, $7) returning id;",
			reporterID, incidentDate, incidentTime, latitude, longitude, incidentType, description).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The user has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteIncident(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	incidentID := params["incidentID"]

	var response = JsonResponse{}

	if incidentID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing incidentid parameter."}
	} else {
		db := setupDB()

		printMessage("Deleting incident from DB")

		_, err := db.Exec("DELETE FROM incidents where id = $1", incidentID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The user has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteIncidents(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Deleting all incidents...")

	_, err := db.Exec("DELETE FROM incidents")

	// check errors
	checkErr(err)

	printMessage("All incidents have been deleted successfully!")

	var response = JsonResponse{Type: "success", Message: "All incidents have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}
