package sms

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestSendMessage(t *testing.T) {

	// Create a mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock SMSSender
	sender := NewMockSMSSender(ctrl)

	// test cases for SendMessage
	tc := []struct {
		to       string
		msg      string
		expErr   error
		mockCall []*gomock.Call
	}{

		// Mock calls for valid phone number and valid message returning nil error
		{to: "1234567890", msg: "Hello", expErr: nil, mockCall: []*gomock.Call{
			sender.EXPECT().Send("1234567890", "Hello").Return(nil),
		}},

		// Mock calls for valid phone number and valid message returning (couldn't send sms) error
		{to: "123213123", msg: "Hi", expErr: errors.New("couldn't send sms"), mockCall: []*gomock.Call{
			sender.EXPECT().Send("123213123", "Hi").Return(errors.New("couldn't send sms")),
		}},

		// Phone number is invalid
		{to: "aavvvsss", msg: "Hi", expErr: errors.New("invalid phone")},
		{to: "12345a67890x", msg: "Hello", expErr: errors.New("invalid phone")},

		// Message length is less than 1 or more than 30
		{to: "1234567890", msg: "", expErr: errors.New("invalid sms message")},
		{to: "1234567890", msg: "qwedrtyuiopasdfghjklzxcvbnm;''.ss", expErr: errors.New("invalid sms message")},
	}

	// Create handler that uses the mock sender
	h := New(sender)
	var err error

	// loop through test cases
	for _, tc := range tc {
		t.Run("Testing case", func(t *testing.T) {

			// actual ouput for the test cases
			err = h.SendMessage(tc.to, tc.msg)

			// check if the error is as expected
			if err != nil && err.Error() != tc.expErr.Error() {
				t.Errorf("Expected error %v, got %v", tc.expErr, err)
			}
		})
	}
}
