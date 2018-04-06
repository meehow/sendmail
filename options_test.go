package sendmail

import (
	"testing"
)

func TestChaningOptions(t *testing.T) {
	m := &Mail{}
	m.SetSendmail("/bin/true")

	if m.sendmail != "/bin/true" {
		t.Errorf("Expected sendmail to be %q, got %q", "/bin/true", m.sendmail)
	}
}
