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

func newGroupRole(db *gorm.DB, opts ...gen.DOOption) groupRole {
	_groupRole := groupRole{}

	_groupRole.groupRoleDo.UseDB(db, opts...)
	_groupRole.groupRoleDo.UseModel(&model.GroupRole{})

	tableName := _groupRole.groupRoleDo.TableName()
	_groupRole.ALL = field.NewAsterisk(tableName)
	_groupRole.RoleID = field.NewInt32(tableName, "role_id")
	_groupRole.GroupID = field.NewInt32(tableName, "group_id")
	_groupRole.PermissionKey = field.NewString(tableName, "permission_key")
	_groupRole.PermissionValue = field.NewInt32(tableName, "permission_value")
	_groupRole.CreateTime = field.NewTime(tableName, "create_time")
	_groupRole.CreateBy = field.NewInt32(tableName, "create_by")
	_groupRole.ChangeTime = field.NewTime(tableName, "change_time")
	_groupRole.ChangeBy = field.NewInt32(tableName, "change_by")

	_groupRole.fillFieldMap()

	return _groupRole
}

type groupRole struct {
	groupRoleDo

	ALL             field.Asterisk
	RoleID          field.Int32
	GroupID         field.Int32
	PermissionKey   field.String
	PermissionValue field.Int32
	CreateTime      field.Time
	CreateBy        field.Int32
	ChangeTime      field.Time
	ChangeBy        field.Int32

	fieldMap map[string]field.Expr
}

func (g groupRole) Table(newTableName string) *groupRole {
	g.groupRoleDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g groupRole) As(alias string) *groupRole {
	g.groupRoleDo.DO = *(g.groupRoleDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *groupRole) updateTableName(table string) *groupRole {
	g.ALL = field.NewAsterisk(table)
	g.RoleID = field.NewInt32(table, "role_id")
	g.GroupID = field.NewInt32(table, "group_id")
	g.PermissionKey = field.NewString(table, "permission_key")
	g.PermissionValue = field.NewInt32(table, "permission_value")
	g.CreateTime = field.NewTime(table, "create_time")
	g.CreateBy = field.NewInt32(table, "create_by")
	g.ChangeTime = field.NewTime(table, "change_time")
	g.ChangeBy = field.NewInt32(table, "change_by")

	g.fillFieldMap()

	return g
}

func (g *groupRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *groupRole) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 8)
	g.fieldMap["role_id"] = g.RoleID
	g.fieldMap["group_id"] = g.GroupID
	g.fieldMap["permission_key"] = g.PermissionKey
	g.fieldMap["permission_value"] = g.PermissionValue
	g.fieldMap["create_time"] = g.CreateTime
	g.fieldMap["create_by"] = g.CreateBy
	g.fieldMap["change_time"] = g.ChangeTime
	g.fieldMap["change_by"] = g.ChangeBy
}

func (g groupRole) clone(db *gorm.DB) groupRole {
	g.groupRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g groupRole) replaceDB(db *gorm.DB) groupRole {
	g.groupRoleDo.ReplaceDB(db)
	return g
}

type groupRoleDo struct{ gen.DO }

type IGroupRoleDo interface {
	gen.SubQuery
	Debug() IGroupRoleDo
	WithContext(ctx context.Context) IGroupRoleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGroupRoleDo
	WriteDB() IGroupRoleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGroupRoleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGroupRoleDo
	Not(conds ...gen.Condition) IGroupRoleDo
	Or(conds ...gen.Condition) IGroupRoleDo
	Select(conds ...field.Expr) IGroupRoleDo
	Where(conds ...gen.Condition) IGroupRoleDo
	Order(conds ...field.Expr) IGroupRoleDo
	Distinct(cols ...field.Expr) IGroupRoleDo
	Omit(cols ...field.Expr) IGroupRoleDo
	Join(table schema.Tabler, on ...field.Expr) IGroupRoleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGroupRoleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGroupRoleDo
	Group(cols ...field.Expr) IGroupRoleDo
	Having(conds ...gen.Condition) IGroupRoleDo
	Limit(limit int) IGroupRoleDo
	Offset(offset int) IGroupRoleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGroupRoleDo
	Unscoped() IGroupRoleDo
	Create(values ...*model.GroupRole) error
	CreateInBatches(values []*model.GroupRole, batchSize int) error
	Save(values ...*model.GroupRole) error
	First() (*model.GroupRole, error)
	Take() (*model.GroupRole, error)
	Last() (*model.GroupRole, error)
	Find() ([]*model.GroupRole, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GroupRole, err error)
	FindInBatches(result *[]*model.GroupRole, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GroupRole) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGroupRoleDo
	Assign(attrs ...field.AssignExpr) IGroupRoleDo
	Joins(fields ...field.RelationField) IGroupRoleDo
	Preload(fields ...field.RelationField) IGroupRoleDo
	FirstOrInit() (*model.GroupRole, error)
	FirstOrCreate() (*model.GroupRole, error)
	FindByPage(offset int, limit int) (result []*model.GroupRole, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGroupRoleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g groupRoleDo) Debug() IGroupRoleDo {
	return g.withDO(g.DO.Debug())
}

func (g groupRoleDo) WithContext(ctx context.Context) IGroupRoleDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g groupRoleDo) ReadDB() IGroupRoleDo {
	return g.Clauses(dbresolver.Read)
}

func (g groupRoleDo) WriteDB() IGroupRoleDo {
	return g.Clauses(dbresolver.Write)
}

func (g groupRoleDo) Session(config *gorm.Session) IGroupRoleDo {
	return g.withDO(g.DO.Session(config))
}

func (g groupRoleDo) Clauses(conds ...clause.Expression) IGroupRoleDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g groupRoleDo) Returning(value interface{}, columns ...string) IGroupRoleDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g groupRoleDo) Not(conds ...gen.Condition) IGroupRoleDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g groupRoleDo) Or(conds ...gen.Condition) IGroupRoleDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g groupRoleDo) Select(conds ...field.Expr) IGroupRoleDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g groupRoleDo) Where(conds ...gen.Condition) IGroupRoleDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g groupRoleDo) Order(conds ...field.Expr) IGroupRoleDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g groupRoleDo) Distinct(cols ...field.Expr) IGroupRoleDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g groupRoleDo) Omit(cols ...field.Expr) IGroupRoleDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g groupRoleDo) Join(table schema.Tabler, on ...field.Expr) IGroupRoleDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g groupRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGroupRoleDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g groupRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) IGroupRoleDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g groupRoleDo) Group(cols ...field.Expr) IGroupRoleDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g groupRoleDo) Having(conds ...gen.Condition) IGroupRoleDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g groupRoleDo) Limit(limit int) IGroupRoleDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g groupRoleDo) Offset(offset int) IGroupRoleDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g groupRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGroupRoleDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g groupRoleDo) Unscoped() IGroupRoleDo {
	return g.withDO(g.DO.Unscoped())
}

func (g groupRoleDo) Create(values ...*model.GroupRole) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g groupRoleDo) CreateInBatches(values []*model.GroupRole, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g groupRoleDo) Save(values ...*model.GroupRole) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g groupRoleDo) First() (*model.GroupRole, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupRole), nil
	}
}

func (g groupRoleDo) Take() (*model.GroupRole, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupRole), nil
	}
}

func (g groupRoleDo) Last() (*model.GroupRole, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupRole), nil
	}
}

func (g groupRoleDo) Find() ([]*model.GroupRole, error) {
	result, err := g.DO.Find()
	return result.([]*model.GroupRole), err
}

func (g groupRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GroupRole, err error) {
	buf := make([]*model.GroupRole, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g groupRoleDo) FindInBatches(result *[]*model.GroupRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g groupRoleDo) Attrs(attrs ...field.AssignExpr) IGroupRoleDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g groupRoleDo) Assign(attrs ...field.AssignExpr) IGroupRoleDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g groupRoleDo) Joins(fields ...field.RelationField) IGroupRoleDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g groupRoleDo) Preload(fields ...field.RelationField) IGroupRoleDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g groupRoleDo) FirstOrInit() (*model.GroupRole, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupRole), nil
	}
}

func (g groupRoleDo) FirstOrCreate() (*model.GroupRole, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupRole), nil
	}
}

func (g groupRoleDo) FindByPage(offset int, limit int) (result []*model.GroupRole, count int64, err error) {
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

func (g groupRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g groupRoleDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g groupRoleDo) Delete(models ...*model.GroupRole) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *groupRoleDo) withDO(do gen.Dao) *groupRoleDo {
	g.DO = *do.(*gen.DO)
	return g
}
