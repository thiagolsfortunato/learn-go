// Unit Test
package address_test

import (
	"testing"
	. "intro-tests/address"
)

type scenarioTest struct { 
	address string
	expected string
}

func TestAddressType (t *testing.T) {

	t.Parallel()

	addressTest := "Avenue 2sd"
	expected := "Avenue"
	received := AddressType(addressTest)

	if expected != received {
		t.Error("The Address Type is different from the expected")
	}
}

func TestAddressTypeScenario (t *testing.T) {

	t.Parallel()

	scenarios := []scenarioTest {
		{"Street Baker", "Street"},
		{"Avenue 3th", "Avenue"},
		{"Road Route 66", "Road"},
		{"Square Time", "Invalid Type"},
		{"", "Invalid Type"},
		{"STREET BAKER", "Street"},
		{"Boulevard Hollywood", "Boulevard"},
	}

	for _, scenario := range scenarios {
		received := AddressType(scenario.address)
		if received != scenario.expected {
			t.Errorf("The type received %s is different from expetected %s", received, scenario.expected)
		}
	}
}
