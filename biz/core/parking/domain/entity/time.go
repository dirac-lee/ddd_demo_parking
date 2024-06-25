package parking_entity

import "time"

type Time struct {
	time.Time
}

func NewTime(timestamp int64) Time {
	return IntoTime(time.Unix(timestamp, 0))
}

func IntoTime(t time.Time) Time {
	return Time{Time: t}
}

func (t Time) Into() time.Time {
	return t.Time
}
