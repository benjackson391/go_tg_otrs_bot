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

func newServicePreference(db *gorm.DB, opts ...gen.DOOption) servicePreference {
	_servicePreference := servicePreference{}

	_servicePreference.servicePreferenceDo.UseDB(db, opts...)
	_servicePreference.servicePreferenceDo.UseModel(&model.ServicePreference{})

	tableName := _servicePreference.servicePreferenceDo.TableName()
	_servicePreference.ALL = field.NewAsterisk(tableName)
	_servicePreference.ServiceID = field.NewInt32(tableName, "service_id")
	_servicePreference.PreferencesKey = field.NewString(tableName, "preferences_key")
	_servicePreference.PreferencesValue = field.NewString(tableName, "preferences_value")

	_servicePreference.fillFieldMap()

	return _servicePreference
}

type servicePreference struct {
	servicePreferenceDo

	ALL              field.Asterisk
	ServiceID        field.Int32
	PreferencesKey   field.String
	PreferencesValue field.String

	fieldMap map[string]field.Expr
}

func (s servicePreference) Table(newTableName string) *servicePreference {
	s.servicePreferenceDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s servicePreference) As(alias string) *servicePreference {
	s.servicePreferenceDo.DO = *(s.servicePreferenceDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *servicePreference) updateTableName(table string) *servicePreference {
	s.ALL = field.NewAsterisk(table)
	s.ServiceID = field.NewInt32(table, "service_id")
	s.PreferencesKey = field.NewString(table, "preferences_key")
	s.PreferencesValue = field.NewString(table, "preferences_value")

	s.fillFieldMap()

	return s
}

func (s *servicePreference) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *servicePreference) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["service_id"] = s.ServiceID
	s.fieldMap["preferences_key"] = s.PreferencesKey
	s.fieldMap["preferences_value"] = s.PreferencesValue
}

func (s servicePreference) clone(db *gorm.DB) servicePreference {
	s.servicePreferenceDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s servicePreference) replaceDB(db *gorm.DB) servicePreference {
	s.servicePreferenceDo.ReplaceDB(db)
	return s
}

type servicePreferenceDo struct{ gen.DO }

type IServicePreferenceDo interface {
	gen.SubQuery
	Debug() IServicePreferenceDo
	WithContext(ctx context.Context) IServicePreferenceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IServicePreferenceDo
	WriteDB() IServicePreferenceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IServicePreferenceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IServicePreferenceDo
	Not(conds ...gen.Condition) IServicePreferenceDo
	Or(conds ...gen.Condition) IServicePreferenceDo
	Select(conds ...field.Expr) IServicePreferenceDo
	Where(conds ...gen.Condition) IServicePreferenceDo
	Order(conds ...field.Expr) IServicePreferenceDo
	Distinct(cols ...field.Expr) IServicePreferenceDo
	Omit(cols ...field.Expr) IServicePreferenceDo
	Join(table schema.Tabler, on ...field.Expr) IServicePreferenceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IServicePreferenceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IServicePreferenceDo
	Group(cols ...field.Expr) IServicePreferenceDo
	Having(conds ...gen.Condition) IServicePreferenceDo
	Limit(limit int) IServicePreferenceDo
	Offset(offset int) IServicePreferenceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IServicePreferenceDo
	Unscoped() IServicePreferenceDo
	Create(values ...*model.ServicePreference) error
	CreateInBatches(values []*model.ServicePreference, batchSize int) error
	Save(values ...*model.ServicePreference) error
	First() (*model.ServicePreference, error)
	Take() (*model.ServicePreference, error)
	Last() (*model.ServicePreference, error)
	Find() ([]*model.ServicePreference, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServicePreference, err error)
	FindInBatches(result *[]*model.ServicePreference, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ServicePreference) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IServicePreferenceDo
	Assign(attrs ...field.AssignExpr) IServicePreferenceDo
	Joins(fields ...field.RelationField) IServicePreferenceDo
	Preload(fields ...field.RelationField) IServicePreferenceDo
	FirstOrInit() (*model.ServicePreference, error)
	FirstOrCreate() (*model.ServicePreference, error)
	FindByPage(offset int, limit int) (result []*model.ServicePreference, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IServicePreferenceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s servicePreferenceDo) Debug() IServicePreferenceDo {
	return s.withDO(s.DO.Debug())
}

func (s servicePreferenceDo) WithContext(ctx context.Context) IServicePreferenceDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s servicePreferenceDo) ReadDB() IServicePreferenceDo {
	return s.Clauses(dbresolver.Read)
}

func (s servicePreferenceDo) WriteDB() IServicePreferenceDo {
	return s.Clauses(dbresolver.Write)
}

func (s servicePreferenceDo) Session(config *gorm.Session) IServicePreferenceDo {
	return s.withDO(s.DO.Session(config))
}

func (s servicePreferenceDo) Clauses(conds ...clause.Expression) IServicePreferenceDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s servicePreferenceDo) Returning(value interface{}, columns ...string) IServicePreferenceDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s servicePreferenceDo) Not(conds ...gen.Condition) IServicePreferenceDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s servicePreferenceDo) Or(conds ...gen.Condition) IServicePreferenceDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s servicePreferenceDo) Select(conds ...field.Expr) IServicePreferenceDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s servicePreferenceDo) Where(conds ...gen.Condition) IServicePreferenceDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s servicePreferenceDo) Order(conds ...field.Expr) IServicePreferenceDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s servicePreferenceDo) Distinct(cols ...field.Expr) IServicePreferenceDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s servicePreferenceDo) Omit(cols ...field.Expr) IServicePreferenceDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s servicePreferenceDo) Join(table schema.Tabler, on ...field.Expr) IServicePreferenceDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s servicePreferenceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IServicePreferenceDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s servicePreferenceDo) RightJoin(table schema.Tabler, on ...field.Expr) IServicePreferenceDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s servicePreferenceDo) Group(cols ...field.Expr) IServicePreferenceDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s servicePreferenceDo) Having(conds ...gen.Condition) IServicePreferenceDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s servicePreferenceDo) Limit(limit int) IServicePreferenceDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s servicePreferenceDo) Offset(offset int) IServicePreferenceDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s servicePreferenceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IServicePreferenceDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s servicePreferenceDo) Unscoped() IServicePreferenceDo {
	return s.withDO(s.DO.Unscoped())
}

func (s servicePreferenceDo) Create(values ...*model.ServicePreference) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s servicePreferenceDo) CreateInBatches(values []*model.ServicePreference, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s servicePreferenceDo) Save(values ...*model.ServicePreference) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s servicePreferenceDo) First() (*model.ServicePreference, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServicePreference), nil
	}
}

func (s servicePreferenceDo) Take() (*model.ServicePreference, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServicePreference), nil
	}
}

func (s servicePreferenceDo) Last() (*model.ServicePreference, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServicePreference), nil
	}
}

func (s servicePreferenceDo) Find() ([]*model.ServicePreference, error) {
	result, err := s.DO.Find()
	return result.([]*model.ServicePreference), err
}

func (s servicePreferenceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServicePreference, err error) {
	buf := make([]*model.ServicePreference, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s servicePreferenceDo) FindInBatches(result *[]*model.ServicePreference, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s servicePreferenceDo) Attrs(attrs ...field.AssignExpr) IServicePreferenceDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s servicePreferenceDo) Assign(attrs ...field.AssignExpr) IServicePreferenceDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s servicePreferenceDo) Joins(fields ...field.RelationField) IServicePreferenceDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s servicePreferenceDo) Preload(fields ...field.RelationField) IServicePreferenceDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s servicePreferenceDo) FirstOrInit() (*model.ServicePreference, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServicePreference), nil
	}
}

func (s servicePreferenceDo) FirstOrCreate() (*model.ServicePreference, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServicePreference), nil
	}
}

func (s servicePreferenceDo) FindByPage(offset int, limit int) (result []*model.ServicePreference, count int64, err error) {
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

func (s servicePreferenceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s servicePreferenceDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s servicePreferenceDo) Delete(models ...*model.ServicePreference) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *servicePreferenceDo) withDO(do gen.Dao) *servicePreferenceDo {
	s.DO = *do.(*gen.DO)
	return s
}