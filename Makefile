hello:
	echo "Hello World"
requirements:
	go mod tidy
clean-packages:
	go clean -modcache
run:
	go run main.go