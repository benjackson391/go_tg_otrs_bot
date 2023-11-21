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

func newOrgstructure(db *gorm.DB, opts ...gen.DOOption) orgstructure {
	_orgstructure := orgstructure{}

	_orgstructure.orgstructureDo.UseDB(db, opts...)
	_orgstructure.orgstructureDo.UseModel(&model.Orgstructure{})

	tableName := _orgstructure.orgstructureDo.TableName()
	_orgstructure.ALL = field.NewAsterisk(tableName)
	_orgstructure.OrgstructureID = field.NewInt32(tableName, "orgstructure_id")
	_orgstructure.Name = field.NewString(tableName, "name")
	_orgstructure.Comments = field.NewString(tableName, "comments")
	_orgstructure.ParentID = field.NewInt32(tableName, "parent_id")
	_orgstructure.OwnerID = field.NewInt32(tableName, "owner_id")
	_orgstructure.Members = field.NewString(tableName, "members")

	_orgstructure.fillFieldMap()

	return _orgstructure
}

type orgstructure struct {
	orgstructureDo

	ALL            field.Asterisk
	OrgstructureID field.Int32
	Name           field.String
	Comments       field.String
	ParentID       field.Int32
	OwnerID        field.Int32
	Members        field.String

	fieldMap map[string]field.Expr
}

func (o orgstructure) Table(newTableName string) *orgstructure {
	o.orgstructureDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o orgstructure) As(alias string) *orgstructure {
	o.orgstructureDo.DO = *(o.orgstructureDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *orgstructure) updateTableName(table string) *orgstructure {
	o.ALL = field.NewAsterisk(table)
	o.OrgstructureID = field.NewInt32(table, "orgstructure_id")
	o.Name = field.NewString(table, "name")
	o.Comments = field.NewString(table, "comments")
	o.ParentID = field.NewInt32(table, "parent_id")
	o.OwnerID = field.NewInt32(table, "owner_id")
	o.Members = field.NewString(table, "members")

	o.fillFieldMap()

	return o
}

func (o *orgstructure) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *orgstructure) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 6)
	o.fieldMap["orgstructure_id"] = o.OrgstructureID
	o.fieldMap["name"] = o.Name
	o.fieldMap["comments"] = o.Comments
	o.fieldMap["parent_id"] = o.ParentID
	o.fieldMap["owner_id"] = o.OwnerID
	o.fieldMap["members"] = o.Members
}

func (o orgstructure) clone(db *gorm.DB) orgstructure {
	o.orgstructureDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o orgstructure) replaceDB(db *gorm.DB) orgstructure {
	o.orgstructureDo.ReplaceDB(db)
	return o
}

type orgstructureDo struct{ gen.DO }

type IOrgstructureDo interface {
	gen.SubQuery
	Debug() IOrgstructureDo
	WithContext(ctx context.Context) IOrgstructureDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOrgstructureDo
	WriteDB() IOrgstructureDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOrgstructureDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOrgstructureDo
	Not(conds ...gen.Condition) IOrgstructureDo
	Or(conds ...gen.Condition) IOrgstructureDo
	Select(conds ...field.Expr) IOrgstructureDo
	Where(conds ...gen.Condition) IOrgstructureDo
	Order(conds ...field.Expr) IOrgstructureDo
	Distinct(cols ...field.Expr) IOrgstructureDo
	Omit(cols ...field.Expr) IOrgstructureDo
	Join(table schema.Tabler, on ...field.Expr) IOrgstructureDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOrgstructureDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOrgstructureDo
	Group(cols ...field.Expr) IOrgstructureDo
	Having(conds ...gen.Condition) IOrgstructureDo
	Limit(limit int) IOrgstructureDo
	Offset(offset int) IOrgstructureDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOrgstructureDo
	Unscoped() IOrgstructureDo
	Create(values ...*model.Orgstructure) error
	CreateInBatches(values []*model.Orgstructure, batchSize int) error
	Save(values ...*model.Orgstructure) error
	First() (*model.Orgstructure, error)
	Take() (*model.Orgstructure, error)
	Last() (*model.Orgstructure, error)
	Find() ([]*model.Orgstructure, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Orgstructure, err error)
	FindInBatches(result *[]*model.Orgstructure, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Orgstructure) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOrgstructureDo
	Assign(attrs ...field.AssignExpr) IOrgstructureDo
	Joins(fields ...field.RelationField) IOrgstructureDo
	Preload(fields ...field.RelationField) IOrgstructureDo
	FirstOrInit() (*model.Orgstructure, error)
	FirstOrCreate() (*model.Orgstructure, error)
	FindByPage(offset int, limit int) (result []*model.Orgstructure, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOrgstructureDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o orgstructureDo) Debug() IOrgstructureDo {
	return o.withDO(o.DO.Debug())
}

func (o orgstructureDo) WithContext(ctx context.Context) IOrgstructureDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o orgstructureDo) ReadDB() IOrgstructureDo {
	return o.Clauses(dbresolver.Read)
}

func (o orgstructureDo) WriteDB() IOrgstructureDo {
	return o.Clauses(dbresolver.Write)
}

func (o orgstructureDo) Session(config *gorm.Session) IOrgstructureDo {
	return o.withDO(o.DO.Session(config))
}

func (o orgstructureDo) Clauses(conds ...clause.Expression) IOrgstructureDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o orgstructureDo) Returning(value interface{}, columns ...string) IOrgstructureDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o orgstructureDo) Not(conds ...gen.Condition) IOrgstructureDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o orgstructureDo) Or(conds ...gen.Condition) IOrgstructureDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o orgstructureDo) Select(conds ...field.Expr) IOrgstructureDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o orgstructureDo) Where(conds ...gen.Condition) IOrgstructureDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o orgstructureDo) Order(conds ...field.Expr) IOrgstructureDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o orgstructureDo) Distinct(cols ...field.Expr) IOrgstructureDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o orgstructureDo) Omit(cols ...field.Expr) IOrgstructureDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o orgstructureDo) Join(table schema.Tabler, on ...field.Expr) IOrgstructureDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o orgstructureDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOrgstructureDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o orgstructureDo) RightJoin(table schema.Tabler, on ...field.Expr) IOrgstructureDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o orgstructureDo) Group(cols ...field.Expr) IOrgstructureDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o orgstructureDo) Having(conds ...gen.Condition) IOrgstructureDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o orgstructureDo) Limit(limit int) IOrgstructureDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o orgstructureDo) Offset(offset int) IOrgstructureDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o orgstructureDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOrgstructureDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o orgstructureDo) Unscoped() IOrgstructureDo {
	return o.withDO(o.DO.Unscoped())
}

func (o orgstructureDo) Create(values ...*model.Orgstructure) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o orgstructureDo) CreateInBatches(values []*model.Orgstructure, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o orgstructureDo) Save(values ...*model.Orgstructure) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o orgstructureDo) First() (*model.Orgstructure, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Orgstructure), nil
	}
}

func (o orgstructureDo) Take() (*model.Orgstructure, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Orgstructure), nil
	}
}

func (o orgstructureDo) Last() (*model.Orgstructure, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Orgstructure), nil
	}
}

func (o orgstructureDo) Find() ([]*model.Orgstructure, error) {
	result, err := o.DO.Find()
	return result.([]*model.Orgstructure), err
}

func (o orgstructureDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Orgstructure, err error) {
	buf := make([]*model.Orgstructure, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o orgstructureDo) FindInBatches(result *[]*model.Orgstructure, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o orgstructureDo) Attrs(attrs ...field.AssignExpr) IOrgstructureDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o orgstructureDo) Assign(attrs ...field.AssignExpr) IOrgstructureDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o orgstructureDo) Joins(fields ...field.RelationField) IOrgstructureDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o orgstructureDo) Preload(fields ...field.RelationField) IOrgstructureDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o orgstructureDo) FirstOrInit() (*model.Orgstructure, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Orgstructure), nil
	}
}

func (o orgstructureDo) FirstOrCreate() (*model.Orgstructure, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Orgstructure), nil
	}
}

func (o orgstructureDo) FindByPage(offset int, limit int) (result []*model.Orgstructure, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o orgstructureDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o orgstructureDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o orgstructureDo) Delete(models ...*model.Orgstructure) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *orgstructureDo) withDO(do gen.Dao) *orgstructureDo {
	o.DO = *do.(*gen.DO)
	return o
}