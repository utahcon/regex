package ipv6

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		address string
		pass    bool
	}{
		{"::", true},
		{"::/0", true},
		{"::/64", true},
		{"::/128", true},
		{"::1/1", true},
		{"::1/24", true},
		{"::1/128", true},
		{"2001:db8::", true},
		{"2001:db8::/1", true},
		{"2001:db8::/64", true},
		{"2001:db8::/128", true},
		{"::1234:5678", true},
		{"::1234:5678/4", true},
		{"::1234:5678/64", true},
		{"::1234:5678/128", true},
		{"2001:db8::1234:5678", true},
		{"2001:db8::1234:5678/1", true},
		{"2001:db8::1234:5678/10", true},
		{"2001:db8::1234:5678/128", true},
		{"2001:db8:1::ab9:C0A8:102", true},
		{"2001:db8:1::ab9:C0A8:102/1", true},
		{"2001:db8:1::ab9:C0A8:102/64", true},
		{"2001:db8:1::ab9:C0A8:102/128", true},
		{"2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF", true},
		{"2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF/1", true},
		{"2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF/10", true},
		{"2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF/128", true},
		{"2001:0db8:0001:0000:0000:0ab9:C0A8:0102", true},
		{"2001:0db8:0001:0000:0000:0ab9:C0A8:0102/1", true},
		{"2001:0db8:0001:0000:0000:0ab9:C0A8:0102/10", true},
		{"2001:0db8:0001:0000:0000:0ab9:C0A8:0102/128", true},
		{"2001:db8:3333:4444:5555:6666:7777:8888", true},
		{"2001:db8:3333:4444:5555:6666:7777:8888/4", true},
		{"2001:db8:3333:4444:5555:6666:7777:8888/64", true},
		{"2001:db8:3333:4444:5555:6666:7777:8888/128", true},

		{":::", false},
		{"2001:db8:1:::ab9:C0A8:102", false},
		{"2001:db8::1234:5678/129", false},
		{"2001:db8:3333:4444:5555:6666:7777:8888/129", false},
		{":::", false},
		{":::/0", false},
		{":::/64", false},
		{":::/128", false},
		{":::1/1", false},
		{"::1/129", false},
		{"2001::db8::", false},
		{"2001:db8::/129", false},
		{":::1234:5678", false},
		{"::1234:5678/129", false},
		{"2001:db8:::1234:5678", false},
		{"2001:db8::1234:5678/129", false},
		{"2001:db8:1:::ab9:C0A8:102/128", false},
		{"2001:db8:1::ab9::C0A8:102/128", false},
		{"2001:db8:1::ab9:C0A8:102/129", false},
		{"2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF/129", false},
		{"2001:0db8:0001:0000:0000:0ab9:C0A8:0102/129", false},
		{"2001:db8:3333:4444:5555:6666:7777:8888/129", false},

	}

	for _, test := range tests {
		result := Validate(test.address)
		if test.pass != result {
			t.Errorf("Address %s recieved %t instead of %t", test.address, result, test.pass)
		}
	}
}
