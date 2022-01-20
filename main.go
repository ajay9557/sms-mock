package main

import (
	"fmt"

	"github.com/zopping/sms-mock/sms"
	"github.com/zopping/sms-mock/twilio"
)

func main() {
	t := twilio.Twilio{
		Config: map[string]string{"accountSid": "test1234", "authToken": "yuyt6566", "from": "test"},
	}
	to := "+918767654545"
	msg := "hey i have recieved sms"
	smsHandler := sms.New(t)
	err := smsHandler.SendMessage(to, msg)
	if err != nil {
		fmt.Println("error while sending sms")
	}
}
