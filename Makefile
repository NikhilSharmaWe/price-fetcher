build: 
	go build -o bin/pricefetcher

run: 
	make build
	./bin/pricefetcher
