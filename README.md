# UUID msgpack
Library to integrate github.com/google/uuid with gopkg.in/vmihailenco/msgpack


Installation
------------
	go get github.com/citilinkru/camunda-client-go/v2

Example
-------
```go
package main

import (
	_ "code.citik.ru/gobase/uuid-msgpack" // Обязательно нужно сделать импорт!
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