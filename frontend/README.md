# Frontend Service

## Overview
Frontend microservice will fetch the random name from backend service based on https://api.randomuser.me.


## Requirements

### Compiling requirement
- NodeJs
- Yarn
- Docker
- GNU Make

### Setup
To run service in your local machine, you need to have `yarn` installed. For more informations on how to setup your environment, please visit [yarn installation](https://classic.yarnpkg.com/en/docs/install#debian-stable) page.

Run `yarn` inside the frontend service directory to fetch all the dependencies. You can run `yarn start` after that. This will open your browser and the frontend service should be visible in `http://localhost:3000`.

### Container Environment Variables
| Environment Variables     | Description                                                                                                                | Default                  |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------|--------------------------|
| `FRONTEND_LISTEN_ADDRESS` | Used to access the frontend service                                                                                        | 8081                     |
| `REACT_APP_BACKEND_URL`   | Used by frontend service to call the API (Please set this envar if you aren't using the default value for backend service) | http://localhost:8080    |

### Build
| make target | Description                                                                                                              |
|-------------|--------------------------------------------------------------------------------------------------------------------------|
| build       | Create the project docker image                                                                                          |
| run         | Run the container. Port `8081` exposed by default. You can change the port by exporting `FRONTEND_LISTEN_ADDRESS` envar  |
| stop        | Stop the service container                                                                                               |
| clean       | Stop and remove the docker images                                                                                        |


> ℹ️ **Please be noted**, when you run this service through `make run` or `docker run`, this service won't be able to fetch the data from the backend. That's because request from this service is proxied through Nginx inside of the container, hence it won't be able to call the backend service. Use `docker-compose` on top of this repository instead.

## Accessing the service
After the container is built and run, you can access the service in your browser:

```bash
http://localhost:8081
```

## License
This project is licensed under the MIT License - see the [License.md](https://github.com/ermusthofa/randomname/blob/master/LICENSE) file for details