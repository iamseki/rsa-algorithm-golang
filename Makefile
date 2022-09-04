## Run with GO
run:
	go run main.go rsa.go

## COMPILE FOR LINUX
compile-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rsa-from-scratch main.go rsa.go 

## COMPILE FOR WINDOWS
compile-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o rsa-from-scratch.exe main.go rsa.go 

## COMPILE FOR MAC
compile-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o rsa-from-scratch-mac main.go rsa.go 