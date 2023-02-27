
.PHONY = build
build:
	docker compose up -d --remove-orphans


.PHONY = rm_containers
rm_containers:
	docker compose down
	docker rmi gocloudplaylist_api
	docker rmi dpage/pgadmin4
	#docker rmi postgresContainerForPlaylistService

restart:
	docker compose down
	docker rmi gocloudplaylist_api
	docker compose up -d


.PHONY = stop
stop:
	docker kill --signal=SIGKILL  gocloudplaylist_api

gen:
	export GO111MODULE=on  # Enable module mode
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	export GO_PATH=~/go
	export PATH=$PATH:/$GO_PATH/bin
	protoc -I api/proto --go_out=pkg/api --go-grpc_out==grpc:pkg/api api/proto/GoCloudPlaylist.proto

re-gen:
	protoc -I api/proto --go_out=pkg/api --go-grpc_out==grpc:pkg/api api/proto/GoCloudPlaylist.proto