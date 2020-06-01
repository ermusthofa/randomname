# Frontend Service

## Overview
Frontend microservice will fetch the random facts from backend service based on http://numbersapi.com.


## Requirements

### Compiling requirement
- NodeJs
- Yarn
- Docker
- GNU Make

### Setup
To run service in your local machine, you need to have `yarn` installed. For more informations on how to setup your environment, please visit [yarn installation](https://classic.yarnpkg.com/en/docs/install#debian-stable) page.

Run `yarn` inside the frontend service directory to fetch all the dependencies. You can run `yarn start` after that. This will open your browser and the frontend service should be visible.

### Environment Variables
| Environment Variables     | Description                                                                                                                | Default |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------|---------|
| `FRONTEND_LISTEN_ADDRESS` | Used to access the frontend service                                                                                        | 8081    |
| `BACKEND_SERVER_ADDRESS`  | Used by frontend service to call the API (Please set this envar if you aren't using the default value for backend service) | 8080    |

### Build
| make target | Description                                                                                                              |
|-------------|--------------------------------------------------------------------------------------------------------------------------|
| build       | Create the project docker image                                                                                          |
| run         | Run the container. Port `8081` exposed by default. You can change the port by exporting `FRONTEND_LISTEN_ADDRESS` envar  |
| stop        | Stop the service container                                                                                               |
| clean       | Stop and remove the docker images                                                                                        |

## Accessing the service
After the container is built and run, you can access the service in your browser:

```bash
http://localhost:8081
```

## License
This project is licensed under the MIT License - see the [License.md](https://github.com/ermusthofa/randomname/blob/master/LICENSE) file for details