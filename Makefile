test:
	go test ./... -v --cover

test-report: 
	go test ./... -v --cover -coverprofile=coverage.out
	go tool cover -html=coverage.out