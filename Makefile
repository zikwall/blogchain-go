PROJECT_NAME=$(shell basename "$(PWD)")
SCRIPT_AUTHOR=Andrey Kapitonov <andrey.kapitonov.96@gmail.com>
SCRIPT_VERSION=0.0.5.dev
SERVICES=\
	common \
	storage

all: tests

common_proto:
	protoc -I . \
    	--go_out=.  \
    	--go_opt=paths=source_relative \
    	--go-grpc_out=. \
    	--go-grpc_opt=paths=source_relative \
    	./src/protobuf/common/*.proto;

$(SERVICES):
	protoc -I ./src/protobuf/$@/ -I . \
		--go_out=./src/protobuf/$@/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=./src/protobuf/$@/ \
		--go-grpc_opt=paths=source_relative \
		$@.proto;

default: common_proto $(SERVICES)

# Download and install golangci-linter
linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.40.1
	mv ./bin/golangci-lint /bin
	rm -rf ./bin

deploy: build-migration-tool migrate-up

tests:
	go test ./... -v

build-migration-tool:
	git clone https://github.com/rubenv/sql-migrate
	cd ./sql-migrate/sql-migrate && go build && cd ../..
	mv ./sql-migrate/sql-migrate/sql-migrate ./cmd
	rm -rf ./sql-migrate

migrate-up:
	./src/cmd/sql-migrate up -config=./cmd/dbconfig.yml

migrate-down:
	./src/cmd/sql-migrate down -config=./cmd/dbconfig.yml

migrate-status:
	./src/cmd/sql-migrate status -config=./cmd/dbconfig.yml

migrate-new:
	./src/cmd/sql-migrate new $(name) -config=./cmd/dbconfig.yml

help:
	@echo -e "Usage: make [target] ...\n"
	@echo -e "build-migration-tool 	: Download & create migration tool"
	@echo -e "migrate-up 			: Apply migrations"
	@echo -e "migrate-down 		: Down migrations"
	@echo -e "migrate-status 		: Status of migrations"
	@echo -e "migrate-new 			: Create new migration by name={name_here}"
	@echo -e '\nProject name : '$(PROJECT_NAME)
	@echo -e "Written by $(SCRIPT_AUTHOR), version $(SCRIPT_VERSION)"
	@echo -e "Please report any bug or error to the author."
