package sms

import(
	"testing"
	//"gomock"
	"errors"
	"github.com/golang/mock/gomock"
)
func TestSendMessage(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc:=NewMockSMSSender(ctrl)
	mock:=New(mockSvc)

	testCases:=[]struct{
		desc string
		to string
		msg string
		err error
		mock []*gomock.Call
	}{
		{
			"success","6303844857","hi", nil,[]*gomock.Call{mockSvc.EXPECT().Send(gomock.Any(),gomock.Any()).Return(nil)},
		},
		{
			desc:"failure",to:"hi",msg:"hi",err: errors.New("invalid phone"),
		},
		{
			desc: "Invalid Msg ochindhi",
	        to:   "6303844857",
	        msg:  "hey i have recieved smsvenruin uiw",
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
	for _,tc:=range testCases{
		err:=mock.SendMessage(tc.to,tc.msg)
		if err!=nil && err.Error()!=tc.err.Error(){
			t.Errorf("expected %v, got %v",tc.err,err)
		}
	}
}

