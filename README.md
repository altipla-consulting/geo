
# geo

[![GoDoc](https://godoc.org/github.com/altipla-consulting/geo?status.svg)](https://godoc.org/github.com/altipla-consulting/geo)

> Customized types and functions for our geo needs.

**NOTE:** If you want a full-fledged geo library we recommend using https://github.com/twpayne/go-geom instead.


### Install

```shell
go get github.com/altipla-consulting/geo
```

This library has no external dependencies outside the Go standard library.


### Usage

You can use the types of this package in your models structs when working with `database/sql`:

```go
type MyModel struct {
  ID          int64      `db:"id,pk"`
  Location        geo.Point `db:"location"`
}
```


### Contributing

You can make pull requests or create issues in GitHub. Any code you send should be formatted using ```gofmt```.


### Running tests

Start the test database:

```shell
docker-compose up -d database
```

Install test libs:

```shell
go get github.com/stretchr/testify
```

Run the tests:

```shell
go test
```

Shutdown the database when finished testing:

```shell
docker-compose stop database
```


### License

[MIT License](LICENSE)
