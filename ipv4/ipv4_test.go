package ipv4

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		address string
		pass    bool
	}{
		{"0.0.0.0", true},
		{"0.0.0.0/1", true},
		{"0.0.0.0/24", true},
		{"0.0.0.0/32", true},
		{"1.1.1.1", true},
		{"1.1.1.1/1", true},
		{"1.1.1.1/24", true},
		{"1.1.1.1/32", true},
		{"12.12.12.12", true},
		{"12.12.12.12/1", true},
		{"12.12.12.12/24", true},
		{"12.12.12.12/32", true},
		{"255.255.255.255", true},
		{"255.255.255.255/1", true},
		{"255.255.255.255/24", true},
		{"255.255.255.255/32", true},
		{"0.0.0.0/33", false},
		{"0.0.0.255/33", false},
		{"256.0.0.0/32", false},
		{"0.256.0.0/32", false},
		{"0.0.256.0/32", false},
		{"0.0.0.256/32", false},
	}

	for _, test := range tests {
		result := Validate(test.address)
		if test.pass != result {
			t.Errorf("Address %s recieved %t instead of %t", test.address, result, test.pass)
		}
	}
}
