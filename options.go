package sendmail

// SetSendmail modifies the path to the sendmail binary.
func (m *Mail) SetSendmail(path string) *Mail {
	m.sendmail = path
	return m
}
