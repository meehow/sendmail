package sendmail

import (
	"io"
	"net/mail"
	"os"
)

// SetSendmail modifies the path to the sendmail binary.
func (m *Mail) SetSendmail(path string) *Mail {
	m.sendmail = path
	return m
}

// SetDebug sets the debug output to stderr if active is true, else it
// removes the debug output. Use SetDebugOutput to set it to something else.
func (m *Mail) SetDebug(active bool) *Mail {
	var out io.Writer
	if active {
		out = os.Stderr
	}
	m.debugOut = out
	return m
}

// SetDebugOutput sets the debug output to the given writer. If w is
// nil, this is equivalent to SetDebug(false).
func (m *Mail) SetDebugOutput(w io.Writer) *Mail {
	m.debugOut = w
	return m
}

// AppendTo adds a recipient to the Mail. The name argument is the
// "proper name" of the recipient and may be empty. The address must be
// in the form "user@domain".
func (m *Mail) AppendTo(name, address string) *Mail {
	m.To = append(m.To, &mail.Address{Name: name, Address: address})
	return m
}

// SetFrom updates the sender's address. Like AppendTo(), name may be
// empty, and address must be in the form "user@domain".
func (m *Mail) SetFrom(name, address string) *Mail {
	m.From = &mail.Address{Name: name, Address: address}
	return m
}

// SetSubject sets the mail subject.
func (m *Mail) SetSubject(subject string) *Mail {
	m.Subject = subject
	return m
}
