db.migrate:
	migrate -url=`echo ${DATABASE_URL}` -path=./db/migrate up

db.rollback:
	migrate -url=`echo ${DATABASE_URL}` -path=./db/migrate down

vet:
	go vet ./...
