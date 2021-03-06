/*
 *
 * Copyright 2015, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"reflect"
	"strings"

	pbdb "bashdrew/yumcodingassignment/addrbookdb"

	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	dbName = "./addrbook.db"
	port   = ":50052"
)

var database *sql.DB

// server is used to implement AddrBookDB Methods.
type server struct{}

// getPersonAll implements AddrBookDB.ReadPersonDB
func getPeople() (people *pbdb.PeopleReplyDB, err error) {
	people = new(pbdb.PeopleReplyDB)

	qry := fmt.Sprintf("SELECT id, firstname, lastname, email, phoneno FROM people")
	peopleDB, errDB := database.Query(qry)
	defer peopleDB.Close()
	if errDB == nil {
		for peopleDB.Next() {
			person := new(pbdb.PersonReplyDB)
			peopleDB.Scan(&person.Id, &person.Firstname, &person.Lastname, &person.Email, &person.Phoneno)
			people.People = append(people.People, person)
		}
	} else {
		log.Printf("DB error: %v", errDB)
		people.Error = errDB.Error()
		err = errDB
	}

	return people, err
}

// getPersonByID implements AddrBookDB.ReadPersonDB
func getPersonByID(id int64) (person *pbdb.PersonReplyDB, err error) {
	person = new(pbdb.PersonReplyDB)

	qry := fmt.Sprintf("SELECT id, firstname, lastname, email, phoneno FROM people WHERE id = %v", id)
	personDB, errDB := database.Query(qry)
	defer personDB.Close()
	if errDB == nil {
		personDB.Next()
		personDB.Scan(&person.Id, &person.Firstname, &person.Lastname, &person.Email, &person.Phoneno)
	} else {
		log.Printf("DB error: %v", errDB)
		err = errDB
	}

	return person, err
}

func generateSetStatement(person *pbdb.PersonRequestDB) (result string) {
	updCols := []string{"id", "firstname", "lastname", "email", "phoneno"}

	v := reflect.ValueOf(*person)
	values := make([]interface{}, v.NumField())

	result = "SET "
	for i := 1; i < v.NumField(); i++ {
		if v.Field(i).Interface().(string) != "" {
			result = result + updCols[i] + " = '" + v.Field(i).Interface().(string) + "', "
			values[i] = v.Field(i).Interface()
		}
	}

	result = strings.TrimRight(result, ", ")

	return
}

// ReadPeopleDB implements AddrBookDB.ReadPeopleDB
func (s *server) ReadPeopleDB(ctx context.Context, in *pbdb.PersonRequestDB) (*pbdb.PeopleReplyDB, error) {
	people, _ := getPeople()

	return people, nil
}

// ReadPersonDB implements AddrBookDB.ReadPersonDB
func (s *server) ReadPersonDB(ctx context.Context, in *pbdb.PersonRequestDB) (*pbdb.PersonReplyDB, error) {
	person, _ := getPersonByID(in.Id)

	return person, nil
}

// CreatePersonDB implements AddrBookDB.CreatePersonDB
func (s *server) CreatePersonDB(ctx context.Context, in *pbdb.PersonRequestDB) (*pbdb.PersonReplyDB, error) {
	person := new(pbdb.PersonReplyDB)

	qry := fmt.Sprintf("INSERT INTO people (id, firstname, lastname, email, phoneno) VALUES (?, ?, ?, ?, ?)")
	statement, errDB := database.Prepare(qry)
	if errDB == nil {
		_, errSt := statement.Exec(in.Id, in.Firstname, in.Lastname, in.Email, in.Phoneno)

		if errSt == nil {
			person, _ = getPersonByID(in.Id)
		} else {
			person.Error = errSt.Error()
			log.Printf("DB error: %v", errSt)
		}
	} else {
		person.Error = errDB.Error()
		log.Printf("DB error: %v", errDB)
	}

	return person, nil
}

// UpdatePersonDB implements AddrBookDB.CreatePersonDB
func (s *server) UpdatePersonDB(ctx context.Context, in *pbdb.PersonRequestDB) (*pbdb.PersonReplyDB, error) {
	person := new(pbdb.PersonReplyDB)

	qrySet := generateSetStatement(in)
	qry := fmt.Sprintf("UPDATE people %v WHERE id = %v", qrySet, in.Id)
	statement, errDB := database.Prepare(qry)
	if errDB == nil {
		_, errSt := statement.Exec()

		if errSt == nil {
			person, _ = getPersonByID(in.Id)
		} else {
			person.Error = errSt.Error()
			log.Printf("DB error: %v", errSt)
		}
	} else {
		person.Error = errDB.Error()
		log.Printf("DB error: %v", errDB)
	}

	return person, nil
}

// UpdatePersonDB implements AddrBookDB.CreatePersonDB
func (s *server) DeletePersonDB(ctx context.Context, in *pbdb.PersonRequestDB) (*pbdb.PersonReplyDB, error) {
	person := new(pbdb.PersonReplyDB)

	qry := fmt.Sprintf("DELETE FROM people WHERE id = %v", in.Id)
	statement, errDB := database.Prepare(qry)
	if errDB == nil {
		_, errSt := statement.Exec()

		if errSt != nil {
			person.Error = errSt.Error()
			log.Printf("DB error: %v", errSt)
		}
	} else {
		person.Error = errDB.Error()
		log.Printf("DB error: %v", errDB)
	}

	return person, nil
}

func openAndCreateDB() (result bool) {
	var errDB error
	qry := "CREATE TABLE IF NOT EXISTS people (" +
		"id BIGINT PRIMARY KEY, " +
		"firstname TEXT, " +
		"lastname TEXT, " +
		"email TEXT, " +
		"phoneno TEXT " +
		")"

	database, errDB = sql.Open("sqlite3", dbName)
	if errDB == nil {
		statement, _ := database.Prepare(qry)
		statement.Exec()
		result = true
	} else {
		log.Fatalf("DB error: %v", errDB)
	}

	return
}

func init() {
	fmt.Println("DB Microservice started...")
}

func main() {
	if !openAndCreateDB() {
		return
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pbdb.RegisterAddrBookDBServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
