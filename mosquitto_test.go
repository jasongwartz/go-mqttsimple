package mqttsimple_test

import (
	"testing"

	"github.com/jasongwartz/go-mqttsimple"
)

func TestCanCreateMosquittoObject(t *testing.T) {
	mosquitto := mqttsimple.Mosquitto{Broker: "test-broker"}

	if mosquitto.Broker == "" {
		t.Fatalf("Couldn't correctly create the mosquitto object: %s", mosquitto)
	}
}
