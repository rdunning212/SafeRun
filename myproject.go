package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var cache redis.Conn

func main() {
	initCache()
	router := mux.NewRouter()

	// Route handles & endpoints

	// User Endpoints
	router.HandleFunc("/users/", GetUsers).Methods("GET")
	router.HandleFunc("/users/{userid}", GetUser).Methods("GET")
	router.HandleFunc("/users/{userID}/teams", GetTeamsByUser).Methods("GET")
	router.HandleFunc("/users/", CreateUser).Methods("POST")
	router.HandleFunc("/users/{userid}/{teamid}", JoinTeam).Methods("POST")
	router.HandleFunc("/users/{userid}/{teamid}", LeaveTeam).Methods("DELETE")
	router.HandleFunc("/users/{userid}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/", DeleteUsers).Methods("DELETE")

	// Incident Endpoints
	router.HandleFunc("/incidents/", GetIncidents).Methods("GET")
	router.HandleFunc("/users/{incidentID}", GetIncident).Methods("GET")
	router.HandleFunc("/incidents/{reporterID}", GetIncidentsByUser).Methods("GET")
	router.HandleFunc("/incidents/", CreateIncident).Methods("POST")
	router.HandleFunc("/incidents/{incidentID}", DeleteIncident).Methods("DELETE")
	router.HandleFunc("/incidents/", DeleteIncidents).Methods("DELETE")

	// Team Endpoints
	router.HandleFunc("/teams/{teamID}/users", GetUsersByTeam).Methods("GET")
	router.HandleFunc("teams/", GetTeams).Methods("GET")
	router.HandleFunc("/teams/{teamid}", GetTeam).Methods("GET")
	router.HandleFunc("/teams/", CreateTeam).Methods("POST")
	router.HandleFunc("/teams/", DeleteTeams).Methods("DELETE")
	router.HandleFunc("/teams/{teamid}", DeleteTeam).Methods("DELETE")

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
