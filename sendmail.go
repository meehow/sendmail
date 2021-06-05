// Package sendmail implements then classic method of sending emails,
// well known from PHP.
package sendmail

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/mail"
	"os"
	"os/exec"
	"strings"
)

var (
	_, debug = os.LookupEnv("DEBUG")

	// Binary points to the sendmail binary.
	Binary = "/usr/sbin/sendmail"
)

// Mail defines basic mail structure and headers
type Mail struct {
	Subject string
	From    *mail.Address
	To      []*mail.Address
	Header  http.Header
	Text    bytes.Buffer
	HTML    bytes.Buffer
}

// Send sends an email, or prints it on stderr,
// when environment variable `DEBUG` is set.
func (m *Mail) Send() error {
	if m.From == nil {
		return errors.New("Missing `From` address")
	}
	if len(m.To) == 0 {
		return errors.New("Missing `To` address")
	}
	if m.Header == nil {
		m.Header = make(http.Header)
	}
	if m.Header.Get("Content-Type") == "" {
		m.Header.Set("Content-Type",  "text/plain; charset=UTF-8")
	}
	m.Header.Set("Subject", mime.QEncoding.Encode("utf-8", m.Subject))
	m.Header.Set("From", m.From.String())
	to := make([]string, len(m.To))
	arg := make([]string, len(m.To))
	for i, t := range m.To {
		to[i] = t.String()
		arg[i] = t.Address
	}
	m.Header.Set("To", strings.Join(to, ", "))
	if debug {
		delimiter := "\n" + strings.Repeat("-", 70)
		fmt.Println(delimiter)
		m.WriteTo(os.Stdout)
		fmt.Println(delimiter)
		return nil
	}
	sendmail := exec.Command(Binary, arg...)
	stdin, err := sendmail.StdinPipe()
	if err != nil {
		return err
	}
	stderr, err := sendmail.StderrPipe()
	if err != nil {
		return err
	}
	if err = sendmail.Start(); err != nil {
		return err
	}
	if err = m.WriteTo(stdin); err != nil {
		return err
	}
	if err = stdin.Close(); err != nil {
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

// WriteTo writes headers and content of the email to io.Writer
func (m *Mail) WriteTo(wr io.Writer) error {
	if err := m.Header.Write(wr); err != nil {
		return err
	}
	if _, err := wr.Write([]byte("\r\n")); err != nil {
		return err
	}
	if _, err := m.Text.WriteTo(wr); err != nil {
		return err
	}
	return nil
}
