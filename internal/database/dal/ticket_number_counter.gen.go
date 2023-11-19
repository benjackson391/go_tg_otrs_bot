// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"tg_bot/internal/database/model"
)

func newTicketNumberCounter(db *gorm.DB, opts ...gen.DOOption) ticketNumberCounter {
	_ticketNumberCounter := ticketNumberCounter{}

	_ticketNumberCounter.ticketNumberCounterDo.UseDB(db, opts...)
	_ticketNumberCounter.ticketNumberCounterDo.UseModel(&model.TicketNumberCounter{})

	tableName := _ticketNumberCounter.ticketNumberCounterDo.TableName()
	_ticketNumberCounter.ALL = field.NewAsterisk(tableName)
	_ticketNumberCounter.ID = field.NewInt64(tableName, "id")
	_ticketNumberCounter.Counter = field.NewInt64(tableName, "counter")
	_ticketNumberCounter.CounterUID = field.NewString(tableName, "counter_uid")
	_ticketNumberCounter.CreateTime = field.NewTime(tableName, "create_time")

	_ticketNumberCounter.fillFieldMap()

	return _ticketNumberCounter
}

type ticketNumberCounter struct {
	ticketNumberCounterDo

	ALL        field.Asterisk
	ID         field.Int64
	Counter    field.Int64
	CounterUID field.String
	CreateTime field.Time

	fieldMap map[string]field.Expr
}

func (t ticketNumberCounter) Table(newTableName string) *ticketNumberCounter {
	t.ticketNumberCounterDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t ticketNumberCounter) As(alias string) *ticketNumberCounter {
	t.ticketNumberCounterDo.DO = *(t.ticketNumberCounterDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *ticketNumberCounter) updateTableName(table string) *ticketNumberCounter {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.Counter = field.NewInt64(table, "counter")
	t.CounterUID = field.NewString(table, "counter_uid")
	t.CreateTime = field.NewTime(table, "create_time")

	t.fillFieldMap()

	return t
}

func (t *ticketNumberCounter) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *ticketNumberCounter) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 4)
	t.fieldMap["id"] = t.ID
	t.fieldMap["counter"] = t.Counter
	t.fieldMap["counter_uid"] = t.CounterUID
	t.fieldMap["create_time"] = t.CreateTime
}

func (t ticketNumberCounter) clone(db *gorm.DB) ticketNumberCounter {
	t.ticketNumberCounterDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t ticketNumberCounter) replaceDB(db *gorm.DB) ticketNumberCounter {
	t.ticketNumberCounterDo.ReplaceDB(db)
	return t
}

type ticketNumberCounterDo struct{ gen.DO }

type ITicketNumberCounterDo interface {
	gen.SubQuery
	Debug() ITicketNumberCounterDo
	WithContext(ctx context.Context) ITicketNumberCounterDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITicketNumberCounterDo
	WriteDB() ITicketNumberCounterDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITicketNumberCounterDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITicketNumberCounterDo
	Not(conds ...gen.Condition) ITicketNumberCounterDo
	Or(conds ...gen.Condition) ITicketNumberCounterDo
	Select(conds ...field.Expr) ITicketNumberCounterDo
	Where(conds ...gen.Condition) ITicketNumberCounterDo
	Order(conds ...field.Expr) ITicketNumberCounterDo
	Distinct(cols ...field.Expr) ITicketNumberCounterDo
	Omit(cols ...field.Expr) ITicketNumberCounterDo
	Join(table schema.Tabler, on ...field.Expr) ITicketNumberCounterDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITicketNumberCounterDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITicketNumberCounterDo
	Group(cols ...field.Expr) ITicketNumberCounterDo
	Having(conds ...gen.Condition) ITicketNumberCounterDo
	Limit(limit int) ITicketNumberCounterDo
	Offset(offset int) ITicketNumberCounterDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITicketNumberCounterDo
	Unscoped() ITicketNumberCounterDo
	Create(values ...*model.TicketNumberCounter) error
	CreateInBatches(values []*model.TicketNumberCounter, batchSize int) error
	Save(values ...*model.TicketNumberCounter) error
	First() (*model.TicketNumberCounter, error)
	Take() (*model.TicketNumberCounter, error)
	Last() (*model.TicketNumberCounter, error)
	Find() ([]*model.TicketNumberCounter, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TicketNumberCounter, err error)
	FindInBatches(result *[]*model.TicketNumberCounter, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TicketNumberCounter) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITicketNumberCounterDo
	Assign(attrs ...field.AssignExpr) ITicketNumberCounterDo
	Joins(fields ...field.RelationField) ITicketNumberCounterDo
	Preload(fields ...field.RelationField) ITicketNumberCounterDo
	FirstOrInit() (*model.TicketNumberCounter, error)
	FirstOrCreate() (*model.TicketNumberCounter, error)
	FindByPage(offset int, limit int) (result []*model.TicketNumberCounter, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITicketNumberCounterDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t ticketNumberCounterDo) Debug() ITicketNumberCounterDo {
	return t.withDO(t.DO.Debug())
}

func (t ticketNumberCounterDo) WithContext(ctx context.Context) ITicketNumberCounterDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t ticketNumberCounterDo) ReadDB() ITicketNumberCounterDo {
	return t.Clauses(dbresolver.Read)
}

func (t ticketNumberCounterDo) WriteDB() ITicketNumberCounterDo {
	return t.Clauses(dbresolver.Write)
}

func (t ticketNumberCounterDo) Session(config *gorm.Session) ITicketNumberCounterDo {
	return t.withDO(t.DO.Session(config))
}

func (t ticketNumberCounterDo) Clauses(conds ...clause.Expression) ITicketNumberCounterDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t ticketNumberCounterDo) Returning(value interface{}, columns ...string) ITicketNumberCounterDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t ticketNumberCounterDo) Not(conds ...gen.Condition) ITicketNumberCounterDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t ticketNumberCounterDo) Or(conds ...gen.Condition) ITicketNumberCounterDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t ticketNumberCounterDo) Select(conds ...field.Expr) ITicketNumberCounterDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t ticketNumberCounterDo) Where(conds ...gen.Condition) ITicketNumberCounterDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t ticketNumberCounterDo) Order(conds ...field.Expr) ITicketNumberCounterDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t ticketNumberCounterDo) Distinct(cols ...field.Expr) ITicketNumberCounterDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t ticketNumberCounterDo) Omit(cols ...field.Expr) ITicketNumberCounterDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t ticketNumberCounterDo) Join(table schema.Tabler, on ...field.Expr) ITicketNumberCounterDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t ticketNumberCounterDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITicketNumberCounterDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t ticketNumberCounterDo) RightJoin(table schema.Tabler, on ...field.Expr) ITicketNumberCounterDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t ticketNumberCounterDo) Group(cols ...field.Expr) ITicketNumberCounterDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t ticketNumberCounterDo) Having(conds ...gen.Condition) ITicketNumberCounterDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t ticketNumberCounterDo) Limit(limit int) ITicketNumberCounterDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t ticketNumberCounterDo) Offset(offset int) ITicketNumberCounterDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t ticketNumberCounterDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITicketNumberCounterDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t ticketNumberCounterDo) Unscoped() ITicketNumberCounterDo {
	return t.withDO(t.DO.Unscoped())
}

func (t ticketNumberCounterDo) Create(values ...*model.TicketNumberCounter) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t ticketNumberCounterDo) CreateInBatches(values []*model.TicketNumberCounter, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t ticketNumberCounterDo) Save(values ...*model.TicketNumberCounter) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t ticketNumberCounterDo) First() (*model.TicketNumberCounter, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TicketNumberCounter), nil
	}
}

func (t ticketNumberCounterDo) Take() (*model.TicketNumberCounter, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TicketNumberCounter), nil
	}
}

func (t ticketNumberCounterDo) Last() (*model.TicketNumberCounter, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TicketNumberCounter), nil
	}
}

func (t ticketNumberCounterDo) Find() ([]*model.TicketNumberCounter, error) {
	result, err := t.DO.Find()
	return result.([]*model.TicketNumberCounter), err
}

func (t ticketNumberCounterDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TicketNumberCounter, err error) {
	buf := make([]*model.TicketNumberCounter, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t ticketNumberCounterDo) FindInBatches(result *[]*model.TicketNumberCounter, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t ticketNumberCounterDo) Attrs(attrs ...field.AssignExpr) ITicketNumberCounterDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t ticketNumberCounterDo) Assign(attrs ...field.AssignExpr) ITicketNumberCounterDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t ticketNumberCounterDo) Joins(fields ...field.RelationField) ITicketNumberCounterDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t ticketNumberCounterDo) Preload(fields ...field.RelationField) ITicketNumberCounterDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t ticketNumberCounterDo) FirstOrInit() (*model.TicketNumberCounter, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TicketNumberCounter), nil
	}
}

func (t ticketNumberCounterDo) FirstOrCreate() (*model.TicketNumberCounter, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TicketNumberCounter), nil
	}
}

func (t ticketNumberCounterDo) FindByPage(offset int, limit int) (result []*model.TicketNumberCounter, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t ticketNumberCounterDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t ticketNumberCounterDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t ticketNumberCounterDo) Delete(models ...*model.TicketNumberCounter) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *ticketNumberCounterDo) withDO(do gen.Dao) *ticketNumberCounterDo {
	t.DO = *do.(*gen.DO)
	return t
}
