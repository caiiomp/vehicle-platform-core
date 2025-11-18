migrate-create:
	migrate create -ext sql -seq -dir db/migrations ${NAME}

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down
