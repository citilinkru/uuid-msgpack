// Package uuidmsgpack register msgpack encoder/decoder for uuid lib
//
// Just import that package in your main and you can encode/decode like any struct
package uuidmsgpack

import (
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/vmihailenco/msgpack.v2"
	"gopkg.in/vmihailenco/msgpack.v2/codes"
	"reflect"
)

// extension id
const extId = 2

// bytes count to read to start decode
const decodeBytesCount = 18

func init() {
	msgpack.Register(reflect.TypeOf((*uuid.UUID)(nil)).Elem(),
		func(e *msgpack.Encoder, v reflect.Value) error {
			id := v.Interface().(uuid.UUID)
			bytes, err := id.MarshalBinary()
			if err != nil {
				return fmt.Errorf("can't marshal binary uuid: %w", err)
			}
			_, err = e.Writer().Write(bytes)
			if err != nil {
				return fmt.Errorf("can't write bytes to writer: %w", err)
			}

			return nil
		},
		func(d *msgpack.Decoder, v reflect.Value) error {
			bytes := make([]byte, decodeBytesCount)
			n, err := d.Buffered().Read(bytes)
			if err != nil {
				return fmt.Errorf("can't read bytes: %w", err)
			}
			if n < decodeBytesCount {
				return fmt.Errorf("invalid bytes count %d instead of %d", n, decodeBytesCount)
			}

			if bytes[0] != codes.FixExt16 {
				return fmt.Errorf("wrong ext len '%d'", bytes[0])
			}

			if bytes[1] != extId {
				return fmt.Errorf("wrong ext id '%d'", bytes[1])
			}

			id, err := uuid.FromBytes(bytes[2:18])
			if err != nil {
				return fmt.Errorf("can't create uuid from bytes: %w", err)
			}

			v.Set(reflect.ValueOf(id))
			return nil
		},
	)
	msgpack.RegisterExt(extId, (*uuid.UUID)(nil))
}
