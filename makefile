run:
	go run cmd/app/main.go -m dev

migrate:
	go run cmd/app/main.go -m dev -act migrate

updateP:
	go run cmd/app/main.go -m dev -act updatePolicies

build:
	go build -o build/example.exe cmd/app/main.go