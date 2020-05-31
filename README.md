# Random Name Microservice
Used to fetch random user information based on https://api.randomuser.me public API.

---

## What this repository covers
- [x] Backend: Get the user information from the https://api.randomuser.me and return the user information via REST API
- [x] Backend: Will remain private. Inaccessible by the public and only accessible from the frontend, it doesn't have a public domain and/or IP
- [x] Backend: Written in Go language
- [x] Frontend: request the user information to the backend and it will display the name information
- [x] Frontend: Written using React Framework
- [x] Frontend: It can be configured to be accessible through a public domain
- [x] Both of the services are containerized
- [x] Microservices deployment to Kubernetes cluster can be done using Helm chart
- [x] The Helm chart can accommodate auto-scaling and easily toggled on and off in a specific environment, the case here is that we need auto-scaling in production but we certainly do not in staging and development
- [x] Makefile that will simplify the building, containerizing, and deployment in various environments (production, staging, etc)

## Setup
---
Environment used during development are as follow:
1. Docker version 19.03.8
   Refer to this [link](https://docs.docker.com/engine/install/) for how to get docker up and running
2. Kubernetes version 1.18.0
   Refer to this [link](https://kubernetes.io/docs/setup/) for setup the kubernetes installation (while we are using `minikube` during development, and please let `ingress-controller` and `metrics-server` addons enabled as we need that addons to get these apps up and running)
3. Go version 1.14.1
   Refer to this [link](https://golang.org/doc/install) how to setup Go
4. Node.js version v12.17.0
   Refer to this [link](https://nodejs.org/en/download/) for how to get Node.js
5. GNU Make 4.2.1
   Refer to this [link](https://www.gnu.org/software/make/) to get more information about GNU Make

## Makefile
---
To get these app ready, you can directly build the containers from the root directory of this repository.
Run `make` to get details information about it's target
```bash
 Choose a target to run:

  build           Build the image with specified target. Accepted value `backend`
                  and `frontend`
  build-all       Build all the images
  chart-review    Review the chart manifest using 'less'.
                  You will be prompted wether you want to enable auto scale or not
  chart-deploy    Deploy the chart to kubernetes cluster.

```

## Helm Chart
---


## Backend
---
Backend microservices will allow frontend service to fetch random name based on https://randomuser.me. For a detailed explanation, please refer to the service README.

## Frontend
---
Frontend microservices will allow user to get random name fetched from backend service. For a detailed explanation, please refer to the service README.

### Author