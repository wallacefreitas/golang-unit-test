package address

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TypeAddress(address string) string {
	typeAddress := []string{"rua", "avenida", "estrada", "rodovia"}

	isValidTypeAddress := false

	addressWithLowerLetter := strings.ToLower(address)
	firstWordAddress := strings.Split(addressWithLowerLetter, " ")[0]

	for _, types := range typeAddress {
		if firstWordAddress == types {
			isValidTypeAddress = true
		}
	}

	if !isValidTypeAddress {
		return "Address type invalid"
	}

	return cases.Title(language.English, cases.Compact).String(firstWordAddress)

}
