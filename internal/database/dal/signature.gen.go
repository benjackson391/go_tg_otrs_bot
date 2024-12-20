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

func newSignature(db *gorm.DB, opts ...gen.DOOption) signature {
	_signature := signature{}

	_signature.signatureDo.UseDB(db, opts...)
	_signature.signatureDo.UseModel(&model.Signature{})

	tableName := _signature.signatureDo.TableName()
	_signature.ALL = field.NewAsterisk(tableName)
	_signature.ID = field.NewInt32(tableName, "id")
	_signature.Name = field.NewString(tableName, "name")
	_signature.Text = field.NewString(tableName, "text")
	_signature.ContentType = field.NewString(tableName, "content_type")
	_signature.Comments = field.NewString(tableName, "comments")
	_signature.ValidID = field.NewInt32(tableName, "valid_id")
	_signature.CreateTime = field.NewTime(tableName, "create_time")
	_signature.CreateBy = field.NewInt32(tableName, "create_by")
	_signature.ChangeTime = field.NewTime(tableName, "change_time")
	_signature.ChangeBy = field.NewInt32(tableName, "change_by")

	_signature.fillFieldMap()

	return _signature
}

type signature struct {
	signatureDo

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

func (s signature) Table(newTableName string) *signature {
	s.signatureDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s signature) As(alias string) *signature {
	s.signatureDo.DO = *(s.signatureDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *signature) updateTableName(table string) *signature {
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

func (s *signature) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *signature) fillFieldMap() {
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

func (s signature) clone(db *gorm.DB) signature {
	s.signatureDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s signature) replaceDB(db *gorm.DB) signature {
	s.signatureDo.ReplaceDB(db)
	return s
}

type signatureDo struct{ gen.DO }

type ISignatureDo interface {
	gen.SubQuery
	Debug() ISignatureDo
	WithContext(ctx context.Context) ISignatureDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISignatureDo
	WriteDB() ISignatureDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISignatureDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISignatureDo
	Not(conds ...gen.Condition) ISignatureDo
	Or(conds ...gen.Condition) ISignatureDo
	Select(conds ...field.Expr) ISignatureDo
	Where(conds ...gen.Condition) ISignatureDo
	Order(conds ...field.Expr) ISignatureDo
	Distinct(cols ...field.Expr) ISignatureDo
	Omit(cols ...field.Expr) ISignatureDo
	Join(table schema.Tabler, on ...field.Expr) ISignatureDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISignatureDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISignatureDo
	Group(cols ...field.Expr) ISignatureDo
	Having(conds ...gen.Condition) ISignatureDo
	Limit(limit int) ISignatureDo
	Offset(offset int) ISignatureDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISignatureDo
	Unscoped() ISignatureDo
	Create(values ...*model.Signature) error
	CreateInBatches(values []*model.Signature, batchSize int) error
	Save(values ...*model.Signature) error
	First() (*model.Signature, error)
	Take() (*model.Signature, error)
	Last() (*model.Signature, error)
	Find() ([]*model.Signature, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Signature, err error)
	FindInBatches(result *[]*model.Signature, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Signature) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISignatureDo
	Assign(attrs ...field.AssignExpr) ISignatureDo
	Joins(fields ...field.RelationField) ISignatureDo
	Preload(fields ...field.RelationField) ISignatureDo
	FirstOrInit() (*model.Signature, error)
	FirstOrCreate() (*model.Signature, error)
	FindByPage(offset int, limit int) (result []*model.Signature, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISignatureDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s signatureDo) Debug() ISignatureDo {
	return s.withDO(s.DO.Debug())
}

func (s signatureDo) WithContext(ctx context.Context) ISignatureDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s signatureDo) ReadDB() ISignatureDo {
	return s.Clauses(dbresolver.Read)
}

func (s signatureDo) WriteDB() ISignatureDo {
	return s.Clauses(dbresolver.Write)
}

func (s signatureDo) Session(config *gorm.Session) ISignatureDo {
	return s.withDO(s.DO.Session(config))
}

func (s signatureDo) Clauses(conds ...clause.Expression) ISignatureDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s signatureDo) Returning(value interface{}, columns ...string) ISignatureDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s signatureDo) Not(conds ...gen.Condition) ISignatureDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s signatureDo) Or(conds ...gen.Condition) ISignatureDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s signatureDo) Select(conds ...field.Expr) ISignatureDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s signatureDo) Where(conds ...gen.Condition) ISignatureDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s signatureDo) Order(conds ...field.Expr) ISignatureDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s signatureDo) Distinct(cols ...field.Expr) ISignatureDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s signatureDo) Omit(cols ...field.Expr) ISignatureDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s signatureDo) Join(table schema.Tabler, on ...field.Expr) ISignatureDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s signatureDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISignatureDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s signatureDo) RightJoin(table schema.Tabler, on ...field.Expr) ISignatureDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s signatureDo) Group(cols ...field.Expr) ISignatureDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s signatureDo) Having(conds ...gen.Condition) ISignatureDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s signatureDo) Limit(limit int) ISignatureDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s signatureDo) Offset(offset int) ISignatureDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s signatureDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISignatureDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s signatureDo) Unscoped() ISignatureDo {
	return s.withDO(s.DO.Unscoped())
}

func (s signatureDo) Create(values ...*model.Signature) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s signatureDo) CreateInBatches(values []*model.Signature, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s signatureDo) Save(values ...*model.Signature) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s signatureDo) First() (*model.Signature, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Signature), nil
	}
}

func (s signatureDo) Take() (*model.Signature, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Signature), nil
	}
}

func (s signatureDo) Last() (*model.Signature, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Signature), nil
	}
}

func (s signatureDo) Find() ([]*model.Signature, error) {
	result, err := s.DO.Find()
	return result.([]*model.Signature), err
}

func (s signatureDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Signature, err error) {
	buf := make([]*model.Signature, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s signatureDo) FindInBatches(result *[]*model.Signature, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s signatureDo) Attrs(attrs ...field.AssignExpr) ISignatureDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s signatureDo) Assign(attrs ...field.AssignExpr) ISignatureDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s signatureDo) Joins(fields ...field.RelationField) ISignatureDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s signatureDo) Preload(fields ...field.RelationField) ISignatureDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s signatureDo) FirstOrInit() (*model.Signature, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Signature), nil
	}
}

func (s signatureDo) FirstOrCreate() (*model.Signature, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Signature), nil
	}
}

func (s signatureDo) FindByPage(offset int, limit int) (result []*model.Signature, count int64, err error) {
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

func (s signatureDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s signatureDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s signatureDo) Delete(models ...*model.Signature) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *signatureDo) withDO(do gen.Dao) *signatureDo {
	s.DO = *do.(*gen.DO)
	return s
}
