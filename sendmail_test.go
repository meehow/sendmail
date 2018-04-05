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

func init() {
	Binary = "/bin/true"
}

func TestSend(tc *testing.T) {
	tc.Run("debug:true", func(t *testing.T) {
		testSend(t, true)
	})
	tc.Run("debug:false", func(t *testing.T) {
		testSend(t, false)
	})
}

func testSend(t *testing.T, withDebug bool) {
	oldDebug := debug
	debug = withDebug
	defer func() { debug = oldDebug }()

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
		t.Errorf("(debug=%v) %v", withDebug, err)
	}
	subject := sm.Header.Get("Subject")
	if subject != "=?utf-8?q?Cze=C5=9B=C4=87?=" {
		t.Errorf("(debug=%v) Wrong `Subject` encoding: %s", withDebug, subject)
	}
	from := sm.Header.Get("From")
	if from != fmt.Sprintf("=?utf-8?q?Micha=C5=82?= <me@%s>", domain) {
		t.Errorf("(debug=%v) Wrong `From` encoding: %s", withDebug, from)
	}
}

func TestFromError(t *testing.T) {
	sm := Mail{
		To: []*mail.Address{maddr("Ktoś", "info@")},
	}
	if sm.Send() == nil {
		t.Errorf("Expected an error because of missing `From` addresses")
	}
}

func TestToError(t *testing.T) {
	sm := Mail{
		From: maddr("Michał", "me@"),
	}
	if sm.Send() == nil {
		t.Errorf("Expected an error because of missing `To` addresses")
	}
}
