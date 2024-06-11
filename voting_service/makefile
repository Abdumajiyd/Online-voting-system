
create:
	migrate create -ext sql -dir migrations -seq create_sql_tables


mig-up:
	migrate -database 'postgres://abdumajiyd:1202@localhost:5432/vote?sslmode=disable' -path migrations up

mig-down:
	migrate -database 'postgres://abdumajiyd:1202@localhost:5432/vote?sslmode=disable' -path migrations down


	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"


gen-proto:

  protoc --go_out=services/ \
       --go-grpc_out=services/ \
       services/protos/cart.proto


gen-protoAll:
  protoc --go_out=./ \
    --go-grpc_out=./ \
    protos/*.proto



gen-proto_curr:
  protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto


CURRENT_DIR=$(shell pwd)


proto-gen:
	./scripts/gen-proto.sh "${CURRENT_DIR}"