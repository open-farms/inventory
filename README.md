# Inventory 📦

> An inventory service for managing farming vehicles, equipment, and resources.

![logo](./static/logo.png)

## What is this?

- A REST API for managing farm inventory via HTTP requests
- A gRPC service for managing farm inventory between internal services

The inventory service enables users to keep track of farm inventory across multiple IoT devices, servers, applications, or databases. This service provides a foundational interface for inventory management and interaction.

## Structure

```shell
/api # protobuf schemas, generated client, generated server
/cmd/inventory # api server
/cmd/inventoryctl # cli tool for service administration  
/config # default application configurations
/ent # database models, migrations, and drivers
/ent/schemas # editable database models
/internal/biz # business logic used by both api and cli
/internal/service # handlers for gRPC and HTTP endpoints
/internal/settings # settings package to read in application configs
/third_party # protobuf libraries
```

## Quickstart ⚡

```shell
# Download all dependencies for the project
make init

# Build it
make build

# Run it
./bin/inventory
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
docker run --rm -v ./config:/data/confifg -p 8000:8000 -p 9000:9000 <your-docker-image-name>
```

## Generate code 🏗️

Generate the protobuf code, openapi spec, 

```
# Generate API files (include: pb.go, http, grpc, validate) by proto
make api
make proto
```

## Services

[View the openapi specification](./openapi.yaml)

[View the .proto definitions](https://github.com/search?q=repo%3Aopen-farms%2Finventory+extension%3Aproto+path%3Aapi%2F&type=Code)

### Vehicles Service

The Vehicle schema and service allow for the management of vehicles on the farm.

```proto
message Vehicle {
	uint64 id = 1;
	string make = 2;
	string model = 3;
	string year = 4;
	bool active = 5;
	repeated string tags = 6;

	enum Condition {
		CONDITION_UNSPECIFIED = 0;
		CONDITION_MINT = 1;
		CONDITION_GOOD = 2;
		CONDITION_POOR = 3;
		CONDITION_BROKEN = 4;
	}

	Condition condition = 7;

	google.type.DateTime created_at = 8;
	google.type.DateTime updated_at = 9;
}
```

### Equipment Service

The Equipment schema and service allow for the management of equipment on the farm.

```proto
message Equipment {
	uint64 id = 1;
	string name = 2;
	repeated string tags = 3;

	enum Condition {
		CONDITION_UNSPECIFIED = 0;
		CONDITION_MINT = 1;
		CONDITION_GOOD = 2;
		CONDITION_POOR = 3;
		CONDITION_BROKEN = 4;
	};

	Condition condition = 4;
	google.type.DateTime created_at = 5;
	google.type.DateTime updated_at = 6;
}
```