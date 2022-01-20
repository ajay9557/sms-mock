package sms

import (
	"testing"
	"errors"
	"github.com/golang/mock/gomock"
	//	"golang.org/x/tools/go/expect"
)

func TestSms(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSender := NewMockSMSSender(ctrl)

	tcs := []struct {
		to, msg    string
		mocksender *gomock.Call
		out        error
	}{
		{to: "9884697050", msg: "Hello", mocksender: mockSender.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil), out: nil},
		{to : "1234@34", msg: "Good morning", mocksender: nil, out : errors.New("invalid phone")},
		{to : "9884697050", msg: "When you have a dream, you've got to grab it and never let go Keep your face always toward the sunshine, and shadows will fall behind you.", mocksender: nil, out : errors.New("invalid sms message")},
	    {to: "9884697050", msg : "Hello", mocksender: mockSender.EXPECT().Send(gomock.Any(), gomock.Any()).Return(errors.New("couldn't send sms")), out : errors.New("couldn't send sms")},
}

	m := New(mockSender)

	for _, tc := range tcs {
		t.Run("test cases", func(t *testing.T) {

			err := m.SendMessage(tc.to, tc.msg)
			//	fmt.Println(err==tc.expected)

			if err != nil && err.Error() != tc.out.Error() {
				t.Errorf("expected is %v but got %T %v", tc.out, err, err)
			}

		})
	}
}
