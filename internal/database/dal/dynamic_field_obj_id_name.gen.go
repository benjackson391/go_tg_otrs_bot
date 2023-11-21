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

func newDynamicFieldObjIDName(db *gorm.DB, opts ...gen.DOOption) dynamicFieldObjIDName {
	_dynamicFieldObjIDName := dynamicFieldObjIDName{}

	_dynamicFieldObjIDName.dynamicFieldObjIDNameDo.UseDB(db, opts...)
	_dynamicFieldObjIDName.dynamicFieldObjIDNameDo.UseModel(&model.DynamicFieldObjIDName{})

	tableName := _dynamicFieldObjIDName.dynamicFieldObjIDNameDo.TableName()
	_dynamicFieldObjIDName.ALL = field.NewAsterisk(tableName)
	_dynamicFieldObjIDName.ObjectID = field.NewInt32(tableName, "object_id")
	_dynamicFieldObjIDName.ObjectName = field.NewString(tableName, "object_name")
	_dynamicFieldObjIDName.ObjectType = field.NewString(tableName, "object_type")

	_dynamicFieldObjIDName.fillFieldMap()

	return _dynamicFieldObjIDName
}

type dynamicFieldObjIDName struct {
	dynamicFieldObjIDNameDo

	ALL        field.Asterisk
	ObjectID   field.Int32
	ObjectName field.String
	ObjectType field.String

	fieldMap map[string]field.Expr
}

func (d dynamicFieldObjIDName) Table(newTableName string) *dynamicFieldObjIDName {
	d.dynamicFieldObjIDNameDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dynamicFieldObjIDName) As(alias string) *dynamicFieldObjIDName {
	d.dynamicFieldObjIDNameDo.DO = *(d.dynamicFieldObjIDNameDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dynamicFieldObjIDName) updateTableName(table string) *dynamicFieldObjIDName {
	d.ALL = field.NewAsterisk(table)
	d.ObjectID = field.NewInt32(table, "object_id")
	d.ObjectName = field.NewString(table, "object_name")
	d.ObjectType = field.NewString(table, "object_type")

	d.fillFieldMap()

	return d
}

func (d *dynamicFieldObjIDName) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dynamicFieldObjIDName) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 3)
	d.fieldMap["object_id"] = d.ObjectID
	d.fieldMap["object_name"] = d.ObjectName
	d.fieldMap["object_type"] = d.ObjectType
}

func (d dynamicFieldObjIDName) clone(db *gorm.DB) dynamicFieldObjIDName {
	d.dynamicFieldObjIDNameDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dynamicFieldObjIDName) replaceDB(db *gorm.DB) dynamicFieldObjIDName {
	d.dynamicFieldObjIDNameDo.ReplaceDB(db)
	return d
}

type dynamicFieldObjIDNameDo struct{ gen.DO }

type IDynamicFieldObjIDNameDo interface {
	gen.SubQuery
	Debug() IDynamicFieldObjIDNameDo
	WithContext(ctx context.Context) IDynamicFieldObjIDNameDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDynamicFieldObjIDNameDo
	WriteDB() IDynamicFieldObjIDNameDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDynamicFieldObjIDNameDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDynamicFieldObjIDNameDo
	Not(conds ...gen.Condition) IDynamicFieldObjIDNameDo
	Or(conds ...gen.Condition) IDynamicFieldObjIDNameDo
	Select(conds ...field.Expr) IDynamicFieldObjIDNameDo
	Where(conds ...gen.Condition) IDynamicFieldObjIDNameDo
	Order(conds ...field.Expr) IDynamicFieldObjIDNameDo
	Distinct(cols ...field.Expr) IDynamicFieldObjIDNameDo
	Omit(cols ...field.Expr) IDynamicFieldObjIDNameDo
	Join(table schema.Tabler, on ...field.Expr) IDynamicFieldObjIDNameDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDynamicFieldObjIDNameDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDynamicFieldObjIDNameDo
	Group(cols ...field.Expr) IDynamicFieldObjIDNameDo
	Having(conds ...gen.Condition) IDynamicFieldObjIDNameDo
	Limit(limit int) IDynamicFieldObjIDNameDo
	Offset(offset int) IDynamicFieldObjIDNameDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDynamicFieldObjIDNameDo
	Unscoped() IDynamicFieldObjIDNameDo
	Create(values ...*model.DynamicFieldObjIDName) error
	CreateInBatches(values []*model.DynamicFieldObjIDName, batchSize int) error
	Save(values ...*model.DynamicFieldObjIDName) error
	First() (*model.DynamicFieldObjIDName, error)
	Take() (*model.DynamicFieldObjIDName, error)
	Last() (*model.DynamicFieldObjIDName, error)
	Find() ([]*model.DynamicFieldObjIDName, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DynamicFieldObjIDName, err error)
	FindInBatches(result *[]*model.DynamicFieldObjIDName, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DynamicFieldObjIDName) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDynamicFieldObjIDNameDo
	Assign(attrs ...field.AssignExpr) IDynamicFieldObjIDNameDo
	Joins(fields ...field.RelationField) IDynamicFieldObjIDNameDo
	Preload(fields ...field.RelationField) IDynamicFieldObjIDNameDo
	FirstOrInit() (*model.DynamicFieldObjIDName, error)
	FirstOrCreate() (*model.DynamicFieldObjIDName, error)
	FindByPage(offset int, limit int) (result []*model.DynamicFieldObjIDName, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDynamicFieldObjIDNameDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dynamicFieldObjIDNameDo) Debug() IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Debug())
}

func (d dynamicFieldObjIDNameDo) WithContext(ctx context.Context) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dynamicFieldObjIDNameDo) ReadDB() IDynamicFieldObjIDNameDo {
	return d.Clauses(dbresolver.Read)
}

func (d dynamicFieldObjIDNameDo) WriteDB() IDynamicFieldObjIDNameDo {
	return d.Clauses(dbresolver.Write)
}

func (d dynamicFieldObjIDNameDo) Session(config *gorm.Session) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Session(config))
}

func (d dynamicFieldObjIDNameDo) Clauses(conds ...clause.Expression) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dynamicFieldObjIDNameDo) Returning(value interface{}, columns ...string) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dynamicFieldObjIDNameDo) Not(conds ...gen.Condition) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dynamicFieldObjIDNameDo) Or(conds ...gen.Condition) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dynamicFieldObjIDNameDo) Select(conds ...field.Expr) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dynamicFieldObjIDNameDo) Where(conds ...gen.Condition) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dynamicFieldObjIDNameDo) Order(conds ...field.Expr) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dynamicFieldObjIDNameDo) Distinct(cols ...field.Expr) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dynamicFieldObjIDNameDo) Omit(cols ...field.Expr) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dynamicFieldObjIDNameDo) Join(table schema.Tabler, on ...field.Expr) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dynamicFieldObjIDNameDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dynamicFieldObjIDNameDo) RightJoin(table schema.Tabler, on ...field.Expr) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dynamicFieldObjIDNameDo) Group(cols ...field.Expr) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dynamicFieldObjIDNameDo) Having(conds ...gen.Condition) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dynamicFieldObjIDNameDo) Limit(limit int) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dynamicFieldObjIDNameDo) Offset(offset int) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dynamicFieldObjIDNameDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dynamicFieldObjIDNameDo) Unscoped() IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dynamicFieldObjIDNameDo) Create(values ...*model.DynamicFieldObjIDName) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dynamicFieldObjIDNameDo) CreateInBatches(values []*model.DynamicFieldObjIDName, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dynamicFieldObjIDNameDo) Save(values ...*model.DynamicFieldObjIDName) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dynamicFieldObjIDNameDo) First() (*model.DynamicFieldObjIDName, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DynamicFieldObjIDName), nil
	}
}

func (d dynamicFieldObjIDNameDo) Take() (*model.DynamicFieldObjIDName, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DynamicFieldObjIDName), nil
	}
}

func (d dynamicFieldObjIDNameDo) Last() (*model.DynamicFieldObjIDName, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DynamicFieldObjIDName), nil
	}
}

func (d dynamicFieldObjIDNameDo) Find() ([]*model.DynamicFieldObjIDName, error) {
	result, err := d.DO.Find()
	return result.([]*model.DynamicFieldObjIDName), err
}

func (d dynamicFieldObjIDNameDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DynamicFieldObjIDName, err error) {
	buf := make([]*model.DynamicFieldObjIDName, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dynamicFieldObjIDNameDo) FindInBatches(result *[]*model.DynamicFieldObjIDName, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dynamicFieldObjIDNameDo) Attrs(attrs ...field.AssignExpr) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dynamicFieldObjIDNameDo) Assign(attrs ...field.AssignExpr) IDynamicFieldObjIDNameDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dynamicFieldObjIDNameDo) Joins(fields ...field.RelationField) IDynamicFieldObjIDNameDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dynamicFieldObjIDNameDo) Preload(fields ...field.RelationField) IDynamicFieldObjIDNameDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dynamicFieldObjIDNameDo) FirstOrInit() (*model.DynamicFieldObjIDName, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DynamicFieldObjIDName), nil
	}
}

func (d dynamicFieldObjIDNameDo) FirstOrCreate() (*model.DynamicFieldObjIDName, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DynamicFieldObjIDName), nil
	}
}

func (d dynamicFieldObjIDNameDo) FindByPage(offset int, limit int) (result []*model.DynamicFieldObjIDName, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dynamicFieldObjIDNameDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dynamicFieldObjIDNameDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dynamicFieldObjIDNameDo) Delete(models ...*model.DynamicFieldObjIDName) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dynamicFieldObjIDNameDo) withDO(do gen.Dao) *dynamicFieldObjIDNameDo {
	d.DO = *do.(*gen.DO)
	return d
}