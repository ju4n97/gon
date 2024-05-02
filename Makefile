DATABASE_URI ?= postgres://root:changeme@172.17.0.1:5432/kitten?search_path=public&sslmode=disable
MIGRATIONS_DIR = "$(PWD)/internals/database/migrations"

dev: 
	go run main.go serve

test:
	go test ./... -v --cover

test-report: 
	go test ./... -v --cover -coverprofile=coverage.out
	go tool cover -html=coverage.out

sqlc: 
	docker run --rm -volume "$(PWD):/src" -w /src sqlc/sqlc generate

sqlc-verify: 
	docker run --rm -volume "$(PWD):/src" -w /src sqlc/sqlc verify

sqlc-vet: 
	docker run --rm -volume "$(PWD):/src" -w /src sqlc/sqlc vet

atlas-migrate-status:
	atlas migrate status --dir "file://internals/database/migrations" --url "$(DATABASE_URI)"

atlas-migrate-create:
	atlas migrate diff $(MIGRATIONS_NAME) --dir "file://internals/database/migrations" --to "file://internals/database/schema.sql" --dev-url "$(DATABASE_URI)" --format '{{ sql . "  " }}'

atlas-migrate-apply:
	atlas migrate apply --dir "file://internals/database/migrations" --url "$(DATABASE_URI)"

atlas-migrate-down:
	atlas migrate down --dir "file://internals/database/migrations" --url "$(DATABASE_URI)" --dev-url "docker://postgres/15/dev?search_path=public"

atlas-schema-inspect-web: 
	atlas schema inspect --url "$(DATABASE_URI)" --web

atlas-schema-clean:
	atlas schema clean --url "$(DATABASE_URI)"

.PHONY: dev \
		test test-report \
        sqlc sqlc-pull sqlc-verify sqlc-vet \
		atlas-migrate-create atlas-migrate-apply atlas-migrate-down atlas-schema-inspect-web atlas-schema-clean