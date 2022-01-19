package sms

import (
	"testing"

	"errors"

	"github.com/golang/mock/gomock"
)

func TestSentMessage(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	mockService := NewMockSMSSender(control)

	mock := New(mockService)

	testCases := []struct {
		desc string
		to   string
		msg  string
		err  error
		mock []*gomock.Call
	}{
		{desc: "Case1", to: "1234567890", msg: "hey!", err: nil, mock: []*gomock.Call{mockService.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)}},
		{desc: "Case2", to: "hi", msg: "hellow", err: errors.New("invalid phone")},
		{desc: "Case3", to: "lllllll", msg: "hellow", err: errors.New("invalid phone")},
		{desc: "Case4", to: "7828789845", msg: "fjlasjflasdlfjalskdjflkasjdfjaslkdfjlksdjflkjsdlkfjas", err: errors.New("invalid sms message")},
	}

	for _, tcs := range testCases {
		err := mock.SendMessage(tcs.to, tcs.msg)
		if err != nil && err.Error() != tcs.err.Error() {
			t.Errorf("expected %v, want %v", err, tcs.err)
		}
	}
}
