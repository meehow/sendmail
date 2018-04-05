package sendmail

import (
	"fmt"
	"io"
	"net/mail"
	"testing"
)

const domain = "example.com"

func maddr(name, address string) *mail.Address {
	return &mail.Address{Name: name, Address: address + domain}
}

func TestSend(t *testing.T) {
	sm := Mail{
		Subject: "Cześć",
		From:    maddr("Michał", "me@"),
		To: []*mail.Address{
			maddr("Ktoś", "info@"),
			maddr("Ktoś2", "info2@"),
		},
	}
	io.WriteString(&sm.Text, ":)\r\n")
	if err := sm.Send(); err != nil {
		t.Fatal(err)
	}
	subject := sm.Header.Get("Subject")
	if subject != "=?utf-8?q?Cze=C5=9B=C4=87?=" {
		t.Fatalf("Wrong `Subject` encoding: %s", subject)
	}
	from := sm.Header.Get("From")
	if from != fmt.Sprintf("=?utf-8?q?Micha=C5=82?= <me@%s>", domain) {
		t.Fatalf("Wrong `From` encoding: %s", from)
	}
}

func TestFromError(t *testing.T) {
	sm := Mail{
		To: []*mail.Address{maddr("Ktoś", "info@")},
	}
	if sm.Send() == nil {
		t.Fatal("Expected an error because of missing `From` addresses")
	}
}

func TestToError(t *testing.T) {
	sm := Mail{
		From: maddr("Michał", "me@"),
	}
	if sm.Send() == nil {
		t.Fatal("Expected an error because of missing `To` addresses")
	}
}
