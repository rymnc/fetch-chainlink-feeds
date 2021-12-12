build:
	go build -o bin/main main.go

clean_run:
	rm output
	go run *.go

run:
	go run *.go

start:
	./bin/main