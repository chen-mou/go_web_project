package time

import (
	"database/sql/driver"
	"strconv"
	"time"
)

type Time struct {
	Time *time.Time
}

type Timestamp struct {
	Val *int64
}

func (time *Timestamp) Scan(value interface{}) error {
	if value == nil {
		time.Val = nil
		return nil
	}
	val := value.(int64)
	time.Val = &val
	return nil
}

func (time Timestamp) Value() (driver.Value, error) {
	if time.Val == nil || *time.Val == 0 {
		return nil, nil
	}
	return *time.Val, nil
}

func (time *Timestamp) MarshalJson() ([]byte, error) {
	return []byte(strconv.FormatInt(*time.Val, 10)), nil
}

func (time *Timestamp) UnmarshalJson(val []byte) error {
	str := string(val)
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	time.Val = &num
	return nil
}
