build-migration-tool:
	git clone https://github.com/rubenv/sql-migrate
	cd ./sql-migrate/sql-migrate && go build && cd ../..
	mv ./sql-migrate/sql-migrate/sql-migrate ./tools
	rm -rf ./sql-migrate