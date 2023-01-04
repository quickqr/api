# TODO: Setup this
#include .env

#build:
#	docker-compose build
#run:
#	docker-compose up -d
#stop:
#	docker-compose down

swagger:
	swag init -g cmd/main.go --md docs/markdown

