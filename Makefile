build:
	go build -o bin/main main.go

clean_output:
	yes | npx prettier --write addresses.json

run:
	rm addresses.json
	go run *.go

start:
	./bin/main