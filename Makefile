# Run the Go program
run:
	go run ./src

# Build the Go program
build:
	go build -o bin ./src

swag:
	swag init --dir src/