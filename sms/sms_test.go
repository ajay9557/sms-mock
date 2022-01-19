package sms

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSendMessage(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockSMSSender(ctrl)
	mock := New(mockSvc)

	testCases := []struct {
		desc string
		to   string
		msg  string
		err  error
		mock []*gomock.Call
	}{
		{
			desc: "Case-1: Successful Sending of message",
			to:   "9871333121",
			msg:  "Hello",
			err:  nil,
			mock: []*gomock.Call{
				mockSvc.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil),
			},
		},
		{
			desc: "Case-2: Failure in validation of phone",
			to:   "hi",
			msg:  "Hello",
			err:  errors.New("invalid phone"),
			mock: nil,
		},
		{
			desc: "Case-3: Failure in validation of message",
			to:   "9832687676",
			msg:  "Hellojiooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooofrmoifnorifnroifrnfrforfnrfrfnrofirnfrfoirnfroinfrofrofirnfroifnrfoirnfoirnfo",
			err:  errors.New("invalid sms message"),
			mock: nil,
		},
		{
			desc: "Case-4: Failure in sending Message",
			to:   "9832687676",
			msg:  "Hello",
			err:  errors.New("couldn't send sms"),
			mock: []*gomock.Call{
				mockSvc.EXPECT().Send(gomock.Any(), gomock.Any()).Return(errors.New("couldn't send sms")),
			},
		},
	}
	for _, tc := range testCases {

		err := mock.SendMessage(tc.to, tc.msg)

		if err != nil && err.Error() != tc.err.Error() {

			t.Errorf("Expected Error %v Got Error: %v", tc.err, err)
		}

	}

}
