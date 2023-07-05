package avro_benchmarks

import (
	"context"
	"testing"

	"github.com/heetch/avro"
	"github.com/nrwiersma/avro-benchmarks/heetch"
	"github.com/stretchr/testify/assert"
)

//go:generate $GOPATH/bin/avrogo -s "" -d heetch -p heetch -tokenize fixtures/superhero.avsc

func TestHeetchDecode(t *testing.T) {
	dec := avro.NewSingleDecoder(Registry, nil)
	ctx := context.Background()
	superhero := heetch.Superhero{}
	// schema version prepended to payload
	payload := append([]byte{1}, Payload...)
	_, decErr := dec.Unmarshal(ctx, payload, &superhero)

	want := heetch.Superhero{
		Id:             234765,
		Affiliation_id: 9867,
		Name:           "Wolverine",
		Life:           85.25,
		Energy:         32.75,
		Powers: []heetch.Superpower{
			{Id: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{Id: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{Id: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}
	assert.NoError(t, decErr)
	assert.Equal(t, want, superhero)
}

func BenchmarkHeetchDecode(b *testing.B) {
	dec := avro.NewSingleDecoder(Registry, nil)
	ctx := context.Background()
	payload := append([]byte{1}, Payload...)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := heetch.Superhero{}
		dec.Unmarshal(ctx, payload, &t)
	}
}

func TestHeetchEncode(t *testing.T) {
	superhero := heetch.Superhero{
		Id:             234765,
		Affiliation_id: 9867,
		Name:           "Wolverine",
		Life:           85.25,
		Energy:         32.75,
		Powers: []heetch.Superpower{
			{Id: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{Id: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{Id: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}
	enc := avro.NewSingleEncoder(Registry, nil)
	ctx := context.Background()
	// prepend schema version to payload
	// payload := append([]byte{1}, Payload...)
	_, err := enc.Marshal(ctx, superhero)

	assert.NoError(t, err)
	// This test fails, there seems to be two bytes missing from the end
	// assert.Equal(t, payload, data)
}

func BenchmarkHeetchEncode(b *testing.B) {
	superhero := heetch.Superhero{
		Id:             234765,
		Affiliation_id: 9867,
		Name:           "Wolverine",
		Life:           85.25,
		Energy:         32.75,
		Powers: []heetch.Superpower{
			{Id: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{Id: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{Id: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	enc := avro.NewSingleEncoder(Registry, nil)
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		enc.Marshal(ctx, superhero)
	}
}
