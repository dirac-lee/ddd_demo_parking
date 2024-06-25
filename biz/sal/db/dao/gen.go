// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:             db,
		AlarmPo:        newAlarmPo(db, opts...),
		DailyRevenuePo: newDailyRevenuePo(db, opts...),
		ParkingPo:      newParkingPo(db, opts...),
		ParkingViewPo:  newParkingViewPo(db, opts...),
		SummaryPo:      newSummaryPo(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AlarmPo        alarmPo
	DailyRevenuePo dailyRevenuePo
	ParkingPo      parkingPo
	ParkingViewPo  parkingViewPo
	SummaryPo      summaryPo
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:             db,
		AlarmPo:        q.AlarmPo.clone(db),
		DailyRevenuePo: q.DailyRevenuePo.clone(db),
		ParkingPo:      q.ParkingPo.clone(db),
		ParkingViewPo:  q.ParkingViewPo.clone(db),
		SummaryPo:      q.SummaryPo.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:             db,
		AlarmPo:        q.AlarmPo.replaceDB(db),
		DailyRevenuePo: q.DailyRevenuePo.replaceDB(db),
		ParkingPo:      q.ParkingPo.replaceDB(db),
		ParkingViewPo:  q.ParkingViewPo.replaceDB(db),
		SummaryPo:      q.SummaryPo.replaceDB(db),
	}
}

type queryCtx struct {
	AlarmPo        IAlarmPoDo
	DailyRevenuePo IDailyRevenuePoDo
	ParkingPo      IParkingPoDo
	ParkingViewPo  IParkingViewPoDo
	SummaryPo      ISummaryPoDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AlarmPo:        q.AlarmPo.WithContext(ctx),
		DailyRevenuePo: q.DailyRevenuePo.WithContext(ctx),
		ParkingPo:      q.ParkingPo.WithContext(ctx),
		ParkingViewPo:  q.ParkingViewPo.WithContext(ctx),
		SummaryPo:      q.SummaryPo.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
