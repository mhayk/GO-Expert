# Dependency Injections

Dependency injection is a technique whereby one object supplies the dependencies of another object. A dependency is an object that can be used (a service). An injection is the passing of a dependency to a dependent object (a client) that would use it. The service is made part of the client's state. Passing the service to the client, rather than allowing a client to build or find the service, is the fundamental requirement of the pattern.

## Google/Wire

Wire is a code generation tool that automates connecting components using Google's DI framework for Go. It generates human-readable code that cleanly separates the injector code from the components that declare their dependencies. This makes the injector code easier to read and maintain.

### Installation

```bash
$ go install github.com/google/wire/cmd/wire@latest
$ ls ~/go/bin
$ wire
$ go run main.go wire_gen.go
```
