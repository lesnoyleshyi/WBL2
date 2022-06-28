package entity

import (
	"testing"
	"time"
)

func TestEvent_String(t *testing.T) { //nolint:paralleltest
	type testCase struct {
		event          Event
		expectedString string
	}

	testCases := []testCase{
		{
			event: Event{
				Datetime:    time.Date(2022, 6, 28, 14, 57, 0, 0, time.UTC),
				Title:       "test_task1",
				Description: "test description 1",
			},
			expectedString: `{"title":"test_task1", "description":"test description 1", "datetime":"2022-06-28T14:57:00Z"}`,
		},
		{
			event: Event{
				Datetime:    time.Date(2044, 1, 1, 13, 37, 59, 590, time.Local),
				Title:       "Test task №2",
				Description: "Test description №2",
			},
			expectedString: `{"title":"Test task №2", "description":"Test description №2", "datetime":"2044-01-01T13:37:59+03:00"}`,
		},
	}

	for _, tc := range testCases {
		if tc.event.String() != tc.expectedString {
			t.Errorf("Wrong event representation!\nExpected %s\nGot\t\t%s",
				tc.expectedString, tc.event.String())
		}
	}
}
