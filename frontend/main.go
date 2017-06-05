package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "bashdrew/yumcodingassignment/yumaddressbook"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"
)

const (
	port           = ":8080"
	restAPIAddress = "localhost:50051"
)

const (
	colID        = "id"
	colFirstName = "firstname"
	colLastName  = "lastname"
	colAddress   = "address"
	colCity      = "city"
	colState     = "state"
)

type Person struct {
	ID        string   `json:"id,omitemty"`
	FirstName string   `json:"firstname,omitemty"`
	LastName  string   `json:"lastname,omitemty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func getRestAPIConn() (conn *grpc.ClientConn, client pb.AddrBookRestAPIClient, err error) {
	conn, err = grpc.Dial(restAPIAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client = pb.NewAddrBookRestAPIClient(conn)

	return
}

// GetPeopleEndpoint ... get person
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPersonEndpoint ... get person
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person *pb.PersonReply
	// Set up a connection to the server.
	conn, c, err := getRestAPIConn()
	defer conn.Close()
	if err == nil {
		// Contact the server and print out its response.
		params := mux.Vars(req)

		person, err = c.GetPerson(context.Background(), &pb.PersonRequest{Id: params[colID]})
		if err != nil {
			log.Fatalf("could not get person: %v", err)
		}
	}

	json.NewEncoder(w).Encode(person)
}

// CreatePersonEndpoint ... get person
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person

	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)

	json.NewEncoder(w).Encode(people)
}

// DeletePersonEndpoint ... get person
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for idx, item := range people {
		if item.ID == params["id"] {
			people = append(people[:idx], people[idx+1:]...)

			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func setupRestAPIConn() (conn *grpc.ClientConn, err error) {
	// Set up a connection to the server
	conn, err = grpc.Dial(restAPIAddress, grpc.WithInsecure())

	return
}
func main() {
	// initialize mock data
	people = append(people, Person{ID: "1", FirstName: "Andrew", LastName: "Amargo", Address: &Address{City: "Plano", State: "Texas"}})
	people = append(people, Person{ID: "2", FirstName: "Cari", LastName: "Amargo"})
	people = append(people, Person{ID: "3", FirstName: "Paulo", LastName: "Amargo", Address: &Address{City: "Plano", State: "TX"}})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{"+colID+"}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(port, router))
}
