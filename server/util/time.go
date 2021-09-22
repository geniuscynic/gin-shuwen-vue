package util

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}

func (date *Time) Scan(value interface{}) (err error) {
	switch vt := value.(type) {
	case string:
		time, _ := time.Parse(timeFormart, vt)

		*date = Time(time)
	default:
		return errors.New("类型处理错误")
	}

	return nil
}

func (date Time) Value() (driver.Value, error) {
	return time.Time(date).Format(timeFormart), nil
}
