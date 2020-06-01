# Helm Chart for Trivia Date Service in Kubernetes


This project provides Helm Charts for deploying a Trivia Date service into any Kubernetes based cluster.


This project provides the following files:

| File                                          | Description                                                                                                                               |
|-----------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| Charts.yaml                                   | The definition file for Trivia Date Service                                                                                               |
| values.yaml                                   | Configurable values that are inserted into the sub-chars template files                                                                   |
| charts/backend/Charts.yaml                    | The definition file for backend service                                                                                                   |
| charts/backend/values.yaml                    | Configurable values if you wish to deploy this service independently                                                                      |
| charts/backend/templates/deployment.yaml      | Template to configure backend service deployment                                                                                          |
| charts/backend/templates/hpa.yaml             | Template to configure auto scale deployment                                                                                               |
| charts/backend/templates/service.yaml         | Template to configure kubernetes service object deployment                                                                                |
| charts/backend/templates/serviceaccount.yaml  | Template to configure kubernetes service account deployment                                                                               |
| charts/frontend/Charts.yaml                   | The definition file for frontend service                                                                                                  |
| charts/frontend/values.yaml                   | Configurable values if you wish to deploy this service independently                                                                      |
| charts/frontend/templates/configmap.yaml      | Template to configure configMap deployment. This is the NGINX configuration to proxy request from frontend service to the backend service |
| charts/frontend/templates/deployment.yaml     | Template to configure frontend service deployment. ConfigMap will be mounted here                                                         |
| charts/frontend/templates/hpa.yaml            | Template to configure auto scale deployment                                                                                               |
| charts/frontend/templates/ingress.yaml        | Template to configure ingress deployment. This allows service to be accessible from outside.                                              |
| charts/frontend/templates/service.yaml        | Template to configure kubernetes service object deployment                                                                                |
| charts/frontend/templates/serviceaccount.yaml | Template to configure kubernetes service account deployment                                                                               |

You should only need to edit the `Chart.yaml` and `values.yaml` files.

## Prerequisites

Using the Helm charts assumes the following prerequisites are completed:  

1. You have a Kubernetes cluster  
  This could be one hosted by a cloud provider or running locally, for example using [Minikube](https://kubernetes.io/docs/setup/minikube/)
  
2. You have `kubectl` installed and configured for your cluster  
  The [Kubernetes command line](https://kubernetes.io/docs/tasks/tools/install-kubectl/) tool, `kubectl`, is used to view and control your Kubernetes cluster.

3. You have the Helm command line installed  
   [Helm](https://helm.sh/docs/intro/using_helm/) provides the command line tool for deploying this service using the Helm chart.

4. You have created a Docker image for Trivia Date Service  



## Configuring the Chart for Trivia Date Service

The following table lists the configurable parameters of the Helm chart and their default values.

**Global**

Configuration in this following table are pushed to all sub-charts during `helm install`.

| Parameter                                              | Description                                                                                 | Default                      |
|--------------------------------------------------------|---------------------------------------------------------------------------------------------|------------------------------|
| `global.autoscaling.enabled`                           | Toggle to enable or disable auto scaling                                                    | `false`                      |
| `global.autoscaling.minReplicas`                       | Minimum replicas if auto scaling is enabled                                                 | `1`                          |
| `global.autoscaling.maxReplicas`                       | Maximum replicas if auto scaling is enabled                                                 | `100`                        |
| `global.autoscaling.targetCPUUtilizationPercentage`    | Target CPU utilization percentage to scale the services                                     | `80`                         |
| `global.autoscaling.targetMemoryUtilizationPercentage` | Target Memory utilization percentage to scale the services                                  | `80`                         |
| `global.replicaCount`                                  | If the autoscaling is enabled, then this configuration will have no effects                 | `1`                          |
| `global.backend.serviceName`                           | Kubernetes object service name to be referred by frontend service to access backend service | `triviadate-backend-service` |




**Backend**

Configuration in this following table are pushed to backend sub-chart during `helm install`.

| Parameter                            | Description                                                                                                            | Default        |
|--------------------------------------|------------------------------------------------------------------------------------------------------------------------|----------------|
| `backend.image.repository`           | Image repository for backend service                                                                                   | `backend`      |
| `backend.image.pullPolicy`           | Image pull policy for backend service                                                                                  | `ifNotPresent` |
| `backend.image.tag`                  | Overrides the image tag which default is the chart appVersion.                                                         | `latest`       |
| `backend.serviceAccount.create`      | Specifies whether a service account should be created                                                                  | `false`        |
| `backend.serviceAccount.annotations` | Annotations to add to the service account                                                                              |                |
| `backend.serviceAccount.name`        | The name of the service account to use. If not set and create is true, a name is generated using the fullname template |                |
| `backend.podAnnotations`             | Annotations to add to the pod template                                                                                 |                |
| `backend.podSecurityContext`         | Security context to add to the pod template                                                                            |                |
| `backend.securityContext`            | Security context to add to the container                                                                               |                |
| `backend.service.type`               | Backend service type                                                                                                   | `ClusterIP`    |
| `backend.service.port`               | Backend service port                                                                                                   | `80`           |
| `backend.resources.limits`           | Resource limits                                                                                                        |                |
| `backend.resources.requests`         | Resource requests                                                                                                      |                |



**Frontend**

Configuration in this following table are pushed to frontend sub-chart during `helm install`.

| Parameter                             | Description                                                                                                            | Default                               |
|---------------------------------------|------------------------------------------------------------------------------------------------------------------------|---------------------------------------|
| `frontend.image.repository`           | Image repository for frontend service                                                                                  | `frontend`                            |
| `frontend.image.pullPolicy`           | Image pull policy for frontend service                                                                                 | `ifNotPresent`                        |
| `frontend.image.tag`                  | Overrides the image tag which default is the chart appVersion.                                                         | `latest`                              |
| `frontend.serviceAccount.create`      | Specifies whether a service account should be created                                                                  | `false`                               |
| `frontend.serviceAccount.annotations` | Annotations to add to the service account                                                                              |                                       |
| `frontend.serviceAccount.name`        | The name of the service account to use. If not set and create is true, a name is generated using the fullname template |                                       |
| `frontend.podAnnotations`             | Annotations to add to the pod template                                                                                 |                                       |
| `frontend.podSecurityContext`         | Security context to add to the pod template                                                                            |                                       |
| `frontend.securityContext`            | Security context to add to the container                                                                               |                                       |
| `frontend.service.type`               | Frontend service type                                                                                                  | `ClusterIP`                           |
| `frontend.service.port`               | Frontend service port                                                                                                  | `80`                                  |
| `frontend.ingress.enabled`            | Toggle to enable or disable ingress object for frontend service                                                        | `true`                                |
| `frontend.ingress.annotations`        | Annotations to add to the ingress                                                                                      |                                       |
| `frontend.ingress.hosts.host`         | Host for this ingress to be routed to                                                                                  | `trivia-date.public-domain.example`   |
| `frontend.ingress.hosts.paths`        | Path for this service to be routed to                                                                                  | `/`                                   |
| `frontend.ingress.tls`                | TLS certificate to secure the ingress                                                                                  |                                       |
| `frontend.resources.limits`           | Resource limits                                                                                                        |                                       |
| `frontend.resources.requests`         | Resource requests                                                                                                      |                                       |



## Using the Chart to deploy your Trivia Date Service to Kubernetes

In order to use the Helm chart to deploy Trivia Date Service in Kubernetes, run the following commands:

1. From this project directory, run:  
   ```sh
   helm install `release-name` .
   ```
   This deploys and runs Trivia Date Service in Kubernetes within `default` namespace, and prints the following text to the console:
   
   ```sh
   NAME: `release-name`
   LAST DEPLOYED: Mon Jun  1 20:20:43 2020
   NAMESPACE: default
   STATUS: deployed
   REVISION: 1
   TEST SUITE: None
   ```
   You can specify which environment to deploy this service with specifying `--namespace namespace`.

   To verify this service is running, run the following command
   ```bash
   kubectl get deployment [--namespace namespace]
   ```
   
2. Open your web browser to http://random-name.public-domain.example<br />
   Note: You need integrate your DNS to resolve this domain (or local hosts file)
   
3. Trivia Date Service should now be visible in your browser.


## Uninstalling Trivia Date Service
If you installed your application with:

```sh
helm install `release-name` .
```
then you can:

* Find the deployment using `helm list --all` and search for an entry with the chart name "`release-name`".
* Remove the application with `helm uninstall release-name`.

## License
This project is licensed under the MIT License - see the [License.md](https://github.com/ermusthofa/randomname/blob/master/LICENSE) file for details