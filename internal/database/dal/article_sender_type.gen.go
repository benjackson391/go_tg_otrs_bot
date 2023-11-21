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

func newArticleSenderType(db *gorm.DB, opts ...gen.DOOption) articleSenderType {
	_articleSenderType := articleSenderType{}

	_articleSenderType.articleSenderTypeDo.UseDB(db, opts...)
	_articleSenderType.articleSenderTypeDo.UseModel(&model.ArticleSenderType{})

	tableName := _articleSenderType.articleSenderTypeDo.TableName()
	_articleSenderType.ALL = field.NewAsterisk(tableName)
	_articleSenderType.ID = field.NewInt32(tableName, "id")
	_articleSenderType.Name = field.NewString(tableName, "name")
	_articleSenderType.Comments = field.NewString(tableName, "comments")
	_articleSenderType.ValidID = field.NewInt32(tableName, "valid_id")
	_articleSenderType.CreateTime = field.NewTime(tableName, "create_time")
	_articleSenderType.CreateBy = field.NewInt32(tableName, "create_by")
	_articleSenderType.ChangeTime = field.NewTime(tableName, "change_time")
	_articleSenderType.ChangeBy = field.NewInt32(tableName, "change_by")

	_articleSenderType.fillFieldMap()

	return _articleSenderType
}

type articleSenderType struct {
	articleSenderTypeDo

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

func (a articleSenderType) Table(newTableName string) *articleSenderType {
	a.articleSenderTypeDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleSenderType) As(alias string) *articleSenderType {
	a.articleSenderTypeDo.DO = *(a.articleSenderTypeDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleSenderType) updateTableName(table string) *articleSenderType {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.Name = field.NewString(table, "name")
	a.Comments = field.NewString(table, "comments")
	a.ValidID = field.NewInt32(table, "valid_id")
	a.CreateTime = field.NewTime(table, "create_time")
	a.CreateBy = field.NewInt32(table, "create_by")
	a.ChangeTime = field.NewTime(table, "change_time")
	a.ChangeBy = field.NewInt32(table, "change_by")

	a.fillFieldMap()

	return a
}

func (a *articleSenderType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleSenderType) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["name"] = a.Name
	a.fieldMap["comments"] = a.Comments
	a.fieldMap["valid_id"] = a.ValidID
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["create_by"] = a.CreateBy
	a.fieldMap["change_time"] = a.ChangeTime
	a.fieldMap["change_by"] = a.ChangeBy
}

func (a articleSenderType) clone(db *gorm.DB) articleSenderType {
	a.articleSenderTypeDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleSenderType) replaceDB(db *gorm.DB) articleSenderType {
	a.articleSenderTypeDo.ReplaceDB(db)
	return a
}

type articleSenderTypeDo struct{ gen.DO }

type IArticleSenderTypeDo interface {
	gen.SubQuery
	Debug() IArticleSenderTypeDo
	WithContext(ctx context.Context) IArticleSenderTypeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArticleSenderTypeDo
	WriteDB() IArticleSenderTypeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArticleSenderTypeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArticleSenderTypeDo
	Not(conds ...gen.Condition) IArticleSenderTypeDo
	Or(conds ...gen.Condition) IArticleSenderTypeDo
	Select(conds ...field.Expr) IArticleSenderTypeDo
	Where(conds ...gen.Condition) IArticleSenderTypeDo
	Order(conds ...field.Expr) IArticleSenderTypeDo
	Distinct(cols ...field.Expr) IArticleSenderTypeDo
	Omit(cols ...field.Expr) IArticleSenderTypeDo
	Join(table schema.Tabler, on ...field.Expr) IArticleSenderTypeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArticleSenderTypeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArticleSenderTypeDo
	Group(cols ...field.Expr) IArticleSenderTypeDo
	Having(conds ...gen.Condition) IArticleSenderTypeDo
	Limit(limit int) IArticleSenderTypeDo
	Offset(offset int) IArticleSenderTypeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleSenderTypeDo
	Unscoped() IArticleSenderTypeDo
	Create(values ...*model.ArticleSenderType) error
	CreateInBatches(values []*model.ArticleSenderType, batchSize int) error
	Save(values ...*model.ArticleSenderType) error
	First() (*model.ArticleSenderType, error)
	Take() (*model.ArticleSenderType, error)
	Last() (*model.ArticleSenderType, error)
	Find() ([]*model.ArticleSenderType, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArticleSenderType, err error)
	FindInBatches(result *[]*model.ArticleSenderType, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ArticleSenderType) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArticleSenderTypeDo
	Assign(attrs ...field.AssignExpr) IArticleSenderTypeDo
	Joins(fields ...field.RelationField) IArticleSenderTypeDo
	Preload(fields ...field.RelationField) IArticleSenderTypeDo
	FirstOrInit() (*model.ArticleSenderType, error)
	FirstOrCreate() (*model.ArticleSenderType, error)
	FindByPage(offset int, limit int) (result []*model.ArticleSenderType, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArticleSenderTypeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a articleSenderTypeDo) Debug() IArticleSenderTypeDo {
	return a.withDO(a.DO.Debug())
}

func (a articleSenderTypeDo) WithContext(ctx context.Context) IArticleSenderTypeDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleSenderTypeDo) ReadDB() IArticleSenderTypeDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleSenderTypeDo) WriteDB() IArticleSenderTypeDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleSenderTypeDo) Session(config *gorm.Session) IArticleSenderTypeDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleSenderTypeDo) Clauses(conds ...clause.Expression) IArticleSenderTypeDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleSenderTypeDo) Returning(value interface{}, columns ...string) IArticleSenderTypeDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleSenderTypeDo) Not(conds ...gen.Condition) IArticleSenderTypeDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleSenderTypeDo) Or(conds ...gen.Condition) IArticleSenderTypeDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleSenderTypeDo) Select(conds ...field.Expr) IArticleSenderTypeDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleSenderTypeDo) Where(conds ...gen.Condition) IArticleSenderTypeDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleSenderTypeDo) Order(conds ...field.Expr) IArticleSenderTypeDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleSenderTypeDo) Distinct(cols ...field.Expr) IArticleSenderTypeDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleSenderTypeDo) Omit(cols ...field.Expr) IArticleSenderTypeDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleSenderTypeDo) Join(table schema.Tabler, on ...field.Expr) IArticleSenderTypeDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleSenderTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArticleSenderTypeDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleSenderTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) IArticleSenderTypeDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleSenderTypeDo) Group(cols ...field.Expr) IArticleSenderTypeDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleSenderTypeDo) Having(conds ...gen.Condition) IArticleSenderTypeDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleSenderTypeDo) Limit(limit int) IArticleSenderTypeDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleSenderTypeDo) Offset(offset int) IArticleSenderTypeDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleSenderTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleSenderTypeDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleSenderTypeDo) Unscoped() IArticleSenderTypeDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleSenderTypeDo) Create(values ...*model.ArticleSenderType) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleSenderTypeDo) CreateInBatches(values []*model.ArticleSenderType, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleSenderTypeDo) Save(values ...*model.ArticleSenderType) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleSenderTypeDo) First() (*model.ArticleSenderType, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleSenderType), nil
	}
}

func (a articleSenderTypeDo) Take() (*model.ArticleSenderType, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleSenderType), nil
	}
}

func (a articleSenderTypeDo) Last() (*model.ArticleSenderType, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleSenderType), nil
	}
}

func (a articleSenderTypeDo) Find() ([]*model.ArticleSenderType, error) {
	result, err := a.DO.Find()
	return result.([]*model.ArticleSenderType), err
}

func (a articleSenderTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArticleSenderType, err error) {
	buf := make([]*model.ArticleSenderType, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleSenderTypeDo) FindInBatches(result *[]*model.ArticleSenderType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleSenderTypeDo) Attrs(attrs ...field.AssignExpr) IArticleSenderTypeDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleSenderTypeDo) Assign(attrs ...field.AssignExpr) IArticleSenderTypeDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleSenderTypeDo) Joins(fields ...field.RelationField) IArticleSenderTypeDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleSenderTypeDo) Preload(fields ...field.RelationField) IArticleSenderTypeDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleSenderTypeDo) FirstOrInit() (*model.ArticleSenderType, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleSenderType), nil
	}
}

func (a articleSenderTypeDo) FirstOrCreate() (*model.ArticleSenderType, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleSenderType), nil
	}
}

func (a articleSenderTypeDo) FindByPage(offset int, limit int) (result []*model.ArticleSenderType, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a articleSenderTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleSenderTypeDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleSenderTypeDo) Delete(models ...*model.ArticleSenderType) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleSenderTypeDo) withDO(do gen.Dao) *articleSenderTypeDo {
	a.DO = *do.(*gen.DO)
	return a
}