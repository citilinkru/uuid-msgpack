# UUID msgpack

![Lint](https://github.com/citilinkru/uuid-msgpack/actions/workflows/golangci-lint.yml/badge.svg?branch=master)
![Tests](https://github.com/citilinkru/uuid-msgpack/actions/workflows/test.yml/badge.svg?branch=master)
[![codecov](https://codecov.io/gh/citilinkru/uuid-msgpack/branch/master/graph/badge.svg)](https://codecov.io/gh/citilinkru/uuid-msgpack)
[![Go Report Card](https://goreportcard.com/badge/github.com/citilinkru/uuid-msgpack)](https://goreportcard.com/report/github.com/citilinkru/uuid-msgpack)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/citilinkru/uuid-msgpack/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/citilinkru/uuid-msgpack?status.svg)](https://godoc.org/github.com/citilinkru/uuid-msgpack)
[![Release](https://img.shields.io/github/release/citilinkru/uuid-msgpack.svg?style=flat-square)](https://github.com/citilinkru/uuid-msgpack/releases/latest)

Library to integrate github.com/google/uuid with gopkg.in/vmihailenco/msgpack


Installation
------------
	go get github.com/citilinkru/uuid-msgpack

Example
-------
```go
package main

import (
	_ "github.com/citilinkru/uuid-msgpack" // This blank import is required to register encoder/decoder!
	"github.com/google/uuid"
	"log"
	"bytes"
	"gopkg.in/vmihailenco/msgpack.v2" 
)

func main() {
	id := uuid.New()
	buf := bytes.NewBuffer(nil)
	encoder := msgpack.NewEncoder(buf)
	err := encoder.Encode(id)
	if err != nil {
		log.Fatalf("can't encode uuid: %s", err)
	}
	reader := bytes.NewReader(buf.Bytes())
	decoder := msgpack.NewDecoder(reader)
	var decodedId uuid.UUID
	err = decoder.Decode(&decodedId)
	if err != nil {
		log.Fatalf("can't decode uuid: %s", err)
	}

	if id.String() == decodedId.String() {
		log.Printf("original id is equal to decoded '%s'\n", id)
	} else {
		log.Printf("decoded id '%s' is not equal to origin '%s'\n", id, decodedId)
	}
	
	log.Println("done")
}
```

Testing
-----------
Unit-tests:
```bash
go test -v -race ./...
```

Run linter:
```bash
go mod vendor \
  && docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.40 golangci-lint run -v \
  && rm -R vendor
```

CONTRIBUTE
-----------
* write code
* run `go fmt ./...`
* run all linters and tests (see above)
* create a PR describing the changes

LICENSE
-----------
MIT

AUTHOR
-----------
Nikita Sapogov <amstaffix@gmail.com>