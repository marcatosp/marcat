build:
	go build -o ./cmd/marcat marcat.go

stop:
	docker stop $(docker ps -q --filter ancestor=marcat)

docker_build:
	docker build -t marcat -f Dockerfile .

docker_gh_push:
	make docker_build
	docker login docker.pkg.github.com -u "${GH_USER}" -p "${GH_MARCAT_PACKAGE_TOKEN}" 1>/dev/null
	docker tag marcat docker.pkg.github.com/marcatosp/marcat/marcat:latest
	docker push docker.pkg.github.com/marcatosp/marcat/marcat:latest

run:
	docker run -d -p 1000:8080 marcat
