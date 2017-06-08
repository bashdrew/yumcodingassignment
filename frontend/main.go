package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	pb "bashdrew/yumcodingassignment/addrbookrestapi"

	"google.golang.org/grpc"

	"github.com/gocarina/gocsv"
	"github.com/gorilla/mux"
)

const (
	port           = ":8082"
	restAPIAddress = "localhost:50051"
)

const (
	colID        = "id"
	colFirstname = "firstname"
	colLastname  = "lastname"
	colEmail     = "email"
	colPhoneno   = "phoneno"
)

func getRestAPIConn() (conn *grpc.ClientConn, client pb.AddrBookRestAPIClient, err error) {
	conn, err = grpc.Dial(restAPIAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client = pb.NewAddrBookRestAPIClient(conn)

	return
}

func getPeopleList() (result *pb.PeopleReply) {
	// Set up a connection to the server.
	conn, c, err := getRestAPIConn()
	defer conn.Close()
	if err == nil {
		// Contact the server and print out its response.
		result, err = c.GetPeople(context.Background(), &pb.PersonRequest{Id: int64(-1)})
		if err != nil {
			log.Fatalf("could not get people: %v", err)
		}
	}

	return
}

// GetPeopleEndpoint ... get people
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	people := getPeopleList()

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
		id64, _ := strconv.ParseInt(params[colID], 10, 64)

		person, err = c.GetPerson(context.Background(), &pb.PersonRequest{Id: id64})
		if err != nil {
			log.Fatalf("could not get person: %v", err)
		}
	}

	json.NewEncoder(w).Encode(person)
}

// CreatePersonEndpoint ... post person
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person *pb.PersonReply
	// Set up a connection to the server.
	conn, c, err := getRestAPIConn()
	defer conn.Close()
	if err == nil {
		// Contact the server and print out its response.
		//		var personDtls common.PersonDetails

		params := mux.Vars(req)
		_ = json.NewDecoder(req.Body).Decode(&person)
		id64, _ := strconv.ParseInt(params[colID], 10, 64)

		person, err = c.PostPerson(context.Background(),
			&pb.PersonRequest{
				Id:        id64,
				Firstname: person.Firstname,
				Lastname:  person.Lastname,
				Email:     person.Email,
				Phoneno:   person.Phoneno,
			})
		if err != nil {
			log.Fatalf("could not post person: %v", err)
		}
	}

	json.NewEncoder(w).Encode(person)
}

// UpdatePersonEndpoint ... put person
func UpdatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person *pb.PersonReply
	// Set up a connection to the server.
	conn, c, err := getRestAPIConn()
	defer conn.Close()
	if err == nil {
		// Contact the server and print out its response.
		//		var personDtls common.PersonDetails

		params := mux.Vars(req)
		_ = json.NewDecoder(req.Body).Decode(&person)
		id64, _ := strconv.ParseInt(params[colID], 10, 64)

		person, err = c.PutPerson(context.Background(),
			&pb.PersonRequest{
				Id:        id64,
				Firstname: person.Firstname,
				Lastname:  person.Lastname,
				Email:     person.Email,
				Phoneno:   person.Phoneno,
			})
		if err != nil {
			log.Fatalf("could not put person: %v", err)
		}
	}

	json.NewEncoder(w).Encode(person)
}

// DeletePersonEndpoint ... delete person
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person *pb.PersonReply
	// Set up a connection to the server.
	conn, c, err := getRestAPIConn()
	defer conn.Close()
	if err == nil {
		// Contact the server and print out its response.
		params := mux.Vars(req)
		id64, _ := strconv.ParseInt(params[colID], 10, 64)

		person, err = c.DeletePerson(context.Background(), &pb.PersonRequest{Id: id64})
		if err != nil {
			log.Fatalf("could not get person: %v", err)
		}
	}

	json.NewEncoder(w).Encode(person)
}

// GetPeopleCSVEndpoint ... get people
func GetPeopleCSVEndpoint(w http.ResponseWriter, req *http.Request) {
	people := getPeopleList()

	gocsv.SetCSVWriter(func(out io.Writer) *csv.Writer {
		return csv.NewWriter(w)
	})
	_, errCSV := gocsv.MarshalBytes(people.People)
	if errCSV != nil {
		log.Fatalf("End point error: %v", errCSV)
	}
}

// PostPeopleCSVEndpoint ... post people
func PostPeopleCSVEndpoint(w http.ResponseWriter, req *http.Request) {
	people := new(pb.PeopleReply)

	gocsv.SetCSVWriter(func(out io.Writer) *csv.Writer {
		return csv.NewWriter(w)
	})
	errCSV := gocsv.Unmarshal(req.Body, &people.People)
	if errCSV != nil {
		log.Fatalf("End point error: %v", errCSV)
	} else {
		// Set up a connection to the server.
		conn, c, err := getRestAPIConn()
		defer conn.Close()
		if err == nil {
			// Contact the server and print out its response.
			people, err = c.PostPeople(context.Background(), people)
			if err != nil {
				log.Fatalf("could not post people: %v", err)
			}
		}
	}

	json.NewEncoder(w).Encode(people)
}

func setupRestAPIConn() (conn *grpc.ClientConn, err error) {
	// Set up a connection to the server
	conn, err = grpc.Dial(restAPIAddress, grpc.WithInsecure())

	return
}

func init() {
	fmt.Println("Frontend Microservice started...")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{"+colID+"}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{"+colID+"}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{"+colID+"}", UpdatePersonEndpoint).Methods("PUT")
	router.HandleFunc("/people/{"+colID+"}", DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/people/csv/download", GetPeopleCSVEndpoint).Methods("GET")
	router.HandleFunc("/people/csv/upload", PostPeopleCSVEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(port, router))
}
