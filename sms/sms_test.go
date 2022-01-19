package sms

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

type Test struct {
	desc      string
	to        string
	msg       string
	expected  error
	mockCalls []*gomock.Call
}

func TestSMSSender(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSMS := NewMockSMSSender(ctrl)
	testSMS := New(mockSMS)

	tests := []Test{
		{desc: "Case1", to: "8320578360", msg: "My message", expected: nil, mockCalls: []*gomock.Call{mockSMS.EXPECT().Send("8320578360", "My message").Return(nil)}},
		{desc: "Case2", to: "8320578360", msg: "", expected: errors.New("couldn't send sms"), mockCalls: []*gomock.Call{mockSMS.EXPECT().Send(gomock.Any(), gomock.Any()).Return(errors.New("couldn't send sms"))}},
		{desc: "Case3", to: "asuidh", msg: "Hello", expected: errors.New("invalid phone"), mockCalls: nil},
		{desc: "Case4", to: "8320578360", msg: "very longggggggggggggggggggggggggggggggggggggggggggggggggg messageeeeeeeeeeeeeeeeeeeeeeeeeeeeee!!!!!!!!!!!!!!!!!", expected: errors.New("invalid sms message"), mockCalls: nil},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := testSMS.SendMessage(test.to, test.msg)

			if errors.Is(err, test.expected) && err != nil && test.expected != nil {
				t.Errorf("Expected: %v, Got: %v", test.expected, err)
			}
		})
	}
}
