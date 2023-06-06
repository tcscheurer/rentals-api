local-up:
	docker-compose up -d && go run ./cmd/api/main.go

db-up:
	docker-compose up -d

db-generate:
	sqlc generate

local-server:
	go run ./cmd/api/main.go

mock-db-gen:
	mockgen -package mockdb -destination db/mock/store.go github.com/tcscheurer/rentals/db/sqlc Querier

test:
	go test ./... -v

.PHONY: local-up db-up test db-generate local-server test 


