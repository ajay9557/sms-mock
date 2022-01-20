package sms

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
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

		{desc: "Success", to: "6303880131", msg: "Hey", err: nil, mock: []*gomock.Call{mockSvc.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)}},
		{desc: "Failure", to: "630*936709", msg: "How are you", err: errors.New("invalid phone"), mock: nil},
		{
			desc: "Invalid Msg",
			to:   "+919849632049",
			msg:  "hey i have recieved smsvenruinuiw",
			mock: nil,
			err:  errors.New("invalid sms message"),
		},
		{
			desc: "Could not send msg",
			to:   "+919908577405",
			msg:  "Hello ,I have received sms",
			mock: []*gomock.Call{mockSvc.EXPECT().Send(gomock.Any(), gomock.Any()).Return(errors.New("couldn't send sms"))},
			err:  errors.New("couldn't send sms"),
		},
	}
	for _, tc := range testCases {
		err := mock.SendMessage(tc.to, tc.msg)
		t.Run(tc.desc, func(t *testing.T) {
			if err != nil && err.Error() != tc.err.Error() {
				t.Error("Expected: ", tc.err, "Obtained: ", err)
			}

		})

	}
}
