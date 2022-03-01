
clean:
	go clean

## install: install go modules
install:
	go mod init oms
	go mod tidy
	go mod vendor
	
run:
	go run main/main.go main/routes.go
