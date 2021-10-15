build:
	go build -o bin/main main.go

clean_output:
	yes | npx prettier --write addresses.json

run:
	go run *.go

start:
	./bin/main