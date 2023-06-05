local-up:
	docker-compose up -d && go run ./cmd/api/main.go

local-down:
	docker-compose down

test:
	go test -v -cover ./...

db-generate:
	sqlc generate

local-server:
	go run ./cmd/api/main.go

mock-db-gen:
	mockgen -package mockdb -destination db/mock/store.go github.com/tcscheurer/rentals/db/sqlc Querier

test:
	go test ./... -v

.PHONY: local-up local-down test db-generate local-server test 


