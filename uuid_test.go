package uuidmsgpack

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/vmihailenco/msgpack.v2"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	id := uuid.New()
	buf := bytes.NewBuffer(nil)
	encoder := msgpack.NewEncoder(buf)
	err := encoder.Encode(id)
	if err != nil {
		assert.NoError(t, err)
	}
	reader := bytes.NewReader(buf.Bytes())
	decoder := msgpack.NewDecoder(reader)
	var decodedId uuid.UUID
	err = decoder.Decode(&decodedId)
	assert.NoError(t, err)
	assert.Equal(t, id.String(), decodedId.String())
}

type customStruct struct {
	a    int
	Uuid uuid.UUID
	b    string
}

func TestEncodeDecodeCustomStruct(t *testing.T) {
	obj := customStruct{Uuid: uuid.New(), a: 1, b: "adsa"}
	buf := bytes.NewBuffer(nil)
	encoder := msgpack.NewEncoder(buf)
	err := encoder.Encode(obj)
	if err != nil {
		assert.NoError(t, err)
	}
	reader := bytes.NewReader(buf.Bytes())
	decoder := msgpack.NewDecoder(reader)
	var decodedObj customStruct
	err = decoder.Decode(&decodedObj)
	assert.NoError(t, err)
	assert.Equal(t, obj.Uuid.String(), decodedObj.Uuid.String())
}
