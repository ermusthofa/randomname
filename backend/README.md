# Backend Service

## Overview
Backend microservice will allow frontend service to fetch random name based on https://api.randomuser.me.


## Requirements

### Compiling requirement
- Go Compiler
- Docker
- GNU Make

### Setup
To run service locally, you can build the project binaries or directly run `go run main.go` inside the `backend` service directory. You can change the default port exposed by the backend service by exporting `SERVER_LISTEN_ADDRESS` as an envar.

This will print the following text to your console
```bash

	Listening [::]:8080

```

For more informations on how to setup your Go environment, please visit [Go Installation](https://golang.org/doc/install) page.

### Build
| make target | Description                                                                                                            |
|-------------|------------------------------------------------------------------------------------------------------------------------|
| build       | Create the project docker image                                                                                        |
| run         | Run the container. Port `8080` exposed by default. You can change the port by exporting `SERVER_LISTEN_ADDRESS` envar  |
| stop        | Stop the service container                                                                                             |
| clean       | Stop and remove the docker images                                                                                      |

## Accessing the service
After the container is built and run, you can access the service in:

```bash
http://localhost:8080
```

Currently there are only two endpoints, which are:
| endpoints        | Description                                                                    |
|------------------|--------------------------------------------------------------------------------|
| `/v1/healthz`    | Used by Kubernetes health check                                                |
| `/v1/randomname` | Used by frontend service to get the random name from https://api.randomuser.me |


## License
This project is licensed under the MIT License - see the [License.md](https://github.com/ermusthofa/randomname/blob/master/LICENSE) file for details