.PHONY = build
build:
	docker compose up -d --remove-orphans

.PHONY = restart
restart:
	docker compose down
	docker rmi gocloudplaylist_api
	docker compose up -d

.PHONY = rm_containers
rm_containers:
	docker compose down
	docker rmi gocloudplaylist_api
	docker rmi dpage/pgadmin4
	rm -rf postgres/

.PHONY = gen
gen:
	export GO111MODULE=on  # Enable module mode
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	export GO_PATH=~/go
	export PATH=$PATH:/$GO_PATH/bin
	protoc -I api/proto --go_out=pkg/api --go-grpc_out==grpc:pkg/api api/proto/GoCloudPlaylist.proto

.PHONY = re-gen
re-gen:
	protoc -I api/proto --go_out=pkg/api --go-grpc_out==grpc:pkg/api api/proto/GoCloudPlaylist.proto

.PHONY = swagger
swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	PATH=$(go env GOPATH)/bin:$PATH
	swag init -g cmd/main.go

.PHONY = test
test:
	go test ./internal/transport/httpHandlers -v -cover
	go test ./internal/transport/gprcEndpoints -v -cover
	go test ./pkg/timeConverting -v -cover

