package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting users...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM users")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var users []User

	// Foreach movie
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		var isCoach bool

		err = rows.Scan(&id, &firstName, &lastName, &isCoach)

		// check errors
		checkErr(err)

		users = append(users, User{UserId: id, FirstName: firstName, LastName: lastName})
	}

	var response = UserResponse{Type: "success", Data: users}

	json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	params := mux.Vars(r)

	targetID := params["userid"]

	printMessage("Getting user...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM users WHERE id IS $1", targetID)

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var users []User

	// Foreach movie
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		var isCoach bool

		err = rows.Scan(&id, &firstName, &lastName, &isCoach)

		// check errors
		checkErr(err)

		users = append(users, User{FirstName: firstName, LastName: lastName, IsCoach: isCoach})
	}

	var response = UserResponse{Type: "success", Data: users}

	json.NewEncoder(w).Encode(response)
}

func GetTeamsByUser(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	params := mux.Vars(r)

	targetID := params["userid"]

	printMessage("Getting user...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM userteam WHERE userID IS $1", targetID)

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var teams []Team

	// Foreach movie
	for rows.Next() {
		var userID int
		var teamID int

		err = rows.Scan(&userID, &teamID)

		checkErr(err)

		teamRows, teamErr := db.Query("SELECT * FROM teams WHERE id IS $1", teamID)

		checkErr(teamErr)

		for teamRows.Next() {
			var id int
			var teamName string
			var description string
			var ownerID int

			err = teamRows.Scan(&id, &teamName, &description, &ownerID)

			// check errors
			checkErr(err)

			teams = append(teams, Team{TeamID: id, TeamName: teamName, TeamDescription: description, OwnerID: ownerID})
		}
	}

	var response = TeamResponse{Type: "success", Data: teams}

	json.NewEncoder(w).Encode(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	isCoach := r.FormValue("isCoach")

	var response = JsonResponse{}

	if firstName == "" || lastName == "" || isCoach == "" {
		response = JsonResponse{Type: "error", Message: "You are missing a creation parameter."}
	} else {
		db := setupDB()

		printMessage("Inserting user into DB")

		fmt.Println("Inserting new user with first name: " + firstName + " and last name: " + lastName)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO users(firstName, lastName, isCoach) VALUES($1, $2, $3) returning id;", firstName, lastName, isCoach).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The user has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID := params["userid"]

	var response = JsonResponse{}

	if userID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing userid parameter."}
	} else {
		db := setupDB()

		printMessage("Deleting user from DB")

		_, err := db.Exec("DELETE FROM users where id = $1", userID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The user has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Deleting all movies...")

	_, err := db.Exec("DELETE FROM users")

	// check errors
	checkErr(err)

	printMessage("All movies have been deleted successfully!")

	var response = JsonResponse{Type: "success", Message: "All users have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}

func JoinTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID := params["userid"]
	teamID := params["teamid"]

	var response = JsonResponse{}

	if userID == "" || teamID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing a creation parameter."}
	} else {
		db := setupDB()

		printMessage("Joining team")

		fmt.Println("Joining team")

		_, err := db.Exec("DELETE FROM userteam where userID = $1 and teamID = $2", userID, teamID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The user has joined the team successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func LeaveTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID := params["userid"]
	teamID := params["teamid"]

	var response = JsonResponse{}

	if userID == "" || teamID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing a creation parameter."}
	} else {
		db := setupDB()

		printMessage("Joining team")

		fmt.Println("Joining team")

		var lastInsertID int
		err := db.QueryRow("INSERT INTO userteam(userID, teamID) VALUES($1, $2) returning userID;", userID, teamID).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The user has joined the team successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
