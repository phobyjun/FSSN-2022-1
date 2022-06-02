# 풀스택서비스네트워킹(FSSN) 프로젝트

## Overview

- python 기반으로 작성된 gRPC 예제를 golang 기반으로 동일하게 개발함
- 예제는 총 4개로 다음과 같음
  - [hello_gRPC](https://github.com/phobyjun/FSSN-2022-1/tree/master/grpc/lec-07-prg-01-hello_gRPC)
  - [bidirectional-streaming](https://github.com/phobyjun/FSSN-2022-1/tree/master/grpc/lec-07-prg-02-bidirectional-streaming)
  - [clientstreaming](https://github.com/phobyjun/FSSN-2022-1/tree/master/grpc/lec-07-prg-03-clientstreaming)
  - [serverstreaming](https://github.com/phobyjun/FSSN-2022-1/tree/master/grpc/lec-07-prg-04-serverstreaming)

## Getting Started

### Requirements

- `Go` >= 1.16
- `protoc` == version3
- Install GO plugins for the protocol compiler:
   ```shell
   # Install protocol compiler plugins
   $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
   $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
   
   # Update PATH for protoc
   $ export PATH="$PATH:$(go env GOPATH)/bin"
   ```

### Run Locally
1. Clone this repository and open a shell in `/grpc/` path
    ```shell
    $ git clone https://github.com/phobyjun/FSSN-2022-1.git
    $ cd ./FSSN-2022-1/grpc
    ```
2. Compile and execute the server code:
    ```shell
    $ go run lec-07-prg-01-hello_gRPC/server/server.go
    ```
3. From a different terminal, compile and execute the client code:
    ```shell
    $ go run lec-07-prg-01-hello_gRPC/client/client.go
    ```
4. Run 1~2 for other three directory
    ```shell
    # For bidirectional streaming
    $ go run lec-07-prg-02-bidirectional-streaming/server/server.go
    $ go run lec-07-prg-02-bidirectional-streaming/client/client.go
    
    # For client streaming
    $ go run lec-07-prg-03-clientstreaming/server/server.go
    $ go run lec-07-prg-03-clientstreaming/client/client.go
    
    # For server streaming
    $ go run lec-07-prg-04-serverstreaming/server/server.go
    $ go run lec-07-prg-04-serverstreaming/client/client.go
    ```
---
> If you want to regenerate gRPC code with `.protoc`, execute next command:
```shell
# if you want to regenerate others, change .protoc file name
$ cd lec-07-prg-01-hello_gRPC
$ protoc -I . --go_out=. --go-grpc_out=. hello_grpc.proto
```

## Demo Videos
- [lec-07-prg-01-hello_gRPC.mov](https://github.com/phobyjun/FSSN-2022-1/blob/master/grpc/lec-07-prg-01-hello_gRPC.mov)
- [lec-07-prg-02-bidirectional-streaming.mov](https://github.com/phobyjun/FSSN-2022-1/blob/master/grpc/lec-07-prg-02-bidirectional-streaming.mov)
- [lec-07-prg-03-clientstreaming.mov](https://github.com/phobyjun/FSSN-2022-1/blob/master/grpc/lec-07-prg-03-clientstreaming.mov)
- [lec-07-prg-04-serverstreaming.mov](https://github.com/phobyjun/FSSN-2022-1/blob/master/grpc/lec-07-prg-04-serverstreaming.mov)