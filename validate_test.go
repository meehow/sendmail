package sendmail

import (
	"testing"
)

var emails = []*struct {
	Address string
	IsValid bool
}{
	{"x@e.com", false},      // NXDOMAIN
	{"x@example.com", true}, // Has IP, but no MX
	{"x@example.org", true}, // Has IP, but no MX
	{"x@github.com", true},  // Has MX
	{"x@example@", false},   // Double @
	{"example", false},      // No @
}

func TestValidate(tc *testing.T) {
	for _, email := range emails {
		tc.Run(email.Address, func(t *testing.T) {
			e := email
			t.Parallel()
			err := Validate(e.Address)
			if err == nil && e.IsValid == false {
				t.Errorf("Email `%s` is valid, but should be invalid", e.Address)
			} else if err != nil && e.IsValid == true {
				t.Errorf("Email `%s` is invalid, but should be valid", e.Address)
			}
		})
	}
}
