// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDailyRevenuePo = "daily_revenue"

// DailyRevenuePo Alarm
type DailyRevenuePo struct {
	Date    string `gorm:"column:date;not null" json:"date"`
	Revenue int64  `gorm:"column:revenue;not null" json:"revenue"`
}

// TableName DailyRevenuePo's table name
func (*DailyRevenuePo) TableName() string {
	return TableNameDailyRevenuePo
}
