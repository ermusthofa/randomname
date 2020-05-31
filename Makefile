SHELL := /bin/bash

CHART_NAME := randomname-chart

RUN=docker run -d --rm \
		-e SERVER_LISTEN_ADDRESS=${SERVER_LISTEN_ADDRESS} \
		-p ${SERVER_LISTEN_ADDRESS}:${SERVER_LISTEN_ADDRESS} \
		--name ${SERVICE}

BUILD=docker build -t

HELM_INSTALL=helm install --generate-name 
HELM_REVIEW=helm install --generate-name --dry-run

help: Makefile
	@echo
	@echo " Choose a target to run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## build : Build the image with specified target. Accepted value `backend`
## : and `frontend`
build:
	@read -p "Which service you wanna build? " INPUT; 													\
		if [[ $$INPUT == "frontend" ]] || [[ $$INPUT == "backend" ]]; then				\
			${BUILD} $$INPUT:latest $$INPUT/ ;																			\
			echo -e "Repository\t\t\tTag\t\t\tImage ID\tCreated\t\t\tSize";					\
			docker images | grep $$INPUT ;																					\
		else																																			\
			echo "Please fill either with 'frontend' or 'backend'";									\
		fi

## build-all : Build all the images
build-all:
	@${BUILD} backend:latest backend
	@${BUILD} frontend:latest frontend
	@echo -e "Repository\t\t\tTag\t\t\tImage ID\tCreated\t\t\tSize"
	@docker images | egrep -e "backend|frontend"

## chart-review : Review the chart manifest using 'less'.
## : You will be prompted wether you want to enable auto scale or not
chart-review:
	@read -p "Which environment do you want to review : " INPUT;											\
		export SCALING_ENABLED=false;																										\
		read -p "Do you want to enable the auto-scaling [Y/N] : " FLAG;									\
		if [[ $$FLAG == "Y" ]] || [[ $$FLAG == "y" ]]; then															\
			export SCALING_ENABLED=true;																									\
		fi;																																							\
																																										\
		${HELM_REVIEW} --set global.autoscaling.enabled=$$SCALING_ENABLED								\
			${CHART_NAME} | less

## chart-deploy : Deploy the chart to kubernetes cluster.
chart-deploy:
	@read -p "Which environment do you want to deploy this apps : " INPUT;						\
		export SCALING_ENABLED=false;																										\
		read -p "Do you want to enable the auto-scaling [Y/N] : " FLAG;									\
		if [[ $$FLAG == "Y" ]] || [[ $$FLAG == "y" ]]; then															\
			export SCALING_ENABLED=true;																									\
		fi;																																							\
																																										\
		${HELM_INSTALL} --set global.autoscaling.enabled=$$SCALING_ENABLED							\
			--namespace $$INPUT ${CHART_NAME}
