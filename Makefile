build:
	@go build -o $(HOME)/go/bin/p2p 

run: build
	$(HOME)/go/bin/p2p

test: 
	@go test ./... --race -v