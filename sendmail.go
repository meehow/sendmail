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
	"os/exec"
	"strings"
)

// SendmailDefault points to the default sendmail binary location.
const SendmailDefault = "/usr/sbin/sendmail"

// Mail defines basic mail structure and headers
type Mail struct {
	Subject string
	From    *mail.Address
	To      []*mail.Address
	Header  http.Header
	Text    bytes.Buffer
	HTML    bytes.Buffer

	sendmailPath string
	sendmailArgs []string
	debugOut     io.Writer
}

// New creates a new Mail instance with the given options.
func New(options ...Option) (m *Mail) {
	m = &Mail{sendmailPath: SendmailDefault}
	for _, option := range options {
		option.execute(m)
	}
	return
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
	m.Header.Set("Subject", mime.QEncoding.Encode("utf-8", m.Subject))
	m.Header.Set("From", m.From.String())

	to := make([]string, len(m.To))
	arg := make([]string, len(m.To))
	for i, t := range m.To {
		to[i] = t.String()
		arg[i] = t.Address
	}
	m.Header.Set("To", strings.Join(to, ", "))
	if m.debugOut != nil {
		_, err := m.WriteTo(m.debugOut)
		return err
	}

	return m.exec(arg...)
}

// exec handles sendmail command invokation.
func (m *Mail) exec(arg ...string) error {
	bin := SendmailDefault
	if m.sendmailPath != "" {
		bin = m.sendmailPath
	}
	args := append(append([]string{}, m.sendmailArgs...), arg...)
	cmd := exec.Command(bin, args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	if _, err = m.WriteTo(stdin); err != nil {
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
	return cmd.Wait()
}

// WriteTo writes headers and content of the email to io.Writer
func (m *Mail) WriteTo(wr io.Writer) (n int64, err error) {
	isText := m.Text.Len() > 0
	isHTML := m.HTML.Len() > 0

	if isText && isHTML {
		err = fmt.Errorf("Multipart mails are not supported yet")
		return
	} else if isHTML {
		m.Header.Set("Content-Type", "text/html; charset=UTF-8")
	} else {
		// also for mails without body
		m.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	}

	w := &writeCounter{w: wr}

	// write header
	if err = m.Header.Write(w); err != nil {
		return
	}
	if _, err = w.Write([]byte("\r\n")); err != nil {
		return
	}

	if isText && isHTML {
		// TODO
	} else if isHTML {
		if _, err = m.HTML.WriteTo(w); err != nil {
			return
		}
	} else {
		if _, err = m.Text.WriteTo(w); err != nil {
			return
		}
	}
	return w.n, nil
}

// writeCounter is an internal type wrapping an io.Writer to work around
// the issue of net.http.Header.Write() not returning the number of bytes
// written.
type writeCounter struct {
	n int64
	w io.Writer
}

// Write satisfies the io.Writer interface. It updates an internal cache
// for the number of bytes written.
func (wc *writeCounter) Write(p []byte) (n int, err error) {
	n, err = wc.w.Write(p)
	if err == nil {
		wc.n += int64(n)
	}
	return
}
