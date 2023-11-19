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

func newSystemMaintenance(db *gorm.DB, opts ...gen.DOOption) systemMaintenance {
	_systemMaintenance := systemMaintenance{}

	_systemMaintenance.systemMaintenanceDo.UseDB(db, opts...)
	_systemMaintenance.systemMaintenanceDo.UseModel(&model.SystemMaintenance{})

	tableName := _systemMaintenance.systemMaintenanceDo.TableName()
	_systemMaintenance.ALL = field.NewAsterisk(tableName)
	_systemMaintenance.ID = field.NewInt32(tableName, "id")
	_systemMaintenance.StartDate = field.NewInt32(tableName, "start_date")
	_systemMaintenance.StopDate = field.NewInt32(tableName, "stop_date")
	_systemMaintenance.Comments = field.NewString(tableName, "comments")
	_systemMaintenance.LoginMessage = field.NewString(tableName, "login_message")
	_systemMaintenance.ShowLoginMessage = field.NewInt32(tableName, "show_login_message")
	_systemMaintenance.NotifyMessage = field.NewString(tableName, "notify_message")
	_systemMaintenance.ValidID = field.NewInt32(tableName, "valid_id")
	_systemMaintenance.CreateTime = field.NewTime(tableName, "create_time")
	_systemMaintenance.CreateBy = field.NewInt32(tableName, "create_by")
	_systemMaintenance.ChangeTime = field.NewTime(tableName, "change_time")
	_systemMaintenance.ChangeBy = field.NewInt32(tableName, "change_by")

	_systemMaintenance.fillFieldMap()

	return _systemMaintenance
}

type systemMaintenance struct {
	systemMaintenanceDo

	ALL              field.Asterisk
	ID               field.Int32
	StartDate        field.Int32
	StopDate         field.Int32
	Comments         field.String
	LoginMessage     field.String
	ShowLoginMessage field.Int32
	NotifyMessage    field.String
	ValidID          field.Int32
	CreateTime       field.Time
	CreateBy         field.Int32
	ChangeTime       field.Time
	ChangeBy         field.Int32

	fieldMap map[string]field.Expr
}

func (s systemMaintenance) Table(newTableName string) *systemMaintenance {
	s.systemMaintenanceDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s systemMaintenance) As(alias string) *systemMaintenance {
	s.systemMaintenanceDo.DO = *(s.systemMaintenanceDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *systemMaintenance) updateTableName(table string) *systemMaintenance {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.StartDate = field.NewInt32(table, "start_date")
	s.StopDate = field.NewInt32(table, "stop_date")
	s.Comments = field.NewString(table, "comments")
	s.LoginMessage = field.NewString(table, "login_message")
	s.ShowLoginMessage = field.NewInt32(table, "show_login_message")
	s.NotifyMessage = field.NewString(table, "notify_message")
	s.ValidID = field.NewInt32(table, "valid_id")
	s.CreateTime = field.NewTime(table, "create_time")
	s.CreateBy = field.NewInt32(table, "create_by")
	s.ChangeTime = field.NewTime(table, "change_time")
	s.ChangeBy = field.NewInt32(table, "change_by")

	s.fillFieldMap()

	return s
}

func (s *systemMaintenance) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *systemMaintenance) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["id"] = s.ID
	s.fieldMap["start_date"] = s.StartDate
	s.fieldMap["stop_date"] = s.StopDate
	s.fieldMap["comments"] = s.Comments
	s.fieldMap["login_message"] = s.LoginMessage
	s.fieldMap["show_login_message"] = s.ShowLoginMessage
	s.fieldMap["notify_message"] = s.NotifyMessage
	s.fieldMap["valid_id"] = s.ValidID
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["change_time"] = s.ChangeTime
	s.fieldMap["change_by"] = s.ChangeBy
}

func (s systemMaintenance) clone(db *gorm.DB) systemMaintenance {
	s.systemMaintenanceDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s systemMaintenance) replaceDB(db *gorm.DB) systemMaintenance {
	s.systemMaintenanceDo.ReplaceDB(db)
	return s
}

type systemMaintenanceDo struct{ gen.DO }

type ISystemMaintenanceDo interface {
	gen.SubQuery
	Debug() ISystemMaintenanceDo
	WithContext(ctx context.Context) ISystemMaintenanceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISystemMaintenanceDo
	WriteDB() ISystemMaintenanceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISystemMaintenanceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISystemMaintenanceDo
	Not(conds ...gen.Condition) ISystemMaintenanceDo
	Or(conds ...gen.Condition) ISystemMaintenanceDo
	Select(conds ...field.Expr) ISystemMaintenanceDo
	Where(conds ...gen.Condition) ISystemMaintenanceDo
	Order(conds ...field.Expr) ISystemMaintenanceDo
	Distinct(cols ...field.Expr) ISystemMaintenanceDo
	Omit(cols ...field.Expr) ISystemMaintenanceDo
	Join(table schema.Tabler, on ...field.Expr) ISystemMaintenanceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISystemMaintenanceDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISystemMaintenanceDo
	Group(cols ...field.Expr) ISystemMaintenanceDo
	Having(conds ...gen.Condition) ISystemMaintenanceDo
	Limit(limit int) ISystemMaintenanceDo
	Offset(offset int) ISystemMaintenanceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemMaintenanceDo
	Unscoped() ISystemMaintenanceDo
	Create(values ...*model.SystemMaintenance) error
	CreateInBatches(values []*model.SystemMaintenance, batchSize int) error
	Save(values ...*model.SystemMaintenance) error
	First() (*model.SystemMaintenance, error)
	Take() (*model.SystemMaintenance, error)
	Last() (*model.SystemMaintenance, error)
	Find() ([]*model.SystemMaintenance, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemMaintenance, err error)
	FindInBatches(result *[]*model.SystemMaintenance, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SystemMaintenance) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISystemMaintenanceDo
	Assign(attrs ...field.AssignExpr) ISystemMaintenanceDo
	Joins(fields ...field.RelationField) ISystemMaintenanceDo
	Preload(fields ...field.RelationField) ISystemMaintenanceDo
	FirstOrInit() (*model.SystemMaintenance, error)
	FirstOrCreate() (*model.SystemMaintenance, error)
	FindByPage(offset int, limit int) (result []*model.SystemMaintenance, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISystemMaintenanceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s systemMaintenanceDo) Debug() ISystemMaintenanceDo {
	return s.withDO(s.DO.Debug())
}

func (s systemMaintenanceDo) WithContext(ctx context.Context) ISystemMaintenanceDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s systemMaintenanceDo) ReadDB() ISystemMaintenanceDo {
	return s.Clauses(dbresolver.Read)
}

func (s systemMaintenanceDo) WriteDB() ISystemMaintenanceDo {
	return s.Clauses(dbresolver.Write)
}

func (s systemMaintenanceDo) Session(config *gorm.Session) ISystemMaintenanceDo {
	return s.withDO(s.DO.Session(config))
}

func (s systemMaintenanceDo) Clauses(conds ...clause.Expression) ISystemMaintenanceDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s systemMaintenanceDo) Returning(value interface{}, columns ...string) ISystemMaintenanceDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s systemMaintenanceDo) Not(conds ...gen.Condition) ISystemMaintenanceDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s systemMaintenanceDo) Or(conds ...gen.Condition) ISystemMaintenanceDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s systemMaintenanceDo) Select(conds ...field.Expr) ISystemMaintenanceDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s systemMaintenanceDo) Where(conds ...gen.Condition) ISystemMaintenanceDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s systemMaintenanceDo) Order(conds ...field.Expr) ISystemMaintenanceDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s systemMaintenanceDo) Distinct(cols ...field.Expr) ISystemMaintenanceDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s systemMaintenanceDo) Omit(cols ...field.Expr) ISystemMaintenanceDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s systemMaintenanceDo) Join(table schema.Tabler, on ...field.Expr) ISystemMaintenanceDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s systemMaintenanceDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISystemMaintenanceDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s systemMaintenanceDo) RightJoin(table schema.Tabler, on ...field.Expr) ISystemMaintenanceDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s systemMaintenanceDo) Group(cols ...field.Expr) ISystemMaintenanceDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s systemMaintenanceDo) Having(conds ...gen.Condition) ISystemMaintenanceDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s systemMaintenanceDo) Limit(limit int) ISystemMaintenanceDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s systemMaintenanceDo) Offset(offset int) ISystemMaintenanceDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s systemMaintenanceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemMaintenanceDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s systemMaintenanceDo) Unscoped() ISystemMaintenanceDo {
	return s.withDO(s.DO.Unscoped())
}

func (s systemMaintenanceDo) Create(values ...*model.SystemMaintenance) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s systemMaintenanceDo) CreateInBatches(values []*model.SystemMaintenance, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s systemMaintenanceDo) Save(values ...*model.SystemMaintenance) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s systemMaintenanceDo) First() (*model.SystemMaintenance, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemMaintenance), nil
	}
}

func (s systemMaintenanceDo) Take() (*model.SystemMaintenance, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemMaintenance), nil
	}
}

func (s systemMaintenanceDo) Last() (*model.SystemMaintenance, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemMaintenance), nil
	}
}

func (s systemMaintenanceDo) Find() ([]*model.SystemMaintenance, error) {
	result, err := s.DO.Find()
	return result.([]*model.SystemMaintenance), err
}

func (s systemMaintenanceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemMaintenance, err error) {
	buf := make([]*model.SystemMaintenance, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s systemMaintenanceDo) FindInBatches(result *[]*model.SystemMaintenance, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s systemMaintenanceDo) Attrs(attrs ...field.AssignExpr) ISystemMaintenanceDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s systemMaintenanceDo) Assign(attrs ...field.AssignExpr) ISystemMaintenanceDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s systemMaintenanceDo) Joins(fields ...field.RelationField) ISystemMaintenanceDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s systemMaintenanceDo) Preload(fields ...field.RelationField) ISystemMaintenanceDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s systemMaintenanceDo) FirstOrInit() (*model.SystemMaintenance, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemMaintenance), nil
	}
}

func (s systemMaintenanceDo) FirstOrCreate() (*model.SystemMaintenance, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemMaintenance), nil
	}
}

func (s systemMaintenanceDo) FindByPage(offset int, limit int) (result []*model.SystemMaintenance, count int64, err error) {
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

func (s systemMaintenanceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s systemMaintenanceDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s systemMaintenanceDo) Delete(models ...*model.SystemMaintenance) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *systemMaintenanceDo) withDO(do gen.Dao) *systemMaintenanceDo {
	s.DO = *do.(*gen.DO)
	return s
}
