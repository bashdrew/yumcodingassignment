# yumcodingassignment

protoc -I addrbookdb/ addrbookdb/addrbookdb.proto --go_out=plugins=grpc:addrbookdb
protoc -I addrbookrestapi/ addrbookrestapi/addrbookrestapi.proto --go_out=plugins=grpc:addrbookrestapi

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X POST http://localhost:8080/people/2 --data '{"firstname":"John","lastname":"Smith","email":"john.smith@nomail.com","phoneno":"123-123-4233"}'

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8080/people