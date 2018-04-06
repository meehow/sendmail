package sendmail

import (
	"bytes"
	"net/mail"
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
