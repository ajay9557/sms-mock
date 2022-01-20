package sms

import (
	"errors"
	"reflect"
	"testing"
	gomock "github.com/golang/mock/gomock"
)	

func TestSendMessage(t *testing.T) {

	ctrl:= gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc:= NewMockSMSSender(ctrl)
	mock:= New(mockSvc)

	testDetails := []struct {
		desc string
		to   string
		msg  string
		mock []*gomock.Call
		err  error
	}{
		{
			desc: "Success case", 
			to: "+918767654545", 
			msg: "hey i have recieved sms",
			mock:[]*gomock.Call{mockSvc.EXPECT().Send(gomock.Any(),gomock.Any()).Return(nil)},
			err: nil,
		},
		{
			desc: "Failure : Invalid phone", 
			to: "+91*767654545",
			msg: "hey i have recieved sms", 
			mock:nil,
			err: errors.New("invalid phone")},
		{
			desc: "Failure : Invalid msg", 
			to: "+918767654545", 
			msg: "hey i have recieved smsvenruinuiw", 
			mock:nil ,
			err: errors.New("invalid sms message"),
		},
		{
			desc: "Failure : Couldnot send message", 
			to: "+918767654545", 
			msg: "hey i have recieved sms",
			mock:[]*gomock.Call{mockSvc.EXPECT().Send(gomock.Any(),gomock.Any()).Return(errors.New("couldn't send sms"))},
			err: errors.New("couldn't send sms"),
		},
	}

	for _, tes := range testDetails {

		err := mock.SendMessage(tes.to,tes.msg)
		t.Run(tes.desc, func(t *testing.T) {
			if err != nil && !reflect.DeepEqual(err, tes.err) {
				t.Error("Expected: ", tes.err, "Obtained: ", err)
			}
		})
	}
}