run:
	cd backend && go run cmd/server/main.go

migrate:
	migrate -path backend/database/postgres/migrations -database "postgres://root:root@host.docker.internal:5432/${table}?sslmode=disable" -verbose ${verbose}

table:
	migrate create -ext sql -dir backend/database/postgres/migrations -seq ${table}

test:
	cd backend && go test -v ./... | { grep -v 'no test files'; true; }

bench:
	cd backend && go test -bench=. ./... | { grep -v 'no test files'; true; }