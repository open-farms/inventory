# Inventory 📦

> An inventory service for managing farming vehicles, equipment, and resources.

## Quickstart ⚡

```shell
make init
```

## Start the service 🏃

Start the service directly with go or containerize it with Docker.

```shell
# Build the binary
VERSION="<x.x.x>" make build

# Run the service
./bin/inventory
```

```shell
# Build the docker image
docker build -t <your-docker-image-name> .

# Run the docker image, exposing http and gRPC services
docker run --rm -p 8000:8000 -p 9000:9000 <your-docker-image-name>
```

## Generate code 🏗️

Generate the protobuf code, openapi spec, 

```
# Generate API files (include: pb.go, http, grpc, validate) by proto
make api
```
