package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting users...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM teams")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var teams []Team

	// Foreach movie
	for rows.Next() {
		var id int
		var teamName string
		var description string
		var ownerID int

		err = rows.Scan(&id, &teamName, &description, &ownerID)

		// check errors
		checkErr(err)

		teams = append(teams, Team{TeamID: id, TeamName: teamName, TeamDescription: description, OwnerID: ownerID})
	}

	var response = TeamResponse{Type: "success", Data: teams}

	json.NewEncoder(w).Encode(response)
}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	params := mux.Vars(r)

	teamID := params["teamID"]

	printMessage("Getting user...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM teams WHERE id IS $1", teamID)

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var teams []Team

	// Foreach movie
	for rows.Next() {
		var id int
		var teamName string
		var description string
		var ownerID int

		err = rows.Scan(&id, &teamName, &description, &ownerID)

		// check errors
		checkErr(err)

		teams = append(teams, Team{TeamID: id, TeamName: teamName, TeamDescription: description, OwnerID: ownerID})
	}

	var response = TeamResponse{Type: "success", Data: teams}

	json.NewEncoder(w).Encode(response)
}

func GetUsersByTeam(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	params := mux.Vars(r)

	targetID := params["teamID"]

	printMessage("Getting user...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM userteam WHERE teamID IS $1", targetID)

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var users []User

	// Foreach movie
	for rows.Next() {
		var userID int
		var teamID int

		err = rows.Scan(&userID, &teamID)

		checkErr(err)

		user_rows, user_err := db.Query("SELECT * FROM users WHERE id IS $1", userID)

		checkErr(user_err)
		for user_rows.Next() {
			var id int
			var firstName string
			var lastName string
			var isCoach bool

			err = rows.Scan(&id, &firstName, &lastName, &isCoach)

			users = append(users, User{UserId: id, FirstName: firstName, LastName: lastName, IsCoach: isCoach})

		}
	}

	var response = UserResponse{Type: "success", Data: users}

	json.NewEncoder(w).Encode(response)

}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	teamName := r.FormValue("teamName")
	description := r.FormValue("description")
	ownerID := r.FormValue("ownerID")

	var response = JsonResponse{}

	if teamName == "" || description == "" || ownerID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing a creation parameter."}
	} else {
		db := setupDB()

		printMessage("Inserting team into DB")

		fmt.Println("Inserting new team with name $1", teamName)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO teams(teamName, description, ownerID) VALUES($1, $2, $3) returning id;",
			teamName, description, ownerID).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The team has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	teamID := params["teamid"]

	var response = JsonResponse{}

	if teamID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing teamid parameter."}
	} else {
		db := setupDB()

		printMessage("Deleting team from DB")

		_, err := db.Exec("DELETE FROM teams where id = $1", teamID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The team has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteTeams(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Deleting all movies...")

	_, err := db.Exec("DELETE FROM teams")

	// check errors
	checkErr(err)

	printMessage("All teams have been deleted successfully!")

	var response = JsonResponse{Type: "success", Message: "All teams have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}
