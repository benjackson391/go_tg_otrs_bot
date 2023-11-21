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

func newNotificationEventMessage(db *gorm.DB, opts ...gen.DOOption) notificationEventMessage {
	_notificationEventMessage := notificationEventMessage{}

	_notificationEventMessage.notificationEventMessageDo.UseDB(db, opts...)
	_notificationEventMessage.notificationEventMessageDo.UseModel(&model.NotificationEventMessage{})

	tableName := _notificationEventMessage.notificationEventMessageDo.TableName()
	_notificationEventMessage.ALL = field.NewAsterisk(tableName)
	_notificationEventMessage.ID = field.NewInt32(tableName, "id")
	_notificationEventMessage.NotificationID = field.NewInt32(tableName, "notification_id")
	_notificationEventMessage.Subject = field.NewString(tableName, "subject")
	_notificationEventMessage.Text = field.NewString(tableName, "text")
	_notificationEventMessage.ContentType = field.NewString(tableName, "content_type")
	_notificationEventMessage.Language = field.NewString(tableName, "language")

	_notificationEventMessage.fillFieldMap()

	return _notificationEventMessage
}

type notificationEventMessage struct {
	notificationEventMessageDo

	ALL            field.Asterisk
	ID             field.Int32
	NotificationID field.Int32
	Subject        field.String
	Text           field.String
	ContentType    field.String
	Language       field.String

	fieldMap map[string]field.Expr
}

func (n notificationEventMessage) Table(newTableName string) *notificationEventMessage {
	n.notificationEventMessageDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n notificationEventMessage) As(alias string) *notificationEventMessage {
	n.notificationEventMessageDo.DO = *(n.notificationEventMessageDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *notificationEventMessage) updateTableName(table string) *notificationEventMessage {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewInt32(table, "id")
	n.NotificationID = field.NewInt32(table, "notification_id")
	n.Subject = field.NewString(table, "subject")
	n.Text = field.NewString(table, "text")
	n.ContentType = field.NewString(table, "content_type")
	n.Language = field.NewString(table, "language")

	n.fillFieldMap()

	return n
}

func (n *notificationEventMessage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *notificationEventMessage) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 6)
	n.fieldMap["id"] = n.ID
	n.fieldMap["notification_id"] = n.NotificationID
	n.fieldMap["subject"] = n.Subject
	n.fieldMap["text"] = n.Text
	n.fieldMap["content_type"] = n.ContentType
	n.fieldMap["language"] = n.Language
}

func (n notificationEventMessage) clone(db *gorm.DB) notificationEventMessage {
	n.notificationEventMessageDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n notificationEventMessage) replaceDB(db *gorm.DB) notificationEventMessage {
	n.notificationEventMessageDo.ReplaceDB(db)
	return n
}

type notificationEventMessageDo struct{ gen.DO }

type INotificationEventMessageDo interface {
	gen.SubQuery
	Debug() INotificationEventMessageDo
	WithContext(ctx context.Context) INotificationEventMessageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() INotificationEventMessageDo
	WriteDB() INotificationEventMessageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) INotificationEventMessageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INotificationEventMessageDo
	Not(conds ...gen.Condition) INotificationEventMessageDo
	Or(conds ...gen.Condition) INotificationEventMessageDo
	Select(conds ...field.Expr) INotificationEventMessageDo
	Where(conds ...gen.Condition) INotificationEventMessageDo
	Order(conds ...field.Expr) INotificationEventMessageDo
	Distinct(cols ...field.Expr) INotificationEventMessageDo
	Omit(cols ...field.Expr) INotificationEventMessageDo
	Join(table schema.Tabler, on ...field.Expr) INotificationEventMessageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INotificationEventMessageDo
	RightJoin(table schema.Tabler, on ...field.Expr) INotificationEventMessageDo
	Group(cols ...field.Expr) INotificationEventMessageDo
	Having(conds ...gen.Condition) INotificationEventMessageDo
	Limit(limit int) INotificationEventMessageDo
	Offset(offset int) INotificationEventMessageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INotificationEventMessageDo
	Unscoped() INotificationEventMessageDo
	Create(values ...*model.NotificationEventMessage) error
	CreateInBatches(values []*model.NotificationEventMessage, batchSize int) error
	Save(values ...*model.NotificationEventMessage) error
	First() (*model.NotificationEventMessage, error)
	Take() (*model.NotificationEventMessage, error)
	Last() (*model.NotificationEventMessage, error)
	Find() ([]*model.NotificationEventMessage, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NotificationEventMessage, err error)
	FindInBatches(result *[]*model.NotificationEventMessage, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.NotificationEventMessage) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INotificationEventMessageDo
	Assign(attrs ...field.AssignExpr) INotificationEventMessageDo
	Joins(fields ...field.RelationField) INotificationEventMessageDo
	Preload(fields ...field.RelationField) INotificationEventMessageDo
	FirstOrInit() (*model.NotificationEventMessage, error)
	FirstOrCreate() (*model.NotificationEventMessage, error)
	FindByPage(offset int, limit int) (result []*model.NotificationEventMessage, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INotificationEventMessageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n notificationEventMessageDo) Debug() INotificationEventMessageDo {
	return n.withDO(n.DO.Debug())
}

func (n notificationEventMessageDo) WithContext(ctx context.Context) INotificationEventMessageDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n notificationEventMessageDo) ReadDB() INotificationEventMessageDo {
	return n.Clauses(dbresolver.Read)
}

func (n notificationEventMessageDo) WriteDB() INotificationEventMessageDo {
	return n.Clauses(dbresolver.Write)
}

func (n notificationEventMessageDo) Session(config *gorm.Session) INotificationEventMessageDo {
	return n.withDO(n.DO.Session(config))
}

func (n notificationEventMessageDo) Clauses(conds ...clause.Expression) INotificationEventMessageDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n notificationEventMessageDo) Returning(value interface{}, columns ...string) INotificationEventMessageDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n notificationEventMessageDo) Not(conds ...gen.Condition) INotificationEventMessageDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n notificationEventMessageDo) Or(conds ...gen.Condition) INotificationEventMessageDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n notificationEventMessageDo) Select(conds ...field.Expr) INotificationEventMessageDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n notificationEventMessageDo) Where(conds ...gen.Condition) INotificationEventMessageDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n notificationEventMessageDo) Order(conds ...field.Expr) INotificationEventMessageDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n notificationEventMessageDo) Distinct(cols ...field.Expr) INotificationEventMessageDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n notificationEventMessageDo) Omit(cols ...field.Expr) INotificationEventMessageDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n notificationEventMessageDo) Join(table schema.Tabler, on ...field.Expr) INotificationEventMessageDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n notificationEventMessageDo) LeftJoin(table schema.Tabler, on ...field.Expr) INotificationEventMessageDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n notificationEventMessageDo) RightJoin(table schema.Tabler, on ...field.Expr) INotificationEventMessageDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n notificationEventMessageDo) Group(cols ...field.Expr) INotificationEventMessageDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n notificationEventMessageDo) Having(conds ...gen.Condition) INotificationEventMessageDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n notificationEventMessageDo) Limit(limit int) INotificationEventMessageDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n notificationEventMessageDo) Offset(offset int) INotificationEventMessageDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n notificationEventMessageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INotificationEventMessageDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n notificationEventMessageDo) Unscoped() INotificationEventMessageDo {
	return n.withDO(n.DO.Unscoped())
}

func (n notificationEventMessageDo) Create(values ...*model.NotificationEventMessage) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n notificationEventMessageDo) CreateInBatches(values []*model.NotificationEventMessage, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n notificationEventMessageDo) Save(values ...*model.NotificationEventMessage) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n notificationEventMessageDo) First() (*model.NotificationEventMessage, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.NotificationEventMessage), nil
	}
}

func (n notificationEventMessageDo) Take() (*model.NotificationEventMessage, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.NotificationEventMessage), nil
	}
}

func (n notificationEventMessageDo) Last() (*model.NotificationEventMessage, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.NotificationEventMessage), nil
	}
}

func (n notificationEventMessageDo) Find() ([]*model.NotificationEventMessage, error) {
	result, err := n.DO.Find()
	return result.([]*model.NotificationEventMessage), err
}

func (n notificationEventMessageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NotificationEventMessage, err error) {
	buf := make([]*model.NotificationEventMessage, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n notificationEventMessageDo) FindInBatches(result *[]*model.NotificationEventMessage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n notificationEventMessageDo) Attrs(attrs ...field.AssignExpr) INotificationEventMessageDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n notificationEventMessageDo) Assign(attrs ...field.AssignExpr) INotificationEventMessageDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n notificationEventMessageDo) Joins(fields ...field.RelationField) INotificationEventMessageDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n notificationEventMessageDo) Preload(fields ...field.RelationField) INotificationEventMessageDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n notificationEventMessageDo) FirstOrInit() (*model.NotificationEventMessage, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.NotificationEventMessage), nil
	}
}

func (n notificationEventMessageDo) FirstOrCreate() (*model.NotificationEventMessage, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.NotificationEventMessage), nil
	}
}

func (n notificationEventMessageDo) FindByPage(offset int, limit int) (result []*model.NotificationEventMessage, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n notificationEventMessageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n notificationEventMessageDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n notificationEventMessageDo) Delete(models ...*model.NotificationEventMessage) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *notificationEventMessageDo) withDO(do gen.Dao) *notificationEventMessageDo {
	n.DO = *do.(*gen.DO)
	return n
}