package sendmail

import (
	"errors"
	"net"
	"strings"
)

// Validate checks if email is valid
func Validate(email string) error {
	emailParts := strings.SplitN(email, "@", 3)
	if len(emailParts) != 2 {
		return errors.New("Email format is incorrect")
	}
	mxlist, _ := net.LookupMX(emailParts[1])
	if len(mxlist) > 0 {
		return nil
	}
	iplist, err := net.LookupIP(emailParts[1])
	if err != nil {
		return err
	}
	if len(iplist) > 0 {
		return nil
	}
	return errors.New("No assigned IP")
}
