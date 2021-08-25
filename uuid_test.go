package uuidmsgpack

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/vmihailenco/msgpack.v2"
	"gopkg.in/vmihailenco/msgpack.v2/codes"
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

func TestErrorCantReadBytes(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	reader := bytes.NewReader(buf.Bytes())
	decoder := msgpack.NewDecoder(reader)
	var decodedId uuid.UUID
	err := decoder.Decode(&decodedId)

	assert.EqualError(t, err, "can't read bytes: EOF")
}

func TestErrorInvalidBytesCount(t *testing.T) {
	buf := bytes.NewBuffer([]byte{0x7d})

	reader := bytes.NewReader(buf.Bytes())
	decoder := msgpack.NewDecoder(reader)
	var decodedId uuid.UUID
	err := decoder.Decode(&decodedId)

	assert.EqualError(t, err, "invalid bytes count 1 instead of 18")
}

func TestErrorWrongExtLen(t *testing.T) {
	buf := bytes.NewBuffer([]byte{
		0x7d, 0x44, 0x48, 0x40,
		0x9d, 0xc0,
		0x11, 0xd1,
		0xb2, 0x45,
		0x5f, 0xfd, 0xce, 0x74, 0xfa, 0xd2,
		0xd3, 0xd4,
	})

	reader := bytes.NewReader(buf.Bytes())
	decoder := msgpack.NewDecoder(reader)
	var decodedId uuid.UUID
	err := decoder.Decode(&decodedId)

	assert.EqualError(t, err, "wrong ext len '125'")
}

func TestErrorWrongExtId(t *testing.T) {
	buf := bytes.NewBuffer([]byte{
		codes.FixExt16, 0x44, 0x48, 0x40,
		0x9d, 0xc0,
		0x11, 0xd1,
		0xb2, 0x45,
		0x5f, 0xfd, 0xce, 0x74, 0xfa, 0xd2,
		0xd3, 0xd4,
	})

	reader := bytes.NewReader(buf.Bytes())
	decoder := msgpack.NewDecoder(reader)
	var decodedId uuid.UUID
	err := decoder.Decode(&decodedId)

	assert.EqualError(t, err, "wrong ext id '68'")
}
