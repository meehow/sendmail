package sendmail

import (
	"fmt"
	"testing"
)

const domain = "example.com"

func TestSend(t *testing.T) {
	subject := "Cześć"
	from := fmt.Sprintf("Fróm <me@%s>", domain)
	to := []string{fmt.Sprintf("Tó <info@%s>", domain)}
	sm, err := New(subject, from, to)
	if err != nil {
		t.Fatal(err)
	}
	sm.Text.WriteString(":)\r\n")
	if err := sm.Send(); err != nil {
		t.Fatal(err)
	}
	if sm.From.String() != fmt.Sprintf("=?utf-8?q?Fr=C3=B3m?= <me@%s>", domain) {
		t.Fatalf("Wrong email encoding: %s", sm.From.String())
	}
}
