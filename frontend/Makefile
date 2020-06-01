SERVICE=frontend
export FRONTEND_LISTEN_ADDRESS?=8081
export REACT_APP_BACKEND_URL?=http://localhost:8080

RUN=docker run -d --rm \
		-e REACT_APP_BACKEND_URL=${REACT_APP_BACKEND_URL} \
		-p ${FRONTEND_LISTEN_ADDRESS}:80 \
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

## run: Run the container. Port 8081 exposed by default.
## : You can change the port by exporting FRONTEND_SERVER_ADDRESS envar
run:
	@${RUN} ${SERVICE}

## stop: Stop the service container
stop:
	@-docker stop ${SERVICE}

clean-image:
	@-docker rmi ${SERVICE}

## clean: Stop and remove the docker images
clean: stop clean-image
