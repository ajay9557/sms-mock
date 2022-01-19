package sms

type SMSSender interface {
	Send(to, message string) error
}
