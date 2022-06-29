package utils

import (
	"WBL2/Tasks/T11/internal/domain/entity"
	domainErrors "WBL2/Tasks/T11/internal/domain/errors"
	"github.com/mailru/easyjson"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetEvent(r io.ReadCloser) (*entity.Event, error) {
	var event entity.Event

	defer func() { _ = r.Close() }()
	if err := easyjson.UnmarshalFromReader(r, &event); err != nil {
		return nil, err
	}

	return &event, nil
}

const queryTimeLayout = `2006-01-02`

func GetPeriod(r *http.Request) (from, to time.Time, err error) {
	defer func() { _ = r.Body.Close() }()

	urlParams := r.URL.Query()
	if !urlParams.Has("date") {
		return from, to, domainErrors.ErrWrongQueryParams
	}

	from, err = time.Parse(queryTimeLayout, urlParams.Get("date"))
	if err != nil {
		return from, to, domainErrors.ErrWrongQueryParams
	}

	switch p := urlParams.Get("period"); p {
	case "month":
		to = from.Add(time.Hour * 24 * 30) //nolint:gomnd
	case "week":
		to = from.Add(time.Hour * 24 * 7) //nolint:gomnd
	case "day", "":
		to = from.Add(time.Hour * 24) //nolint:gomnd
	default:
		return from, to, domainErrors.ErrWrongQueryParams
	}

	return from, to, nil
}

func ArrayPresenter(events []entity.Event) string {
	resp := new(strings.Builder)

	resp.Write([]byte("["))
	for i, event := range events {
		resp.WriteString(event.String())
		if i < len(events)-1 {
			resp.WriteByte(',')
		}
	}

	resp.WriteByte(']')

	return resp.String()
}
