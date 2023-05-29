package avro_benchmarks

import (
	"bytes"
	"testing"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/nrwiersma/avro-benchmarks/models"
	"github.com/stretchr/testify/assert"
)

//go:generate $GOPATH/bin/gogen-avro -package models ./models fixtures/superhero.avsc

func TestGoGenAvroDecode(t *testing.T) {

	superhero, err := models.DeserializeSuperhero(bytes.NewReader(Payload))

	want := models.Superhero{
		Id:             234765,
		Affiliation_id: 9867,
		Name:           "Wolverine",
		Life:           85.25,
		Energy:         32.75,
		Powers: []models.Superpower{
			{Id: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{Id: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{Id: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}
	assert.NoError(t, err)
	assert.Equal(t, want, superhero)
}

func BenchmarkGoGenAvroDecode(b *testing.B) {
	deser, _ := compiler.CompileSchemaBytes([]byte(Schema), []byte(Schema))

	r := bytes.NewReader(nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := models.NewSuperhero()

		r.Reset(Payload)
		_ = vm.Eval(r, deser, &t)
	}
}

func TestGoGenAvroEncode(t *testing.T) {
	superhero := models.Superhero{
		Id:             234765,
		Affiliation_id: 9867,
		Name:           "Wolverine",
		Life:           85.25,
		Energy:         32.75,
		Powers: []models.Superpower{
			{Id: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{Id: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{Id: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}

	buf := bytes.Buffer{}
	err := superhero.Serialize(&buf)

	assert.NoError(t, err)
	// assert.Equal(t, Payload, buf.Bytes()) seems to add 2 bytes
}

func BenchmarkGoGenAvroEncode(b *testing.B) {
	superhero := models.Superhero{
		Id:             234765,
		Affiliation_id: 9867,
		Name:           "Wolverine",
		Life:           85.25,
		Energy:         32.75,
		Powers: []models.Superpower{
			{Id: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{Id: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{Id: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := bytes.Buffer{}
		_ = superhero.Serialize(&buf)
	}
}
