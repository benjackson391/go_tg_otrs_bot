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

func newTgUser(db *gorm.DB, opts ...gen.DOOption) tgUser {
	_tgUser := tgUser{}

	_tgUser.tgUserDo.UseDB(db, opts...)
	_tgUser.tgUserDo.UseModel(&model.TgUser{})

	tableName := _tgUser.tgUserDo.TableName()
	_tgUser.ALL = field.NewAsterisk(tableName)
	_tgUser.TelegramLogin = field.NewString(tableName, "TelegramLogin")
	_tgUser.CustomerUserLogin = field.NewString(tableName, "CustomerUserLogin")

	_tgUser.fillFieldMap()

	return _tgUser
}

type tgUser struct {
	tgUserDo

	ALL               field.Asterisk
	TelegramLogin     field.String
	CustomerUserLogin field.String

	fieldMap map[string]field.Expr
}

func (t tgUser) Table(newTableName string) *tgUser {
	t.tgUserDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tgUser) As(alias string) *tgUser {
	t.tgUserDo.DO = *(t.tgUserDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tgUser) updateTableName(table string) *tgUser {
	t.ALL = field.NewAsterisk(table)
	t.TelegramLogin = field.NewString(table, "TelegramLogin")
	t.CustomerUserLogin = field.NewString(table, "CustomerUserLogin")

	t.fillFieldMap()

	return t
}

func (t *tgUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tgUser) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 2)
	t.fieldMap["TelegramLogin"] = t.TelegramLogin
	t.fieldMap["CustomerUserLogin"] = t.CustomerUserLogin
}

func (t tgUser) clone(db *gorm.DB) tgUser {
	t.tgUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tgUser) replaceDB(db *gorm.DB) tgUser {
	t.tgUserDo.ReplaceDB(db)
	return t
}

type tgUserDo struct{ gen.DO }

type ITgUserDo interface {
	gen.SubQuery
	Debug() ITgUserDo
	WithContext(ctx context.Context) ITgUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITgUserDo
	WriteDB() ITgUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITgUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITgUserDo
	Not(conds ...gen.Condition) ITgUserDo
	Or(conds ...gen.Condition) ITgUserDo
	Select(conds ...field.Expr) ITgUserDo
	Where(conds ...gen.Condition) ITgUserDo
	Order(conds ...field.Expr) ITgUserDo
	Distinct(cols ...field.Expr) ITgUserDo
	Omit(cols ...field.Expr) ITgUserDo
	Join(table schema.Tabler, on ...field.Expr) ITgUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITgUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITgUserDo
	Group(cols ...field.Expr) ITgUserDo
	Having(conds ...gen.Condition) ITgUserDo
	Limit(limit int) ITgUserDo
	Offset(offset int) ITgUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITgUserDo
	Unscoped() ITgUserDo
	Create(values ...*model.TgUser) error
	CreateInBatches(values []*model.TgUser, batchSize int) error
	Save(values ...*model.TgUser) error
	First() (*model.TgUser, error)
	Take() (*model.TgUser, error)
	Last() (*model.TgUser, error)
	Find() ([]*model.TgUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TgUser, err error)
	FindInBatches(result *[]*model.TgUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TgUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITgUserDo
	Assign(attrs ...field.AssignExpr) ITgUserDo
	Joins(fields ...field.RelationField) ITgUserDo
	Preload(fields ...field.RelationField) ITgUserDo
	FirstOrInit() (*model.TgUser, error)
	FirstOrCreate() (*model.TgUser, error)
	FindByPage(offset int, limit int) (result []*model.TgUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITgUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tgUserDo) Debug() ITgUserDo {
	return t.withDO(t.DO.Debug())
}

func (t tgUserDo) WithContext(ctx context.Context) ITgUserDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tgUserDo) ReadDB() ITgUserDo {
	return t.Clauses(dbresolver.Read)
}

func (t tgUserDo) WriteDB() ITgUserDo {
	return t.Clauses(dbresolver.Write)
}

func (t tgUserDo) Session(config *gorm.Session) ITgUserDo {
	return t.withDO(t.DO.Session(config))
}

func (t tgUserDo) Clauses(conds ...clause.Expression) ITgUserDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tgUserDo) Returning(value interface{}, columns ...string) ITgUserDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tgUserDo) Not(conds ...gen.Condition) ITgUserDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tgUserDo) Or(conds ...gen.Condition) ITgUserDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tgUserDo) Select(conds ...field.Expr) ITgUserDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tgUserDo) Where(conds ...gen.Condition) ITgUserDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tgUserDo) Order(conds ...field.Expr) ITgUserDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tgUserDo) Distinct(cols ...field.Expr) ITgUserDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tgUserDo) Omit(cols ...field.Expr) ITgUserDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tgUserDo) Join(table schema.Tabler, on ...field.Expr) ITgUserDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tgUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITgUserDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tgUserDo) RightJoin(table schema.Tabler, on ...field.Expr) ITgUserDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tgUserDo) Group(cols ...field.Expr) ITgUserDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tgUserDo) Having(conds ...gen.Condition) ITgUserDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tgUserDo) Limit(limit int) ITgUserDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tgUserDo) Offset(offset int) ITgUserDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tgUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITgUserDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tgUserDo) Unscoped() ITgUserDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tgUserDo) Create(values ...*model.TgUser) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tgUserDo) CreateInBatches(values []*model.TgUser, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tgUserDo) Save(values ...*model.TgUser) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tgUserDo) First() (*model.TgUser, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TgUser), nil
	}
}

func (t tgUserDo) Take() (*model.TgUser, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TgUser), nil
	}
}

func (t tgUserDo) Last() (*model.TgUser, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TgUser), nil
	}
}

func (t tgUserDo) Find() ([]*model.TgUser, error) {
	result, err := t.DO.Find()
	return result.([]*model.TgUser), err
}

func (t tgUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TgUser, err error) {
	buf := make([]*model.TgUser, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tgUserDo) FindInBatches(result *[]*model.TgUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tgUserDo) Attrs(attrs ...field.AssignExpr) ITgUserDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tgUserDo) Assign(attrs ...field.AssignExpr) ITgUserDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tgUserDo) Joins(fields ...field.RelationField) ITgUserDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tgUserDo) Preload(fields ...field.RelationField) ITgUserDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tgUserDo) FirstOrInit() (*model.TgUser, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TgUser), nil
	}
}

func (t tgUserDo) FirstOrCreate() (*model.TgUser, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TgUser), nil
	}
}

func (t tgUserDo) FindByPage(offset int, limit int) (result []*model.TgUser, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tgUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tgUserDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tgUserDo) Delete(models ...*model.TgUser) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tgUserDo) withDO(do gen.Dao) *tgUserDo {
	t.DO = *do.(*gen.DO)
	return t
}