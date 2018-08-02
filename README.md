# Go sendmail

[![GoDoc](https://godoc.org/github.com/meehow/sendmail?status.svg)](https://godoc.org/github.com/meehow/sendmail)
[![Build Status](https://travis-ci.org/meehow/sendmail.svg?branch=master)](https://travis-ci.org/meehow/sendmail)


This package implements the classic method of sending emails, well known
from PHP. It's stupid simple and it works not only with Sendmail, but also
with other MTAs, like [Postfix][], [sSMTP][], or [mhsendmail][], which
provide a compatible interface.

[Postfix]:    http://www.postfix.org/sendmail.1.html
[sSMTP]:      https://wiki.debian.org/sSMTP
[mhsendmail]: https://github.com/mailhog/mhsendmail

* it separates email headers from email body,
* encodes UTF-8 headers like `Subject`, `From`, `To`
* makes it easy to use [text/template](https://golang.org/pkg/text/template)
* doesn't require any SMTP configuration,
* can write email body to a custom `io.Writer` to simplify testing
* by default, it just uses `/usr/sbin/sendmail` (but can be changed if need be)


## Installation

```
go get -u github.com/meehow/sendmail
```

## Usage

```go
package main

import (
	"io"
	"log"
	"net/mail"

	"github.com/meehow/sendmail"
)

func main() {
	sm := sendmail.Mail{
		Subject: "Cześć",
		From:    &mail.Address{"Michał", "me@example.com"},
		To: []*mail.Address{
			{"Ktoś", "info@example.com"},
			{"Ktoś2", "info2@example.com"},
		},
	}
	io.WriteString(&sm.Text, ":)\r\n")
	if err := sm.Send(); err != nil {
		log.Println(err)
	}
}
```

Instead of `io.WriteString`, you can [execute a template][template]:

[template]: https://golang.org/pkg/text/template/#Template.Execute

```go
tpl := template.Must(template.New("email").Parse(`Hello {{.Name}}!`))
tpl.ExecuteTemplate(&sm.Text, "email", &struct{ Name string }{"Dominik"})
```


## ToDo

- [x] HTML emails
- [ ] multipart emails (HTML + Text)
- [ ] attachments
- [ ] inline attachments
