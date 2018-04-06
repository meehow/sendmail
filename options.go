package sendmail

import (
	"io"
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
