migrate-create:
	migrate create -ext sql -seq -dir db/migrations ${NAME}

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

generate-mocks:
	go install github.com/vektra/mockery/v2@v2.53.3
	mockery --dir src/core/_interfaces/ --all --output src/core/_mocks

