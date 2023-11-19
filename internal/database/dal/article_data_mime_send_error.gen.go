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

func newArticleDataMimeSendError(db *gorm.DB, opts ...gen.DOOption) articleDataMimeSendError {
	_articleDataMimeSendError := articleDataMimeSendError{}

	_articleDataMimeSendError.articleDataMimeSendErrorDo.UseDB(db, opts...)
	_articleDataMimeSendError.articleDataMimeSendErrorDo.UseModel(&model.ArticleDataMimeSendError{})

	tableName := _articleDataMimeSendError.articleDataMimeSendErrorDo.TableName()
	_articleDataMimeSendError.ALL = field.NewAsterisk(tableName)
	_articleDataMimeSendError.ID = field.NewInt64(tableName, "id")
	_articleDataMimeSendError.ArticleID = field.NewInt64(tableName, "article_id")
	_articleDataMimeSendError.MessageID = field.NewString(tableName, "message_id")
	_articleDataMimeSendError.LogMessage = field.NewString(tableName, "log_message")
	_articleDataMimeSendError.CreateTime = field.NewTime(tableName, "create_time")

	_articleDataMimeSendError.fillFieldMap()

	return _articleDataMimeSendError
}

type articleDataMimeSendError struct {
	articleDataMimeSendErrorDo

	ALL        field.Asterisk
	ID         field.Int64
	ArticleID  field.Int64
	MessageID  field.String
	LogMessage field.String
	CreateTime field.Time

	fieldMap map[string]field.Expr
}

func (a articleDataMimeSendError) Table(newTableName string) *articleDataMimeSendError {
	a.articleDataMimeSendErrorDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleDataMimeSendError) As(alias string) *articleDataMimeSendError {
	a.articleDataMimeSendErrorDo.DO = *(a.articleDataMimeSendErrorDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleDataMimeSendError) updateTableName(table string) *articleDataMimeSendError {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.ArticleID = field.NewInt64(table, "article_id")
	a.MessageID = field.NewString(table, "message_id")
	a.LogMessage = field.NewString(table, "log_message")
	a.CreateTime = field.NewTime(table, "create_time")

	a.fillFieldMap()

	return a
}

func (a *articleDataMimeSendError) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleDataMimeSendError) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 5)
	a.fieldMap["id"] = a.ID
	a.fieldMap["article_id"] = a.ArticleID
	a.fieldMap["message_id"] = a.MessageID
	a.fieldMap["log_message"] = a.LogMessage
	a.fieldMap["create_time"] = a.CreateTime
}

func (a articleDataMimeSendError) clone(db *gorm.DB) articleDataMimeSendError {
	a.articleDataMimeSendErrorDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleDataMimeSendError) replaceDB(db *gorm.DB) articleDataMimeSendError {
	a.articleDataMimeSendErrorDo.ReplaceDB(db)
	return a
}

type articleDataMimeSendErrorDo struct{ gen.DO }

type IArticleDataMimeSendErrorDo interface {
	gen.SubQuery
	Debug() IArticleDataMimeSendErrorDo
	WithContext(ctx context.Context) IArticleDataMimeSendErrorDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArticleDataMimeSendErrorDo
	WriteDB() IArticleDataMimeSendErrorDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArticleDataMimeSendErrorDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArticleDataMimeSendErrorDo
	Not(conds ...gen.Condition) IArticleDataMimeSendErrorDo
	Or(conds ...gen.Condition) IArticleDataMimeSendErrorDo
	Select(conds ...field.Expr) IArticleDataMimeSendErrorDo
	Where(conds ...gen.Condition) IArticleDataMimeSendErrorDo
	Order(conds ...field.Expr) IArticleDataMimeSendErrorDo
	Distinct(cols ...field.Expr) IArticleDataMimeSendErrorDo
	Omit(cols ...field.Expr) IArticleDataMimeSendErrorDo
	Join(table schema.Tabler, on ...field.Expr) IArticleDataMimeSendErrorDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArticleDataMimeSendErrorDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArticleDataMimeSendErrorDo
	Group(cols ...field.Expr) IArticleDataMimeSendErrorDo
	Having(conds ...gen.Condition) IArticleDataMimeSendErrorDo
	Limit(limit int) IArticleDataMimeSendErrorDo
	Offset(offset int) IArticleDataMimeSendErrorDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleDataMimeSendErrorDo
	Unscoped() IArticleDataMimeSendErrorDo
	Create(values ...*model.ArticleDataMimeSendError) error
	CreateInBatches(values []*model.ArticleDataMimeSendError, batchSize int) error
	Save(values ...*model.ArticleDataMimeSendError) error
	First() (*model.ArticleDataMimeSendError, error)
	Take() (*model.ArticleDataMimeSendError, error)
	Last() (*model.ArticleDataMimeSendError, error)
	Find() ([]*model.ArticleDataMimeSendError, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArticleDataMimeSendError, err error)
	FindInBatches(result *[]*model.ArticleDataMimeSendError, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ArticleDataMimeSendError) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArticleDataMimeSendErrorDo
	Assign(attrs ...field.AssignExpr) IArticleDataMimeSendErrorDo
	Joins(fields ...field.RelationField) IArticleDataMimeSendErrorDo
	Preload(fields ...field.RelationField) IArticleDataMimeSendErrorDo
	FirstOrInit() (*model.ArticleDataMimeSendError, error)
	FirstOrCreate() (*model.ArticleDataMimeSendError, error)
	FindByPage(offset int, limit int) (result []*model.ArticleDataMimeSendError, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArticleDataMimeSendErrorDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a articleDataMimeSendErrorDo) Debug() IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Debug())
}

func (a articleDataMimeSendErrorDo) WithContext(ctx context.Context) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleDataMimeSendErrorDo) ReadDB() IArticleDataMimeSendErrorDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleDataMimeSendErrorDo) WriteDB() IArticleDataMimeSendErrorDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleDataMimeSendErrorDo) Session(config *gorm.Session) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleDataMimeSendErrorDo) Clauses(conds ...clause.Expression) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleDataMimeSendErrorDo) Returning(value interface{}, columns ...string) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleDataMimeSendErrorDo) Not(conds ...gen.Condition) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleDataMimeSendErrorDo) Or(conds ...gen.Condition) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleDataMimeSendErrorDo) Select(conds ...field.Expr) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleDataMimeSendErrorDo) Where(conds ...gen.Condition) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleDataMimeSendErrorDo) Order(conds ...field.Expr) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleDataMimeSendErrorDo) Distinct(cols ...field.Expr) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleDataMimeSendErrorDo) Omit(cols ...field.Expr) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleDataMimeSendErrorDo) Join(table schema.Tabler, on ...field.Expr) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleDataMimeSendErrorDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleDataMimeSendErrorDo) RightJoin(table schema.Tabler, on ...field.Expr) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleDataMimeSendErrorDo) Group(cols ...field.Expr) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleDataMimeSendErrorDo) Having(conds ...gen.Condition) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleDataMimeSendErrorDo) Limit(limit int) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleDataMimeSendErrorDo) Offset(offset int) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleDataMimeSendErrorDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleDataMimeSendErrorDo) Unscoped() IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleDataMimeSendErrorDo) Create(values ...*model.ArticleDataMimeSendError) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleDataMimeSendErrorDo) CreateInBatches(values []*model.ArticleDataMimeSendError, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleDataMimeSendErrorDo) Save(values ...*model.ArticleDataMimeSendError) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleDataMimeSendErrorDo) First() (*model.ArticleDataMimeSendError, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleDataMimeSendError), nil
	}
}

func (a articleDataMimeSendErrorDo) Take() (*model.ArticleDataMimeSendError, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleDataMimeSendError), nil
	}
}

func (a articleDataMimeSendErrorDo) Last() (*model.ArticleDataMimeSendError, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleDataMimeSendError), nil
	}
}

func (a articleDataMimeSendErrorDo) Find() ([]*model.ArticleDataMimeSendError, error) {
	result, err := a.DO.Find()
	return result.([]*model.ArticleDataMimeSendError), err
}

func (a articleDataMimeSendErrorDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArticleDataMimeSendError, err error) {
	buf := make([]*model.ArticleDataMimeSendError, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleDataMimeSendErrorDo) FindInBatches(result *[]*model.ArticleDataMimeSendError, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleDataMimeSendErrorDo) Attrs(attrs ...field.AssignExpr) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleDataMimeSendErrorDo) Assign(attrs ...field.AssignExpr) IArticleDataMimeSendErrorDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleDataMimeSendErrorDo) Joins(fields ...field.RelationField) IArticleDataMimeSendErrorDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleDataMimeSendErrorDo) Preload(fields ...field.RelationField) IArticleDataMimeSendErrorDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleDataMimeSendErrorDo) FirstOrInit() (*model.ArticleDataMimeSendError, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleDataMimeSendError), nil
	}
}

func (a articleDataMimeSendErrorDo) FirstOrCreate() (*model.ArticleDataMimeSendError, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArticleDataMimeSendError), nil
	}
}

func (a articleDataMimeSendErrorDo) FindByPage(offset int, limit int) (result []*model.ArticleDataMimeSendError, count int64, err error) {
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

func (a articleDataMimeSendErrorDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleDataMimeSendErrorDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleDataMimeSendErrorDo) Delete(models ...*model.ArticleDataMimeSendError) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleDataMimeSendErrorDo) withDO(do gen.Dao) *articleDataMimeSendErrorDo {
	a.DO = *do.(*gen.DO)
	return a
}
