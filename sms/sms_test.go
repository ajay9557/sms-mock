package sms

import (
	"errors"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_SendMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockSMSSender(ctrl)
	mock := New(mockSvc)

	testcases := []struct {
		desc   string
		to     string
		msg    string
		experr error
		mock   []*gomock.Call
	}{
		{"invalid-Phone", "8083a", "hi", errors.New("invalid phone"), nil},
		{"invalid message", "8083860404", "abcdefghijklmnopqrstuvwxyzabcdefgh", errors.New("invalid sms message"), nil},
		{"could not send", "8083860404", "hi", errors.New("couldn't send sms"), []*gomock.Call{
			mockSvc.EXPECT().Send(gomock.Any(), gomock.Any()).Return(errors.New("couldn't send sms")),
		}},
		{"success", "8083860404", "hi", nil, []*gomock.Call{mockSvc.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)}},
	}

	for _, tcs := range testcases {
		err := mock.SendMessage(tcs.to, tcs.msg)

		if err != nil && !reflect.DeepEqual(err, tcs.experr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.experr, err)
		}
	}
}
func TestValidatePhone(t *testing.T) {
	testcases := []struct {
		desc string
		to   string
		exp  bool
	}{
		{"invalid-Phone", "8083a", false},
		{"valid- phone", "8083860404", true},
		//{"invalid message", "8083860404", "abcdefghijklmnopqrstuvwxyzabcdefgh", errors.New("invalid sms message")},

	}

	for _, tcs := range testcases {
		out := validatePhone(tcs.to)

		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}
func TestValidateMessage(t *testing.T) {
	testcases := []struct {
		desc string
		msg  string
		exp  bool
	}{
		{"invalid message", "abcdefghijklmnopqrstuvwxyzabcdefgh", false},
		{"valid msg", "hi", true},
	}

	for _, tcs := range testcases {
		out := validateMessage(tcs.msg)

		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}
