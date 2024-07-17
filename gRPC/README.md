# gRPC

## Protocol buffer compiler
### Linux
```shell
$ sudo apt install protobuf-compiler
```
### MacOS
```shell
$ brew install protobuf
$ protoc --version
libprotoc 27.1
```

## Go plugins
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Generate Go code
```shell
$ protoc --go_out=. --go-grpc_out=. proto/course_category.proto
$ go mod tidy
```

## gRPC Client
```shell
$ brew tap ktr0731/evans
$ brew install evans
$ evans -r repl
> package pb
> service CourseCategoryService
> call CreateCategory
> call GetCourseCategory
```