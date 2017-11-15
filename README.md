# Demo app

This application is intended for demo-ing how to monitor a Go application getting data out of postgres and with an Nginx server in front, serving static assets.

The Datadog agent is deployed as a daemonset. Nginx, the Go app, and postgres have annotations that allow the agent to automatically discover and monitor them.
The Go app exposes some `expvar` metrics.

## How to run it

- configure kubectl to connect to a Kubernetes cluster ([minikube](https://github.com/kubernetes/minikube) is supported)
- get a Datadog API key and replace the place holder in monitoring/agent-ds.yaml (or replace monitoring/agent-ds.yaml with your favorite monitoring agent)
- run `chmod +x provision.sh && ./provision.sh`
- profit!
- ... and tear it down with cleanup.sh

## Build the images

The following instruction include building the nginx and app images, but they are available on docker hub as well: https://hub.docker.com/r/hkaj/demo-app/ and https://hub.docker.com/r/hkaj/nginx/

## Nginx

```bash
cd nginx && docker build -t $TAG_NAME .
docker run -d --name nginx $TAG_NAME
```

## App

```bash
cd app && docker build -t $TAG_NAME .
docker run -d --name demo $TAG_NAME .
```

## postgres

```bash
docker run --name postgres -e POSTGRES_USER=foo -e POSTGRES_PASSWORD=bar -e POSTGRES_DB=app -d -v /path/to/init-app-db.sh:/docker-entrypoint-initdb.d/init-app-db.sh postgres:9
```
