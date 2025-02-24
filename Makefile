docker_run = docker run --rm -it -v .:/work -v /tmp/notes:/tmp/notes -v ${COTONETES_GOCACHE}:/cotonetes_gocache -v ${COTONETES_GOMODCACHE}:/cotonetes_gomodcache -e GOCACHE=/cotonetes_gocache -e GOMODCACHE=/cotonetes_gomodcache cotonetes/golang:v1

build:
	docker build -t cotonetes/golang:v1 .

sh:
	$(docker_run) bash

run:
	$(docker_run) go run latex_to_cotonetes.go
