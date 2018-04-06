package sendmail

import (
	"bytes"
	"testing"
)

func TestChaningOptions(t *testing.T) {
	var buf bytes.Buffer
	m := &Mail{}

	if m.sendmail != "" {
		t.Errorf("Expected initial sendmail to be empty, got %q", m.sendmail)
	}
	if m.debugOut != nil {
		t.Errorf("Expected initial debugOut to be nil, got %T", m.debugOut)
	}

	m.SetSendmail("/bin/true").SetDebugOutput(&buf)

	if m.sendmail != "/bin/true" {
		t.Errorf("Expected sendmail to be %q, got %q", "/bin/true", m.sendmail)
	}
	if m.debugOut != &buf {
		t.Errorf("Expected debugOut to be %T (buf), got %T", &buf, m.debugOut)
	}
}
