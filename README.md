# logging

`logging` is a Golang library that defines a very simple logger interface conforming with logging best practices and
[Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). It's especially
useful as a standardized logger for microservices and libraries.

The logger has 4 levels:

- Debug
- Info
- Warn
- Error

All levels have 3 types of functions:

- Simple — analogous to log.Print
- According to a format — analogous to log.Printf
- With key-value pairs

It's compatible with [zap](https://github.com/uber-go/zap) out of the box. If you want to use this interface with
another logger you can make a wraper for it that implements Logger.

## Installation

```sh
go get -u github.com/happybydefault/logging
```

## More information

- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Go Microservice with Clean Architecture: Application Logging](https://medium.com/@jfeng45/go-microservice-with-clean-architecture-application-logging-b43dc5839bce)
