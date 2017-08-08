// This package implements classic, well known from PHP, method of sending emails.
package sendmail

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	_, debug = os.LookupEnv("DEBUG")
)

// Send sends an email, or prints it on stderr,
// when environment variable `DEBUG` is set.
func Send(from string, to []string, msg []byte) error {
	if debug {
		var delimiter = strings.Repeat("-", 70)
		fmt.Println(delimiter)
		fmt.Printf("E-MAIL FROM %s TO %s\n", from, strings.Join(to, ", "))
		fmt.Println(delimiter)
		fmt.Println(string(msg))
		fmt.Println(delimiter)
		return nil
	}
	arg := []string{"-f", from}
	arg = append(arg, to...)
	sendmail := exec.Command("/usr/sbin/sendmail", arg...)
	stdin, err := sendmail.StdinPipe()
	if err != nil {
		return err
	}
	stderr, err := sendmail.StderrPipe()
	if err != nil {
		return err
	}
	if err := sendmail.Start(); err != nil {
		return err
	}
	if _, err := stdin.Write(msg); err != nil {
		return err
	}
	if err := stdin.Close(); err != nil {
		return err
	}
	out, err := ioutil.ReadAll(stderr)
	if err != nil {
		return err
	}
	if len(out) != 0 {
		return errors.New(string(out))
	}
	return sendmail.Wait()
}
