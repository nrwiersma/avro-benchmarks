package avro_benchmarks

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/avro.v0"
)

func TestGoAvroDecode(t *testing.T) {
	schema := avro.MustParseSchema(Schema)

	reader := avro.NewSpecificDatumReader()
	reader.SetSchema(schema)

	decoder := avro.NewBinaryDecoder(Payload)

	superhero := Superhero{}
	err := reader.Read(&superhero, decoder)

	want := Superhero{
		ID:            234765,
		AffiliationID: 9867,
		Name:          "Wolverine",
		Life:          85.25,
		Energy:        32.75,
		Powers: []*Superpower{
			{ID: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{ID: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{ID: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}
	assert.NoError(t, err)
	assert.Equal(t, want, superhero)
}

func BenchmarkGoAvroDecode(b *testing.B) {
	schema := avro.MustParseSchema(Schema)

	reader := avro.NewSpecificDatumReader()
	reader.SetSchema(schema)

	superhero := Superhero{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder := avro.NewBinaryDecoder(Payload)
		_ = reader.Read(&superhero, decoder)
	}
}

func TestGoAvroEncode(t *testing.T) {
	schema := avro.MustParseSchema(Schema)

	writer := avro.NewSpecificDatumWriter()
	writer.SetSchema(schema)

	buf := &bytes.Buffer{}
	encoder := avro.NewBinaryEncoder(buf)

	superhero := Superhero{
		ID:            234765,
		AffiliationID: 9867,
		Name:          "Wolverine",
		Life:          85.25,
		Energy:        32.75,
		Powers: []*Superpower{
			{ID: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{ID: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{ID: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}

	err := writer.Write(&superhero, encoder)

	assert.NoError(t, err)
	assert.Equal(t, PayloadSimple, buf.Bytes())
}

func BenchmarkGoAvroEncode(b *testing.B) {
	schema := avro.MustParseSchema(Schema)

	writer := avro.NewSpecificDatumWriter()
	writer.SetSchema(schema)

	buf := &bytes.Buffer{}
	encoder := avro.NewBinaryEncoder(buf)

	superhero := Superhero{
		ID:            234765,
		AffiliationID: 9867,
		Name:          "Wolverine",
		Life:          85.25,
		Energy:        32.75,
		Powers: []*Superpower{
			{ID: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{ID: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{ID: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = writer.Write(&superhero, encoder)
	}
}
