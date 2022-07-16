SECURE=true
ADDRESS=hexcloud-j6feiuh7aa-ue.a.run.app:443

build:
	go build -o bin/nb nb.go

protoc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	protoc --go_out=./internal/pkg  --go-grpc_out=./internal/pkg ./api/hexagon.proto

run-help:
	go run nb.go -h

run-help-repo:
	go run nb.go repo -h

run-help-map:
	go run nb.go map -h

#
# Repository
#
run-repo-add:
	go run nb.go repo add 1000-0000-0000-0000 --secure=$(SECURE) --addr=$(ADDRESS)

run-repo-add-file:
	go run nb.go repo add file hexagons.csv --secure=$(SECURE) --addr=$(ADDRESS)

w	run-repo-add-data:
	go run nb.go repo add data 1000-0000-0000-0000 biome swamp --secure=$(SECURE) --addr=$(ADDRESS)

run-repo-add-data-file:
	go run nb.go repo add data file hexagon_data.csv --secure=$(SECURE) --addr=$(ADDRESS)

run-repo-get:
	go run nb.go repo get 0000-0000-0000-0000 --secure=$(SECURE) --addr=$(ADDRESS)

run-repo-get-data:
	go run nb.go repo get data 1000-0000-0000-0000 biome --secure=$(SECURE) --addr=$(ADDRESS)

run-repo-del:
	go run nb.go repo del 1000-0000-0000-0000 --secure=$(SECURE) --addr=$(ADDRESS)

run-repo-del-data:
	go run nb.go repo del data 1000-0000-0000-0000 biome --secure=$(SECURE) --addr=$(ADDRESS)

#
# Map
#
run-map-add:
	go run nb.go map add --secure=$(SECURE) --addr=$(ADDRESS) -- 0 -5 5 0000-0000-0000-0000

run-map-add-data:
	go run nb.go map add data --secure=$(SECURE) --addr=$(ADDRESS) -- 0 -5 5 animal cat

run-map-get:
	go run nb.go map get --radius=2 --secure=$(SECURE) --addr=$(ADDRESS) -- 0 0 0

run-map-update:
	go run nb.go map update --secure=$(SECURE) --addr=$(ADDRESS) -- 0 -5 5 0000-0000-0000-0003

run-map-update-data:
	go run nb.go map update data --secure=$(SECURE) --addr=$(ADDRESS) -- 0 -5 5 animal dog

run-map-update-data-1:
	go run nb.go map update data --secure=$(SECURE) --addr=$(ADDRESS) -- 0 0 0 direction SE

run-map-remove:
	go run nb.go map remove --secure=$(SECURE) --addr=$(ADDRESS) -- 0 -5 5

run-map-remove-data:
	go run nb.go map remove data --secure=$(SECURE) --addr=$(ADDRESS) -- 0 -5 5 animal

#
# Status
#
run-status-server:
	go run nb.go status server --secure=$(SECURE) --addr=$(ADDRESS)

run-status-storage:
	go run nb.go status storage --secure=$(SECURE) --addr=$(ADDRESS)

run-status-clients:
	go run nb.go status clients --secure=$(SECURE) --addr=$(ADDRESS)
