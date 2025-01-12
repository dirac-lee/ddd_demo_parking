// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAlarmPo = "alarm"

// AlarmPo Alarm
type AlarmPo struct {
	ID      int64     `gorm:"column:id;primaryKey" json:"id"`
	Plate   string    `gorm:"column:plate;not null" json:"plate"`
	Message string    `gorm:"column:message;not null" json:"message"`
	Time    time.Time `gorm:"column:time;not null;default:CURRENT_TIMESTAMP" json:"time"`
}

// TableName AlarmPo's table name
func (*AlarmPo) TableName() string {
	return TableNameAlarmPo
}
