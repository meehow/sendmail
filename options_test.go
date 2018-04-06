package sendmail

import (
	"bytes"
	"net/mail"
	"os"
	"testing"
)

func TestChaningOptions(t *testing.T) {
	var buf bytes.Buffer
	m := &Mail{
		To: []*mail.Address{
			&mail.Address{Name: "Michał", Address: "me@example.com"},
		},
	}
	if m.Subject != "" {
		t.Errorf("Expected subject to be empty, got %q", m.Subject)
	}
	if len(m.To) != 1 {
		t.Errorf("Expected len(To) to be 1, got %d: %+v", len(m.To), m.To)
	}
	if m.From != nil {
		t.Errorf("Expected From address to be nil, got %s", m.From)
	}
	if m.sendmail != "" {
		t.Errorf("Expected initial sendmail to be empty, got %q", m.sendmail)
	}
	if m.debugOut != nil {
		t.Errorf("Expected initial debugOut to be nil, got %T", m.debugOut)
	}

	m.SetSubject("Test subject").
		SetFrom("Dominik", "dominik@example.org").
		AppendTo("Dominik2", "dominik2@example.org").
		SetDebugOutput(&buf).
		SetSendmail("/bin/true")

	if m.Subject != "Test subject" {
		t.Errorf("Expected subject to be %q, got %q", "Test subject", m.Subject)
	}
	if len(m.To) != 2 {
		t.Errorf("Expected len(To) to be 2, got %d: %+v", len(m.To), m.To)
	}
	if m.From == nil || m.From.Address != "dominik@example.org" {
		expected := mail.Address{Name: "Dominik", Address: "dominik@example.org"}
		t.Errorf("Expected From address to be %s, got %s", expected, m.From)
	}
	if m.sendmail != "/bin/true" {
		t.Errorf("Expected sendmail to be %q, got %q", "/bin/true", m.sendmail)
	}
	if m.debugOut != &buf {
		t.Errorf("Expected debugOut to be %T (buf), got %T", &buf, m.debugOut)
	}
}

func TestOptions(t *testing.T) {
	m := &Mail{}
	var o Option

	o = Sendmail("/foo/bar")
	if o.execute(m); m.sendmail != "/foo/bar" {
		t.Errorf("Expected sendmail to be %q, got %q", "/foo/bar", m.sendmail)
	}

	o = Debug(true)
	if o.execute(m); m.debugOut != os.Stderr {
		t.Errorf("Expected debugOut to be %T (stderr), got %T", os.Stderr, m.debugOut)
	}

	o = Debug(false)
	if o.execute(m); m.debugOut != nil {
		t.Errorf("Expected debugOut to be nil, got %T", m.debugOut)
	}

	var buf bytes.Buffer
	o = DebugOutput(&buf)
	if o.execute(m); m.debugOut != &buf {
		t.Errorf("Expected debugOut to be %T (buf), got %T", &buf, m.debugOut)
	}

	o = DebugOutput(nil)
	if o.execute(m); m.debugOut != nil {
		t.Errorf("Expected debugOut to be nil, got %T", m.debugOut)
	}

	// To() appends list
	o = To("Ktoś", "info@example.com")
	if o.execute(m); len(m.To) != 1 {
		t.Errorf("Expected len(To) to be 1, got %d: %+v", len(m.To), m.To)
	}
	o = To("Ktoś2", "info2@example.com")
	if o.execute(m); len(m.To) != 2 {
		t.Errorf("Expected len(To) to be 2, got %d: %+v", len(m.To), m.To)
	}

	// From() updates current sender
	o = From("Michał", "me@example.com")
	if o.execute(m); m.From == nil || m.From.Address != "me@example.com" {
		expected := mail.Address{Name: "Michał", Address: "me@example.com"}
		t.Errorf("Expected From address to be %s, got %s", expected, m.From)
	}
	o = From("Michał", "me@example.com")
	if o.execute(m); m.From == nil || m.From.Address != "me@example.com" {
		expected := mail.Address{Name: "Michał", Address: "me@example.com"}
		t.Errorf("Expected From address to be %s, got %s", expected, m.From)
	}

	// Subject() updates current subject
	o = Subject("Cześć")
	if o.execute(m); m.Subject != "Cześć" {
		t.Errorf("Expected Subject to be %q, got %q", "Cześć", m.Subject)
	}
	o = Subject("Test")
	if o.execute(m); m.Subject != "Test" {
		t.Errorf("Expected Subject to be %q, got %q", "Test", m.Subject)
	}
}
