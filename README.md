# Kubernetes Status API

| This project is intended as an exercise.

## Description

This repo contains an HTTP API service that expose the status of a Kubernetes cluster.

## Considerations

I will be using [Go](https://go.dev/) and [Okteto](https://www.okteto.com/). Go is a good choice when it comes to build backend services thanks to its type-safety and rich ecosystem of modules that will help me in this project. Also, the Kubernetes client is very DevX friendly to use in Go.

To implement the HTTP web server and also to keep the project structure well organised easily extensible in the future, I will be using [Gin](https://github.com/gin-gonic/gin).

I am assuming that the service will run inside the Kubernetes cluster itself. If the code run elsewhere, some changes are required to retrieve the K8S's config.

## Development

As requirement to run this, you need a Kubernetes cluster. In my case I am using Okteto Cloud. Check out [their Go sample project](https://www.okteto.com/docs/samples/golang/) to know more about Okteto.

To launch the project in Okteto, follow these steps:

1. Get the [Okteto CLI](https://www.okteto.com/docs/getting-started/#installing-okteto-cli)
1. `okteto context use https://<YOUR_NAMESPACE>.okteto.net`
1. `okteto deploy --build`
1. `okteto up`

### Publish to Docker

```bash
make docker-publish TAG=1.0.0
```

### Future improvements

- [ ] Create Swagger/OpenAPI specs and use them for request/response validation and potentially for code generation
- [ ] Setup GitHub Actions