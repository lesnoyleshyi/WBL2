package utils

import (
	"WBL2/Tasks/T11/internal/domain/entity"
	domainErrors "WBL2/Tasks/T11/internal/domain/errors"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

type getPeriodTestCase struct {
	urlString    string
	expectedFrom time.Time
	expectedTo   time.Time
	expectedErr  error
}

type getEventTestCase struct {
	bodyContent   string
	expectedEvent entity.Event
	expectedErr   error
}

type arrayPresenter struct {
	array                  []entity.Event
	expectedRepresentation string
}

func TestGetPeriod(t *testing.T) {
	testCases := []getPeriodTestCase{
		{
			urlString:    `localhost/events_for_day?date=2022-06-29`,
			expectedFrom: time.Date(2022, 6, 29, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2022, 6, 30, 0, 0, 0, 0, time.UTC),
			expectedErr:  nil,
		},
		{
			urlString:    `localhost/events_for_day?date=2022-06-29&period=day`,
			expectedFrom: time.Date(2022, 6, 29, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2022, 6, 30, 0, 0, 0, 0, time.UTC),
			expectedErr:  nil,
		},
		{
			urlString:    `localhost/events_for_day`,
			expectedFrom: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedErr:  domainErrors.ErrWrongQueryParams,
		},
		{
			urlString:    `localhost/events_for_day?date=2022-06-29&period=week`,
			expectedFrom: time.Date(2022, 6, 29, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2022, 7, 6, 0, 0, 0, 0, time.UTC),
			expectedErr:  nil,
		},
		{
			urlString:    `localhost/events_for_day?date=2022-06-29&period=month`,
			expectedFrom: time.Date(2022, 6, 29, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(2022, 7, 29, 0, 0, 0, 0, time.UTC),
			expectedErr:  nil,
		},
		{
			urlString:    `localhost/events_for_day?date=2022-06-29&period=year`,
			expectedFrom: time.Date(2022, 6, 29, 0, 0, 0, 0, time.UTC),
			expectedTo:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedErr:  domainErrors.ErrWrongQueryParams,
		},
	}

	for _, tc := range testCases {
		u, err := url.Parse(tc.urlString)
		if err != nil {
			t.Fatal("wrong input")
		}
		req := http.Request{ //nolint:exhaustruct
			URL:        u,
			Body:       io.NopCloser(strings.NewReader("")),
			RequestURI: strings.TrimPrefix(tc.urlString, "localhost/"),
		}

		from, to, err := GetPeriod(&req)
		if from != tc.expectedFrom {
			t.Errorf("Wrong from value.\tCase\t\"%s\"\nExpected:\t%s\nGot:\t\t%s\n",
				tc.urlString, tc.expectedFrom, from)
		}
		if to != tc.expectedTo {
			t.Errorf("Wrong to value.\t\tCase\t\"%s\"\nExpected:\t%s\nGot:\t\t%s\n",
				tc.urlString, tc.expectedTo, to)
		}
		if !errors.Is(err, tc.expectedErr) {
			t.Errorf("Wrong err value.\t\tCase\t\"%s\"\nExpected: %s\nGot:\t\t%s\n",
				tc.urlString, tc.expectedErr, err)
		}
	}
}

func TestGetEvent(t *testing.T) {
	testCases := []getEventTestCase{
		{
			bodyContent: `{"datetime": "2022-06-29T15:28:00Z", "title": "test title 1", "description": "something"}`,
			expectedEvent: entity.Event{
				Datetime:    time.Date(2022, 6, 29, 15, 28, 0, 0, time.UTC),
				Title:       "test title 1",
				Description: "something",
			},
			expectedErr: nil,
		},
		{
			bodyContent: `{"datetime": "2022-06-29T15:28:00Z", "title": "test title 1"}`,
			expectedEvent: entity.Event{
				Datetime:    time.Date(2022, 6, 29, 15, 28, 0, 0, time.UTC),
				Title:       "test title 1",
				Description: "",
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		event, err := GetEvent(io.NopCloser(strings.NewReader(tc.bodyContent)))

		if !errors.Is(err, tc.expectedErr) {
			t.Errorf("Wrong err value.\tCase\t\"%s\"\nExpected:\t%s\nGot:\t\t%s\n",
				tc.bodyContent, tc.expectedErr, err)
		}

		if !reflect.DeepEqual(tc.expectedEvent, *event) {
			t.Errorf("Wrong event.\t\tCase\t\"%s\"\nExpected:\t%s\nGot:\t\t%s\n",
				tc.bodyContent, tc.expectedEvent, event)
		}
	}
}

func TestArrayPresenter(t *testing.T) {
	testCases := []arrayPresenter{
		{
			array: []entity.Event{
				{
					Datetime:    time.Date(2022, 6, 29, 15, 28, 0, 0, time.UTC),
					Title:       "test title",
					Description: "test desc",
				},
				{
					Datetime:    time.Date(2022, 6, 30, 0, 0, 0, 0, time.UTC),
					Title:       "real title",
					Description: "real desc",
				},
			},
			expectedRepresentation: "[" +
				"{\"title\":\"test title\", \"description\":\"test desc\", \"datetime\":\"2022-06-29T15:28:00Z\"}" +
				"," +
				"{\"title\":\"real title\", \"description\":\"real desc\", \"datetime\":\"2022-06-30T00:00:00Z\"}" +
				"]",
		},
		{
			array: []entity.Event{
				{
					Datetime:    time.Date(2022, 6, 29, 15, 28, 0, 0, time.UTC),
					Title:       "test title",
					Description: "test desc",
				},
			},
			expectedRepresentation: "[" +
				"{\"title\":\"test title\", \"description\":\"test desc\", \"datetime\":\"2022-06-29T15:28:00Z\"}" +
				"]",
		},
	}

	for _, tc := range testCases {
		myResult := ArrayPresenter(tc.array)
		if tc.expectedRepresentation != myResult {
			t.Errorf("Wrong result.\tCase\t\"%s\"\nExpected:\t%s\nGot:\t\t%s\n",
				tc.array, tc.expectedRepresentation, myResult)
		}
	}
}
