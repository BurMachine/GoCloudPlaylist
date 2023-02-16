
.PHONY = build
build:
	docker build .
	docker compose up -d --remove-orphans


.PHONY = rm_containers
rm_containers:
	docker compose down
	docker rmi gocloudplaylist_api
	docker rmi dpage/pgadmin4
	docker rmi postgresContainerForPlaylistService

restart:
	docker compose down
	docker rmi gocloudplaylist_api
	docker compose up -d
