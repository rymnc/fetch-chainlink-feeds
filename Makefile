build:
	go build -o bin/main main.go

clean_output:
	yes | npx prettier --write addresses.json

run:
	go run main.go
	make clean_output

start:
	./bin/main
	make clean_output