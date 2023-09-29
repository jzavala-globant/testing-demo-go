# testing-demo-go

An example of how to perform tests using either test tables strategy or suite.

## *Testfiles

**This repository includes multiple testing files for the same service (`bookstore.go`). This is intended ONLY as an example of multiple strategies, but in practice, there should only be one (ideally `bookstore_test.go`).**

## Commands

- Run the project locally (default port 8080):
```sh
make run
```

Now you can test the API requesting this URL in any browser or any other requester tool:
```
http://localhost:8080/books
```
---
- Run the **go test** on the `./internal/` folder. This also generates the `output/heatmap.svg`:
```sh
make test
```
---
- Autogenerates the mock for the **services** into `./internal/services/mocks`:
```sh
make mock
```
---
- Generates the `output/graph.svg` with the package dependencies graph:
```sh
make show-dependencies
```
---
- Run the **go test** on the `./internal/` folder and displays a table with the sorted results of the impact coverage:
```sh
make coverage
```

## Dependencies

Use `go install` get the tools to generate the coverage outputs:

- [**goda**: to show dependencies between packages](https://github.com/loov/goda)
- [**go-cover-treemap**: to get the coverage heatmap](https://github.com/nikolaydubina/go-cover-treemap)
- [**go-coverage**: to get table of coverage impact](https://github.com/gojek/go-coverage)
