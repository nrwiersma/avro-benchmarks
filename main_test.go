package avro_benchmarks

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var Schema string

var Payload []byte

// The spec for Avro make blocks (arrays) size optional,
// this payload does not include it.
var PayloadSimple []byte

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
	schema, err := ioutil.ReadFile("fixtures/superhero.avsc")
	if err != nil {
		log.Fatal(err)
	}
	Schema = string(schema)

	// Read in the payload
	payload, err := ioutil.ReadFile("fixtures/superhero.bin")
	if err != nil {
		log.Fatal(err)
	}
	Payload = payload

	// Read in the payload
	payload, err = ioutil.ReadFile("fixtures/superhero-simple-block.bin")
	if err != nil {
		log.Fatal(err)
	}
	PayloadSimple = payload

	os.Exit(m.Run())
}
