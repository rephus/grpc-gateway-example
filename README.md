# Simple gRPC gateway on Go 
This is a simple template service file to help you building a service in Go using 
gRPC but also keeping the API compatible with REST so you can use it standalone for 
web development without any API layer. 

If you want to know more, please visit the official documentation of 
[gPRC](http://www.grpc.io/docs/quickstart/go.html#build-the-example) and 
[gRPC-gateway](https://github.com/grpc-ecosystem/grpc-gateway)

## Quick start, run this example 
If you just want to test how this works, without making any changes to the proto file, or the codebase, do: 

```
> go get -u github.com/rephus/grpc-gateway-example

> cd $GOPATH/src/github.com/rephus/grpc-gateway-example
> go run server/main.go & # Run gRPC server
> go run main.go & # Run Gateway

> curl localhost:8080/get?name=rephus
{"message":"Received GET method rephus"}

> curl -X POST localhost:8080/post -d '{"name": "rephus"}'
{"message":"Received POST method rephus"}
```

## Install dependencies for gRPC gateway 
```
    apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool
    
    mkdir -p /tmp 
    cd /tmp 
    git clone https://github.com/google/protobuf 
    cd protobuf 
    ./autogen.sh 
    ./configure 
    make 
    make check 
    make install

    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway 
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger 
    go get -u github.com/golang/protobuf/protoc-gen-go
```

## Run gRPC server
    go run greeter_server/main.go

## Run REST gateway

Before running gateway, make sure gRPC server is running, 
because the gateway connects to it and converts the JSON

    go run main.go 

## Run gRPC client 

    go run greeter_client/main.go

## Build protoc

If you want to generate everything, just run `make`

### Generate gRPC stub (template.pb.go)
protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --go_out=plugins=grpc:.   template/template.proto

### Generate gateway (template.pb.gw.go)
protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --grpc-gateway_out=logtostderr=true:.  template/template.proto

## Test Gateway 

Make sure you're running `server/main.go` and `main.go` at the same time.

```
GET
> curl localhost:8080/get?name=rephus
{"message":"Received GET method rephus"}
```

```
POST 
> curl -X POST localhost:8080/post -d '{"name": "rephus"}'
{"message":"Received POST method rephus"}
```

You can also access `http://localhost:8080/swagger/template.swagger.json`
to get the API implementation details

## Troubleshooting

### protoc: error while loading shared libraries

```
protoc: error while loading shared libraries: libprotoc.so.13: cannot open shared object file: No such file or directory`
```

Make sure your protoc is working in /usr/bin/protoc and not in /usr/local/bin/protoc

### protoc-gen-go: program not found or is not executable

If any of the protoc-get-* (like protoc-gen-go , protoc-gen-grpc-gateway or protoc-gen-swagger...) commands fail when building the proto file, follow these steps: 

* Make sure you installed the package `go get -u github.com/golang/protobuf/protoc-gen-go`
* That will create a binary on $GOPATH/bin/protoc-gen-go
* Make sure you have `export PATH=$PATH:$GOPATH/bin`
