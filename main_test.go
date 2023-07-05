package avro_benchmarks

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/heetch/avro"
)

var Schema string

var Payload []byte

// The spec for Avro make blocks (arrays) size optional,
// this payload does not include it.
var PayloadSimple []byte

// Heetch specific in-memory schema registry, implements DecodingRegistry and
// EncodingRegistry by associating a single-byte schema ID with schemas.
type memRegistry map[int64]*avro.Type

var Registry memRegistry

type Superhero struct {
	ID            int32         `avro:"id"`
	AffiliationID int32         `avro:"affiliation_id"`
	Name          string        `avro:"name"`
	Life          float32       `avro:"life"`
	Energy        float32       `avro:"energy"`
	Powers        []*Superpower `avro:"powers"`
}

type Superpower struct {
	ID      int32   `avro:"id"`
	Name    string  `avro:"name"`
	Damage  float32 `avro:"damage"`
	Energy  float32 `avro:"energy"`
	Passive bool    `avro:"passive"`
}

func TestMain(m *testing.M) {
	// Read in the schema
	schema, err := os.ReadFile("fixtures/superhero.avsc")
	if err != nil {
		log.Fatal(err)
	}
	Schema = string(schema)

	// Read in the payload
	payload, err := os.ReadFile("fixtures/superhero.bin")
	if err != nil {
		log.Fatal(err)
	}
	Payload = payload

	// Read in the payload
	payload, err = os.ReadFile("fixtures/superhero-simple-block.bin")
	if err != nil {
		log.Fatal(err)
	}
	PayloadSimple = payload

	// Initialise the in-memory registry for Heetch
	bytes, err := os.ReadFile("fixtures/superhero.avsc")
	if err != nil {
		panic(fmt.Sprintf("failed reading schema file: %v", err))
	}
	Registry = map[int64]*avro.Type{1: mustParseType(string(bytes))}

	os.Exit(m.Run())
}

// Heetch specific methods for the schema registry (almost copied verbatim from
// their tests):
func (m memRegistry) DecodeSchemaID(msg []byte) (int64, []byte) {
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
		sType := s.Name()
		schemaType := schema.Name()
		if sType == schemaType {
			return id, nil
		}
	}
	return 0, fmt.Errorf("schema not found")
}

func mustParseType(s string) *avro.Type {
	t, err := avro.ParseType(s)
	if err != nil {
		panic(err)
	}
	return t
}

// End heetch specifics
