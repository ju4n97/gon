DATABASE_URI ?= postgres://postgres:postgres@172.17.0.1:5432/gokit?search_path=public&sslmode=disable
MIGRATIONS_DIR = "$(PWD)/internal/db/migrations"

.PHONY: dev
dev: 
	go run main.go serve

.PHONY: test
test:
	go test ./... -v --cover

.PHONY: test-report
test-report: 
	go test ./... -v --cover -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: sqlc-generate
sqlc-generate: 	
	docker run --rm -v "$(PWD):/src" -w /src sqlc/sqlc generate

.PHONY: sqlc-verify
sqlc-verify: 
	docker run --rm -v "$(PWD):/src" -w /src sqlc/sqlc verify

.PHONY: sqlc-vet
sqlc-vet: 
	docker run --rm -v "$(PWD):/src" -w /src sqlc/sqlc vet

.PHONY: atlas-migrate-status
atlas-migrate-status:
	atlas migrate status --dir "file://internal/db/migrations" --url "$(DATABASE_URI)"

.PHONY: atlas-migrate-create
atlas-migrate-create:
	atlas migrate diff $(MIGRATIONS_NAME) --dir "file://internal/db/migrations" --to "file://internal/db/schema.sql" --dev-url "$(DATABASE_URI)" --format '{{ sql . "  " }}'

.PHONY: atlas-migrate-apply
atlas-migrate-apply:
	atlas migrate apply --dir "file://internal/db/migrations" --url "$(DATABASE_URI)"

.PHONY: atlas-migrate-down
atlas-migrate-down:
	atlas migrate down --dir "file://internal/db/migrations" --url "$(DATABASE_URI)" --dev-url "docker://postgres/15/dev?search_path=public"

.PHONY: atlas-schema-inspect-web
atlas-schema-inspect-web: 
	atlas schema inspect --url "$(DATABASE_URI)" --web

.PHONY: atlas-schema-clean
atlas-schema-clean:
	atlas schema clean --url "$(DATABASE_URI)"
