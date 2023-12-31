package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"firstGO/pkg/model"
)

func newSubject(db *gorm.DB, opts ...gen.DOOption) subject {
	_subject := subject{}

	_subject.subjectDo.UseDB(db, opts...)
	_subject.subjectDo.UseModel(&model.Subject{})

	tableName := _subject.subjectDo.TableName()
	_subject.ALL = field.NewAsterisk(tableName)
	_subject.SubjectID = field.NewInt64(tableName, "SubjectId")
	_subject.Name = field.NewString(tableName, "Name")

	_subject.fillFieldMap()

	return _subject
}

type subject struct {
	subjectDo

	ALL       field.Asterisk
	SubjectID field.Int64
	Name      field.String

	fieldMap map[string]field.Expr
}

func (s subject) Table(newTableName string) *subject {
	s.subjectDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s subject) As(alias string) *subject {
	s.subjectDo.DO = *(s.subjectDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *subject) updateTableName(table string) *subject {
	s.ALL = field.NewAsterisk(table)
	s.SubjectID = field.NewInt64(table, "SubjectId")
	s.Name = field.NewString(table, "Name")

	s.fillFieldMap()

	return s
}

func (s *subject) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *subject) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 2)
	s.fieldMap["SubjectId"] = s.SubjectID
	s.fieldMap["Name"] = s.Name
}

func (s subject) clone(db *gorm.DB) subject {
	s.subjectDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s subject) replaceDB(db *gorm.DB) subject {
	s.subjectDo.ReplaceDB(db)
	return s
}

type subjectDo struct{ gen.DO }

type ISubjectDo interface {
	gen.SubQuery
	Debug() ISubjectDo
	WithContext(ctx context.Context) ISubjectDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISubjectDo
	WriteDB() ISubjectDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISubjectDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISubjectDo
	Not(conds ...gen.Condition) ISubjectDo
	Or(conds ...gen.Condition) ISubjectDo
	Select(conds ...field.Expr) ISubjectDo
	Where(conds ...gen.Condition) ISubjectDo
	Order(conds ...field.Expr) ISubjectDo
	Distinct(cols ...field.Expr) ISubjectDo
	Omit(cols ...field.Expr) ISubjectDo
	Join(table schema.Tabler, on ...field.Expr) ISubjectDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISubjectDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISubjectDo
	Group(cols ...field.Expr) ISubjectDo
	Having(conds ...gen.Condition) ISubjectDo
	Limit(limit int) ISubjectDo
	Offset(offset int) ISubjectDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISubjectDo
	Unscoped() ISubjectDo
	Create(values ...*model.Subject) error
	CreateInBatches(values []*model.Subject, batchSize int) error
	Save(values ...*model.Subject) error
	First() (*model.Subject, error)
	Take() (*model.Subject, error)
	Last() (*model.Subject, error)
	Find() ([]*model.Subject, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Subject, err error)
	FindInBatches(result *[]*model.Subject, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Subject) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISubjectDo
	Assign(attrs ...field.AssignExpr) ISubjectDo
	Joins(fields ...field.RelationField) ISubjectDo
	Preload(fields ...field.RelationField) ISubjectDo
	FirstOrInit() (*model.Subject, error)
	FirstOrCreate() (*model.Subject, error)
	FindByPage(offset int, limit int) (result []*model.Subject, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISubjectDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s subjectDo) Debug() ISubjectDo {
	return s.withDO(s.DO.Debug())
}

func (s subjectDo) WithContext(ctx context.Context) ISubjectDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s subjectDo) ReadDB() ISubjectDo {
	return s.Clauses(dbresolver.Read)
}

func (s subjectDo) WriteDB() ISubjectDo {
	return s.Clauses(dbresolver.Write)
}

func (s subjectDo) Session(config *gorm.Session) ISubjectDo {
	return s.withDO(s.DO.Session(config))
}

func (s subjectDo) Clauses(conds ...clause.Expression) ISubjectDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s subjectDo) Returning(value interface{}, columns ...string) ISubjectDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s subjectDo) Not(conds ...gen.Condition) ISubjectDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s subjectDo) Or(conds ...gen.Condition) ISubjectDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s subjectDo) Select(conds ...field.Expr) ISubjectDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s subjectDo) Where(conds ...gen.Condition) ISubjectDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s subjectDo) Order(conds ...field.Expr) ISubjectDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s subjectDo) Distinct(cols ...field.Expr) ISubjectDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s subjectDo) Omit(cols ...field.Expr) ISubjectDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s subjectDo) Join(table schema.Tabler, on ...field.Expr) ISubjectDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s subjectDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISubjectDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s subjectDo) RightJoin(table schema.Tabler, on ...field.Expr) ISubjectDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s subjectDo) Group(cols ...field.Expr) ISubjectDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s subjectDo) Having(conds ...gen.Condition) ISubjectDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s subjectDo) Limit(limit int) ISubjectDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s subjectDo) Offset(offset int) ISubjectDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s subjectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISubjectDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s subjectDo) Unscoped() ISubjectDo {
	return s.withDO(s.DO.Unscoped())
}

func (s subjectDo) Create(values ...*model.Subject) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s subjectDo) CreateInBatches(values []*model.Subject, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s subjectDo) Save(values ...*model.Subject) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s subjectDo) First() (*model.Subject, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Subject), nil
	}
}

func (s subjectDo) Take() (*model.Subject, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Subject), nil
	}
}

func (s subjectDo) Last() (*model.Subject, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Subject), nil
	}
}

func (s subjectDo) Find() ([]*model.Subject, error) {
	result, err := s.DO.Find()
	return result.([]*model.Subject), err
}

func (s subjectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Subject, err error) {
	buf := make([]*model.Subject, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s subjectDo) FindInBatches(result *[]*model.Subject, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s subjectDo) Attrs(attrs ...field.AssignExpr) ISubjectDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s subjectDo) Assign(attrs ...field.AssignExpr) ISubjectDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s subjectDo) Joins(fields ...field.RelationField) ISubjectDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s subjectDo) Preload(fields ...field.RelationField) ISubjectDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s subjectDo) FirstOrInit() (*model.Subject, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Subject), nil
	}
}

func (s subjectDo) FirstOrCreate() (*model.Subject, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Subject), nil
	}
}

func (s subjectDo) FindByPage(offset int, limit int) (result []*model.Subject, count int64, err error) {
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

func (s subjectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s subjectDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s subjectDo) Delete(models ...*model.Subject) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *subjectDo) withDO(do gen.Dao) *subjectDo {
	s.DO = *do.(*gen.DO)
	return s
}
