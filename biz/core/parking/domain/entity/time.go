package parking_entity

import "time"

type Time struct {
	time.Time
}

func NewTime(t time.Time) Time {
	return Time{Time: t}
}

func (t Time) Into() time.Time {
	return t.Time
}
