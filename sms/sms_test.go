/*package sms

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
	mockCalls *gomock.Call
}

func TestSMSSender(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSMS := NewMockSMSSender(ctrl)
	testSMS := New(mockSMS)

	tests := []Test{
		{desc: "Case1", to: "8320578360", msg: "My message", expected: nil, mockCalls: mockSMS.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)},
		{desc: "Case2", to: "8320578360", msg: "", expected: errors.New("couldn't send sms"), mockCalls: mockSMS.EXPECT().Send(gomock.Any(), gomock.Any()).Return(errors.New("couldn't send sms"))},
		{desc: "Case3", to: "asuidh", msg: "Hello", expected: errors.New("invalid phone"), mockCalls: mockSMS.EXPECT().Send(gomock.Any(), gomock.Any()).Return(errors.New("Invalid Mobile no."))},
		{desc: "Case4", to: "8320578360", msg: "very longggggggggggggggggggggggggggggggggggggggggggggggggg messageeeeeeeeeeeeeeeeeeeeeeeeeeeeee!!!!!!!!!!!!!!!!!", expected: errors.New("invalid sms message"), mockCalls: nil},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := testSMS.SendMessage(test.to, test.msg)

			// TODO: Check for the Error
			if errors.Is(err, test.expected) && err != nil && test.expected != nil {
				t.Errorf("Expected: %v, Got: %v", test.expected, err)
			}
		})
	}
}

*/

package sms

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestSendMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ImockSvc := NewMockSMSSender(ctrl)
	hmock := New(ImockSvc)
	testCases := []struct {
		desc string
		to   string
		msg  string
		err  error
		mock []*gomock.Call
	}{
		{
			desc: "Success",
			to:   "+919908577405",
			msg:  "Hello ,I have received sms",
			mock: []*gomock.Call{ImockSvc.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)},
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
			mock: []*gomock.Call{ImockSvc.EXPECT().Send(gomock.Any(), gomock.Any()).Return(errors.New("couldn't send sms"))},
			err:  errors.New("couldn't send sms"),
		},
	}
	for _, tc := range testCases {
		err := hmock.SendMessage(tc.to, tc.msg)
		t.Run(tc.desc, func(t *testing.T) {
			if err != nil && err.Error() != tc.err.Error() {
				t.Error("Expected: ", tc.err, "Obtained: ", err)
			}

		})

	}
}
