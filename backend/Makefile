SERVICE=backend
export SERVER_LISTEN_ADDRESS?=8080

RUN=docker run -d --rm \
		-e SERVER_LISTEN_ADDRESS=${SERVER_LISTEN_ADDRESS} \
		-p ${SERVER_LISTEN_ADDRESS}:${SERVER_LISTEN_ADDRESS} \
		--name ${SERVICE}

MAKEFLAGS += --silent

all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(SERVICE)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## build: Create the project docker image
build:
	docker build -t ${SERVICE} .

## run: Run the container. Port 8080 exposed by default.
## : You can change the port by exporting SERVER_LISTEN_ADDRESS envar
run:
	@${RUN} ${SERVICE}

## stop: Stop the service container
stop:
	@-docker stop ${SERVICE}

clean-image:
	@-docker rmi ${SERVICE}

## clean: Stop and remove the docker images
clean: stop clean-image
