include .env
export $(shell sed 's/=.*//' .env)

NAME := rift

all: test

test: generate
	[ -f .env ] && source .env; go test -v ./...

generate:
	golem generate

build: generate
	go build

install: build
	go install

server:
	go run main.go server

docker:
	docker build --build-arg="GITHUB_TOKEN=$(GITHUB_TOKEN)" -t $(NAME) .

docker-run:
	docker run -d --rm --name $(NAME) -p $(PORT):$(PORT) $(NAME)

dotenv:
	npx dotenv-vault local build

clean:
	rm -rf static/assets/*

ui: clean
	cd ui && npm run build

deps:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/dashotv/golem@latest

.PHONY: server receiver test deps docker docker-run
