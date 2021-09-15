# Go Playground

Yet another playground project. This time to learn a bit of Go! Know the business details about this project looking at the [docs](./docs) folder.

Soon I'll give you more information.

## How to manage dependencies

If you want to install packages:

    go get -u gorm.io/gorm
    go get -u gorm.io/driver/postgres

You can know how it works with the following command:

    go help get

Let's say [you install a package, and you want to remove it](https://stackoverflow.com/a/57186981/3899136), given it's not being used anywhere, you can execute the command below:

    go mod tidy

## Links

Official:

- [The Go Programming Language Specification](https://golang.org/ref/spec)
- [Effective Go](https://golang.org/doc/effective_go)

Others:

- [How to Structure a Golang Project](https://qvault.io/golang/golang-project-structure/)
