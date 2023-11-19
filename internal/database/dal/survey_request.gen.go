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

func newSurveyRequest(db *gorm.DB, opts ...gen.DOOption) surveyRequest {
	_surveyRequest := surveyRequest{}

	_surveyRequest.surveyRequestDo.UseDB(db, opts...)
	_surveyRequest.surveyRequestDo.UseModel(&model.SurveyRequest{})

	tableName := _surveyRequest.surveyRequestDo.TableName()
	_surveyRequest.ALL = field.NewAsterisk(tableName)
	_surveyRequest.ID = field.NewInt64(tableName, "id")
	_surveyRequest.TicketID = field.NewInt64(tableName, "ticket_id")
	_surveyRequest.SurveyID = field.NewInt64(tableName, "survey_id")
	_surveyRequest.ValidID = field.NewInt32(tableName, "valid_id")
	_surveyRequest.PublicSurveyKey = field.NewString(tableName, "public_survey_key")
	_surveyRequest.SendTo = field.NewString(tableName, "send_to")
	_surveyRequest.SendTime = field.NewTime(tableName, "send_time")
	_surveyRequest.VoteTime = field.NewTime(tableName, "vote_time")
	_surveyRequest.CreateTime = field.NewTime(tableName, "create_time")

	_surveyRequest.fillFieldMap()

	return _surveyRequest
}

type surveyRequest struct {
	surveyRequestDo

	ALL             field.Asterisk
	ID              field.Int64
	TicketID        field.Int64
	SurveyID        field.Int64
	ValidID         field.Int32
	PublicSurveyKey field.String
	SendTo          field.String
	SendTime        field.Time
	VoteTime        field.Time
	CreateTime      field.Time

	fieldMap map[string]field.Expr
}

func (s surveyRequest) Table(newTableName string) *surveyRequest {
	s.surveyRequestDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s surveyRequest) As(alias string) *surveyRequest {
	s.surveyRequestDo.DO = *(s.surveyRequestDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *surveyRequest) updateTableName(table string) *surveyRequest {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.TicketID = field.NewInt64(table, "ticket_id")
	s.SurveyID = field.NewInt64(table, "survey_id")
	s.ValidID = field.NewInt32(table, "valid_id")
	s.PublicSurveyKey = field.NewString(table, "public_survey_key")
	s.SendTo = field.NewString(table, "send_to")
	s.SendTime = field.NewTime(table, "send_time")
	s.VoteTime = field.NewTime(table, "vote_time")
	s.CreateTime = field.NewTime(table, "create_time")

	s.fillFieldMap()

	return s
}

func (s *surveyRequest) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *surveyRequest) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["id"] = s.ID
	s.fieldMap["ticket_id"] = s.TicketID
	s.fieldMap["survey_id"] = s.SurveyID
	s.fieldMap["valid_id"] = s.ValidID
	s.fieldMap["public_survey_key"] = s.PublicSurveyKey
	s.fieldMap["send_to"] = s.SendTo
	s.fieldMap["send_time"] = s.SendTime
	s.fieldMap["vote_time"] = s.VoteTime
	s.fieldMap["create_time"] = s.CreateTime
}

func (s surveyRequest) clone(db *gorm.DB) surveyRequest {
	s.surveyRequestDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s surveyRequest) replaceDB(db *gorm.DB) surveyRequest {
	s.surveyRequestDo.ReplaceDB(db)
	return s
}

type surveyRequestDo struct{ gen.DO }

type ISurveyRequestDo interface {
	gen.SubQuery
	Debug() ISurveyRequestDo
	WithContext(ctx context.Context) ISurveyRequestDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISurveyRequestDo
	WriteDB() ISurveyRequestDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISurveyRequestDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISurveyRequestDo
	Not(conds ...gen.Condition) ISurveyRequestDo
	Or(conds ...gen.Condition) ISurveyRequestDo
	Select(conds ...field.Expr) ISurveyRequestDo
	Where(conds ...gen.Condition) ISurveyRequestDo
	Order(conds ...field.Expr) ISurveyRequestDo
	Distinct(cols ...field.Expr) ISurveyRequestDo
	Omit(cols ...field.Expr) ISurveyRequestDo
	Join(table schema.Tabler, on ...field.Expr) ISurveyRequestDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISurveyRequestDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISurveyRequestDo
	Group(cols ...field.Expr) ISurveyRequestDo
	Having(conds ...gen.Condition) ISurveyRequestDo
	Limit(limit int) ISurveyRequestDo
	Offset(offset int) ISurveyRequestDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISurveyRequestDo
	Unscoped() ISurveyRequestDo
	Create(values ...*model.SurveyRequest) error
	CreateInBatches(values []*model.SurveyRequest, batchSize int) error
	Save(values ...*model.SurveyRequest) error
	First() (*model.SurveyRequest, error)
	Take() (*model.SurveyRequest, error)
	Last() (*model.SurveyRequest, error)
	Find() ([]*model.SurveyRequest, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SurveyRequest, err error)
	FindInBatches(result *[]*model.SurveyRequest, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SurveyRequest) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISurveyRequestDo
	Assign(attrs ...field.AssignExpr) ISurveyRequestDo
	Joins(fields ...field.RelationField) ISurveyRequestDo
	Preload(fields ...field.RelationField) ISurveyRequestDo
	FirstOrInit() (*model.SurveyRequest, error)
	FirstOrCreate() (*model.SurveyRequest, error)
	FindByPage(offset int, limit int) (result []*model.SurveyRequest, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISurveyRequestDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s surveyRequestDo) Debug() ISurveyRequestDo {
	return s.withDO(s.DO.Debug())
}

func (s surveyRequestDo) WithContext(ctx context.Context) ISurveyRequestDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s surveyRequestDo) ReadDB() ISurveyRequestDo {
	return s.Clauses(dbresolver.Read)
}

func (s surveyRequestDo) WriteDB() ISurveyRequestDo {
	return s.Clauses(dbresolver.Write)
}

func (s surveyRequestDo) Session(config *gorm.Session) ISurveyRequestDo {
	return s.withDO(s.DO.Session(config))
}

func (s surveyRequestDo) Clauses(conds ...clause.Expression) ISurveyRequestDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s surveyRequestDo) Returning(value interface{}, columns ...string) ISurveyRequestDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s surveyRequestDo) Not(conds ...gen.Condition) ISurveyRequestDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s surveyRequestDo) Or(conds ...gen.Condition) ISurveyRequestDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s surveyRequestDo) Select(conds ...field.Expr) ISurveyRequestDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s surveyRequestDo) Where(conds ...gen.Condition) ISurveyRequestDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s surveyRequestDo) Order(conds ...field.Expr) ISurveyRequestDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s surveyRequestDo) Distinct(cols ...field.Expr) ISurveyRequestDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s surveyRequestDo) Omit(cols ...field.Expr) ISurveyRequestDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s surveyRequestDo) Join(table schema.Tabler, on ...field.Expr) ISurveyRequestDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s surveyRequestDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISurveyRequestDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s surveyRequestDo) RightJoin(table schema.Tabler, on ...field.Expr) ISurveyRequestDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s surveyRequestDo) Group(cols ...field.Expr) ISurveyRequestDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s surveyRequestDo) Having(conds ...gen.Condition) ISurveyRequestDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s surveyRequestDo) Limit(limit int) ISurveyRequestDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s surveyRequestDo) Offset(offset int) ISurveyRequestDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s surveyRequestDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISurveyRequestDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s surveyRequestDo) Unscoped() ISurveyRequestDo {
	return s.withDO(s.DO.Unscoped())
}

func (s surveyRequestDo) Create(values ...*model.SurveyRequest) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s surveyRequestDo) CreateInBatches(values []*model.SurveyRequest, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s surveyRequestDo) Save(values ...*model.SurveyRequest) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s surveyRequestDo) First() (*model.SurveyRequest, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SurveyRequest), nil
	}
}

func (s surveyRequestDo) Take() (*model.SurveyRequest, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SurveyRequest), nil
	}
}

func (s surveyRequestDo) Last() (*model.SurveyRequest, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SurveyRequest), nil
	}
}

func (s surveyRequestDo) Find() ([]*model.SurveyRequest, error) {
	result, err := s.DO.Find()
	return result.([]*model.SurveyRequest), err
}

func (s surveyRequestDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SurveyRequest, err error) {
	buf := make([]*model.SurveyRequest, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s surveyRequestDo) FindInBatches(result *[]*model.SurveyRequest, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s surveyRequestDo) Attrs(attrs ...field.AssignExpr) ISurveyRequestDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s surveyRequestDo) Assign(attrs ...field.AssignExpr) ISurveyRequestDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s surveyRequestDo) Joins(fields ...field.RelationField) ISurveyRequestDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s surveyRequestDo) Preload(fields ...field.RelationField) ISurveyRequestDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s surveyRequestDo) FirstOrInit() (*model.SurveyRequest, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SurveyRequest), nil
	}
}

func (s surveyRequestDo) FirstOrCreate() (*model.SurveyRequest, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SurveyRequest), nil
	}
}

func (s surveyRequestDo) FindByPage(offset int, limit int) (result []*model.SurveyRequest, count int64, err error) {
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

func (s surveyRequestDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s surveyRequestDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s surveyRequestDo) Delete(models ...*model.SurveyRequest) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *surveyRequestDo) withDO(do gen.Dao) *surveyRequestDo {
	s.DO = *do.(*gen.DO)
	return s
}
