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

func newPostmasterFilter(db *gorm.DB, opts ...gen.DOOption) postmasterFilter {
	_postmasterFilter := postmasterFilter{}

	_postmasterFilter.postmasterFilterDo.UseDB(db, opts...)
	_postmasterFilter.postmasterFilterDo.UseModel(&model.PostmasterFilter{})

	tableName := _postmasterFilter.postmasterFilterDo.TableName()
	_postmasterFilter.ALL = field.NewAsterisk(tableName)
	_postmasterFilter.FName = field.NewString(tableName, "f_name")
	_postmasterFilter.FStop = field.NewInt32(tableName, "f_stop")
	_postmasterFilter.FType = field.NewString(tableName, "f_type")
	_postmasterFilter.FKey = field.NewString(tableName, "f_key")
	_postmasterFilter.FValue = field.NewString(tableName, "f_value")
	_postmasterFilter.FNot = field.NewInt32(tableName, "f_not")

	_postmasterFilter.fillFieldMap()

	return _postmasterFilter
}

type postmasterFilter struct {
	postmasterFilterDo

	ALL    field.Asterisk
	FName  field.String
	FStop  field.Int32
	FType  field.String
	FKey   field.String
	FValue field.String
	FNot   field.Int32

	fieldMap map[string]field.Expr
}

func (p postmasterFilter) Table(newTableName string) *postmasterFilter {
	p.postmasterFilterDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p postmasterFilter) As(alias string) *postmasterFilter {
	p.postmasterFilterDo.DO = *(p.postmasterFilterDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *postmasterFilter) updateTableName(table string) *postmasterFilter {
	p.ALL = field.NewAsterisk(table)
	p.FName = field.NewString(table, "f_name")
	p.FStop = field.NewInt32(table, "f_stop")
	p.FType = field.NewString(table, "f_type")
	p.FKey = field.NewString(table, "f_key")
	p.FValue = field.NewString(table, "f_value")
	p.FNot = field.NewInt32(table, "f_not")

	p.fillFieldMap()

	return p
}

func (p *postmasterFilter) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *postmasterFilter) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["f_name"] = p.FName
	p.fieldMap["f_stop"] = p.FStop
	p.fieldMap["f_type"] = p.FType
	p.fieldMap["f_key"] = p.FKey
	p.fieldMap["f_value"] = p.FValue
	p.fieldMap["f_not"] = p.FNot
}

func (p postmasterFilter) clone(db *gorm.DB) postmasterFilter {
	p.postmasterFilterDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p postmasterFilter) replaceDB(db *gorm.DB) postmasterFilter {
	p.postmasterFilterDo.ReplaceDB(db)
	return p
}

type postmasterFilterDo struct{ gen.DO }

type IPostmasterFilterDo interface {
	gen.SubQuery
	Debug() IPostmasterFilterDo
	WithContext(ctx context.Context) IPostmasterFilterDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPostmasterFilterDo
	WriteDB() IPostmasterFilterDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPostmasterFilterDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPostmasterFilterDo
	Not(conds ...gen.Condition) IPostmasterFilterDo
	Or(conds ...gen.Condition) IPostmasterFilterDo
	Select(conds ...field.Expr) IPostmasterFilterDo
	Where(conds ...gen.Condition) IPostmasterFilterDo
	Order(conds ...field.Expr) IPostmasterFilterDo
	Distinct(cols ...field.Expr) IPostmasterFilterDo
	Omit(cols ...field.Expr) IPostmasterFilterDo
	Join(table schema.Tabler, on ...field.Expr) IPostmasterFilterDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPostmasterFilterDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPostmasterFilterDo
	Group(cols ...field.Expr) IPostmasterFilterDo
	Having(conds ...gen.Condition) IPostmasterFilterDo
	Limit(limit int) IPostmasterFilterDo
	Offset(offset int) IPostmasterFilterDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPostmasterFilterDo
	Unscoped() IPostmasterFilterDo
	Create(values ...*model.PostmasterFilter) error
	CreateInBatches(values []*model.PostmasterFilter, batchSize int) error
	Save(values ...*model.PostmasterFilter) error
	First() (*model.PostmasterFilter, error)
	Take() (*model.PostmasterFilter, error)
	Last() (*model.PostmasterFilter, error)
	Find() ([]*model.PostmasterFilter, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PostmasterFilter, err error)
	FindInBatches(result *[]*model.PostmasterFilter, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PostmasterFilter) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPostmasterFilterDo
	Assign(attrs ...field.AssignExpr) IPostmasterFilterDo
	Joins(fields ...field.RelationField) IPostmasterFilterDo
	Preload(fields ...field.RelationField) IPostmasterFilterDo
	FirstOrInit() (*model.PostmasterFilter, error)
	FirstOrCreate() (*model.PostmasterFilter, error)
	FindByPage(offset int, limit int) (result []*model.PostmasterFilter, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPostmasterFilterDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p postmasterFilterDo) Debug() IPostmasterFilterDo {
	return p.withDO(p.DO.Debug())
}

func (p postmasterFilterDo) WithContext(ctx context.Context) IPostmasterFilterDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p postmasterFilterDo) ReadDB() IPostmasterFilterDo {
	return p.Clauses(dbresolver.Read)
}

func (p postmasterFilterDo) WriteDB() IPostmasterFilterDo {
	return p.Clauses(dbresolver.Write)
}

func (p postmasterFilterDo) Session(config *gorm.Session) IPostmasterFilterDo {
	return p.withDO(p.DO.Session(config))
}

func (p postmasterFilterDo) Clauses(conds ...clause.Expression) IPostmasterFilterDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p postmasterFilterDo) Returning(value interface{}, columns ...string) IPostmasterFilterDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p postmasterFilterDo) Not(conds ...gen.Condition) IPostmasterFilterDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p postmasterFilterDo) Or(conds ...gen.Condition) IPostmasterFilterDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p postmasterFilterDo) Select(conds ...field.Expr) IPostmasterFilterDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p postmasterFilterDo) Where(conds ...gen.Condition) IPostmasterFilterDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p postmasterFilterDo) Order(conds ...field.Expr) IPostmasterFilterDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p postmasterFilterDo) Distinct(cols ...field.Expr) IPostmasterFilterDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p postmasterFilterDo) Omit(cols ...field.Expr) IPostmasterFilterDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p postmasterFilterDo) Join(table schema.Tabler, on ...field.Expr) IPostmasterFilterDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p postmasterFilterDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPostmasterFilterDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p postmasterFilterDo) RightJoin(table schema.Tabler, on ...field.Expr) IPostmasterFilterDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p postmasterFilterDo) Group(cols ...field.Expr) IPostmasterFilterDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p postmasterFilterDo) Having(conds ...gen.Condition) IPostmasterFilterDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p postmasterFilterDo) Limit(limit int) IPostmasterFilterDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p postmasterFilterDo) Offset(offset int) IPostmasterFilterDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p postmasterFilterDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPostmasterFilterDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p postmasterFilterDo) Unscoped() IPostmasterFilterDo {
	return p.withDO(p.DO.Unscoped())
}

func (p postmasterFilterDo) Create(values ...*model.PostmasterFilter) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p postmasterFilterDo) CreateInBatches(values []*model.PostmasterFilter, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p postmasterFilterDo) Save(values ...*model.PostmasterFilter) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p postmasterFilterDo) First() (*model.PostmasterFilter, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostmasterFilter), nil
	}
}

func (p postmasterFilterDo) Take() (*model.PostmasterFilter, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostmasterFilter), nil
	}
}

func (p postmasterFilterDo) Last() (*model.PostmasterFilter, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostmasterFilter), nil
	}
}

func (p postmasterFilterDo) Find() ([]*model.PostmasterFilter, error) {
	result, err := p.DO.Find()
	return result.([]*model.PostmasterFilter), err
}

func (p postmasterFilterDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PostmasterFilter, err error) {
	buf := make([]*model.PostmasterFilter, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p postmasterFilterDo) FindInBatches(result *[]*model.PostmasterFilter, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p postmasterFilterDo) Attrs(attrs ...field.AssignExpr) IPostmasterFilterDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p postmasterFilterDo) Assign(attrs ...field.AssignExpr) IPostmasterFilterDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p postmasterFilterDo) Joins(fields ...field.RelationField) IPostmasterFilterDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p postmasterFilterDo) Preload(fields ...field.RelationField) IPostmasterFilterDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p postmasterFilterDo) FirstOrInit() (*model.PostmasterFilter, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostmasterFilter), nil
	}
}

func (p postmasterFilterDo) FirstOrCreate() (*model.PostmasterFilter, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostmasterFilter), nil
	}
}

func (p postmasterFilterDo) FindByPage(offset int, limit int) (result []*model.PostmasterFilter, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p postmasterFilterDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p postmasterFilterDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p postmasterFilterDo) Delete(models ...*model.PostmasterFilter) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *postmasterFilterDo) withDO(do gen.Dao) *postmasterFilterDo {
	p.DO = *do.(*gen.DO)
	return p
}