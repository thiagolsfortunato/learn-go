package address

import (
	"strings"
)

// AddressType checks if an address has a valid type and returns it
func AddressType(address string) string {
	validTypes := []string{"road", "street", "avenue", "boulevard"}

	firstAddressWorld := strings.Split(strings.ToLower(address), " ")[0]

	addressIsValid := false

	for _, t := range validTypes {
		if t == firstAddressWorld {
			addressIsValid = true
		}
	}

	if addressIsValid {
		return strings.Title(firstAddressWorld)
	}

	return "Invalid Type"
}
