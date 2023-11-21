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

func newSalutation(db *gorm.DB, opts ...gen.DOOption) salutation {
	_salutation := salutation{}

	_salutation.salutationDo.UseDB(db, opts...)
	_salutation.salutationDo.UseModel(&model.Salutation{})

	tableName := _salutation.salutationDo.TableName()
	_salutation.ALL = field.NewAsterisk(tableName)
	_salutation.ID = field.NewInt32(tableName, "id")
	_salutation.Name = field.NewString(tableName, "name")
	_salutation.Text = field.NewString(tableName, "text")
	_salutation.ContentType = field.NewString(tableName, "content_type")
	_salutation.Comments = field.NewString(tableName, "comments")
	_salutation.ValidID = field.NewInt32(tableName, "valid_id")
	_salutation.CreateTime = field.NewTime(tableName, "create_time")
	_salutation.CreateBy = field.NewInt32(tableName, "create_by")
	_salutation.ChangeTime = field.NewTime(tableName, "change_time")
	_salutation.ChangeBy = field.NewInt32(tableName, "change_by")

	_salutation.fillFieldMap()

	return _salutation
}

type salutation struct {
	salutationDo

	ALL         field.Asterisk
	ID          field.Int32
	Name        field.String
	Text        field.String
	ContentType field.String
	Comments    field.String
	ValidID     field.Int32
	CreateTime  field.Time
	CreateBy    field.Int32
	ChangeTime  field.Time
	ChangeBy    field.Int32

	fieldMap map[string]field.Expr
}

func (s salutation) Table(newTableName string) *salutation {
	s.salutationDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s salutation) As(alias string) *salutation {
	s.salutationDo.DO = *(s.salutationDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *salutation) updateTableName(table string) *salutation {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.Name = field.NewString(table, "name")
	s.Text = field.NewString(table, "text")
	s.ContentType = field.NewString(table, "content_type")
	s.Comments = field.NewString(table, "comments")
	s.ValidID = field.NewInt32(table, "valid_id")
	s.CreateTime = field.NewTime(table, "create_time")
	s.CreateBy = field.NewInt32(table, "create_by")
	s.ChangeTime = field.NewTime(table, "change_time")
	s.ChangeBy = field.NewInt32(table, "change_by")

	s.fillFieldMap()

	return s
}

func (s *salutation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *salutation) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["text"] = s.Text
	s.fieldMap["content_type"] = s.ContentType
	s.fieldMap["comments"] = s.Comments
	s.fieldMap["valid_id"] = s.ValidID
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["change_time"] = s.ChangeTime
	s.fieldMap["change_by"] = s.ChangeBy
}

func (s salutation) clone(db *gorm.DB) salutation {
	s.salutationDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s salutation) replaceDB(db *gorm.DB) salutation {
	s.salutationDo.ReplaceDB(db)
	return s
}

type salutationDo struct{ gen.DO }

type ISalutationDo interface {
	gen.SubQuery
	Debug() ISalutationDo
	WithContext(ctx context.Context) ISalutationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISalutationDo
	WriteDB() ISalutationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISalutationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISalutationDo
	Not(conds ...gen.Condition) ISalutationDo
	Or(conds ...gen.Condition) ISalutationDo
	Select(conds ...field.Expr) ISalutationDo
	Where(conds ...gen.Condition) ISalutationDo
	Order(conds ...field.Expr) ISalutationDo
	Distinct(cols ...field.Expr) ISalutationDo
	Omit(cols ...field.Expr) ISalutationDo
	Join(table schema.Tabler, on ...field.Expr) ISalutationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISalutationDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISalutationDo
	Group(cols ...field.Expr) ISalutationDo
	Having(conds ...gen.Condition) ISalutationDo
	Limit(limit int) ISalutationDo
	Offset(offset int) ISalutationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISalutationDo
	Unscoped() ISalutationDo
	Create(values ...*model.Salutation) error
	CreateInBatches(values []*model.Salutation, batchSize int) error
	Save(values ...*model.Salutation) error
	First() (*model.Salutation, error)
	Take() (*model.Salutation, error)
	Last() (*model.Salutation, error)
	Find() ([]*model.Salutation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Salutation, err error)
	FindInBatches(result *[]*model.Salutation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Salutation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISalutationDo
	Assign(attrs ...field.AssignExpr) ISalutationDo
	Joins(fields ...field.RelationField) ISalutationDo
	Preload(fields ...field.RelationField) ISalutationDo
	FirstOrInit() (*model.Salutation, error)
	FirstOrCreate() (*model.Salutation, error)
	FindByPage(offset int, limit int) (result []*model.Salutation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISalutationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s salutationDo) Debug() ISalutationDo {
	return s.withDO(s.DO.Debug())
}

func (s salutationDo) WithContext(ctx context.Context) ISalutationDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s salutationDo) ReadDB() ISalutationDo {
	return s.Clauses(dbresolver.Read)
}

func (s salutationDo) WriteDB() ISalutationDo {
	return s.Clauses(dbresolver.Write)
}

func (s salutationDo) Session(config *gorm.Session) ISalutationDo {
	return s.withDO(s.DO.Session(config))
}

func (s salutationDo) Clauses(conds ...clause.Expression) ISalutationDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s salutationDo) Returning(value interface{}, columns ...string) ISalutationDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s salutationDo) Not(conds ...gen.Condition) ISalutationDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s salutationDo) Or(conds ...gen.Condition) ISalutationDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s salutationDo) Select(conds ...field.Expr) ISalutationDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s salutationDo) Where(conds ...gen.Condition) ISalutationDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s salutationDo) Order(conds ...field.Expr) ISalutationDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s salutationDo) Distinct(cols ...field.Expr) ISalutationDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s salutationDo) Omit(cols ...field.Expr) ISalutationDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s salutationDo) Join(table schema.Tabler, on ...field.Expr) ISalutationDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s salutationDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISalutationDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s salutationDo) RightJoin(table schema.Tabler, on ...field.Expr) ISalutationDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s salutationDo) Group(cols ...field.Expr) ISalutationDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s salutationDo) Having(conds ...gen.Condition) ISalutationDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s salutationDo) Limit(limit int) ISalutationDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s salutationDo) Offset(offset int) ISalutationDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s salutationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISalutationDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s salutationDo) Unscoped() ISalutationDo {
	return s.withDO(s.DO.Unscoped())
}

func (s salutationDo) Create(values ...*model.Salutation) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s salutationDo) CreateInBatches(values []*model.Salutation, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s salutationDo) Save(values ...*model.Salutation) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s salutationDo) First() (*model.Salutation, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salutation), nil
	}
}

func (s salutationDo) Take() (*model.Salutation, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salutation), nil
	}
}

func (s salutationDo) Last() (*model.Salutation, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salutation), nil
	}
}

func (s salutationDo) Find() ([]*model.Salutation, error) {
	result, err := s.DO.Find()
	return result.([]*model.Salutation), err
}

func (s salutationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Salutation, err error) {
	buf := make([]*model.Salutation, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s salutationDo) FindInBatches(result *[]*model.Salutation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s salutationDo) Attrs(attrs ...field.AssignExpr) ISalutationDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s salutationDo) Assign(attrs ...field.AssignExpr) ISalutationDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s salutationDo) Joins(fields ...field.RelationField) ISalutationDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s salutationDo) Preload(fields ...field.RelationField) ISalutationDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s salutationDo) FirstOrInit() (*model.Salutation, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salutation), nil
	}
}

func (s salutationDo) FirstOrCreate() (*model.Salutation, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Salutation), nil
	}
}

func (s salutationDo) FindByPage(offset int, limit int) (result []*model.Salutation, count int64, err error) {
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

func (s salutationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s salutationDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s salutationDo) Delete(models ...*model.Salutation) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *salutationDo) withDO(do gen.Dao) *salutationDo {
	s.DO = *do.(*gen.DO)
	return s
}