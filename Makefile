PROJECTNAME=$(shell basename "$(PWD)")
SCRIPT_AUTHOR=Andrey Kapitonov <andrey.kapitonov.96@gmail.com>
SCRIPT_VERSION=0.0.1.dev

# ENV
TEST_DB=${MYSQL_DATABASE}
TEST_USER=${MYSQL_USER}
TEST_PASS=${MYSQL_PASSWORD}

all: database test

deploy: build-migration-tool migrate-up

tests:
	go test ./... -v

database:
	mysql -u${TEST_USER} -p${TEST_PASS} -e "drop database if exists ${TEST_DB}; create database ${TEST_DB};"
	mysql -u${TEST_USER} -p${TEST_PASS} ${TEST_DB} < ci/.teamcity/dump.sql

build-migration-tool:
	git clone https://github.com/rubenv/sql-migrate
	cd ./sql-migrate/sql-migrate && go build && cd ../..
	mv ./sql-migrate/sql-migrate/sql-migrate ./src/cmd
	rm -rf ./sql-migrate

migrate-up:
	./src/cmd/sql-migrate up -config=./src/cmd/dbconfig.yml

migrate-down:
	./src/cmd/sql-migrate down -config=./src/cmd/dbconfig.yml

migrate-status:
	./src/cmd/sql-migrate status -config=./src/cmd/dbconfig.yml

migrate-new:
	./src/cmd/sql-migrate new $(name) -config=./src/cmd/dbconfig.yml

help:
	@echo -e "Usage: make [target] ...\n"
	@echo -e "build-migration-tool 	: Download & create migration tool"
	@echo -e "migrate-up 			: Apply migrations"
	@echo -e "migrate-down 		: Down migrations"
	@echo -e "migrate-status 		: Status of migrations"
	@echo -e "migrate-new 			: Create new migration by name={name_here}"
	@echo -e '\nProject name : '$(PROJECTNAME)
	@echo -e "Written by $(SCRIPT_AUTHOR), version $(SCRIPT_VERSION)"
	@echo -e "Please report any bug or error to the author."
