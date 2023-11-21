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

func newSysconfigDefaultVersion(db *gorm.DB, opts ...gen.DOOption) sysconfigDefaultVersion {
	_sysconfigDefaultVersion := sysconfigDefaultVersion{}

	_sysconfigDefaultVersion.sysconfigDefaultVersionDo.UseDB(db, opts...)
	_sysconfigDefaultVersion.sysconfigDefaultVersionDo.UseModel(&model.SysconfigDefaultVersion{})

	tableName := _sysconfigDefaultVersion.sysconfigDefaultVersionDo.TableName()
	_sysconfigDefaultVersion.ALL = field.NewAsterisk(tableName)
	_sysconfigDefaultVersion.ID = field.NewInt32(tableName, "id")
	_sysconfigDefaultVersion.SysconfigDefaultID = field.NewInt32(tableName, "sysconfig_default_id")
	_sysconfigDefaultVersion.Name = field.NewString(tableName, "name")
	_sysconfigDefaultVersion.Description = field.NewBytes(tableName, "description")
	_sysconfigDefaultVersion.Navigation = field.NewString(tableName, "navigation")
	_sysconfigDefaultVersion.IsInvisible = field.NewInt32(tableName, "is_invisible")
	_sysconfigDefaultVersion.IsReadonly = field.NewInt32(tableName, "is_readonly")
	_sysconfigDefaultVersion.IsRequired = field.NewInt32(tableName, "is_required")
	_sysconfigDefaultVersion.IsValid = field.NewInt32(tableName, "is_valid")
	_sysconfigDefaultVersion.HasConfiglevel = field.NewInt32(tableName, "has_configlevel")
	_sysconfigDefaultVersion.UserModificationPossible = field.NewInt32(tableName, "user_modification_possible")
	_sysconfigDefaultVersion.UserModificationActive = field.NewInt32(tableName, "user_modification_active")
	_sysconfigDefaultVersion.UserPreferencesGroup = field.NewString(tableName, "user_preferences_group")
	_sysconfigDefaultVersion.XMLContentRaw = field.NewBytes(tableName, "xml_content_raw")
	_sysconfigDefaultVersion.XMLContentParsed = field.NewBytes(tableName, "xml_content_parsed")
	_sysconfigDefaultVersion.XMLFilename = field.NewString(tableName, "xml_filename")
	_sysconfigDefaultVersion.EffectiveValue = field.NewBytes(tableName, "effective_value")
	_sysconfigDefaultVersion.CreateTime = field.NewTime(tableName, "create_time")
	_sysconfigDefaultVersion.CreateBy = field.NewInt32(tableName, "create_by")
	_sysconfigDefaultVersion.ChangeTime = field.NewTime(tableName, "change_time")
	_sysconfigDefaultVersion.ChangeBy = field.NewInt32(tableName, "change_by")

	_sysconfigDefaultVersion.fillFieldMap()

	return _sysconfigDefaultVersion
}

type sysconfigDefaultVersion struct {
	sysconfigDefaultVersionDo

	ALL                      field.Asterisk
	ID                       field.Int32
	SysconfigDefaultID       field.Int32
	Name                     field.String
	Description              field.Bytes
	Navigation               field.String
	IsInvisible              field.Int32
	IsReadonly               field.Int32
	IsRequired               field.Int32
	IsValid                  field.Int32
	HasConfiglevel           field.Int32
	UserModificationPossible field.Int32
	UserModificationActive   field.Int32
	UserPreferencesGroup     field.String
	XMLContentRaw            field.Bytes
	XMLContentParsed         field.Bytes
	XMLFilename              field.String
	EffectiveValue           field.Bytes
	CreateTime               field.Time
	CreateBy                 field.Int32
	ChangeTime               field.Time
	ChangeBy                 field.Int32

	fieldMap map[string]field.Expr
}

func (s sysconfigDefaultVersion) Table(newTableName string) *sysconfigDefaultVersion {
	s.sysconfigDefaultVersionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysconfigDefaultVersion) As(alias string) *sysconfigDefaultVersion {
	s.sysconfigDefaultVersionDo.DO = *(s.sysconfigDefaultVersionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysconfigDefaultVersion) updateTableName(table string) *sysconfigDefaultVersion {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.SysconfigDefaultID = field.NewInt32(table, "sysconfig_default_id")
	s.Name = field.NewString(table, "name")
	s.Description = field.NewBytes(table, "description")
	s.Navigation = field.NewString(table, "navigation")
	s.IsInvisible = field.NewInt32(table, "is_invisible")
	s.IsReadonly = field.NewInt32(table, "is_readonly")
	s.IsRequired = field.NewInt32(table, "is_required")
	s.IsValid = field.NewInt32(table, "is_valid")
	s.HasConfiglevel = field.NewInt32(table, "has_configlevel")
	s.UserModificationPossible = field.NewInt32(table, "user_modification_possible")
	s.UserModificationActive = field.NewInt32(table, "user_modification_active")
	s.UserPreferencesGroup = field.NewString(table, "user_preferences_group")
	s.XMLContentRaw = field.NewBytes(table, "xml_content_raw")
	s.XMLContentParsed = field.NewBytes(table, "xml_content_parsed")
	s.XMLFilename = field.NewString(table, "xml_filename")
	s.EffectiveValue = field.NewBytes(table, "effective_value")
	s.CreateTime = field.NewTime(table, "create_time")
	s.CreateBy = field.NewInt32(table, "create_by")
	s.ChangeTime = field.NewTime(table, "change_time")
	s.ChangeBy = field.NewInt32(table, "change_by")

	s.fillFieldMap()

	return s
}

func (s *sysconfigDefaultVersion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysconfigDefaultVersion) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 21)
	s.fieldMap["id"] = s.ID
	s.fieldMap["sysconfig_default_id"] = s.SysconfigDefaultID
	s.fieldMap["name"] = s.Name
	s.fieldMap["description"] = s.Description
	s.fieldMap["navigation"] = s.Navigation
	s.fieldMap["is_invisible"] = s.IsInvisible
	s.fieldMap["is_readonly"] = s.IsReadonly
	s.fieldMap["is_required"] = s.IsRequired
	s.fieldMap["is_valid"] = s.IsValid
	s.fieldMap["has_configlevel"] = s.HasConfiglevel
	s.fieldMap["user_modification_possible"] = s.UserModificationPossible
	s.fieldMap["user_modification_active"] = s.UserModificationActive
	s.fieldMap["user_preferences_group"] = s.UserPreferencesGroup
	s.fieldMap["xml_content_raw"] = s.XMLContentRaw
	s.fieldMap["xml_content_parsed"] = s.XMLContentParsed
	s.fieldMap["xml_filename"] = s.XMLFilename
	s.fieldMap["effective_value"] = s.EffectiveValue
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["change_time"] = s.ChangeTime
	s.fieldMap["change_by"] = s.ChangeBy
}

func (s sysconfigDefaultVersion) clone(db *gorm.DB) sysconfigDefaultVersion {
	s.sysconfigDefaultVersionDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysconfigDefaultVersion) replaceDB(db *gorm.DB) sysconfigDefaultVersion {
	s.sysconfigDefaultVersionDo.ReplaceDB(db)
	return s
}

type sysconfigDefaultVersionDo struct{ gen.DO }

type ISysconfigDefaultVersionDo interface {
	gen.SubQuery
	Debug() ISysconfigDefaultVersionDo
	WithContext(ctx context.Context) ISysconfigDefaultVersionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysconfigDefaultVersionDo
	WriteDB() ISysconfigDefaultVersionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysconfigDefaultVersionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysconfigDefaultVersionDo
	Not(conds ...gen.Condition) ISysconfigDefaultVersionDo
	Or(conds ...gen.Condition) ISysconfigDefaultVersionDo
	Select(conds ...field.Expr) ISysconfigDefaultVersionDo
	Where(conds ...gen.Condition) ISysconfigDefaultVersionDo
	Order(conds ...field.Expr) ISysconfigDefaultVersionDo
	Distinct(cols ...field.Expr) ISysconfigDefaultVersionDo
	Omit(cols ...field.Expr) ISysconfigDefaultVersionDo
	Join(table schema.Tabler, on ...field.Expr) ISysconfigDefaultVersionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysconfigDefaultVersionDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysconfigDefaultVersionDo
	Group(cols ...field.Expr) ISysconfigDefaultVersionDo
	Having(conds ...gen.Condition) ISysconfigDefaultVersionDo
	Limit(limit int) ISysconfigDefaultVersionDo
	Offset(offset int) ISysconfigDefaultVersionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysconfigDefaultVersionDo
	Unscoped() ISysconfigDefaultVersionDo
	Create(values ...*model.SysconfigDefaultVersion) error
	CreateInBatches(values []*model.SysconfigDefaultVersion, batchSize int) error
	Save(values ...*model.SysconfigDefaultVersion) error
	First() (*model.SysconfigDefaultVersion, error)
	Take() (*model.SysconfigDefaultVersion, error)
	Last() (*model.SysconfigDefaultVersion, error)
	Find() ([]*model.SysconfigDefaultVersion, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysconfigDefaultVersion, err error)
	FindInBatches(result *[]*model.SysconfigDefaultVersion, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysconfigDefaultVersion) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysconfigDefaultVersionDo
	Assign(attrs ...field.AssignExpr) ISysconfigDefaultVersionDo
	Joins(fields ...field.RelationField) ISysconfigDefaultVersionDo
	Preload(fields ...field.RelationField) ISysconfigDefaultVersionDo
	FirstOrInit() (*model.SysconfigDefaultVersion, error)
	FirstOrCreate() (*model.SysconfigDefaultVersion, error)
	FindByPage(offset int, limit int) (result []*model.SysconfigDefaultVersion, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysconfigDefaultVersionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysconfigDefaultVersionDo) Debug() ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Debug())
}

func (s sysconfigDefaultVersionDo) WithContext(ctx context.Context) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysconfigDefaultVersionDo) ReadDB() ISysconfigDefaultVersionDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysconfigDefaultVersionDo) WriteDB() ISysconfigDefaultVersionDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysconfigDefaultVersionDo) Session(config *gorm.Session) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysconfigDefaultVersionDo) Clauses(conds ...clause.Expression) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysconfigDefaultVersionDo) Returning(value interface{}, columns ...string) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysconfigDefaultVersionDo) Not(conds ...gen.Condition) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysconfigDefaultVersionDo) Or(conds ...gen.Condition) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysconfigDefaultVersionDo) Select(conds ...field.Expr) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysconfigDefaultVersionDo) Where(conds ...gen.Condition) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysconfigDefaultVersionDo) Order(conds ...field.Expr) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysconfigDefaultVersionDo) Distinct(cols ...field.Expr) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysconfigDefaultVersionDo) Omit(cols ...field.Expr) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysconfigDefaultVersionDo) Join(table schema.Tabler, on ...field.Expr) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysconfigDefaultVersionDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysconfigDefaultVersionDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysconfigDefaultVersionDo) Group(cols ...field.Expr) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysconfigDefaultVersionDo) Having(conds ...gen.Condition) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysconfigDefaultVersionDo) Limit(limit int) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysconfigDefaultVersionDo) Offset(offset int) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysconfigDefaultVersionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysconfigDefaultVersionDo) Unscoped() ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysconfigDefaultVersionDo) Create(values ...*model.SysconfigDefaultVersion) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysconfigDefaultVersionDo) CreateInBatches(values []*model.SysconfigDefaultVersion, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysconfigDefaultVersionDo) Save(values ...*model.SysconfigDefaultVersion) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysconfigDefaultVersionDo) First() (*model.SysconfigDefaultVersion, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysconfigDefaultVersion), nil
	}
}

func (s sysconfigDefaultVersionDo) Take() (*model.SysconfigDefaultVersion, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysconfigDefaultVersion), nil
	}
}

func (s sysconfigDefaultVersionDo) Last() (*model.SysconfigDefaultVersion, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysconfigDefaultVersion), nil
	}
}

func (s sysconfigDefaultVersionDo) Find() ([]*model.SysconfigDefaultVersion, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysconfigDefaultVersion), err
}

func (s sysconfigDefaultVersionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysconfigDefaultVersion, err error) {
	buf := make([]*model.SysconfigDefaultVersion, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysconfigDefaultVersionDo) FindInBatches(result *[]*model.SysconfigDefaultVersion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysconfigDefaultVersionDo) Attrs(attrs ...field.AssignExpr) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysconfigDefaultVersionDo) Assign(attrs ...field.AssignExpr) ISysconfigDefaultVersionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysconfigDefaultVersionDo) Joins(fields ...field.RelationField) ISysconfigDefaultVersionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysconfigDefaultVersionDo) Preload(fields ...field.RelationField) ISysconfigDefaultVersionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysconfigDefaultVersionDo) FirstOrInit() (*model.SysconfigDefaultVersion, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysconfigDefaultVersion), nil
	}
}

func (s sysconfigDefaultVersionDo) FirstOrCreate() (*model.SysconfigDefaultVersion, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysconfigDefaultVersion), nil
	}
}

func (s sysconfigDefaultVersionDo) FindByPage(offset int, limit int) (result []*model.SysconfigDefaultVersion, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysconfigDefaultVersionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysconfigDefaultVersionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysconfigDefaultVersionDo) Delete(models ...*model.SysconfigDefaultVersion) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysconfigDefaultVersionDo) withDO(do gen.Dao) *sysconfigDefaultVersionDo {
	s.DO = *do.(*gen.DO)
	return s
}