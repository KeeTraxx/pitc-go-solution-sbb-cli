# pitc-go-solution-sbb-cli

## How to run locally
1. `go get -v ./...`
2. `go build .`
3. `./pitc-go-solution-sbb-cli`

Oneliner:
`go get -v ./... && go build . && ./$(basename $(pwd))`

## How to run in a docker container
`docker run -ti -w=/go/src/app -p 1323:1323 -v $(pwd):/go/src/app golang:1.8 go get -v ./... && go build . && ./app`
