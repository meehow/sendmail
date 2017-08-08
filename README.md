Go sendmail
===========

This package implements classic, well known from PHP, method of sending emails.
It's stupid simple and it works not only with Sendmail,
but also with other MTAs, like [Postfix](http://www.postfix.org/sendmail.1.html)
or [sSMTP](https://wiki.debian.org/sSMTP), which provide compatibility interface.

* doesn't require any SMTP configuration,
* just uses `/usr/sbin/sendmail` command which is present on most of the systems,
* outputs emails to _stdout_ when environment variable `DEBUG` is set.

Installation
------------
```
go get -u github.com/meehow/sendmail
```

Usage
-----
```go
msg := []byte("Subject: Hi\r\nHi Bob.")
err := sendmail.Send("me@example.com", []string{"bob@example.com"}, msg)
if err != nil {
	log.Println(err)
}
```
