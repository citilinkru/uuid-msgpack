package uuidmsgpack

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/vmihailenco/msgpack.v2"
	"os"
)

func ExampleInit() {
	id := uuid.New()
	buf := bytes.NewBuffer(nil)
	encoder := msgpack.NewEncoder(buf)
	err := encoder.Encode(id)
	if err != nil {
		fmt.Printf("can't encode uuid: %s", err)
		os.Exit(1)
	}
	reader := bytes.NewReader(buf.Bytes())
	decoder := msgpack.NewDecoder(reader)
	var decodedId uuid.UUID
	err = decoder.Decode(&decodedId)
	if err != nil {
		fmt.Printf("can't decode uuid: %s", err)
		os.Exit(1)
	}

	fmt.Printf("original id is equal to decoded: %t", id.String() == decodedId.String())

	// Output: original id is equal to decoded: true
}
