# Blog Server
## Introduction
This is a blog backend, it provides file server, post, and user service. It is a microservice archetecture witch is written in Golang.

Now, this server is deloyed on GKE cluster, you can click  to check the server's health status

## Feature
File service can handle image upload and serve, which can support posts' image and users' avators.

Post service can handle create, update, delete, list. Besides, we use JWT token to controll the authorization.

User service show user's information and return suitable JWT token according user's identity.

- Use gRPC for communication between microservice.
- Use gRPC-gateway to serve RESTful APIs and gRPC APIs, and it also checks authentication before any request.
- Use MongoDB to store user, file and post information.
- Use Kubenetes to manage our service. You can check helm and k8s directories for more information.

## Code generation
Some modules like gRPC need you to make command for generating necessary node.
```
make dc.generate
```

## Style check
You can make the following command to check whole project style.
```
make dc.lint
```

## Build image
You can build your own blog server image by following command.
```
make dc.image
```

## CI/CD
- CI/CD run in the github action.
- CI workflow work [here](.github/workflows/main.yaml).
- We will add CD as soon as possible.

## Future work
1. Finish CD to update service when pusing to default branch.
2. Enhance search to make partial serach work.
3. Add monitor services like Prometheus, Jaeger or Grafana.
4. Add Redis to make response faster.
