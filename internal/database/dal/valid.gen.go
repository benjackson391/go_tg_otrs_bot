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

func newValid(db *gorm.DB, opts ...gen.DOOption) valid {
	_valid := valid{}

	_valid.validDo.UseDB(db, opts...)
	_valid.validDo.UseModel(&model.Valid{})

	tableName := _valid.validDo.TableName()
	_valid.ALL = field.NewAsterisk(tableName)
	_valid.ID = field.NewInt32(tableName, "id")
	_valid.Name = field.NewString(tableName, "name")
	_valid.CreateTime = field.NewTime(tableName, "create_time")
	_valid.CreateBy = field.NewInt32(tableName, "create_by")
	_valid.ChangeTime = field.NewTime(tableName, "change_time")
	_valid.ChangeBy = field.NewInt32(tableName, "change_by")

	_valid.fillFieldMap()

	return _valid
}

type valid struct {
	validDo

	ALL        field.Asterisk
	ID         field.Int32
	Name       field.String
	CreateTime field.Time
	CreateBy   field.Int32
	ChangeTime field.Time
	ChangeBy   field.Int32

	fieldMap map[string]field.Expr
}

func (v valid) Table(newTableName string) *valid {
	v.validDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v valid) As(alias string) *valid {
	v.validDo.DO = *(v.validDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *valid) updateTableName(table string) *valid {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewInt32(table, "id")
	v.Name = field.NewString(table, "name")
	v.CreateTime = field.NewTime(table, "create_time")
	v.CreateBy = field.NewInt32(table, "create_by")
	v.ChangeTime = field.NewTime(table, "change_time")
	v.ChangeBy = field.NewInt32(table, "change_by")

	v.fillFieldMap()

	return v
}

func (v *valid) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *valid) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 6)
	v.fieldMap["id"] = v.ID
	v.fieldMap["name"] = v.Name
	v.fieldMap["create_time"] = v.CreateTime
	v.fieldMap["create_by"] = v.CreateBy
	v.fieldMap["change_time"] = v.ChangeTime
	v.fieldMap["change_by"] = v.ChangeBy
}

func (v valid) clone(db *gorm.DB) valid {
	v.validDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v valid) replaceDB(db *gorm.DB) valid {
	v.validDo.ReplaceDB(db)
	return v
}

type validDo struct{ gen.DO }

type IValidDo interface {
	gen.SubQuery
	Debug() IValidDo
	WithContext(ctx context.Context) IValidDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IValidDo
	WriteDB() IValidDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IValidDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IValidDo
	Not(conds ...gen.Condition) IValidDo
	Or(conds ...gen.Condition) IValidDo
	Select(conds ...field.Expr) IValidDo
	Where(conds ...gen.Condition) IValidDo
	Order(conds ...field.Expr) IValidDo
	Distinct(cols ...field.Expr) IValidDo
	Omit(cols ...field.Expr) IValidDo
	Join(table schema.Tabler, on ...field.Expr) IValidDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IValidDo
	RightJoin(table schema.Tabler, on ...field.Expr) IValidDo
	Group(cols ...field.Expr) IValidDo
	Having(conds ...gen.Condition) IValidDo
	Limit(limit int) IValidDo
	Offset(offset int) IValidDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IValidDo
	Unscoped() IValidDo
	Create(values ...*model.Valid) error
	CreateInBatches(values []*model.Valid, batchSize int) error
	Save(values ...*model.Valid) error
	First() (*model.Valid, error)
	Take() (*model.Valid, error)
	Last() (*model.Valid, error)
	Find() ([]*model.Valid, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Valid, err error)
	FindInBatches(result *[]*model.Valid, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Valid) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IValidDo
	Assign(attrs ...field.AssignExpr) IValidDo
	Joins(fields ...field.RelationField) IValidDo
	Preload(fields ...field.RelationField) IValidDo
	FirstOrInit() (*model.Valid, error)
	FirstOrCreate() (*model.Valid, error)
	FindByPage(offset int, limit int) (result []*model.Valid, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IValidDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v validDo) Debug() IValidDo {
	return v.withDO(v.DO.Debug())
}

func (v validDo) WithContext(ctx context.Context) IValidDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v validDo) ReadDB() IValidDo {
	return v.Clauses(dbresolver.Read)
}

func (v validDo) WriteDB() IValidDo {
	return v.Clauses(dbresolver.Write)
}

func (v validDo) Session(config *gorm.Session) IValidDo {
	return v.withDO(v.DO.Session(config))
}

func (v validDo) Clauses(conds ...clause.Expression) IValidDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v validDo) Returning(value interface{}, columns ...string) IValidDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v validDo) Not(conds ...gen.Condition) IValidDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v validDo) Or(conds ...gen.Condition) IValidDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v validDo) Select(conds ...field.Expr) IValidDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v validDo) Where(conds ...gen.Condition) IValidDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v validDo) Order(conds ...field.Expr) IValidDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v validDo) Distinct(cols ...field.Expr) IValidDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v validDo) Omit(cols ...field.Expr) IValidDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v validDo) Join(table schema.Tabler, on ...field.Expr) IValidDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v validDo) LeftJoin(table schema.Tabler, on ...field.Expr) IValidDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v validDo) RightJoin(table schema.Tabler, on ...field.Expr) IValidDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v validDo) Group(cols ...field.Expr) IValidDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v validDo) Having(conds ...gen.Condition) IValidDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v validDo) Limit(limit int) IValidDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v validDo) Offset(offset int) IValidDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v validDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IValidDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v validDo) Unscoped() IValidDo {
	return v.withDO(v.DO.Unscoped())
}

func (v validDo) Create(values ...*model.Valid) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v validDo) CreateInBatches(values []*model.Valid, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v validDo) Save(values ...*model.Valid) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v validDo) First() (*model.Valid, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Valid), nil
	}
}

func (v validDo) Take() (*model.Valid, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Valid), nil
	}
}

func (v validDo) Last() (*model.Valid, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Valid), nil
	}
}

func (v validDo) Find() ([]*model.Valid, error) {
	result, err := v.DO.Find()
	return result.([]*model.Valid), err
}

func (v validDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Valid, err error) {
	buf := make([]*model.Valid, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v validDo) FindInBatches(result *[]*model.Valid, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v validDo) Attrs(attrs ...field.AssignExpr) IValidDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v validDo) Assign(attrs ...field.AssignExpr) IValidDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v validDo) Joins(fields ...field.RelationField) IValidDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v validDo) Preload(fields ...field.RelationField) IValidDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v validDo) FirstOrInit() (*model.Valid, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Valid), nil
	}
}

func (v validDo) FirstOrCreate() (*model.Valid, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Valid), nil
	}
}

func (v validDo) FindByPage(offset int, limit int) (result []*model.Valid, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v validDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v validDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v validDo) Delete(models ...*model.Valid) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *validDo) withDO(do gen.Dao) *validDo {
	v.DO = *do.(*gen.DO)
	return v
}
