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

func newCommunicationLog(db *gorm.DB, opts ...gen.DOOption) communicationLog {
	_communicationLog := communicationLog{}

	_communicationLog.communicationLogDo.UseDB(db, opts...)
	_communicationLog.communicationLogDo.UseModel(&model.CommunicationLog{})

	tableName := _communicationLog.communicationLogDo.TableName()
	_communicationLog.ALL = field.NewAsterisk(tableName)
	_communicationLog.ID = field.NewInt64(tableName, "id")
	_communicationLog.InsertFingerprint = field.NewString(tableName, "insert_fingerprint")
	_communicationLog.Transport = field.NewString(tableName, "transport")
	_communicationLog.Direction = field.NewString(tableName, "direction")
	_communicationLog.Status = field.NewString(tableName, "status")
	_communicationLog.AccountType = field.NewString(tableName, "account_type")
	_communicationLog.AccountID = field.NewString(tableName, "account_id")
	_communicationLog.StartTime = field.NewTime(tableName, "start_time")
	_communicationLog.EndTime = field.NewTime(tableName, "end_time")

	_communicationLog.fillFieldMap()

	return _communicationLog
}

type communicationLog struct {
	communicationLogDo

	ALL               field.Asterisk
	ID                field.Int64
	InsertFingerprint field.String
	Transport         field.String
	Direction         field.String
	Status            field.String
	AccountType       field.String
	AccountID         field.String
	StartTime         field.Time
	EndTime           field.Time

	fieldMap map[string]field.Expr
}

func (c communicationLog) Table(newTableName string) *communicationLog {
	c.communicationLogDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c communicationLog) As(alias string) *communicationLog {
	c.communicationLogDo.DO = *(c.communicationLogDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *communicationLog) updateTableName(table string) *communicationLog {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.InsertFingerprint = field.NewString(table, "insert_fingerprint")
	c.Transport = field.NewString(table, "transport")
	c.Direction = field.NewString(table, "direction")
	c.Status = field.NewString(table, "status")
	c.AccountType = field.NewString(table, "account_type")
	c.AccountID = field.NewString(table, "account_id")
	c.StartTime = field.NewTime(table, "start_time")
	c.EndTime = field.NewTime(table, "end_time")

	c.fillFieldMap()

	return c
}

func (c *communicationLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *communicationLog) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["insert_fingerprint"] = c.InsertFingerprint
	c.fieldMap["transport"] = c.Transport
	c.fieldMap["direction"] = c.Direction
	c.fieldMap["status"] = c.Status
	c.fieldMap["account_type"] = c.AccountType
	c.fieldMap["account_id"] = c.AccountID
	c.fieldMap["start_time"] = c.StartTime
	c.fieldMap["end_time"] = c.EndTime
}

func (c communicationLog) clone(db *gorm.DB) communicationLog {
	c.communicationLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c communicationLog) replaceDB(db *gorm.DB) communicationLog {
	c.communicationLogDo.ReplaceDB(db)
	return c
}

type communicationLogDo struct{ gen.DO }

type ICommunicationLogDo interface {
	gen.SubQuery
	Debug() ICommunicationLogDo
	WithContext(ctx context.Context) ICommunicationLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICommunicationLogDo
	WriteDB() ICommunicationLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICommunicationLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICommunicationLogDo
	Not(conds ...gen.Condition) ICommunicationLogDo
	Or(conds ...gen.Condition) ICommunicationLogDo
	Select(conds ...field.Expr) ICommunicationLogDo
	Where(conds ...gen.Condition) ICommunicationLogDo
	Order(conds ...field.Expr) ICommunicationLogDo
	Distinct(cols ...field.Expr) ICommunicationLogDo
	Omit(cols ...field.Expr) ICommunicationLogDo
	Join(table schema.Tabler, on ...field.Expr) ICommunicationLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICommunicationLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICommunicationLogDo
	Group(cols ...field.Expr) ICommunicationLogDo
	Having(conds ...gen.Condition) ICommunicationLogDo
	Limit(limit int) ICommunicationLogDo
	Offset(offset int) ICommunicationLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICommunicationLogDo
	Unscoped() ICommunicationLogDo
	Create(values ...*model.CommunicationLog) error
	CreateInBatches(values []*model.CommunicationLog, batchSize int) error
	Save(values ...*model.CommunicationLog) error
	First() (*model.CommunicationLog, error)
	Take() (*model.CommunicationLog, error)
	Last() (*model.CommunicationLog, error)
	Find() ([]*model.CommunicationLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommunicationLog, err error)
	FindInBatches(result *[]*model.CommunicationLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CommunicationLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICommunicationLogDo
	Assign(attrs ...field.AssignExpr) ICommunicationLogDo
	Joins(fields ...field.RelationField) ICommunicationLogDo
	Preload(fields ...field.RelationField) ICommunicationLogDo
	FirstOrInit() (*model.CommunicationLog, error)
	FirstOrCreate() (*model.CommunicationLog, error)
	FindByPage(offset int, limit int) (result []*model.CommunicationLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICommunicationLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c communicationLogDo) Debug() ICommunicationLogDo {
	return c.withDO(c.DO.Debug())
}

func (c communicationLogDo) WithContext(ctx context.Context) ICommunicationLogDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c communicationLogDo) ReadDB() ICommunicationLogDo {
	return c.Clauses(dbresolver.Read)
}

func (c communicationLogDo) WriteDB() ICommunicationLogDo {
	return c.Clauses(dbresolver.Write)
}

func (c communicationLogDo) Session(config *gorm.Session) ICommunicationLogDo {
	return c.withDO(c.DO.Session(config))
}

func (c communicationLogDo) Clauses(conds ...clause.Expression) ICommunicationLogDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c communicationLogDo) Returning(value interface{}, columns ...string) ICommunicationLogDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c communicationLogDo) Not(conds ...gen.Condition) ICommunicationLogDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c communicationLogDo) Or(conds ...gen.Condition) ICommunicationLogDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c communicationLogDo) Select(conds ...field.Expr) ICommunicationLogDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c communicationLogDo) Where(conds ...gen.Condition) ICommunicationLogDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c communicationLogDo) Order(conds ...field.Expr) ICommunicationLogDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c communicationLogDo) Distinct(cols ...field.Expr) ICommunicationLogDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c communicationLogDo) Omit(cols ...field.Expr) ICommunicationLogDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c communicationLogDo) Join(table schema.Tabler, on ...field.Expr) ICommunicationLogDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c communicationLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICommunicationLogDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c communicationLogDo) RightJoin(table schema.Tabler, on ...field.Expr) ICommunicationLogDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c communicationLogDo) Group(cols ...field.Expr) ICommunicationLogDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c communicationLogDo) Having(conds ...gen.Condition) ICommunicationLogDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c communicationLogDo) Limit(limit int) ICommunicationLogDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c communicationLogDo) Offset(offset int) ICommunicationLogDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c communicationLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICommunicationLogDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c communicationLogDo) Unscoped() ICommunicationLogDo {
	return c.withDO(c.DO.Unscoped())
}

func (c communicationLogDo) Create(values ...*model.CommunicationLog) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c communicationLogDo) CreateInBatches(values []*model.CommunicationLog, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c communicationLogDo) Save(values ...*model.CommunicationLog) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c communicationLogDo) First() (*model.CommunicationLog, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommunicationLog), nil
	}
}

func (c communicationLogDo) Take() (*model.CommunicationLog, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommunicationLog), nil
	}
}

func (c communicationLogDo) Last() (*model.CommunicationLog, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommunicationLog), nil
	}
}

func (c communicationLogDo) Find() ([]*model.CommunicationLog, error) {
	result, err := c.DO.Find()
	return result.([]*model.CommunicationLog), err
}

func (c communicationLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommunicationLog, err error) {
	buf := make([]*model.CommunicationLog, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c communicationLogDo) FindInBatches(result *[]*model.CommunicationLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c communicationLogDo) Attrs(attrs ...field.AssignExpr) ICommunicationLogDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c communicationLogDo) Assign(attrs ...field.AssignExpr) ICommunicationLogDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c communicationLogDo) Joins(fields ...field.RelationField) ICommunicationLogDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c communicationLogDo) Preload(fields ...field.RelationField) ICommunicationLogDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c communicationLogDo) FirstOrInit() (*model.CommunicationLog, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommunicationLog), nil
	}
}

func (c communicationLogDo) FirstOrCreate() (*model.CommunicationLog, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommunicationLog), nil
	}
}

func (c communicationLogDo) FindByPage(offset int, limit int) (result []*model.CommunicationLog, count int64, err error) {
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

func (c communicationLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c communicationLogDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c communicationLogDo) Delete(models ...*model.CommunicationLog) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *communicationLogDo) withDO(do gen.Dao) *communicationLogDo {
	c.DO = *do.(*gen.DO)
	return c
}
