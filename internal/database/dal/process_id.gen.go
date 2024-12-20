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

func newProcessID(db *gorm.DB, opts ...gen.DOOption) processID {
	_processID := processID{}

	_processID.processIDDo.UseDB(db, opts...)
	_processID.processIDDo.UseModel(&model.ProcessID{})

	tableName := _processID.processIDDo.TableName()
	_processID.ALL = field.NewAsterisk(tableName)
	_processID.ProcessName = field.NewString(tableName, "process_name")
	_processID.ProcessID = field.NewString(tableName, "process_id")
	_processID.ProcessHost = field.NewString(tableName, "process_host")
	_processID.ProcessCreate = field.NewInt32(tableName, "process_create")
	_processID.ProcessChange = field.NewInt32(tableName, "process_change")

	_processID.fillFieldMap()

	return _processID
}

type processID struct {
	processIDDo

	ALL           field.Asterisk
	ProcessName   field.String
	ProcessID     field.String
	ProcessHost   field.String
	ProcessCreate field.Int32
	ProcessChange field.Int32

	fieldMap map[string]field.Expr
}

func (p processID) Table(newTableName string) *processID {
	p.processIDDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p processID) As(alias string) *processID {
	p.processIDDo.DO = *(p.processIDDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *processID) updateTableName(table string) *processID {
	p.ALL = field.NewAsterisk(table)
	p.ProcessName = field.NewString(table, "process_name")
	p.ProcessID = field.NewString(table, "process_id")
	p.ProcessHost = field.NewString(table, "process_host")
	p.ProcessCreate = field.NewInt32(table, "process_create")
	p.ProcessChange = field.NewInt32(table, "process_change")

	p.fillFieldMap()

	return p
}

func (p *processID) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *processID) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 5)
	p.fieldMap["process_name"] = p.ProcessName
	p.fieldMap["process_id"] = p.ProcessID
	p.fieldMap["process_host"] = p.ProcessHost
	p.fieldMap["process_create"] = p.ProcessCreate
	p.fieldMap["process_change"] = p.ProcessChange
}

func (p processID) clone(db *gorm.DB) processID {
	p.processIDDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p processID) replaceDB(db *gorm.DB) processID {
	p.processIDDo.ReplaceDB(db)
	return p
}

type processIDDo struct{ gen.DO }

type IProcessIDDo interface {
	gen.SubQuery
	Debug() IProcessIDDo
	WithContext(ctx context.Context) IProcessIDDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProcessIDDo
	WriteDB() IProcessIDDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProcessIDDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProcessIDDo
	Not(conds ...gen.Condition) IProcessIDDo
	Or(conds ...gen.Condition) IProcessIDDo
	Select(conds ...field.Expr) IProcessIDDo
	Where(conds ...gen.Condition) IProcessIDDo
	Order(conds ...field.Expr) IProcessIDDo
	Distinct(cols ...field.Expr) IProcessIDDo
	Omit(cols ...field.Expr) IProcessIDDo
	Join(table schema.Tabler, on ...field.Expr) IProcessIDDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProcessIDDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProcessIDDo
	Group(cols ...field.Expr) IProcessIDDo
	Having(conds ...gen.Condition) IProcessIDDo
	Limit(limit int) IProcessIDDo
	Offset(offset int) IProcessIDDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProcessIDDo
	Unscoped() IProcessIDDo
	Create(values ...*model.ProcessID) error
	CreateInBatches(values []*model.ProcessID, batchSize int) error
	Save(values ...*model.ProcessID) error
	First() (*model.ProcessID, error)
	Take() (*model.ProcessID, error)
	Last() (*model.ProcessID, error)
	Find() ([]*model.ProcessID, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProcessID, err error)
	FindInBatches(result *[]*model.ProcessID, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProcessID) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProcessIDDo
	Assign(attrs ...field.AssignExpr) IProcessIDDo
	Joins(fields ...field.RelationField) IProcessIDDo
	Preload(fields ...field.RelationField) IProcessIDDo
	FirstOrInit() (*model.ProcessID, error)
	FirstOrCreate() (*model.ProcessID, error)
	FindByPage(offset int, limit int) (result []*model.ProcessID, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProcessIDDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p processIDDo) Debug() IProcessIDDo {
	return p.withDO(p.DO.Debug())
}

func (p processIDDo) WithContext(ctx context.Context) IProcessIDDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p processIDDo) ReadDB() IProcessIDDo {
	return p.Clauses(dbresolver.Read)
}

func (p processIDDo) WriteDB() IProcessIDDo {
	return p.Clauses(dbresolver.Write)
}

func (p processIDDo) Session(config *gorm.Session) IProcessIDDo {
	return p.withDO(p.DO.Session(config))
}

func (p processIDDo) Clauses(conds ...clause.Expression) IProcessIDDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p processIDDo) Returning(value interface{}, columns ...string) IProcessIDDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p processIDDo) Not(conds ...gen.Condition) IProcessIDDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p processIDDo) Or(conds ...gen.Condition) IProcessIDDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p processIDDo) Select(conds ...field.Expr) IProcessIDDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p processIDDo) Where(conds ...gen.Condition) IProcessIDDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p processIDDo) Order(conds ...field.Expr) IProcessIDDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p processIDDo) Distinct(cols ...field.Expr) IProcessIDDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p processIDDo) Omit(cols ...field.Expr) IProcessIDDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p processIDDo) Join(table schema.Tabler, on ...field.Expr) IProcessIDDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p processIDDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProcessIDDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p processIDDo) RightJoin(table schema.Tabler, on ...field.Expr) IProcessIDDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p processIDDo) Group(cols ...field.Expr) IProcessIDDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p processIDDo) Having(conds ...gen.Condition) IProcessIDDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p processIDDo) Limit(limit int) IProcessIDDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p processIDDo) Offset(offset int) IProcessIDDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p processIDDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProcessIDDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p processIDDo) Unscoped() IProcessIDDo {
	return p.withDO(p.DO.Unscoped())
}

func (p processIDDo) Create(values ...*model.ProcessID) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p processIDDo) CreateInBatches(values []*model.ProcessID, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p processIDDo) Save(values ...*model.ProcessID) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p processIDDo) First() (*model.ProcessID, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessID), nil
	}
}

func (p processIDDo) Take() (*model.ProcessID, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessID), nil
	}
}

func (p processIDDo) Last() (*model.ProcessID, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessID), nil
	}
}

func (p processIDDo) Find() ([]*model.ProcessID, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProcessID), err
}

func (p processIDDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProcessID, err error) {
	buf := make([]*model.ProcessID, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p processIDDo) FindInBatches(result *[]*model.ProcessID, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p processIDDo) Attrs(attrs ...field.AssignExpr) IProcessIDDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p processIDDo) Assign(attrs ...field.AssignExpr) IProcessIDDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p processIDDo) Joins(fields ...field.RelationField) IProcessIDDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p processIDDo) Preload(fields ...field.RelationField) IProcessIDDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p processIDDo) FirstOrInit() (*model.ProcessID, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessID), nil
	}
}

func (p processIDDo) FirstOrCreate() (*model.ProcessID, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessID), nil
	}
}

func (p processIDDo) FindByPage(offset int, limit int) (result []*model.ProcessID, count int64, err error) {
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

func (p processIDDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p processIDDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p processIDDo) Delete(models ...*model.ProcessID) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *processIDDo) withDO(do gen.Dao) *processIDDo {
	p.DO = *do.(*gen.DO)
	return p
}
