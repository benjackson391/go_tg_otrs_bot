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

func newSearchProfile(db *gorm.DB, opts ...gen.DOOption) searchProfile {
	_searchProfile := searchProfile{}

	_searchProfile.searchProfileDo.UseDB(db, opts...)
	_searchProfile.searchProfileDo.UseModel(&model.SearchProfile{})

	tableName := _searchProfile.searchProfileDo.TableName()
	_searchProfile.ALL = field.NewAsterisk(tableName)
	_searchProfile.Login = field.NewString(tableName, "login")
	_searchProfile.ProfileName = field.NewString(tableName, "profile_name")
	_searchProfile.ProfileType = field.NewString(tableName, "profile_type")
	_searchProfile.ProfileKey = field.NewString(tableName, "profile_key")
	_searchProfile.ProfileValue = field.NewString(tableName, "profile_value")

	_searchProfile.fillFieldMap()

	return _searchProfile
}

type searchProfile struct {
	searchProfileDo

	ALL          field.Asterisk
	Login        field.String
	ProfileName  field.String
	ProfileType  field.String
	ProfileKey   field.String
	ProfileValue field.String

	fieldMap map[string]field.Expr
}

func (s searchProfile) Table(newTableName string) *searchProfile {
	s.searchProfileDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s searchProfile) As(alias string) *searchProfile {
	s.searchProfileDo.DO = *(s.searchProfileDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *searchProfile) updateTableName(table string) *searchProfile {
	s.ALL = field.NewAsterisk(table)
	s.Login = field.NewString(table, "login")
	s.ProfileName = field.NewString(table, "profile_name")
	s.ProfileType = field.NewString(table, "profile_type")
	s.ProfileKey = field.NewString(table, "profile_key")
	s.ProfileValue = field.NewString(table, "profile_value")

	s.fillFieldMap()

	return s
}

func (s *searchProfile) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *searchProfile) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["login"] = s.Login
	s.fieldMap["profile_name"] = s.ProfileName
	s.fieldMap["profile_type"] = s.ProfileType
	s.fieldMap["profile_key"] = s.ProfileKey
	s.fieldMap["profile_value"] = s.ProfileValue
}

func (s searchProfile) clone(db *gorm.DB) searchProfile {
	s.searchProfileDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s searchProfile) replaceDB(db *gorm.DB) searchProfile {
	s.searchProfileDo.ReplaceDB(db)
	return s
}

type searchProfileDo struct{ gen.DO }

type ISearchProfileDo interface {
	gen.SubQuery
	Debug() ISearchProfileDo
	WithContext(ctx context.Context) ISearchProfileDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISearchProfileDo
	WriteDB() ISearchProfileDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISearchProfileDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISearchProfileDo
	Not(conds ...gen.Condition) ISearchProfileDo
	Or(conds ...gen.Condition) ISearchProfileDo
	Select(conds ...field.Expr) ISearchProfileDo
	Where(conds ...gen.Condition) ISearchProfileDo
	Order(conds ...field.Expr) ISearchProfileDo
	Distinct(cols ...field.Expr) ISearchProfileDo
	Omit(cols ...field.Expr) ISearchProfileDo
	Join(table schema.Tabler, on ...field.Expr) ISearchProfileDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISearchProfileDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISearchProfileDo
	Group(cols ...field.Expr) ISearchProfileDo
	Having(conds ...gen.Condition) ISearchProfileDo
	Limit(limit int) ISearchProfileDo
	Offset(offset int) ISearchProfileDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISearchProfileDo
	Unscoped() ISearchProfileDo
	Create(values ...*model.SearchProfile) error
	CreateInBatches(values []*model.SearchProfile, batchSize int) error
	Save(values ...*model.SearchProfile) error
	First() (*model.SearchProfile, error)
	Take() (*model.SearchProfile, error)
	Last() (*model.SearchProfile, error)
	Find() ([]*model.SearchProfile, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SearchProfile, err error)
	FindInBatches(result *[]*model.SearchProfile, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SearchProfile) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISearchProfileDo
	Assign(attrs ...field.AssignExpr) ISearchProfileDo
	Joins(fields ...field.RelationField) ISearchProfileDo
	Preload(fields ...field.RelationField) ISearchProfileDo
	FirstOrInit() (*model.SearchProfile, error)
	FirstOrCreate() (*model.SearchProfile, error)
	FindByPage(offset int, limit int) (result []*model.SearchProfile, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISearchProfileDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s searchProfileDo) Debug() ISearchProfileDo {
	return s.withDO(s.DO.Debug())
}

func (s searchProfileDo) WithContext(ctx context.Context) ISearchProfileDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s searchProfileDo) ReadDB() ISearchProfileDo {
	return s.Clauses(dbresolver.Read)
}

func (s searchProfileDo) WriteDB() ISearchProfileDo {
	return s.Clauses(dbresolver.Write)
}

func (s searchProfileDo) Session(config *gorm.Session) ISearchProfileDo {
	return s.withDO(s.DO.Session(config))
}

func (s searchProfileDo) Clauses(conds ...clause.Expression) ISearchProfileDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s searchProfileDo) Returning(value interface{}, columns ...string) ISearchProfileDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s searchProfileDo) Not(conds ...gen.Condition) ISearchProfileDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s searchProfileDo) Or(conds ...gen.Condition) ISearchProfileDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s searchProfileDo) Select(conds ...field.Expr) ISearchProfileDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s searchProfileDo) Where(conds ...gen.Condition) ISearchProfileDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s searchProfileDo) Order(conds ...field.Expr) ISearchProfileDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s searchProfileDo) Distinct(cols ...field.Expr) ISearchProfileDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s searchProfileDo) Omit(cols ...field.Expr) ISearchProfileDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s searchProfileDo) Join(table schema.Tabler, on ...field.Expr) ISearchProfileDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s searchProfileDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISearchProfileDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s searchProfileDo) RightJoin(table schema.Tabler, on ...field.Expr) ISearchProfileDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s searchProfileDo) Group(cols ...field.Expr) ISearchProfileDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s searchProfileDo) Having(conds ...gen.Condition) ISearchProfileDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s searchProfileDo) Limit(limit int) ISearchProfileDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s searchProfileDo) Offset(offset int) ISearchProfileDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s searchProfileDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISearchProfileDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s searchProfileDo) Unscoped() ISearchProfileDo {
	return s.withDO(s.DO.Unscoped())
}

func (s searchProfileDo) Create(values ...*model.SearchProfile) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s searchProfileDo) CreateInBatches(values []*model.SearchProfile, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s searchProfileDo) Save(values ...*model.SearchProfile) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s searchProfileDo) First() (*model.SearchProfile, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchProfile), nil
	}
}

func (s searchProfileDo) Take() (*model.SearchProfile, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchProfile), nil
	}
}

func (s searchProfileDo) Last() (*model.SearchProfile, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchProfile), nil
	}
}

func (s searchProfileDo) Find() ([]*model.SearchProfile, error) {
	result, err := s.DO.Find()
	return result.([]*model.SearchProfile), err
}

func (s searchProfileDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SearchProfile, err error) {
	buf := make([]*model.SearchProfile, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s searchProfileDo) FindInBatches(result *[]*model.SearchProfile, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s searchProfileDo) Attrs(attrs ...field.AssignExpr) ISearchProfileDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s searchProfileDo) Assign(attrs ...field.AssignExpr) ISearchProfileDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s searchProfileDo) Joins(fields ...field.RelationField) ISearchProfileDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s searchProfileDo) Preload(fields ...field.RelationField) ISearchProfileDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s searchProfileDo) FirstOrInit() (*model.SearchProfile, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchProfile), nil
	}
}

func (s searchProfileDo) FirstOrCreate() (*model.SearchProfile, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchProfile), nil
	}
}

func (s searchProfileDo) FindByPage(offset int, limit int) (result []*model.SearchProfile, count int64, err error) {
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

func (s searchProfileDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s searchProfileDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s searchProfileDo) Delete(models ...*model.SearchProfile) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *searchProfileDo) withDO(do gen.Dao) *searchProfileDo {
	s.DO = *do.(*gen.DO)
	return s
}