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
	mail, err := New(subject, from, to)
	if err != nil {
		t.Fatal(err)
	}
	mail.Text = []byte(":)")
	if err := mail.Send(); err != nil {
		t.Fatal(err)
	}
	if mail.From.String() != fmt.Sprintf("=?utf-8?q?Fr=C3=B3m?= <me@%s>", domain) {
		t.Fatalf("Wrong email encoding: %s", mail.From.String())
	}
}
