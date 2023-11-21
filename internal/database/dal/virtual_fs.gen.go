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

func newVirtualF(db *gorm.DB, opts ...gen.DOOption) virtualF {
	_virtualF := virtualF{}

	_virtualF.virtualFDo.UseDB(db, opts...)
	_virtualF.virtualFDo.UseModel(&model.VirtualF{})

	tableName := _virtualF.virtualFDo.TableName()
	_virtualF.ALL = field.NewAsterisk(tableName)
	_virtualF.ID = field.NewInt64(tableName, "id")
	_virtualF.Filename = field.NewString(tableName, "filename")
	_virtualF.Backend = field.NewString(tableName, "backend")
	_virtualF.BackendKey = field.NewString(tableName, "backend_key")
	_virtualF.CreateTime = field.NewTime(tableName, "create_time")

	_virtualF.fillFieldMap()

	return _virtualF
}

type virtualF struct {
	virtualFDo

	ALL        field.Asterisk
	ID         field.Int64
	Filename   field.String
	Backend    field.String
	BackendKey field.String
	CreateTime field.Time

	fieldMap map[string]field.Expr
}

func (v virtualF) Table(newTableName string) *virtualF {
	v.virtualFDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v virtualF) As(alias string) *virtualF {
	v.virtualFDo.DO = *(v.virtualFDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *virtualF) updateTableName(table string) *virtualF {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewInt64(table, "id")
	v.Filename = field.NewString(table, "filename")
	v.Backend = field.NewString(table, "backend")
	v.BackendKey = field.NewString(table, "backend_key")
	v.CreateTime = field.NewTime(table, "create_time")

	v.fillFieldMap()

	return v
}

func (v *virtualF) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *virtualF) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 5)
	v.fieldMap["id"] = v.ID
	v.fieldMap["filename"] = v.Filename
	v.fieldMap["backend"] = v.Backend
	v.fieldMap["backend_key"] = v.BackendKey
	v.fieldMap["create_time"] = v.CreateTime
}

func (v virtualF) clone(db *gorm.DB) virtualF {
	v.virtualFDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v virtualF) replaceDB(db *gorm.DB) virtualF {
	v.virtualFDo.ReplaceDB(db)
	return v
}

type virtualFDo struct{ gen.DO }

type IVirtualFDo interface {
	gen.SubQuery
	Debug() IVirtualFDo
	WithContext(ctx context.Context) IVirtualFDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVirtualFDo
	WriteDB() IVirtualFDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVirtualFDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVirtualFDo
	Not(conds ...gen.Condition) IVirtualFDo
	Or(conds ...gen.Condition) IVirtualFDo
	Select(conds ...field.Expr) IVirtualFDo
	Where(conds ...gen.Condition) IVirtualFDo
	Order(conds ...field.Expr) IVirtualFDo
	Distinct(cols ...field.Expr) IVirtualFDo
	Omit(cols ...field.Expr) IVirtualFDo
	Join(table schema.Tabler, on ...field.Expr) IVirtualFDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVirtualFDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVirtualFDo
	Group(cols ...field.Expr) IVirtualFDo
	Having(conds ...gen.Condition) IVirtualFDo
	Limit(limit int) IVirtualFDo
	Offset(offset int) IVirtualFDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVirtualFDo
	Unscoped() IVirtualFDo
	Create(values ...*model.VirtualF) error
	CreateInBatches(values []*model.VirtualF, batchSize int) error
	Save(values ...*model.VirtualF) error
	First() (*model.VirtualF, error)
	Take() (*model.VirtualF, error)
	Last() (*model.VirtualF, error)
	Find() ([]*model.VirtualF, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VirtualF, err error)
	FindInBatches(result *[]*model.VirtualF, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.VirtualF) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVirtualFDo
	Assign(attrs ...field.AssignExpr) IVirtualFDo
	Joins(fields ...field.RelationField) IVirtualFDo
	Preload(fields ...field.RelationField) IVirtualFDo
	FirstOrInit() (*model.VirtualF, error)
	FirstOrCreate() (*model.VirtualF, error)
	FindByPage(offset int, limit int) (result []*model.VirtualF, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVirtualFDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v virtualFDo) Debug() IVirtualFDo {
	return v.withDO(v.DO.Debug())
}

func (v virtualFDo) WithContext(ctx context.Context) IVirtualFDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v virtualFDo) ReadDB() IVirtualFDo {
	return v.Clauses(dbresolver.Read)
}

func (v virtualFDo) WriteDB() IVirtualFDo {
	return v.Clauses(dbresolver.Write)
}

func (v virtualFDo) Session(config *gorm.Session) IVirtualFDo {
	return v.withDO(v.DO.Session(config))
}

func (v virtualFDo) Clauses(conds ...clause.Expression) IVirtualFDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v virtualFDo) Returning(value interface{}, columns ...string) IVirtualFDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v virtualFDo) Not(conds ...gen.Condition) IVirtualFDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v virtualFDo) Or(conds ...gen.Condition) IVirtualFDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v virtualFDo) Select(conds ...field.Expr) IVirtualFDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v virtualFDo) Where(conds ...gen.Condition) IVirtualFDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v virtualFDo) Order(conds ...field.Expr) IVirtualFDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v virtualFDo) Distinct(cols ...field.Expr) IVirtualFDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v virtualFDo) Omit(cols ...field.Expr) IVirtualFDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v virtualFDo) Join(table schema.Tabler, on ...field.Expr) IVirtualFDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v virtualFDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVirtualFDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v virtualFDo) RightJoin(table schema.Tabler, on ...field.Expr) IVirtualFDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v virtualFDo) Group(cols ...field.Expr) IVirtualFDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v virtualFDo) Having(conds ...gen.Condition) IVirtualFDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v virtualFDo) Limit(limit int) IVirtualFDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v virtualFDo) Offset(offset int) IVirtualFDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v virtualFDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVirtualFDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v virtualFDo) Unscoped() IVirtualFDo {
	return v.withDO(v.DO.Unscoped())
}

func (v virtualFDo) Create(values ...*model.VirtualF) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v virtualFDo) CreateInBatches(values []*model.VirtualF, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v virtualFDo) Save(values ...*model.VirtualF) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v virtualFDo) First() (*model.VirtualF, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VirtualF), nil
	}
}

func (v virtualFDo) Take() (*model.VirtualF, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VirtualF), nil
	}
}

func (v virtualFDo) Last() (*model.VirtualF, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VirtualF), nil
	}
}

func (v virtualFDo) Find() ([]*model.VirtualF, error) {
	result, err := v.DO.Find()
	return result.([]*model.VirtualF), err
}

func (v virtualFDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VirtualF, err error) {
	buf := make([]*model.VirtualF, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v virtualFDo) FindInBatches(result *[]*model.VirtualF, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v virtualFDo) Attrs(attrs ...field.AssignExpr) IVirtualFDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v virtualFDo) Assign(attrs ...field.AssignExpr) IVirtualFDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v virtualFDo) Joins(fields ...field.RelationField) IVirtualFDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v virtualFDo) Preload(fields ...field.RelationField) IVirtualFDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v virtualFDo) FirstOrInit() (*model.VirtualF, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VirtualF), nil
	}
}

func (v virtualFDo) FirstOrCreate() (*model.VirtualF, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VirtualF), nil
	}
}

func (v virtualFDo) FindByPage(offset int, limit int) (result []*model.VirtualF, count int64, err error) {
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

func (v virtualFDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v virtualFDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v virtualFDo) Delete(models ...*model.VirtualF) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *virtualFDo) withDO(do gen.Dao) *virtualFDo {
	v.DO = *do.(*gen.DO)
	return v
}