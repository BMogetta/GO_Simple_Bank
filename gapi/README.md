# **[gRPC](https://grpc.io/docs/)**

In gRPC, a client application can directly call a method on a server application on a different machine as if it were a local object, making it easier for you to create distributed applications and services. As in many RPC systems, gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types. On the server side, the server implements this interface and runs a gRPC server to handle client calls. On the client side, the client has a stub (referred to as just a client in some languages) that provides the same methods as the server.


---

## **[evans:](https://github.com/ktr0731/evans) gRPC client for testing**

Evans has been created to use easier than other existing gRPC clients.
```
curl -Lo evans.tgz https://github.com/ktr0731/evans/releases/download/v0.10.9/evans_linux_amd64.tar.gz

tar -xf evans.tgz

sudo install evans /usr/local/bin/

evans --version

```
  
---
## **[gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/)**

```sh
  go install \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
  google.golang.org/protobuf/cmd/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc
```