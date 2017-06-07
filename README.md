# yumcodingassignment

Packages used:
    gorilla/mux
    mattn/go-sqlite3
    gocarina/gocsv

Source files:
    dbconn/main.go - DB connections service
    frontend/main.go - Front-end service
    restapi/main.go - REST API service

Commands used:
    protoc -I addrbookdb/ addrbookdb/addrbookdb.proto --go_out=plugins=grpc:addrbookdb
    protoc -I addrbookrestapi/ addrbookrestapi/addrbookrestapi.proto --go_out=plugins=grpc:addrbookrestapi

Test commands:
    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST http://localhost:8082/people/1 --data '{"firstname":"John","lastname":"Smith","email":"john.smith@nomail.com","phoneno":"123-123-4233"}'

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8082/people

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8082/people/1

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X PUT http://localhost:8082/people/1 --data '{"firstname":"John","lastname":"Smith","email":"john.smith@nomail.com","phoneno":"123-123-4233"}'

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X DELETE http://localhost:8082/people/1
