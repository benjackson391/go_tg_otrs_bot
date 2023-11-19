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

func newGiDebuggerEntry(db *gorm.DB, opts ...gen.DOOption) giDebuggerEntry {
	_giDebuggerEntry := giDebuggerEntry{}

	_giDebuggerEntry.giDebuggerEntryDo.UseDB(db, opts...)
	_giDebuggerEntry.giDebuggerEntryDo.UseModel(&model.GiDebuggerEntry{})

	tableName := _giDebuggerEntry.giDebuggerEntryDo.TableName()
	_giDebuggerEntry.ALL = field.NewAsterisk(tableName)
	_giDebuggerEntry.ID = field.NewInt64(tableName, "id")
	_giDebuggerEntry.CommunicationID = field.NewString(tableName, "communication_id")
	_giDebuggerEntry.CommunicationType = field.NewString(tableName, "communication_type")
	_giDebuggerEntry.RemoteIP = field.NewString(tableName, "remote_ip")
	_giDebuggerEntry.WebserviceID = field.NewInt32(tableName, "webservice_id")
	_giDebuggerEntry.CreateTime = field.NewTime(tableName, "create_time")

	_giDebuggerEntry.fillFieldMap()

	return _giDebuggerEntry
}

type giDebuggerEntry struct {
	giDebuggerEntryDo

	ALL               field.Asterisk
	ID                field.Int64
	CommunicationID   field.String
	CommunicationType field.String
	RemoteIP          field.String
	WebserviceID      field.Int32
	CreateTime        field.Time

	fieldMap map[string]field.Expr
}

func (g giDebuggerEntry) Table(newTableName string) *giDebuggerEntry {
	g.giDebuggerEntryDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g giDebuggerEntry) As(alias string) *giDebuggerEntry {
	g.giDebuggerEntryDo.DO = *(g.giDebuggerEntryDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *giDebuggerEntry) updateTableName(table string) *giDebuggerEntry {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt64(table, "id")
	g.CommunicationID = field.NewString(table, "communication_id")
	g.CommunicationType = field.NewString(table, "communication_type")
	g.RemoteIP = field.NewString(table, "remote_ip")
	g.WebserviceID = field.NewInt32(table, "webservice_id")
	g.CreateTime = field.NewTime(table, "create_time")

	g.fillFieldMap()

	return g
}

func (g *giDebuggerEntry) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *giDebuggerEntry) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 6)
	g.fieldMap["id"] = g.ID
	g.fieldMap["communication_id"] = g.CommunicationID
	g.fieldMap["communication_type"] = g.CommunicationType
	g.fieldMap["remote_ip"] = g.RemoteIP
	g.fieldMap["webservice_id"] = g.WebserviceID
	g.fieldMap["create_time"] = g.CreateTime
}

func (g giDebuggerEntry) clone(db *gorm.DB) giDebuggerEntry {
	g.giDebuggerEntryDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g giDebuggerEntry) replaceDB(db *gorm.DB) giDebuggerEntry {
	g.giDebuggerEntryDo.ReplaceDB(db)
	return g
}

type giDebuggerEntryDo struct{ gen.DO }

type IGiDebuggerEntryDo interface {
	gen.SubQuery
	Debug() IGiDebuggerEntryDo
	WithContext(ctx context.Context) IGiDebuggerEntryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGiDebuggerEntryDo
	WriteDB() IGiDebuggerEntryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGiDebuggerEntryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGiDebuggerEntryDo
	Not(conds ...gen.Condition) IGiDebuggerEntryDo
	Or(conds ...gen.Condition) IGiDebuggerEntryDo
	Select(conds ...field.Expr) IGiDebuggerEntryDo
	Where(conds ...gen.Condition) IGiDebuggerEntryDo
	Order(conds ...field.Expr) IGiDebuggerEntryDo
	Distinct(cols ...field.Expr) IGiDebuggerEntryDo
	Omit(cols ...field.Expr) IGiDebuggerEntryDo
	Join(table schema.Tabler, on ...field.Expr) IGiDebuggerEntryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGiDebuggerEntryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGiDebuggerEntryDo
	Group(cols ...field.Expr) IGiDebuggerEntryDo
	Having(conds ...gen.Condition) IGiDebuggerEntryDo
	Limit(limit int) IGiDebuggerEntryDo
	Offset(offset int) IGiDebuggerEntryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGiDebuggerEntryDo
	Unscoped() IGiDebuggerEntryDo
	Create(values ...*model.GiDebuggerEntry) error
	CreateInBatches(values []*model.GiDebuggerEntry, batchSize int) error
	Save(values ...*model.GiDebuggerEntry) error
	First() (*model.GiDebuggerEntry, error)
	Take() (*model.GiDebuggerEntry, error)
	Last() (*model.GiDebuggerEntry, error)
	Find() ([]*model.GiDebuggerEntry, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GiDebuggerEntry, err error)
	FindInBatches(result *[]*model.GiDebuggerEntry, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GiDebuggerEntry) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGiDebuggerEntryDo
	Assign(attrs ...field.AssignExpr) IGiDebuggerEntryDo
	Joins(fields ...field.RelationField) IGiDebuggerEntryDo
	Preload(fields ...field.RelationField) IGiDebuggerEntryDo
	FirstOrInit() (*model.GiDebuggerEntry, error)
	FirstOrCreate() (*model.GiDebuggerEntry, error)
	FindByPage(offset int, limit int) (result []*model.GiDebuggerEntry, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGiDebuggerEntryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g giDebuggerEntryDo) Debug() IGiDebuggerEntryDo {
	return g.withDO(g.DO.Debug())
}

func (g giDebuggerEntryDo) WithContext(ctx context.Context) IGiDebuggerEntryDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g giDebuggerEntryDo) ReadDB() IGiDebuggerEntryDo {
	return g.Clauses(dbresolver.Read)
}

func (g giDebuggerEntryDo) WriteDB() IGiDebuggerEntryDo {
	return g.Clauses(dbresolver.Write)
}

func (g giDebuggerEntryDo) Session(config *gorm.Session) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Session(config))
}

func (g giDebuggerEntryDo) Clauses(conds ...clause.Expression) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g giDebuggerEntryDo) Returning(value interface{}, columns ...string) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g giDebuggerEntryDo) Not(conds ...gen.Condition) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g giDebuggerEntryDo) Or(conds ...gen.Condition) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g giDebuggerEntryDo) Select(conds ...field.Expr) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g giDebuggerEntryDo) Where(conds ...gen.Condition) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g giDebuggerEntryDo) Order(conds ...field.Expr) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g giDebuggerEntryDo) Distinct(cols ...field.Expr) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g giDebuggerEntryDo) Omit(cols ...field.Expr) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g giDebuggerEntryDo) Join(table schema.Tabler, on ...field.Expr) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g giDebuggerEntryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGiDebuggerEntryDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g giDebuggerEntryDo) RightJoin(table schema.Tabler, on ...field.Expr) IGiDebuggerEntryDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g giDebuggerEntryDo) Group(cols ...field.Expr) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g giDebuggerEntryDo) Having(conds ...gen.Condition) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g giDebuggerEntryDo) Limit(limit int) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g giDebuggerEntryDo) Offset(offset int) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g giDebuggerEntryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g giDebuggerEntryDo) Unscoped() IGiDebuggerEntryDo {
	return g.withDO(g.DO.Unscoped())
}

func (g giDebuggerEntryDo) Create(values ...*model.GiDebuggerEntry) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g giDebuggerEntryDo) CreateInBatches(values []*model.GiDebuggerEntry, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g giDebuggerEntryDo) Save(values ...*model.GiDebuggerEntry) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g giDebuggerEntryDo) First() (*model.GiDebuggerEntry, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiDebuggerEntry), nil
	}
}

func (g giDebuggerEntryDo) Take() (*model.GiDebuggerEntry, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiDebuggerEntry), nil
	}
}

func (g giDebuggerEntryDo) Last() (*model.GiDebuggerEntry, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiDebuggerEntry), nil
	}
}

func (g giDebuggerEntryDo) Find() ([]*model.GiDebuggerEntry, error) {
	result, err := g.DO.Find()
	return result.([]*model.GiDebuggerEntry), err
}

func (g giDebuggerEntryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GiDebuggerEntry, err error) {
	buf := make([]*model.GiDebuggerEntry, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g giDebuggerEntryDo) FindInBatches(result *[]*model.GiDebuggerEntry, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g giDebuggerEntryDo) Attrs(attrs ...field.AssignExpr) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g giDebuggerEntryDo) Assign(attrs ...field.AssignExpr) IGiDebuggerEntryDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g giDebuggerEntryDo) Joins(fields ...field.RelationField) IGiDebuggerEntryDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g giDebuggerEntryDo) Preload(fields ...field.RelationField) IGiDebuggerEntryDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g giDebuggerEntryDo) FirstOrInit() (*model.GiDebuggerEntry, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiDebuggerEntry), nil
	}
}

func (g giDebuggerEntryDo) FirstOrCreate() (*model.GiDebuggerEntry, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GiDebuggerEntry), nil
	}
}

func (g giDebuggerEntryDo) FindByPage(offset int, limit int) (result []*model.GiDebuggerEntry, count int64, err error) {
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

func (g giDebuggerEntryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g giDebuggerEntryDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g giDebuggerEntryDo) Delete(models ...*model.GiDebuggerEntry) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *giDebuggerEntryDo) withDO(do gen.Dao) *giDebuggerEntryDo {
	g.DO = *do.(*gen.DO)
	return g
}
