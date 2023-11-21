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

func newGiDebuggerEntryContent(db *gorm.DB, opts ...gen.DOOption) giDebuggerEntryContent {
	_giDebuggerEntryContent := giDebuggerEntryContent{}

	_giDebuggerEntryContent.giDebuggerEntryContentDo.UseDB(db, opts...)
	_giDebuggerEntryContent.giDebuggerEntryContentDo.UseModel(&model.GiDebuggerEntryContent{})

	tableName := _giDebuggerEntryContent.giDebuggerEntryContentDo.TableName()
	_giDebuggerEntryContent.ALL = field.NewAsterisk(tableName)
	_giDebuggerEntryContent.ID = field.NewInt64(tableName, "id")
	_giDebuggerEntryContent.GiDebuggerEntryID = field.NewInt64(tableName, "gi_debugger_entry_id")
	_giDebuggerEntryContent.DebugLevel = field.NewString(tableName, "debug_level")
	_giDebuggerEntryContent.Subject = field.NewString(tableName, "subject")
	_giDebuggerEntryContent.Content = field.NewBytes(tableName, "content")
	_giDebuggerEntryContent.CreateTime = field.NewTime(tableName, "create_time")

	_giDebuggerEntryContent.fillFieldMap()

	return _giDebuggerEntryContent
}

type giDebuggerEntryContent struct {
	giDebuggerEntryContentDo

	ALL               field.Asterisk
	ID                field.Int64
	GiDebuggerEntryID field.Int64
	DebugLevel        field.String
	Subject           field.String
	Content           field.Bytes
	CreateTime        field.Time

	fieldMap map[string]field.Expr
}

func (g giDebuggerEntryContent) Table(newTableName string) *giDebuggerEntryContent {
	g.giDebuggerEntryContentDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g giDebuggerEntryContent) As(alias string) *giDebuggerEntryContent {
	g.giDebuggerEntryContentDo.DO = *(g.giDebuggerEntryContentDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *giDebuggerEntryContent) updateTableName(table string) *giDebuggerEntryContent {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt64(table, "id")
	g.GiDebuggerEntryID = field.NewInt64(table, "gi_debugger_entry_id")
	g.DebugLevel = field.NewString(table, "debug_level")
	g.Subject = field.NewString(table, "subject")
	g.Content = field.NewBytes(table, "content")
	g.CreateTime = field.NewTime(table, "create_time")

	g.fillFieldMap()

	return g
}

func (g *giDebuggerEntryContent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *giDebuggerEntryContent) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 6)
	g.fieldMap["id"] = g.ID
	g.fieldMap["gi_debugger_entry_id"] = g.GiDebuggerEntryID
	g.fieldMap["debug_level"] = g.DebugLevel
	g.fieldMap["subject"] = g.Subject
	g.fieldMap["content"] = g.Content
	g.fieldMap["create_time"] = g.CreateTime
}

func (g giDebuggerEntryContent) clone(db *gorm.DB) giDebuggerEntryContent {
	g.giDebuggerEntryContentDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g giDebuggerEntryContent) replaceDB(db *gorm.DB) giDebuggerEntryContent {
	g.giDebuggerEntryContentDo.ReplaceDB(db)
	return g
}

type giDebuggerEntryContentDo struct{ gen.DO }

type IGiDebuggerEntryContentDo interface {
	gen.SubQuery
	Debug() IGiDebuggerEntryContentDo
	WithContext(ctx context.Context) IGiDebuggerEntryContentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGiDebuggerEntryContentDo
	WriteDB() IGiDebuggerEntryContentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGiDebuggerEntryContentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGiDebuggerEntryContentDo
	Not(conds ...gen.Condition) IGiDebuggerEntryContentDo
	Or(conds ...gen.Condition) IGiDebuggerEntryContentDo
	Select(conds ...field.Expr) IGiDebuggerEntryContentDo
	Where(conds ...gen.Condition) IGiDebuggerEntryContentDo
	Order(conds ...field.Expr) IGiDebuggerEntryContentDo
	Distinct(cols ...field.Expr) IGiDebuggerEntryContentDo
	Omit(cols ...field.Expr) IGiDebuggerEntryContentDo
	Join(table schema.Tabler, on ...field.Expr) IGiDebuggerEntryContentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGiDebuggerEntryContentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGiDebuggerEntryContentDo
	Group(cols ...field.Expr) IGiDebuggerEntryContentDo
	Having(conds ...gen.Condition) IGiDebuggerEntryContentDo
	Limit(limit int) IGiDebuggerEntryContentDo
	Offset(offset int) IGiDebuggerEntryContentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGiDebuggerEntryContentDo
	Unscoped() IGiDebuggerEntryContentDo
	Create(values ...*model.GiDebuggerEntryContent) error
	CreateInBatches(values []*model.GiDebuggerEntryContent, batchSize int) error
	Save(values ...*model.GiDebuggerEntryContent) error
	First() (*model.GiDebuggerEntryContent, error)
	Take() (*model.GiDebuggerEntryContent, error)
	Last() (*model.GiDebuggerEntryContent, error)
	Find() ([]*model.GiDebuggerEntryContent, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GiDebuggerEntryContent, err error)
	FindInBatches(result *[]*model.GiDebuggerEntryContent, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GiDebuggerEntryContent) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGiDebuggerEntryContentDo
	Assign(attrs ...field.AssignExpr) IGiDebuggerEntryContentDo
	Joins(fields ...field.RelationField) IGiDebuggerEntryContentDo
	Preload(fields ...field.RelationField) IGiDebuggerEntryContentDo
	FirstOrInit() (*model.GiDebuggerEntryContent, error)
	FirstOrCreate() (*model.GiDebuggerEntryContent, error)
	FindByPage(offset int, limit int) (result []*model.GiDebuggerEntryContent, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGiDebuggerEntryContentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g giDebuggerEntryContentDo) Debug() IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Debug())
}

func (g giDebuggerEntryContentDo) WithContext(ctx context.Context) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g giDebuggerEntryContentDo) ReadDB() IGiDebuggerEntryContentDo {
	return g.Clauses(dbresolver.Read)
}

func (g giDebuggerEntryContentDo) WriteDB() IGiDebuggerEntryContentDo {
	return g.Clauses(dbresolver.Write)
}

func (g giDebuggerEntryContentDo) Session(config *gorm.Session) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Session(config))
}

func (g giDebuggerEntryContentDo) Clauses(conds ...clause.Expression) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g giDebuggerEntryContentDo) Returning(value interface{}, columns ...string) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g giDebuggerEntryContentDo) Not(conds ...gen.Condition) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g giDebuggerEntryContentDo) Or(conds ...gen.Condition) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g giDebuggerEntryContentDo) Select(conds ...field.Expr) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g giDebuggerEntryContentDo) Where(conds ...gen.Condition) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g giDebuggerEntryContentDo) Order(conds ...field.Expr) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g giDebuggerEntryContentDo) Distinct(cols ...field.Expr) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g giDebuggerEntryContentDo) Omit(cols ...field.Expr) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g giDebuggerEntryContentDo) Join(table schema.Tabler, on ...field.Expr) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g giDebuggerEntryContentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g giDebuggerEntryContentDo) RightJoin(table schema.Tabler, on ...field.Expr) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g giDebuggerEntryContentDo) Group(cols ...field.Expr) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g giDebuggerEntryContentDo) Having(conds ...gen.Condition) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g giDebuggerEntryContentDo) Limit(limit int) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g giDebuggerEntryContentDo) Offset(offset int) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g giDebuggerEntryContentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g giDebuggerEntryContentDo) Unscoped() IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Unscoped())
}

func (g giDebuggerEntryContentDo) Create(values ...*model.GiDebuggerEntryContent) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g giDebuggerEntryContentDo) CreateInBatches(values []*model.GiDebuggerEntryContent, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g giDebuggerEntryContentDo) Save(values ...*model.GiDebuggerEntryContent) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g giDebuggerEntryContentDo) First() (*model.GiDebuggerEntryContent, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiDebuggerEntryContent), nil
	}
}

func (g giDebuggerEntryContentDo) Take() (*model.GiDebuggerEntryContent, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiDebuggerEntryContent), nil
	}
}

func (g giDebuggerEntryContentDo) Last() (*model.GiDebuggerEntryContent, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiDebuggerEntryContent), nil
	}
}

func (g giDebuggerEntryContentDo) Find() ([]*model.GiDebuggerEntryContent, error) {
	result, err := g.DO.Find()
	return result.([]*model.GiDebuggerEntryContent), err
}

func (g giDebuggerEntryContentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GiDebuggerEntryContent, err error) {
	buf := make([]*model.GiDebuggerEntryContent, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g giDebuggerEntryContentDo) FindInBatches(result *[]*model.GiDebuggerEntryContent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g giDebuggerEntryContentDo) Attrs(attrs ...field.AssignExpr) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g giDebuggerEntryContentDo) Assign(attrs ...field.AssignExpr) IGiDebuggerEntryContentDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g giDebuggerEntryContentDo) Joins(fields ...field.RelationField) IGiDebuggerEntryContentDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g giDebuggerEntryContentDo) Preload(fields ...field.RelationField) IGiDebuggerEntryContentDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g giDebuggerEntryContentDo) FirstOrInit() (*model.GiDebuggerEntryContent, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiDebuggerEntryContent), nil
	}
}

func (g giDebuggerEntryContentDo) FirstOrCreate() (*model.GiDebuggerEntryContent, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiDebuggerEntryContent), nil
	}
}

func (g giDebuggerEntryContentDo) FindByPage(offset int, limit int) (result []*model.GiDebuggerEntryContent, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g giDebuggerEntryContentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g giDebuggerEntryContentDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g giDebuggerEntryContentDo) Delete(models ...*model.GiDebuggerEntryContent) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *giDebuggerEntryContentDo) withDO(do gen.Dao) *giDebuggerEntryContentDo {
	g.DO = *do.(*gen.DO)
	return g
}