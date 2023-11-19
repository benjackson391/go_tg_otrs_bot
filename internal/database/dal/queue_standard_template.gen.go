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

func newQueueStandardTemplate(db *gorm.DB, opts ...gen.DOOption) queueStandardTemplate {
	_queueStandardTemplate := queueStandardTemplate{}

	_queueStandardTemplate.queueStandardTemplateDo.UseDB(db, opts...)
	_queueStandardTemplate.queueStandardTemplateDo.UseModel(&model.QueueStandardTemplate{})

	tableName := _queueStandardTemplate.queueStandardTemplateDo.TableName()
	_queueStandardTemplate.ALL = field.NewAsterisk(tableName)
	_queueStandardTemplate.QueueID = field.NewInt32(tableName, "queue_id")
	_queueStandardTemplate.StandardTemplateID = field.NewInt32(tableName, "standard_template_id")
	_queueStandardTemplate.CreateTime = field.NewTime(tableName, "create_time")
	_queueStandardTemplate.CreateBy = field.NewInt32(tableName, "create_by")
	_queueStandardTemplate.ChangeTime = field.NewTime(tableName, "change_time")
	_queueStandardTemplate.ChangeBy = field.NewInt32(tableName, "change_by")

	_queueStandardTemplate.fillFieldMap()

	return _queueStandardTemplate
}

type queueStandardTemplate struct {
	queueStandardTemplateDo

	ALL                field.Asterisk
	QueueID            field.Int32
	StandardTemplateID field.Int32
	CreateTime         field.Time
	CreateBy           field.Int32
	ChangeTime         field.Time
	ChangeBy           field.Int32

	fieldMap map[string]field.Expr
}

func (q queueStandardTemplate) Table(newTableName string) *queueStandardTemplate {
	q.queueStandardTemplateDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q queueStandardTemplate) As(alias string) *queueStandardTemplate {
	q.queueStandardTemplateDo.DO = *(q.queueStandardTemplateDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *queueStandardTemplate) updateTableName(table string) *queueStandardTemplate {
	q.ALL = field.NewAsterisk(table)
	q.QueueID = field.NewInt32(table, "queue_id")
	q.StandardTemplateID = field.NewInt32(table, "standard_template_id")
	q.CreateTime = field.NewTime(table, "create_time")
	q.CreateBy = field.NewInt32(table, "create_by")
	q.ChangeTime = field.NewTime(table, "change_time")
	q.ChangeBy = field.NewInt32(table, "change_by")

	q.fillFieldMap()

	return q
}

func (q *queueStandardTemplate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *queueStandardTemplate) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 6)
	q.fieldMap["queue_id"] = q.QueueID
	q.fieldMap["standard_template_id"] = q.StandardTemplateID
	q.fieldMap["create_time"] = q.CreateTime
	q.fieldMap["create_by"] = q.CreateBy
	q.fieldMap["change_time"] = q.ChangeTime
	q.fieldMap["change_by"] = q.ChangeBy
}

func (q queueStandardTemplate) clone(db *gorm.DB) queueStandardTemplate {
	q.queueStandardTemplateDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q queueStandardTemplate) replaceDB(db *gorm.DB) queueStandardTemplate {
	q.queueStandardTemplateDo.ReplaceDB(db)
	return q
}

type queueStandardTemplateDo struct{ gen.DO }

type IQueueStandardTemplateDo interface {
	gen.SubQuery
	Debug() IQueueStandardTemplateDo
	WithContext(ctx context.Context) IQueueStandardTemplateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IQueueStandardTemplateDo
	WriteDB() IQueueStandardTemplateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IQueueStandardTemplateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IQueueStandardTemplateDo
	Not(conds ...gen.Condition) IQueueStandardTemplateDo
	Or(conds ...gen.Condition) IQueueStandardTemplateDo
	Select(conds ...field.Expr) IQueueStandardTemplateDo
	Where(conds ...gen.Condition) IQueueStandardTemplateDo
	Order(conds ...field.Expr) IQueueStandardTemplateDo
	Distinct(cols ...field.Expr) IQueueStandardTemplateDo
	Omit(cols ...field.Expr) IQueueStandardTemplateDo
	Join(table schema.Tabler, on ...field.Expr) IQueueStandardTemplateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IQueueStandardTemplateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IQueueStandardTemplateDo
	Group(cols ...field.Expr) IQueueStandardTemplateDo
	Having(conds ...gen.Condition) IQueueStandardTemplateDo
	Limit(limit int) IQueueStandardTemplateDo
	Offset(offset int) IQueueStandardTemplateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IQueueStandardTemplateDo
	Unscoped() IQueueStandardTemplateDo
	Create(values ...*model.QueueStandardTemplate) error
	CreateInBatches(values []*model.QueueStandardTemplate, batchSize int) error
	Save(values ...*model.QueueStandardTemplate) error
	First() (*model.QueueStandardTemplate, error)
	Take() (*model.QueueStandardTemplate, error)
	Last() (*model.QueueStandardTemplate, error)
	Find() ([]*model.QueueStandardTemplate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QueueStandardTemplate, err error)
	FindInBatches(result *[]*model.QueueStandardTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.QueueStandardTemplate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IQueueStandardTemplateDo
	Assign(attrs ...field.AssignExpr) IQueueStandardTemplateDo
	Joins(fields ...field.RelationField) IQueueStandardTemplateDo
	Preload(fields ...field.RelationField) IQueueStandardTemplateDo
	FirstOrInit() (*model.QueueStandardTemplate, error)
	FirstOrCreate() (*model.QueueStandardTemplate, error)
	FindByPage(offset int, limit int) (result []*model.QueueStandardTemplate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IQueueStandardTemplateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (q queueStandardTemplateDo) Debug() IQueueStandardTemplateDo {
	return q.withDO(q.DO.Debug())
}

func (q queueStandardTemplateDo) WithContext(ctx context.Context) IQueueStandardTemplateDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q queueStandardTemplateDo) ReadDB() IQueueStandardTemplateDo {
	return q.Clauses(dbresolver.Read)
}

func (q queueStandardTemplateDo) WriteDB() IQueueStandardTemplateDo {
	return q.Clauses(dbresolver.Write)
}

func (q queueStandardTemplateDo) Session(config *gorm.Session) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Session(config))
}

func (q queueStandardTemplateDo) Clauses(conds ...clause.Expression) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q queueStandardTemplateDo) Returning(value interface{}, columns ...string) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q queueStandardTemplateDo) Not(conds ...gen.Condition) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q queueStandardTemplateDo) Or(conds ...gen.Condition) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q queueStandardTemplateDo) Select(conds ...field.Expr) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q queueStandardTemplateDo) Where(conds ...gen.Condition) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q queueStandardTemplateDo) Order(conds ...field.Expr) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q queueStandardTemplateDo) Distinct(cols ...field.Expr) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q queueStandardTemplateDo) Omit(cols ...field.Expr) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q queueStandardTemplateDo) Join(table schema.Tabler, on ...field.Expr) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q queueStandardTemplateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IQueueStandardTemplateDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q queueStandardTemplateDo) RightJoin(table schema.Tabler, on ...field.Expr) IQueueStandardTemplateDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q queueStandardTemplateDo) Group(cols ...field.Expr) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q queueStandardTemplateDo) Having(conds ...gen.Condition) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q queueStandardTemplateDo) Limit(limit int) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q queueStandardTemplateDo) Offset(offset int) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q queueStandardTemplateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q queueStandardTemplateDo) Unscoped() IQueueStandardTemplateDo {
	return q.withDO(q.DO.Unscoped())
}

func (q queueStandardTemplateDo) Create(values ...*model.QueueStandardTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q queueStandardTemplateDo) CreateInBatches(values []*model.QueueStandardTemplate, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q queueStandardTemplateDo) Save(values ...*model.QueueStandardTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q queueStandardTemplateDo) First() (*model.QueueStandardTemplate, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueueStandardTemplate), nil
	}
}

func (q queueStandardTemplateDo) Take() (*model.QueueStandardTemplate, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueueStandardTemplate), nil
	}
}

func (q queueStandardTemplateDo) Last() (*model.QueueStandardTemplate, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueueStandardTemplate), nil
	}
}

func (q queueStandardTemplateDo) Find() ([]*model.QueueStandardTemplate, error) {
	result, err := q.DO.Find()
	return result.([]*model.QueueStandardTemplate), err
}

func (q queueStandardTemplateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QueueStandardTemplate, err error) {
	buf := make([]*model.QueueStandardTemplate, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q queueStandardTemplateDo) FindInBatches(result *[]*model.QueueStandardTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q queueStandardTemplateDo) Attrs(attrs ...field.AssignExpr) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q queueStandardTemplateDo) Assign(attrs ...field.AssignExpr) IQueueStandardTemplateDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q queueStandardTemplateDo) Joins(fields ...field.RelationField) IQueueStandardTemplateDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q queueStandardTemplateDo) Preload(fields ...field.RelationField) IQueueStandardTemplateDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q queueStandardTemplateDo) FirstOrInit() (*model.QueueStandardTemplate, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueueStandardTemplate), nil
	}
}

func (q queueStandardTemplateDo) FirstOrCreate() (*model.QueueStandardTemplate, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueueStandardTemplate), nil
	}
}

func (q queueStandardTemplateDo) FindByPage(offset int, limit int) (result []*model.QueueStandardTemplate, count int64, err error) {
	result, err = q.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = q.Offset(-1).Limit(-1).Count()
	return
}

func (q queueStandardTemplateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q queueStandardTemplateDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q queueStandardTemplateDo) Delete(models ...*model.QueueStandardTemplate) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *queueStandardTemplateDo) withDO(do gen.Dao) *queueStandardTemplateDo {
	q.DO = *do.(*gen.DO)
	return q
}
