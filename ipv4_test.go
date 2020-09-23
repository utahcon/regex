package ipv4

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		address string
		pass    bool
	}{
		{"0.0.0.0", true},
		{"0.0.0.0/1", true},
		{"0.0.0.0/32", true},
		{"1.1.1.1", true},
		{"12.12.12.12", true},
		{"20.20.20.255", true},
		{"88.1.1.1", true},
		{"123.0.0.0", true},
		{"220.220.220.220", true},
		{"240.240.240.240", true},
		{"255.255.255.255", true},
		{"255.255.255.255/1", true},
		{"255.255.255.255/31", true},
		{"0.0.0.0/33", false},
		{"1.1.1.1/33", false},
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
