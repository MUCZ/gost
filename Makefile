
all: 
	go mod tidy
	go build  .

docker:
	GOOS=linux GOARCH=amd64 go build .
	docker build -t gost .

run:
	docker run -it -p 8080:8080 gost