[![HexWorld hexcloud](https://github.com/3vilM33pl3/hexrepo/actions/workflows/build-hexcloud.yml/badge.svg)](https://github.com/3vilM33pl3/hexrepo/actions/workflows/build-hexcloud.yml)

Backend server and main access point via gRPC API for the game.

## Build and test
`hexcloud` uses sqlite3 package which uses cgo. For this to compile you need to install the `gcc` compiler
and set the `CGO_ENABLED=1` ennvironment variable.

To build and test:
```shell
go build -v ./...
go test -v ./...
```

## Run
### Help
```shell
go run .\cmd\hexcloud\hexcloud.go -h 
or 
hexcloud -h 

Usage of hexcloud:
  -address string
        address and port number to listen on (default "0.0.0.0:8080")
  -local
        running local
```
### Start server
Start the server listening on `localhost:8080`
```shell
go run .\cmd\hexcloud\hexcloud.go  
or 
hexcloud 
```

### Insert data
```shell
go run .\cmd\hexcli\nb.go repo add file .\test\data_repo.csv --secure=false --addr=localhost:8080
go run .\cmd\hexcli\nb.go repo add data file .\test\data_hexdata.csv --secure=false --addr=localhost:8080
```

## Docker
```bash
docker build -f ./deploy/hexcloud/Dockerfile -t hexcloud .
```

## Update gRPC proto
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc --proto_path=../../ --go_opt=Mapi/hexagon.proto=/hexcloud --go-grpc_opt=Mapi/hexagon.proto=/hexcloud --go_out=./internal/pkg --go-grpc_out=./internal/pkg ../../api/hexagon.proto
```

## Program Flow
Program flow of two of the [API calls](./api/hexagon.proto) (MapAdd and MapGet)
![UML Sequence Diagram](./images/hexcloud.svg)

