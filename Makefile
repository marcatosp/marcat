build:
	go build -o cmd/marcat marcat.go

run:
	cmd/marcat

docker:
	docker build -t marcat -f Dockerfile .

run:
	docker run -d -p 8080:8080 marcat
