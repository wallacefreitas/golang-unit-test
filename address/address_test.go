package address

import "testing"

type testScenary struct {
	addressSending string
	returnExpected string
}

func TestTypeAddress(t *testing.T) {
	testsScenary := []testScenary{
		{"Rua Teste 1", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia Teste 1", "Rodovia"},
		// {"Praça Teste 1", "Address type invalid"},
		{"Estrada Teste 1", "Estrada"},
		{"RUA TESTE 2", "Rua"},
		{"AVENIDA TESTE 2", "Avenida"},
		// {"", "Address type invalid"},
	}

	for _, scenary := range testsScenary {
		typeAddressReceived := TypeAddress(scenary.addressSending)

		if typeAddressReceived != scenary.returnExpected {
			t.Errorf("The received type %s is different to expected %s", typeAddressReceived, scenary.returnExpected)
		}
	}
}
