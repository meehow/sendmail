package sendmail

import "testing"

func TestSend(t *testing.T) {
	from := "test@example.com"
	to := []string{"test@example.com"}
	msg := []byte("Subject: Hi\r\n:)")
	if err := Send(from, to, msg); err != nil {
		t.Error(err)
	}
}
