# yumcodingassignment

Packages used:
    gorilla/mux
    mattn/go-sqlite3
    gocarina/gocsv

Source files:
    dbconn/main.go - DB connections service
    frontend/main.go - Front-end service
    restapi/main.go - REST API service/Business Logic

Commands used:
    protoc -I addrbookdb/ addrbookdb/addrbookdb.proto --go_out=plugins=grpc:addrbookdb
    protoc -I addrbookrestapi/ addrbookrestapi/addrbookrestapi.proto --go_out=plugins=grpc:addrbookrestapi

Test commands:
    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST http://localhost:8082/people/1 --data '{"firstname":"John","lastname":"Smith","email":"john.smith@nomail.com","phoneno":"123-123-4233"}'

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8082/people

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8082/people/1

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X PUT http://localhost:8082/people/1 --data '{"firstname":"John","lastname":"Smith","email":"john.smith@nomail.com","phoneno":"123-123-4233"}'

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X DELETE http://localhost:8082/people/1

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST http://localhost:8082/people/csv/upload --data '
    Id,Firstname,Lastname,Email,Phoneno,Error
    2,John,Smith,john.smith@nomail.com,123-123-4233,
    3,Mark,Smith,mark.smith@nomail.com,123-123-4233,
    10,Luke,Smith,luke.smith@nomail.com,123-123-4233,
    1,Matthew,Smith,matthew.smith@nomail.com,123-123-4233,'

    curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8082/people/csv/download

Mock Test:

    go get github.com/golang/mock/gomock
    go get github.com/golang/mock/mockgen
    
    mockgen bashdrew/yumcodingassignment/addrbookrestapi AddrBookRestAPIClient > mock_addrbookrestapi/addrbookrestapi_mock.go

    mockgen bashdrew/yumcodingassignment/addrbookdb AddrBookDBClient > mock_addrbookdb/addrbookdb_mock.go