hello:
	echo "Hello World"
requirements:
	go mod tidy
clean-packages:
	go clean -modcache
build:
	go build main.go
run:
	go run main.go
postgresql:
	docker compose up