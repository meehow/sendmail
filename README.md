Go sendmail [![Build Status](https://travis-ci.org/meehow/sendmail.svg?branch=master)](https://travis-ci.org/meehow/sendmail)
===========

This package implements classic, well known from PHP, method of sending emails.
It's stupid simple and it works not only with Sendmail,
but also with other MTAs, like [Postfix](http://www.postfix.org/sendmail.1.html)
or [sSMTP](https://wiki.debian.org/sSMTP), which provide compatibility interface.

* it separates email headers from email body,
* encodes UTF-8 headers like `Subject`, `From`, `To`
* makes it easy to use [text/template](https://golang.org/pkg/text/template)
* doesn't require any SMTP configuration,
* just uses `/usr/sbin/sendmail` command which is present on most of the systems,
  * if not, just update `sendmail.Binary`
* outputs emails to _stdout_ when environment variable `DEBUG` is set.

Installation
------------
```
go get -u github.com/meehow/sendmail
```

Usage
-----
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


ToDo
----

* HTML emails
