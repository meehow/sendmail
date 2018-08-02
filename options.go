package sendmail

import (
	"io"
	"net/mail"
	"os"
)

// SetSendmail modifies the path to the sendmail binary. You can pass
// additional arguments, if you need to.
func (m *Mail) SetSendmail(path string, args ...string) *Mail {
	m.sendmailPath = path
	m.sendmailArgs = args
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

// AppendTo adds a recipient to the Mail.
func (m *Mail) AppendTo(toAddress ...*mail.Address) *Mail {
	m.To = append(m.To, toAddress...)
	return m
}

// AppendCC adds a carbon-copy recipient to the Mail.
func (m *Mail) AppendCC(ccAddress ...*mail.Address) *Mail {
	m.CC = append(m.CC, ccAddress...)
	return m
}

// AppendBCC adds a blind carbon-copy recipient to the Mail.
func (m *Mail) AppendBCC(bccAddress ...*mail.Address) *Mail {
	m.BCC = append(m.BCC, bccAddress...)
	return m
}

// SetFrom updates (replaces) the sender's address.
func (m *Mail) SetFrom(fromAddress *mail.Address) *Mail {
	m.From = fromAddress
	return m
}

// SetSubject sets the mail subject.
func (m *Mail) SetSubject(subject string) *Mail {
	m.Subject = subject
	return m
}

// Option is used in the Mail constructor.
type Option interface {
	execute(*Mail)
}

type optionFunc func(*Mail)

func (o optionFunc) execute(m *Mail) { o(m) }

// Sendmail modifies the path to the sendmail binary.
func Sendmail(path string, args ...string) Option {
	return optionFunc(func(m *Mail) { m.SetSendmail(path, args...) })
}

// Debug sets the debug output to stderr if active is true, else it
// removes the debug output. Use SetDebugOutput to set it to something else.
func Debug(active bool) Option {
	return optionFunc(func(m *Mail) { m.SetDebug(active) })
}

// DebugOutput sets the debug output to the given writer. If w is nil,
// this is equivalent to SetDebug(false).
func DebugOutput(w io.Writer) Option {
	return optionFunc(func(m *Mail) { m.SetDebugOutput(w) })
}

// To adds a recipient to the Mail.
func To(address *mail.Address) Option {
	return optionFunc(func(m *Mail) { m.AppendTo(address) })
}

// From sets the sender's address.
func From(fromAddress *mail.Address) Option {
	return optionFunc(func(m *Mail) { m.SetFrom(fromAddress) })
}

// Subject sets the mail subject.
func Subject(subject string) Option {
	return optionFunc(func(m *Mail) { m.SetSubject(subject) })
}
