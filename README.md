# gRPC studying

## Tutorial package :information_source:

This *package* were based on **[Protocol Buffers Basics: Go Tutorial](https://developers.google.com/protocol-buffers/docs/gotutorial)** with _addition_ of gRPC server up and running.
### Generate Protobuf files

- To generate messages
    -   `protoc --go_out=. --go_opt=paths=source_relative tutorial/proto/addressbook.proto`
    ***or*** with ***option go_package = "tutorial/proto";*** setted in proto file
    -   `protoc --go-grpc_out=.  tutorial/proto/addressbook.proto`    
- To generate services
    -   `protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative tutorial/proto/addressbook.proto` 
    ***or***
    -   `protoc --go-grpc_out=. tutorial/proto/addressbook.proto`

### Server

Implementations of interfaces generated in `addressbook_grpc.pb.go` accordingly to services definitions in proto file.

### Testing

As the gRPC server deppends only on interfaces we can easily testing its methods in isolation.

### Client

The client uses the proto genereated file to make calls into server methods.

---

## Pattern 1 - Unary RPC



