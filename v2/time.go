package infobip

import (
	"strings"
	"time"
)

type Time struct {
	T time.Time
}

const INFOBIP_TIME_FORMAT = "2006-01-02T15:04:05.000-0700"

func TimeNow() Time {
	s := time.Now().Format(INFOBIP_TIME_FORMAT)
	t, _ := time.Parse(INFOBIP_TIME_FORMAT, s)
	ibtime := Time{t}
	return ibtime
}

func (ibtime Time) MarshalJSON() ([]byte, error) {
	s := time.Time(ibtime.T).Format(INFOBIP_TIME_FORMAT)
	return []byte(`"` + s + `"`), nil
}

func (ibtime *Time) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	t, err := time.Parse(INFOBIP_TIME_FORMAT, s)
	if err != nil {
		return err
	}
	*ibtime = Time{t}
	return nil
}

func (ibtime Time) String() string {
	s := ibtime.T.Format(INFOBIP_TIME_FORMAT)
	return s
}
