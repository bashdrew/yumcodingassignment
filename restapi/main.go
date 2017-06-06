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
	"fmt"
	"log"
	"net"

	pbdb "bashdrew/yumcodingassignment/addrbookdb"
	pb "bashdrew/yumcodingassignment/addrbookrestapi"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port      = ":50051"
	dbAddress = "localhost:50052"
)

// server is used to implement AddrBookRestAPI.GetPerson.
type server struct{}

func getDBConn() (conn *grpc.ClientConn, client pbdb.AddrBookDBClient, err error) {
	conn, err = grpc.Dial(dbAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client = pbdb.NewAddrBookDBClient(conn)

	return
}

// GetPeople implements AddrBookRestAPI.GetPerson
func (s *server) GetPeople(ctx context.Context, in *pb.PersonRequest) (*pb.PeopleReply, error) {
	var peopleDB *pbdb.PeopleReplyDB
	// Set up a connection to the server.
	conn, c, err := getDBConn()
	defer conn.Close()

	if err == nil {
		// Contact the server and print out its response.
		peopleDB, err = c.ReadPeopleDB(context.Background(), &pbdb.PersonRequestDB{Id: in.Id})
		if err != nil {
			log.Fatalf("could not get person from DB: %v", err)
		}
	}

	people := new(pb.PeopleReply)
	for _, personDB := range peopleDB.People {
		person := &pb.PersonReply{
			Id:        personDB.Id,
			Firstname: personDB.Firstname,
			Lastname:  personDB.Lastname,
			Email:     personDB.Email,
			Phoneno:   personDB.Phoneno,
		}

		people.People = append(people.People, person)
	}
	return people, nil
}

// GetPerson implements AddrBookRestAPI.GetPerson
func (s *server) GetPerson(ctx context.Context, in *pb.PersonRequest) (*pb.PersonReply, error) {
	var person *pbdb.PersonReplyDB
	// Set up a connection to the server.
	conn, c, err := getDBConn()
	defer conn.Close()

	if err == nil {
		// Contact the server and print out its response.
		person, err = c.ReadPersonDB(context.Background(), &pbdb.PersonRequestDB{Id: in.Id})
		if err != nil {
			log.Fatalf("could not get person from DB: %v", err)
		}
	}

	return &pb.PersonReply{
		Id:        person.Id,
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Email:     person.Email,
		Phoneno:   person.Phoneno,
	}, nil
}

// PostPerson implements AddrBookRestAPI.PostPerson
func (s *server) PostPerson(ctx context.Context, in *pb.PersonReply) (*pb.PersonReply, error) {
	var person *pbdb.PersonReplyDB
	// Set up a connection to the server.
	conn, c, err := getDBConn()
	defer conn.Close()

	if err == nil {
		// Contact the server and print out its response.
		person, err = c.CreatePersonDB(context.Background(),
			&pbdb.PersonReplyDB{
				Id:        in.Id,
				Firstname: in.Firstname,
				Lastname:  in.Lastname,
				Email:     in.Email,
				Phoneno:   in.Phoneno,
			})
		if err != nil {
			log.Fatalf("could not post person to DB: %v", err)
		}
	}

	return &pb.PersonReply{
		Id:        person.Id,
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Email:     person.Email,
		Phoneno:   person.Phoneno,
	}, nil
}

// PutPerson implements AddrBookRestAPI.PutPerson
func (s *server) PutPerson(ctx context.Context, in *pb.PersonReply) (*pb.PersonReply, error) {
	var person *pbdb.PersonReplyDB
	// Set up a connection to the server.
	conn, c, err := getDBConn()
	defer conn.Close()

	if err == nil {
		// Contact the server and print out its response.
		person, err = c.UpdatePersonDB(context.Background(),
			&pbdb.PersonReplyDB{
				Id:        in.Id,
				Firstname: in.Firstname,
				Lastname:  in.Lastname,
				Email:     in.Email,
				Phoneno:   in.Phoneno,
			})
		if err != nil {
			log.Fatalf("could not put person to DB: %v", err)
		}
	}

	return &pb.PersonReply{
		Id:        person.Id,
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Email:     person.Email,
		Phoneno:   person.Phoneno,
	}, nil
}

// DeletePerson implements AddrBookRestAPI.DeletePerson
func (s *server) DeletePerson(ctx context.Context, in *pb.PersonRequest) (*pb.PersonReply, error) {
	var person *pbdb.PersonReplyDB
	// Set up a connection to the server.
	conn, c, err := getDBConn()
	defer conn.Close()

	if err == nil {
		// Contact the server and print out its response.
		person, err = c.DeletePersonDB(context.Background(), &pbdb.PersonRequestDB{Id: in.Id})
		if err != nil {
			log.Fatalf("could not get person from DB: %v", err)
		}
	}

	return &pb.PersonReply{
		Id:        person.Id,
		Firstname: person.Firstname,
		Lastname:  person.Lastname,
		Email:     person.Email,
		Phoneno:   person.Phoneno,
	}, nil
}

func init() {
	fmt.Println("REST API Microservice started...")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAddrBookRestAPIServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
