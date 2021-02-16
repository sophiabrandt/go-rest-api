# go-rest-api

go-rest-api is a playground for me to explore how to write Go micro-services from scratch.

## Installation

The project uses Go modules, so you'll need Go >= 1.11.

You'll also need [SQLite](https://www.sqlite.org/index.html).

```bash
git clone https://github.com/sophiabrandt/go-rest-api.git
```

## Usage

1. Migrate and seed the SQLite database.

```bash
go run ./cmd/admin -action="migrate"
go run ./cmd/admin -action="seed"
```

Or use [`Makefile`](Makefile):

```bash
make seed
```

2. Run the server. Default port is 4000, you can change it via command line flag.

```bash
go run ./cmd/server
# or go run ./cmd/server -addr=":8000"
```

Or use [`Makefile`](Makefile):

```bash
make run
```

Example endpoints:

* [`http://localhost:4000/api/health`](http://localhost:4000/api/health)
* [`http://localhost:4000/api/books`](http://localhost:4000/api/books)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

Apache 2.0 License, see [`LICENSE`](LICENSE).

## Acknowledgments

- [https://github.com/dlsniper/gopherconuk](https://github.com/dlsniper/gopherconuk)
- [How I write HTTP services after eight years.](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html)
- [http.Handler and Error Handling in Go](https://blog.questionable.services/article/http-handler-error-handling-revisited/)
- [Develop A Production Ready REST API in Go](https://tutorialedge.net/courses/go-rest-api-course/)
- [https://github.com/ardanlabs/service](https://github.com/ardanlabs/service)
- [Graceful shutdown of Golang servers using Context and OS signals](https://archive.is/Mf0dJ)
- [Learning Cloud Native Go](https://learning-cloud-native-go.github.io/)
