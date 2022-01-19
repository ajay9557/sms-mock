package sms

import (
	"errors"
	"regexp"
)

type handler struct {
	sender SMSSender
}

func New(s SMSSender) *handler {
	return &handler{sender: s}
}

func (h *handler) SendMessage(to, msg string) error {
	// validate phone
	isValid := validatePhone(to)
	if !isValid {
		return errors.New("invalid phone")
	}
	// validate message
	isValid = validateMessage(msg)
	if !isValid {
		return errors.New("invalid sms message")
	}
	// use sms sender and send an sms,
	err := h.sender.Send(to, msg)
	if err != nil {
		return errors.New("couldn't send sms")
	}
	return nil
}

func validatePhone(number string) bool {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if re.MatchString(number) {
		return true
	}
	return false
}

func validateMessage(msg string) bool {
	if len([]rune(msg)) > 30 || len([]rune(msg)) < 1 {
		return false
	}
	return true
}
