package sms

import (
	"errors"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestSendMessage(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockSMSSender(ctrl)
	mock := New(mockSvc)

	tcs := []struct {
		desc string
		to   string
		msg  string
		mock []*gomock.Call
		err  error
	}{
		{
			desc: "Success",
			to:   "+919908577405",
			msg:  "Hello ,I have received sms",
			mock: []*gomock.Call{mockSvc.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)},
			err:  nil,
		},
		{
			desc: "Invalid phone",
			to:   "+9198495758*",
			msg:  "I didn't receive",
			mock: nil,
			err:  errors.New("invalid phone")},
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

	for _, tc := range tcs {

		err := mock.SendMessage(tc.to, tc.msg)
		t.Run(tc.desc, func(t *testing.T) {
			if err != nil && !reflect.DeepEqual(err, tc.err) {
				t.Error("Expected: ", tc.err, "Obtained: ", err)
			}
		})
	}
}
