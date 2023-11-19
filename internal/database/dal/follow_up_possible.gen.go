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

func newFollowUpPossible(db *gorm.DB, opts ...gen.DOOption) followUpPossible {
	_followUpPossible := followUpPossible{}

	_followUpPossible.followUpPossibleDo.UseDB(db, opts...)
	_followUpPossible.followUpPossibleDo.UseModel(&model.FollowUpPossible{})

	tableName := _followUpPossible.followUpPossibleDo.TableName()
	_followUpPossible.ALL = field.NewAsterisk(tableName)
	_followUpPossible.ID = field.NewInt32(tableName, "id")
	_followUpPossible.Name = field.NewString(tableName, "name")
	_followUpPossible.Comments = field.NewString(tableName, "comments")
	_followUpPossible.ValidID = field.NewInt32(tableName, "valid_id")
	_followUpPossible.CreateTime = field.NewTime(tableName, "create_time")
	_followUpPossible.CreateBy = field.NewInt32(tableName, "create_by")
	_followUpPossible.ChangeTime = field.NewTime(tableName, "change_time")
	_followUpPossible.ChangeBy = field.NewInt32(tableName, "change_by")

	_followUpPossible.fillFieldMap()

	return _followUpPossible
}

type followUpPossible struct {
	followUpPossibleDo

	ALL        field.Asterisk
	ID         field.Int32
	Name       field.String
	Comments   field.String
	ValidID    field.Int32
	CreateTime field.Time
	CreateBy   field.Int32
	ChangeTime field.Time
	ChangeBy   field.Int32

	fieldMap map[string]field.Expr
}

func (f followUpPossible) Table(newTableName string) *followUpPossible {
	f.followUpPossibleDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f followUpPossible) As(alias string) *followUpPossible {
	f.followUpPossibleDo.DO = *(f.followUpPossibleDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *followUpPossible) updateTableName(table string) *followUpPossible {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt32(table, "id")
	f.Name = field.NewString(table, "name")
	f.Comments = field.NewString(table, "comments")
	f.ValidID = field.NewInt32(table, "valid_id")
	f.CreateTime = field.NewTime(table, "create_time")
	f.CreateBy = field.NewInt32(table, "create_by")
	f.ChangeTime = field.NewTime(table, "change_time")
	f.ChangeBy = field.NewInt32(table, "change_by")

	f.fillFieldMap()

	return f
}

func (f *followUpPossible) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *followUpPossible) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 8)
	f.fieldMap["id"] = f.ID
	f.fieldMap["name"] = f.Name
	f.fieldMap["comments"] = f.Comments
	f.fieldMap["valid_id"] = f.ValidID
	f.fieldMap["create_time"] = f.CreateTime
	f.fieldMap["create_by"] = f.CreateBy
	f.fieldMap["change_time"] = f.ChangeTime
	f.fieldMap["change_by"] = f.ChangeBy
}

func (f followUpPossible) clone(db *gorm.DB) followUpPossible {
	f.followUpPossibleDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f followUpPossible) replaceDB(db *gorm.DB) followUpPossible {
	f.followUpPossibleDo.ReplaceDB(db)
	return f
}

type followUpPossibleDo struct{ gen.DO }

type IFollowUpPossibleDo interface {
	gen.SubQuery
	Debug() IFollowUpPossibleDo
	WithContext(ctx context.Context) IFollowUpPossibleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFollowUpPossibleDo
	WriteDB() IFollowUpPossibleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFollowUpPossibleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFollowUpPossibleDo
	Not(conds ...gen.Condition) IFollowUpPossibleDo
	Or(conds ...gen.Condition) IFollowUpPossibleDo
	Select(conds ...field.Expr) IFollowUpPossibleDo
	Where(conds ...gen.Condition) IFollowUpPossibleDo
	Order(conds ...field.Expr) IFollowUpPossibleDo
	Distinct(cols ...field.Expr) IFollowUpPossibleDo
	Omit(cols ...field.Expr) IFollowUpPossibleDo
	Join(table schema.Tabler, on ...field.Expr) IFollowUpPossibleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFollowUpPossibleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFollowUpPossibleDo
	Group(cols ...field.Expr) IFollowUpPossibleDo
	Having(conds ...gen.Condition) IFollowUpPossibleDo
	Limit(limit int) IFollowUpPossibleDo
	Offset(offset int) IFollowUpPossibleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFollowUpPossibleDo
	Unscoped() IFollowUpPossibleDo
	Create(values ...*model.FollowUpPossible) error
	CreateInBatches(values []*model.FollowUpPossible, batchSize int) error
	Save(values ...*model.FollowUpPossible) error
	First() (*model.FollowUpPossible, error)
	Take() (*model.FollowUpPossible, error)
	Last() (*model.FollowUpPossible, error)
	Find() ([]*model.FollowUpPossible, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FollowUpPossible, err error)
	FindInBatches(result *[]*model.FollowUpPossible, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FollowUpPossible) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFollowUpPossibleDo
	Assign(attrs ...field.AssignExpr) IFollowUpPossibleDo
	Joins(fields ...field.RelationField) IFollowUpPossibleDo
	Preload(fields ...field.RelationField) IFollowUpPossibleDo
	FirstOrInit() (*model.FollowUpPossible, error)
	FirstOrCreate() (*model.FollowUpPossible, error)
	FindByPage(offset int, limit int) (result []*model.FollowUpPossible, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFollowUpPossibleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f followUpPossibleDo) Debug() IFollowUpPossibleDo {
	return f.withDO(f.DO.Debug())
}

func (f followUpPossibleDo) WithContext(ctx context.Context) IFollowUpPossibleDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f followUpPossibleDo) ReadDB() IFollowUpPossibleDo {
	return f.Clauses(dbresolver.Read)
}

func (f followUpPossibleDo) WriteDB() IFollowUpPossibleDo {
	return f.Clauses(dbresolver.Write)
}

func (f followUpPossibleDo) Session(config *gorm.Session) IFollowUpPossibleDo {
	return f.withDO(f.DO.Session(config))
}

func (f followUpPossibleDo) Clauses(conds ...clause.Expression) IFollowUpPossibleDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f followUpPossibleDo) Returning(value interface{}, columns ...string) IFollowUpPossibleDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f followUpPossibleDo) Not(conds ...gen.Condition) IFollowUpPossibleDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f followUpPossibleDo) Or(conds ...gen.Condition) IFollowUpPossibleDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f followUpPossibleDo) Select(conds ...field.Expr) IFollowUpPossibleDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f followUpPossibleDo) Where(conds ...gen.Condition) IFollowUpPossibleDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f followUpPossibleDo) Order(conds ...field.Expr) IFollowUpPossibleDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f followUpPossibleDo) Distinct(cols ...field.Expr) IFollowUpPossibleDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f followUpPossibleDo) Omit(cols ...field.Expr) IFollowUpPossibleDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f followUpPossibleDo) Join(table schema.Tabler, on ...field.Expr) IFollowUpPossibleDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f followUpPossibleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFollowUpPossibleDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f followUpPossibleDo) RightJoin(table schema.Tabler, on ...field.Expr) IFollowUpPossibleDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f followUpPossibleDo) Group(cols ...field.Expr) IFollowUpPossibleDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f followUpPossibleDo) Having(conds ...gen.Condition) IFollowUpPossibleDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f followUpPossibleDo) Limit(limit int) IFollowUpPossibleDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f followUpPossibleDo) Offset(offset int) IFollowUpPossibleDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f followUpPossibleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFollowUpPossibleDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f followUpPossibleDo) Unscoped() IFollowUpPossibleDo {
	return f.withDO(f.DO.Unscoped())
}

func (f followUpPossibleDo) Create(values ...*model.FollowUpPossible) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f followUpPossibleDo) CreateInBatches(values []*model.FollowUpPossible, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f followUpPossibleDo) Save(values ...*model.FollowUpPossible) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f followUpPossibleDo) First() (*model.FollowUpPossible, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowUpPossible), nil
	}
}

func (f followUpPossibleDo) Take() (*model.FollowUpPossible, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowUpPossible), nil
	}
}

func (f followUpPossibleDo) Last() (*model.FollowUpPossible, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowUpPossible), nil
	}
}

func (f followUpPossibleDo) Find() ([]*model.FollowUpPossible, error) {
	result, err := f.DO.Find()
	return result.([]*model.FollowUpPossible), err
}

func (f followUpPossibleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FollowUpPossible, err error) {
	buf := make([]*model.FollowUpPossible, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f followUpPossibleDo) FindInBatches(result *[]*model.FollowUpPossible, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f followUpPossibleDo) Attrs(attrs ...field.AssignExpr) IFollowUpPossibleDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f followUpPossibleDo) Assign(attrs ...field.AssignExpr) IFollowUpPossibleDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f followUpPossibleDo) Joins(fields ...field.RelationField) IFollowUpPossibleDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f followUpPossibleDo) Preload(fields ...field.RelationField) IFollowUpPossibleDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f followUpPossibleDo) FirstOrInit() (*model.FollowUpPossible, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowUpPossible), nil
	}
}

func (f followUpPossibleDo) FirstOrCreate() (*model.FollowUpPossible, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowUpPossible), nil
	}
}

func (f followUpPossibleDo) FindByPage(offset int, limit int) (result []*model.FollowUpPossible, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f followUpPossibleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f followUpPossibleDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f followUpPossibleDo) Delete(models ...*model.FollowUpPossible) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *followUpPossibleDo) withDO(do gen.Dao) *followUpPossibleDo {
	f.DO = *do.(*gen.DO)
	return f
}
