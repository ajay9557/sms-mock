package sms

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSendMessage(t *testing.T) {
	Ctrl := gomock.NewController(t)

	defer Ctrl.Finish()

	mockSvc := NewMockSMSSender(Ctrl)

	testCases := []struct {
		desc, to, msg string
		err           error
		mock          []*gomock.Call
	}{
		{desc: "success 1", to: "9876543210", msg: "Hello", err: nil, mock: []*gomock.Call{mockSvc.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)}},
		{desc: "failure 1", to: "93939c", msg: "Hello", err: errors.New("invalid phone")},
		{desc: "success 2", to: "9876543210", msg: "Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.", mock: nil, err: errors.New("invalid sms message")},
		{desc: " failure 2", to: "9876543210", msg: "Hello", err: errors.New("couldn't send sms"), mock: []*gomock.Call{mockSvc.EXPECT().Send(gomock.Any(), gomock.Any()).Return(errors.New("couldn't send sms"))}},
	}

	m := New(mockSvc)
	for _, tcs := range testCases {
		t.Run(tcs.desc, func(t *testing.T) {
			err := m.SendMessage(tcs.to, tcs.msg)

			if err != nil && err.Error() != tcs.err.Error() {
				t.Errorf("expected %v, got %v", tcs.err, err)
			}
		})
	}
}
