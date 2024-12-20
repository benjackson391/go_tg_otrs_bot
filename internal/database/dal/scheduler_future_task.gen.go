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

func newSchedulerFutureTask(db *gorm.DB, opts ...gen.DOOption) schedulerFutureTask {
	_schedulerFutureTask := schedulerFutureTask{}

	_schedulerFutureTask.schedulerFutureTaskDo.UseDB(db, opts...)
	_schedulerFutureTask.schedulerFutureTaskDo.UseModel(&model.SchedulerFutureTask{})

	tableName := _schedulerFutureTask.schedulerFutureTaskDo.TableName()
	_schedulerFutureTask.ALL = field.NewAsterisk(tableName)
	_schedulerFutureTask.ID = field.NewInt64(tableName, "id")
	_schedulerFutureTask.Ident = field.NewInt64(tableName, "ident")
	_schedulerFutureTask.ExecutionTime = field.NewTime(tableName, "execution_time")
	_schedulerFutureTask.Name = field.NewString(tableName, "name")
	_schedulerFutureTask.TaskType = field.NewString(tableName, "task_type")
	_schedulerFutureTask.TaskData = field.NewBytes(tableName, "task_data")
	_schedulerFutureTask.Attempts = field.NewInt32(tableName, "attempts")
	_schedulerFutureTask.LockKey = field.NewInt64(tableName, "lock_key")
	_schedulerFutureTask.LockTime = field.NewTime(tableName, "lock_time")
	_schedulerFutureTask.CreateTime = field.NewTime(tableName, "create_time")

	_schedulerFutureTask.fillFieldMap()

	return _schedulerFutureTask
}

type schedulerFutureTask struct {
	schedulerFutureTaskDo

	ALL           field.Asterisk
	ID            field.Int64
	Ident         field.Int64
	ExecutionTime field.Time
	Name          field.String
	TaskType      field.String
	TaskData      field.Bytes
	Attempts      field.Int32
	LockKey       field.Int64
	LockTime      field.Time
	CreateTime    field.Time

	fieldMap map[string]field.Expr
}

func (s schedulerFutureTask) Table(newTableName string) *schedulerFutureTask {
	s.schedulerFutureTaskDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s schedulerFutureTask) As(alias string) *schedulerFutureTask {
	s.schedulerFutureTaskDo.DO = *(s.schedulerFutureTaskDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *schedulerFutureTask) updateTableName(table string) *schedulerFutureTask {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Ident = field.NewInt64(table, "ident")
	s.ExecutionTime = field.NewTime(table, "execution_time")
	s.Name = field.NewString(table, "name")
	s.TaskType = field.NewString(table, "task_type")
	s.TaskData = field.NewBytes(table, "task_data")
	s.Attempts = field.NewInt32(table, "attempts")
	s.LockKey = field.NewInt64(table, "lock_key")
	s.LockTime = field.NewTime(table, "lock_time")
	s.CreateTime = field.NewTime(table, "create_time")

	s.fillFieldMap()

	return s
}

func (s *schedulerFutureTask) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *schedulerFutureTask) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["id"] = s.ID
	s.fieldMap["ident"] = s.Ident
	s.fieldMap["execution_time"] = s.ExecutionTime
	s.fieldMap["name"] = s.Name
	s.fieldMap["task_type"] = s.TaskType
	s.fieldMap["task_data"] = s.TaskData
	s.fieldMap["attempts"] = s.Attempts
	s.fieldMap["lock_key"] = s.LockKey
	s.fieldMap["lock_time"] = s.LockTime
	s.fieldMap["create_time"] = s.CreateTime
}

func (s schedulerFutureTask) clone(db *gorm.DB) schedulerFutureTask {
	s.schedulerFutureTaskDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s schedulerFutureTask) replaceDB(db *gorm.DB) schedulerFutureTask {
	s.schedulerFutureTaskDo.ReplaceDB(db)
	return s
}

type schedulerFutureTaskDo struct{ gen.DO }

type ISchedulerFutureTaskDo interface {
	gen.SubQuery
	Debug() ISchedulerFutureTaskDo
	WithContext(ctx context.Context) ISchedulerFutureTaskDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISchedulerFutureTaskDo
	WriteDB() ISchedulerFutureTaskDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISchedulerFutureTaskDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISchedulerFutureTaskDo
	Not(conds ...gen.Condition) ISchedulerFutureTaskDo
	Or(conds ...gen.Condition) ISchedulerFutureTaskDo
	Select(conds ...field.Expr) ISchedulerFutureTaskDo
	Where(conds ...gen.Condition) ISchedulerFutureTaskDo
	Order(conds ...field.Expr) ISchedulerFutureTaskDo
	Distinct(cols ...field.Expr) ISchedulerFutureTaskDo
	Omit(cols ...field.Expr) ISchedulerFutureTaskDo
	Join(table schema.Tabler, on ...field.Expr) ISchedulerFutureTaskDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISchedulerFutureTaskDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISchedulerFutureTaskDo
	Group(cols ...field.Expr) ISchedulerFutureTaskDo
	Having(conds ...gen.Condition) ISchedulerFutureTaskDo
	Limit(limit int) ISchedulerFutureTaskDo
	Offset(offset int) ISchedulerFutureTaskDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISchedulerFutureTaskDo
	Unscoped() ISchedulerFutureTaskDo
	Create(values ...*model.SchedulerFutureTask) error
	CreateInBatches(values []*model.SchedulerFutureTask, batchSize int) error
	Save(values ...*model.SchedulerFutureTask) error
	First() (*model.SchedulerFutureTask, error)
	Take() (*model.SchedulerFutureTask, error)
	Last() (*model.SchedulerFutureTask, error)
	Find() ([]*model.SchedulerFutureTask, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SchedulerFutureTask, err error)
	FindInBatches(result *[]*model.SchedulerFutureTask, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SchedulerFutureTask) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISchedulerFutureTaskDo
	Assign(attrs ...field.AssignExpr) ISchedulerFutureTaskDo
	Joins(fields ...field.RelationField) ISchedulerFutureTaskDo
	Preload(fields ...field.RelationField) ISchedulerFutureTaskDo
	FirstOrInit() (*model.SchedulerFutureTask, error)
	FirstOrCreate() (*model.SchedulerFutureTask, error)
	FindByPage(offset int, limit int) (result []*model.SchedulerFutureTask, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISchedulerFutureTaskDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s schedulerFutureTaskDo) Debug() ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Debug())
}

func (s schedulerFutureTaskDo) WithContext(ctx context.Context) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s schedulerFutureTaskDo) ReadDB() ISchedulerFutureTaskDo {
	return s.Clauses(dbresolver.Read)
}

func (s schedulerFutureTaskDo) WriteDB() ISchedulerFutureTaskDo {
	return s.Clauses(dbresolver.Write)
}

func (s schedulerFutureTaskDo) Session(config *gorm.Session) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Session(config))
}

func (s schedulerFutureTaskDo) Clauses(conds ...clause.Expression) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s schedulerFutureTaskDo) Returning(value interface{}, columns ...string) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s schedulerFutureTaskDo) Not(conds ...gen.Condition) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s schedulerFutureTaskDo) Or(conds ...gen.Condition) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s schedulerFutureTaskDo) Select(conds ...field.Expr) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s schedulerFutureTaskDo) Where(conds ...gen.Condition) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s schedulerFutureTaskDo) Order(conds ...field.Expr) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s schedulerFutureTaskDo) Distinct(cols ...field.Expr) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s schedulerFutureTaskDo) Omit(cols ...field.Expr) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s schedulerFutureTaskDo) Join(table schema.Tabler, on ...field.Expr) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s schedulerFutureTaskDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s schedulerFutureTaskDo) RightJoin(table schema.Tabler, on ...field.Expr) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s schedulerFutureTaskDo) Group(cols ...field.Expr) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s schedulerFutureTaskDo) Having(conds ...gen.Condition) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s schedulerFutureTaskDo) Limit(limit int) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s schedulerFutureTaskDo) Offset(offset int) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s schedulerFutureTaskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s schedulerFutureTaskDo) Unscoped() ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Unscoped())
}

func (s schedulerFutureTaskDo) Create(values ...*model.SchedulerFutureTask) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s schedulerFutureTaskDo) CreateInBatches(values []*model.SchedulerFutureTask, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s schedulerFutureTaskDo) Save(values ...*model.SchedulerFutureTask) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s schedulerFutureTaskDo) First() (*model.SchedulerFutureTask, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SchedulerFutureTask), nil
	}
}

func (s schedulerFutureTaskDo) Take() (*model.SchedulerFutureTask, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SchedulerFutureTask), nil
	}
}

func (s schedulerFutureTaskDo) Last() (*model.SchedulerFutureTask, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SchedulerFutureTask), nil
	}
}

func (s schedulerFutureTaskDo) Find() ([]*model.SchedulerFutureTask, error) {
	result, err := s.DO.Find()
	return result.([]*model.SchedulerFutureTask), err
}

func (s schedulerFutureTaskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SchedulerFutureTask, err error) {
	buf := make([]*model.SchedulerFutureTask, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s schedulerFutureTaskDo) FindInBatches(result *[]*model.SchedulerFutureTask, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s schedulerFutureTaskDo) Attrs(attrs ...field.AssignExpr) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s schedulerFutureTaskDo) Assign(attrs ...field.AssignExpr) ISchedulerFutureTaskDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s schedulerFutureTaskDo) Joins(fields ...field.RelationField) ISchedulerFutureTaskDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s schedulerFutureTaskDo) Preload(fields ...field.RelationField) ISchedulerFutureTaskDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s schedulerFutureTaskDo) FirstOrInit() (*model.SchedulerFutureTask, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SchedulerFutureTask), nil
	}
}

func (s schedulerFutureTaskDo) FirstOrCreate() (*model.SchedulerFutureTask, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SchedulerFutureTask), nil
	}
}

func (s schedulerFutureTaskDo) FindByPage(offset int, limit int) (result []*model.SchedulerFutureTask, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s schedulerFutureTaskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s schedulerFutureTaskDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s schedulerFutureTaskDo) Delete(models ...*model.SchedulerFutureTask) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *schedulerFutureTaskDo) withDO(do gen.Dao) *schedulerFutureTaskDo {
	s.DO = *do.(*gen.DO)
	return s
}
