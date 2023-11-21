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

func newCalendar(db *gorm.DB, opts ...gen.DOOption) calendar {
	_calendar := calendar{}

	_calendar.calendarDo.UseDB(db, opts...)
	_calendar.calendarDo.UseModel(&model.Calendar{})

	tableName := _calendar.calendarDo.TableName()
	_calendar.ALL = field.NewAsterisk(tableName)
	_calendar.ID = field.NewInt64(tableName, "id")
	_calendar.GroupID = field.NewInt32(tableName, "group_id")
	_calendar.Name = field.NewString(tableName, "name")
	_calendar.SaltString = field.NewString(tableName, "salt_string")
	_calendar.Color = field.NewString(tableName, "color")
	_calendar.TicketAppointments = field.NewBytes(tableName, "ticket_appointments")
	_calendar.ValidID = field.NewInt32(tableName, "valid_id")
	_calendar.CreateTime = field.NewTime(tableName, "create_time")
	_calendar.CreateBy = field.NewInt32(tableName, "create_by")
	_calendar.ChangeTime = field.NewTime(tableName, "change_time")
	_calendar.ChangeBy = field.NewInt32(tableName, "change_by")

	_calendar.fillFieldMap()

	return _calendar
}

type calendar struct {
	calendarDo

	ALL                field.Asterisk
	ID                 field.Int64
	GroupID            field.Int32
	Name               field.String
	SaltString         field.String
	Color              field.String
	TicketAppointments field.Bytes
	ValidID            field.Int32
	CreateTime         field.Time
	CreateBy           field.Int32
	ChangeTime         field.Time
	ChangeBy           field.Int32

	fieldMap map[string]field.Expr
}

func (c calendar) Table(newTableName string) *calendar {
	c.calendarDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c calendar) As(alias string) *calendar {
	c.calendarDo.DO = *(c.calendarDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *calendar) updateTableName(table string) *calendar {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.GroupID = field.NewInt32(table, "group_id")
	c.Name = field.NewString(table, "name")
	c.SaltString = field.NewString(table, "salt_string")
	c.Color = field.NewString(table, "color")
	c.TicketAppointments = field.NewBytes(table, "ticket_appointments")
	c.ValidID = field.NewInt32(table, "valid_id")
	c.CreateTime = field.NewTime(table, "create_time")
	c.CreateBy = field.NewInt32(table, "create_by")
	c.ChangeTime = field.NewTime(table, "change_time")
	c.ChangeBy = field.NewInt32(table, "change_by")

	c.fillFieldMap()

	return c
}

func (c *calendar) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *calendar) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 11)
	c.fieldMap["id"] = c.ID
	c.fieldMap["group_id"] = c.GroupID
	c.fieldMap["name"] = c.Name
	c.fieldMap["salt_string"] = c.SaltString
	c.fieldMap["color"] = c.Color
	c.fieldMap["ticket_appointments"] = c.TicketAppointments
	c.fieldMap["valid_id"] = c.ValidID
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["create_by"] = c.CreateBy
	c.fieldMap["change_time"] = c.ChangeTime
	c.fieldMap["change_by"] = c.ChangeBy
}

func (c calendar) clone(db *gorm.DB) calendar {
	c.calendarDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c calendar) replaceDB(db *gorm.DB) calendar {
	c.calendarDo.ReplaceDB(db)
	return c
}

type calendarDo struct{ gen.DO }

type ICalendarDo interface {
	gen.SubQuery
	Debug() ICalendarDo
	WithContext(ctx context.Context) ICalendarDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICalendarDo
	WriteDB() ICalendarDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICalendarDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICalendarDo
	Not(conds ...gen.Condition) ICalendarDo
	Or(conds ...gen.Condition) ICalendarDo
	Select(conds ...field.Expr) ICalendarDo
	Where(conds ...gen.Condition) ICalendarDo
	Order(conds ...field.Expr) ICalendarDo
	Distinct(cols ...field.Expr) ICalendarDo
	Omit(cols ...field.Expr) ICalendarDo
	Join(table schema.Tabler, on ...field.Expr) ICalendarDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICalendarDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICalendarDo
	Group(cols ...field.Expr) ICalendarDo
	Having(conds ...gen.Condition) ICalendarDo
	Limit(limit int) ICalendarDo
	Offset(offset int) ICalendarDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICalendarDo
	Unscoped() ICalendarDo
	Create(values ...*model.Calendar) error
	CreateInBatches(values []*model.Calendar, batchSize int) error
	Save(values ...*model.Calendar) error
	First() (*model.Calendar, error)
	Take() (*model.Calendar, error)
	Last() (*model.Calendar, error)
	Find() ([]*model.Calendar, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Calendar, err error)
	FindInBatches(result *[]*model.Calendar, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Calendar) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICalendarDo
	Assign(attrs ...field.AssignExpr) ICalendarDo
	Joins(fields ...field.RelationField) ICalendarDo
	Preload(fields ...field.RelationField) ICalendarDo
	FirstOrInit() (*model.Calendar, error)
	FirstOrCreate() (*model.Calendar, error)
	FindByPage(offset int, limit int) (result []*model.Calendar, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICalendarDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c calendarDo) Debug() ICalendarDo {
	return c.withDO(c.DO.Debug())
}

func (c calendarDo) WithContext(ctx context.Context) ICalendarDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c calendarDo) ReadDB() ICalendarDo {
	return c.Clauses(dbresolver.Read)
}

func (c calendarDo) WriteDB() ICalendarDo {
	return c.Clauses(dbresolver.Write)
}

func (c calendarDo) Session(config *gorm.Session) ICalendarDo {
	return c.withDO(c.DO.Session(config))
}

func (c calendarDo) Clauses(conds ...clause.Expression) ICalendarDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c calendarDo) Returning(value interface{}, columns ...string) ICalendarDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c calendarDo) Not(conds ...gen.Condition) ICalendarDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c calendarDo) Or(conds ...gen.Condition) ICalendarDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c calendarDo) Select(conds ...field.Expr) ICalendarDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c calendarDo) Where(conds ...gen.Condition) ICalendarDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c calendarDo) Order(conds ...field.Expr) ICalendarDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c calendarDo) Distinct(cols ...field.Expr) ICalendarDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c calendarDo) Omit(cols ...field.Expr) ICalendarDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c calendarDo) Join(table schema.Tabler, on ...field.Expr) ICalendarDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c calendarDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICalendarDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c calendarDo) RightJoin(table schema.Tabler, on ...field.Expr) ICalendarDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c calendarDo) Group(cols ...field.Expr) ICalendarDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c calendarDo) Having(conds ...gen.Condition) ICalendarDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c calendarDo) Limit(limit int) ICalendarDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c calendarDo) Offset(offset int) ICalendarDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c calendarDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICalendarDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c calendarDo) Unscoped() ICalendarDo {
	return c.withDO(c.DO.Unscoped())
}

func (c calendarDo) Create(values ...*model.Calendar) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c calendarDo) CreateInBatches(values []*model.Calendar, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c calendarDo) Save(values ...*model.Calendar) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c calendarDo) First() (*model.Calendar, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Calendar), nil
	}
}

func (c calendarDo) Take() (*model.Calendar, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Calendar), nil
	}
}

func (c calendarDo) Last() (*model.Calendar, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Calendar), nil
	}
}

func (c calendarDo) Find() ([]*model.Calendar, error) {
	result, err := c.DO.Find()
	return result.([]*model.Calendar), err
}

func (c calendarDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Calendar, err error) {
	buf := make([]*model.Calendar, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c calendarDo) FindInBatches(result *[]*model.Calendar, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c calendarDo) Attrs(attrs ...field.AssignExpr) ICalendarDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c calendarDo) Assign(attrs ...field.AssignExpr) ICalendarDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c calendarDo) Joins(fields ...field.RelationField) ICalendarDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c calendarDo) Preload(fields ...field.RelationField) ICalendarDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c calendarDo) FirstOrInit() (*model.Calendar, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Calendar), nil
	}
}

func (c calendarDo) FirstOrCreate() (*model.Calendar, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Calendar), nil
	}
}

func (c calendarDo) FindByPage(offset int, limit int) (result []*model.Calendar, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c calendarDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c calendarDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c calendarDo) Delete(models ...*model.Calendar) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *calendarDo) withDO(do gen.Dao) *calendarDo {
	c.DO = *do.(*gen.DO)
	return c
}