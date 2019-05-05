# gRPC Microservices

[![Go Report Card](https://goreportcard.com/badge/github.com/178inaba/grpc-microservices)](https://goreportcard.com/report/github.com/178inaba/grpc-microservices)

## Command to generate code

```console
$ protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/activity/activity.proto
$ protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/user/user.proto
$ protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/project/project.proto
$ protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/task/task.proto
```

## Environment switching

Use `GRPC_MICROSERVICES_ENVIRONMENT`.
Everything except **production** is development.

```console
$ GRPC_MICROSERVICES_ENVIRONMENT=production docker-compose up
```
