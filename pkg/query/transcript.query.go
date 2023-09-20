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

func newTranscript(db *gorm.DB, opts ...gen.DOOption) transcript {
	_transcript := transcript{}

	_transcript.transcriptDo.UseDB(db, opts...)
	_transcript.transcriptDo.UseModel(&model.Transcript{})

	tableName := _transcript.transcriptDo.TableName()
	_transcript.ALL = field.NewAsterisk(tableName)
	_transcript.TranscriptID = field.NewInt64(tableName, "TranscriptId")
	_transcript.StudentID = field.NewInt64(tableName, "StudentId")
	_transcript.ClassID = field.NewInt64(tableName, "ClassId")
	_transcript.Rank = field.NewInt64(tableName, "Rank")
	_transcript.Description = field.NewString(tableName, "Description")

	_transcript.fillFieldMap()

	return _transcript
}

type transcript struct {
	transcriptDo

	ALL          field.Asterisk
	TranscriptID field.Int64
	StudentID    field.Int64
	ClassID      field.Int64
	Rank         field.Int64
	Description  field.String

	fieldMap map[string]field.Expr
}

func (t transcript) Table(newTableName string) *transcript {
	t.transcriptDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t transcript) As(alias string) *transcript {
	t.transcriptDo.DO = *(t.transcriptDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *transcript) updateTableName(table string) *transcript {
	t.ALL = field.NewAsterisk(table)
	t.TranscriptID = field.NewInt64(table, "TranscriptId")
	t.StudentID = field.NewInt64(table, "StudentId")
	t.ClassID = field.NewInt64(table, "ClassId")
	t.Rank = field.NewInt64(table, "Rank")
	t.Description = field.NewString(table, "Description")

	t.fillFieldMap()

	return t
}

func (t *transcript) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *transcript) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["TranscriptId"] = t.TranscriptID
	t.fieldMap["StudentId"] = t.StudentID
	t.fieldMap["ClassId"] = t.ClassID
	t.fieldMap["Rank"] = t.Rank
	t.fieldMap["Description"] = t.Description
}

func (t transcript) clone(db *gorm.DB) transcript {
	t.transcriptDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t transcript) replaceDB(db *gorm.DB) transcript {
	t.transcriptDo.ReplaceDB(db)
	return t
}

type transcriptDo struct{ gen.DO }

type ITranscriptDo interface {
	gen.SubQuery
	Debug() ITranscriptDo
	WithContext(ctx context.Context) ITranscriptDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITranscriptDo
	WriteDB() ITranscriptDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITranscriptDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITranscriptDo
	Not(conds ...gen.Condition) ITranscriptDo
	Or(conds ...gen.Condition) ITranscriptDo
	Select(conds ...field.Expr) ITranscriptDo
	Where(conds ...gen.Condition) ITranscriptDo
	Order(conds ...field.Expr) ITranscriptDo
	Distinct(cols ...field.Expr) ITranscriptDo
	Omit(cols ...field.Expr) ITranscriptDo
	Join(table schema.Tabler, on ...field.Expr) ITranscriptDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITranscriptDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITranscriptDo
	Group(cols ...field.Expr) ITranscriptDo
	Having(conds ...gen.Condition) ITranscriptDo
	Limit(limit int) ITranscriptDo
	Offset(offset int) ITranscriptDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITranscriptDo
	Unscoped() ITranscriptDo
	Create(values ...*model.Transcript) error
	CreateInBatches(values []*model.Transcript, batchSize int) error
	Save(values ...*model.Transcript) error
	First() (*model.Transcript, error)
	Take() (*model.Transcript, error)
	Last() (*model.Transcript, error)
	Find() ([]*model.Transcript, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Transcript, err error)
	FindInBatches(result *[]*model.Transcript, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Transcript) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITranscriptDo
	Assign(attrs ...field.AssignExpr) ITranscriptDo
	Joins(fields ...field.RelationField) ITranscriptDo
	Preload(fields ...field.RelationField) ITranscriptDo
	FirstOrInit() (*model.Transcript, error)
	FirstOrCreate() (*model.Transcript, error)
	FindByPage(offset int, limit int) (result []*model.Transcript, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITranscriptDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t transcriptDo) Debug() ITranscriptDo {
	return t.withDO(t.DO.Debug())
}

func (t transcriptDo) WithContext(ctx context.Context) ITranscriptDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t transcriptDo) ReadDB() ITranscriptDo {
	return t.Clauses(dbresolver.Read)
}

func (t transcriptDo) WriteDB() ITranscriptDo {
	return t.Clauses(dbresolver.Write)
}

func (t transcriptDo) Session(config *gorm.Session) ITranscriptDo {
	return t.withDO(t.DO.Session(config))
}

func (t transcriptDo) Clauses(conds ...clause.Expression) ITranscriptDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t transcriptDo) Returning(value interface{}, columns ...string) ITranscriptDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t transcriptDo) Not(conds ...gen.Condition) ITranscriptDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t transcriptDo) Or(conds ...gen.Condition) ITranscriptDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t transcriptDo) Select(conds ...field.Expr) ITranscriptDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t transcriptDo) Where(conds ...gen.Condition) ITranscriptDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t transcriptDo) Order(conds ...field.Expr) ITranscriptDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t transcriptDo) Distinct(cols ...field.Expr) ITranscriptDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t transcriptDo) Omit(cols ...field.Expr) ITranscriptDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t transcriptDo) Join(table schema.Tabler, on ...field.Expr) ITranscriptDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t transcriptDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITranscriptDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t transcriptDo) RightJoin(table schema.Tabler, on ...field.Expr) ITranscriptDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t transcriptDo) Group(cols ...field.Expr) ITranscriptDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t transcriptDo) Having(conds ...gen.Condition) ITranscriptDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t transcriptDo) Limit(limit int) ITranscriptDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t transcriptDo) Offset(offset int) ITranscriptDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t transcriptDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITranscriptDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t transcriptDo) Unscoped() ITranscriptDo {
	return t.withDO(t.DO.Unscoped())
}

func (t transcriptDo) Create(values ...*model.Transcript) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t transcriptDo) CreateInBatches(values []*model.Transcript, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t transcriptDo) Save(values ...*model.Transcript) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t transcriptDo) First() (*model.Transcript, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Transcript), nil
	}
}

func (t transcriptDo) Take() (*model.Transcript, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Transcript), nil
	}
}

func (t transcriptDo) Last() (*model.Transcript, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Transcript), nil
	}
}

func (t transcriptDo) Find() ([]*model.Transcript, error) {
	result, err := t.DO.Find()
	return result.([]*model.Transcript), err
}

func (t transcriptDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Transcript, err error) {
	buf := make([]*model.Transcript, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t transcriptDo) FindInBatches(result *[]*model.Transcript, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t transcriptDo) Attrs(attrs ...field.AssignExpr) ITranscriptDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t transcriptDo) Assign(attrs ...field.AssignExpr) ITranscriptDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t transcriptDo) Joins(fields ...field.RelationField) ITranscriptDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t transcriptDo) Preload(fields ...field.RelationField) ITranscriptDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t transcriptDo) FirstOrInit() (*model.Transcript, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Transcript), nil
	}
}

func (t transcriptDo) FirstOrCreate() (*model.Transcript, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Transcript), nil
	}
}

func (t transcriptDo) FindByPage(offset int, limit int) (result []*model.Transcript, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t transcriptDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t transcriptDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t transcriptDo) Delete(models ...*model.Transcript) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *transcriptDo) withDO(do gen.Dao) *transcriptDo {
	t.DO = *do.(*gen.DO)
	return t
}
