# syn-api

##About

A demo project for the purposes of benchmarking a Golang web CRUD API, with a backing MongoDB database.

##Running the application

###Database
There is a simple `docker-compose` configuration included in the project that will run a MongoDB container to serve as a local test database. 

Open a terminal and run `docker-compose up` to run the container.

###Server
Running the server with hot-reloading requires an external tool called [CompileDaemon](https://pkg.go.dev/github.com/githubnemo/CompileDaemon#section-readme) to watch for files changes. You can download and install CompileDaemon globally by first running `go get github.com/githubnemo/CompileDaemon`, followed by `go install github.com/githubnemo/CompileDaemon`. Once you have CompileDaemon installed, run `CompileDaemon -command="./syn-api"` to run the compiled executable of the application.

To run the server without hot-reloading, run `go run .`.

##Web Framework

###Fiber

[Fiber](https://docs.gofiber.io/) is a simple, fast and versatile web framework, inspired by Express.js. It is built on top of another Go module called [fasthttp](https://pkg.go.dev/github.com/valyala/fasthttp), which was designed to provide more performance than Go's built in [net/http](https://pkg.go.dev/net/http) module.

This gives Fiber a performance edge over other frameworks in it's class (especially when you enable it's Prefork mode), while it still maintains a focus on simplicity and ease-of-use.

It offers some great features:

* Out-of-the-box concurrency utilizing Go's built in concurrency superpowers

* Minimal allocation of memory resources thanks to a focus on immutability, allowing route handlers to share and re-use common resources

* Some pre-packaged middleware, including logging, rate limiting, sessions, and a version of the popular Node-based middleware, Helmet

##Dependencies

###mongo-driver
The standard MongoDB driver for Go. Documentation is available [here](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#section-documentation).