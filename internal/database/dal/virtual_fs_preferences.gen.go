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

func newVirtualFsPreference(db *gorm.DB, opts ...gen.DOOption) virtualFsPreference {
	_virtualFsPreference := virtualFsPreference{}

	_virtualFsPreference.virtualFsPreferenceDo.UseDB(db, opts...)
	_virtualFsPreference.virtualFsPreferenceDo.UseModel(&model.VirtualFsPreference{})

	tableName := _virtualFsPreference.virtualFsPreferenceDo.TableName()
	_virtualFsPreference.ALL = field.NewAsterisk(tableName)
	_virtualFsPreference.VirtualFsID = field.NewInt64(tableName, "virtual_fs_id")
	_virtualFsPreference.PreferencesKey = field.NewString(tableName, "preferences_key")
	_virtualFsPreference.PreferencesValue = field.NewString(tableName, "preferences_value")

	_virtualFsPreference.fillFieldMap()

	return _virtualFsPreference
}

type virtualFsPreference struct {
	virtualFsPreferenceDo

	ALL              field.Asterisk
	VirtualFsID      field.Int64
	PreferencesKey   field.String
	PreferencesValue field.String

	fieldMap map[string]field.Expr
}

func (v virtualFsPreference) Table(newTableName string) *virtualFsPreference {
	v.virtualFsPreferenceDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v virtualFsPreference) As(alias string) *virtualFsPreference {
	v.virtualFsPreferenceDo.DO = *(v.virtualFsPreferenceDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *virtualFsPreference) updateTableName(table string) *virtualFsPreference {
	v.ALL = field.NewAsterisk(table)
	v.VirtualFsID = field.NewInt64(table, "virtual_fs_id")
	v.PreferencesKey = field.NewString(table, "preferences_key")
	v.PreferencesValue = field.NewString(table, "preferences_value")

	v.fillFieldMap()

	return v
}

func (v *virtualFsPreference) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *virtualFsPreference) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 3)
	v.fieldMap["virtual_fs_id"] = v.VirtualFsID
	v.fieldMap["preferences_key"] = v.PreferencesKey
	v.fieldMap["preferences_value"] = v.PreferencesValue
}

func (v virtualFsPreference) clone(db *gorm.DB) virtualFsPreference {
	v.virtualFsPreferenceDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v virtualFsPreference) replaceDB(db *gorm.DB) virtualFsPreference {
	v.virtualFsPreferenceDo.ReplaceDB(db)
	return v
}

type virtualFsPreferenceDo struct{ gen.DO }

type IVirtualFsPreferenceDo interface {
	gen.SubQuery
	Debug() IVirtualFsPreferenceDo
	WithContext(ctx context.Context) IVirtualFsPreferenceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVirtualFsPreferenceDo
	WriteDB() IVirtualFsPreferenceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVirtualFsPreferenceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVirtualFsPreferenceDo
	Not(conds ...gen.Condition) IVirtualFsPreferenceDo
	Or(conds ...gen.Condition) IVirtualFsPreferenceDo
	Select(conds ...field.Expr) IVirtualFsPreferenceDo
	Where(conds ...gen.Condition) IVirtualFsPreferenceDo
	Order(conds ...field.Expr) IVirtualFsPreferenceDo
	Distinct(cols ...field.Expr) IVirtualFsPreferenceDo
	Omit(cols ...field.Expr) IVirtualFsPreferenceDo
	Join(table schema.Tabler, on ...field.Expr) IVirtualFsPreferenceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVirtualFsPreferenceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVirtualFsPreferenceDo
	Group(cols ...field.Expr) IVirtualFsPreferenceDo
	Having(conds ...gen.Condition) IVirtualFsPreferenceDo
	Limit(limit int) IVirtualFsPreferenceDo
	Offset(offset int) IVirtualFsPreferenceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVirtualFsPreferenceDo
	Unscoped() IVirtualFsPreferenceDo
	Create(values ...*model.VirtualFsPreference) error
	CreateInBatches(values []*model.VirtualFsPreference, batchSize int) error
	Save(values ...*model.VirtualFsPreference) error
	First() (*model.VirtualFsPreference, error)
	Take() (*model.VirtualFsPreference, error)
	Last() (*model.VirtualFsPreference, error)
	Find() ([]*model.VirtualFsPreference, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VirtualFsPreference, err error)
	FindInBatches(result *[]*model.VirtualFsPreference, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.VirtualFsPreference) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVirtualFsPreferenceDo
	Assign(attrs ...field.AssignExpr) IVirtualFsPreferenceDo
	Joins(fields ...field.RelationField) IVirtualFsPreferenceDo
	Preload(fields ...field.RelationField) IVirtualFsPreferenceDo
	FirstOrInit() (*model.VirtualFsPreference, error)
	FirstOrCreate() (*model.VirtualFsPreference, error)
	FindByPage(offset int, limit int) (result []*model.VirtualFsPreference, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVirtualFsPreferenceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v virtualFsPreferenceDo) Debug() IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Debug())
}

func (v virtualFsPreferenceDo) WithContext(ctx context.Context) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v virtualFsPreferenceDo) ReadDB() IVirtualFsPreferenceDo {
	return v.Clauses(dbresolver.Read)
}

func (v virtualFsPreferenceDo) WriteDB() IVirtualFsPreferenceDo {
	return v.Clauses(dbresolver.Write)
}

func (v virtualFsPreferenceDo) Session(config *gorm.Session) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Session(config))
}

func (v virtualFsPreferenceDo) Clauses(conds ...clause.Expression) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v virtualFsPreferenceDo) Returning(value interface{}, columns ...string) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v virtualFsPreferenceDo) Not(conds ...gen.Condition) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v virtualFsPreferenceDo) Or(conds ...gen.Condition) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v virtualFsPreferenceDo) Select(conds ...field.Expr) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v virtualFsPreferenceDo) Where(conds ...gen.Condition) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v virtualFsPreferenceDo) Order(conds ...field.Expr) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v virtualFsPreferenceDo) Distinct(cols ...field.Expr) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v virtualFsPreferenceDo) Omit(cols ...field.Expr) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v virtualFsPreferenceDo) Join(table schema.Tabler, on ...field.Expr) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v virtualFsPreferenceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v virtualFsPreferenceDo) RightJoin(table schema.Tabler, on ...field.Expr) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v virtualFsPreferenceDo) Group(cols ...field.Expr) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v virtualFsPreferenceDo) Having(conds ...gen.Condition) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v virtualFsPreferenceDo) Limit(limit int) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v virtualFsPreferenceDo) Offset(offset int) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v virtualFsPreferenceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v virtualFsPreferenceDo) Unscoped() IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Unscoped())
}

func (v virtualFsPreferenceDo) Create(values ...*model.VirtualFsPreference) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v virtualFsPreferenceDo) CreateInBatches(values []*model.VirtualFsPreference, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v virtualFsPreferenceDo) Save(values ...*model.VirtualFsPreference) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v virtualFsPreferenceDo) First() (*model.VirtualFsPreference, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VirtualFsPreference), nil
	}
}

func (v virtualFsPreferenceDo) Take() (*model.VirtualFsPreference, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VirtualFsPreference), nil
	}
}

func (v virtualFsPreferenceDo) Last() (*model.VirtualFsPreference, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VirtualFsPreference), nil
	}
}

func (v virtualFsPreferenceDo) Find() ([]*model.VirtualFsPreference, error) {
	result, err := v.DO.Find()
	return result.([]*model.VirtualFsPreference), err
}

func (v virtualFsPreferenceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VirtualFsPreference, err error) {
	buf := make([]*model.VirtualFsPreference, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v virtualFsPreferenceDo) FindInBatches(result *[]*model.VirtualFsPreference, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v virtualFsPreferenceDo) Attrs(attrs ...field.AssignExpr) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v virtualFsPreferenceDo) Assign(attrs ...field.AssignExpr) IVirtualFsPreferenceDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v virtualFsPreferenceDo) Joins(fields ...field.RelationField) IVirtualFsPreferenceDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v virtualFsPreferenceDo) Preload(fields ...field.RelationField) IVirtualFsPreferenceDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v virtualFsPreferenceDo) FirstOrInit() (*model.VirtualFsPreference, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VirtualFsPreference), nil
	}
}

func (v virtualFsPreferenceDo) FirstOrCreate() (*model.VirtualFsPreference, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VirtualFsPreference), nil
	}
}

func (v virtualFsPreferenceDo) FindByPage(offset int, limit int) (result []*model.VirtualFsPreference, count int64, err error) {
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

func (v virtualFsPreferenceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v virtualFsPreferenceDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v virtualFsPreferenceDo) Delete(models ...*model.VirtualFsPreference) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *virtualFsPreferenceDo) withDO(do gen.Dao) *virtualFsPreferenceDo {
	v.DO = *do.(*gen.DO)
	return v
}
