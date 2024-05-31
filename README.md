# optional

[![Go Doc](https://pkg.go.dev/badge/github.com/HarutakaMatsumoto/optional_go)](https://pkg.go.dev/github.com/HarutakaMatsumoto/optional_go)
[![Go Report Card](https://goreportcard.com/badge/github.com/HarutakaMatsumoto/optional_go)](https://goreportcard.com/report/github.com/HarutakaMatsumoto/optional_go)

Go module optional provides a way to handle optional values in Go.
Optional is underlain by a pointer of a type.
Therefore, you must explicitly dereference the pointer to access its value like `(*optional)`.
This constraint prevents the nil pointer dereference.

## Features
- The most simple API
- Comparable with nil

## Prerequisites

- **[Go](https://go.dev)**: **1.22+**.

## Installation

Simply append the following import to your code, and then `go [build|run|test]`
will automatically fetch the necessary dependencies:

```go
import "github.com/HarutakaMatsumoto/optional_go"
```

## Contribution

You are very welcome to:

- Create pull requests of any kind
- Let me know if you are using this library and find it useful
- Open issues with request for support because they will help you and many others
