.PHONY: all update clean

all: main.zip

main.zip: main
	build-lambda-zip -output main.zip main

main: main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -ldflags '-d -s -w' -a -tags
	netgo -installsuffix netgo -o main main.go

	update:main.zip
		aws lambda update-function-code --function-name hello --zip-file fileb://main.zip

	clean:
	rm -f main.zip main