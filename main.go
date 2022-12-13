package main

import (
	"context"
	"fmt"

	"github.com/heetch/avro"
)

type memRegistry map[int64]*avro.Type

func (m memRegistry) DecodeSchemaID(msg []byte) (id int64, header []byte) {
	if len(msg) < 1 {
		return 0, nil
	}
	return int64(msg[0]), msg[1:]
}

func (m memRegistry) SchemaForID(ctx context.Context, id int64) (*avro.Type, error) {
	t, ok := m[id]
	if !ok {
		return nil, fmt.Errorf("schema not found for id %d", id)
	}
	return t, nil
}

func (m memRegistry) AppendSchemaID(buf []byte, id int64) []byte {
	if id < 0 || id > 256 {
		panic("schema ID out of range")
	}
	return append(buf, byte(id))
}

func (m memRegistry) IDForSchema(ctx context.Context, schema *avro.Type) (int64, error) {
	for id, s := range m {
		if s.String() == schema.String() {
			return id, nil
		}
	}
	return 0, fmt.Errorf("schema not found")
}

func main() {
	avroType, err := avro.TypeOf(Data{})
	if err != nil {
		panic(err)
	}
	registry := memRegistry{
		1: avroType,
	}
	enc := avro.NewSingleEncoder(registry, nil)
	data, err := enc.Marshal(context.Background(), Data{})
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
