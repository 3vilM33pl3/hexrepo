# Overview
Command line interface to manipulate game content

## Commands
### Help commands
    nb -h
    nb repo -h
    nb map -h

### Hexagon manipulation
    nb repo add [ref]
    nb repo add data [ref] [key] [value]
    nb repo add data file [file.csv]
    nb repo get [ref]
    nb repo del [ref]
    nb repo del data [ref] [key]


### Adding/deleting/updating hexagons on map
    nb map add [0,0,0] [ref]
    nb map add data [0,0,0] [key] [value]
    nb map add data file [file.csv]
    nb map get [0,0,0]
    nb map get data [0,0,0] [key]
    nb map update [0,0,0] [ref]
    nb map update data [0,0,0] [key] [value]
    nb map del [0,0,0]
    nb map del data [0,0,0]

### Bulk operations
    nb repo del data all
    nb map del data all
    nb repo add file [file.csv]
    nb map add file [file.csv]
    nb grap add file [graph.gv] 

#### CSV data file example      
    -1,0,direction,S
    0,0,direction,SE
    0,1,link,-1:1
    -1,1,link,-1:0

The `link` keyword is used to link hexagons together (unidirectional). 
The value is the x and y coordinates of the linked hexagon separated by a colon. 

### Status hexagon network (storage server, meta data server, connected clients)
    nb status server
    nb status storage
    nb status clients

### Package content
    nb content pack [dir]
    nb content upload [pack]
    nb content download [ref]

## Use with server
Run [`hexcloud`](https://github.com/3vilM33pl3/hexcloud) server. The server listens to 0.0.0.0:8080 by default.

`go run .\cmd\hexcloud\hexcloud.go --local=true`

Test with status command:

`go run nb.go status server`

## Update gRPC proto
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go
install google.golang.org/grpc/cmd/protoc-gen-go-grpc
protoc --proto_path=../../ --go_opt=Mapi/hexagon.proto=/hexcloud --go-grpc_opt=Mapi/hexagon.proto=/hexcloud --go_out=./internal/pkg --go-grpc_out=./internal/pkg ../../api/hexagon.proto
```

## Program Flow
Program flow of two of the [API call](./api/hexagon.proto) (MapAdd)
![UML Sequence Diagram](./images/hexcli.svg)