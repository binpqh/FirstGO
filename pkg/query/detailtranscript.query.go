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

func newDetailTranscript(db *gorm.DB, opts ...gen.DOOption) detailTranscript {
	_detailTranscript := detailTranscript{}

	_detailTranscript.detailTranscriptDo.UseDB(db, opts...)
	_detailTranscript.detailTranscriptDo.UseModel(&model.DetailTranscript{})

	tableName := _detailTranscript.detailTranscriptDo.TableName()
	_detailTranscript.ALL = field.NewAsterisk(tableName)
	_detailTranscript.TranscriptID = field.NewInt64(tableName, "TranscriptId")
	_detailTranscript.SubjectID = field.NewInt64(tableName, "SubjectId")
	_detailTranscript.Score = field.NewFloat64(tableName, "Score")

	_detailTranscript.fillFieldMap()

	return _detailTranscript
}

type detailTranscript struct {
	detailTranscriptDo

	ALL          field.Asterisk
	TranscriptID field.Int64
	SubjectID    field.Int64
	Score        field.Float64

	fieldMap map[string]field.Expr
}

func (d detailTranscript) Table(newTableName string) *detailTranscript {
	d.detailTranscriptDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d detailTranscript) As(alias string) *detailTranscript {
	d.detailTranscriptDo.DO = *(d.detailTranscriptDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *detailTranscript) updateTableName(table string) *detailTranscript {
	d.ALL = field.NewAsterisk(table)
	d.TranscriptID = field.NewInt64(table, "TranscriptId")
	d.SubjectID = field.NewInt64(table, "SubjectId")
	d.Score = field.NewFloat64(table, "Score")

	d.fillFieldMap()

	return d
}

func (d *detailTranscript) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *detailTranscript) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 3)
	d.fieldMap["TranscriptId"] = d.TranscriptID
	d.fieldMap["SubjectId"] = d.SubjectID
	d.fieldMap["Score"] = d.Score
}

func (d detailTranscript) clone(db *gorm.DB) detailTranscript {
	d.detailTranscriptDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d detailTranscript) replaceDB(db *gorm.DB) detailTranscript {
	d.detailTranscriptDo.ReplaceDB(db)
	return d
}

type detailTranscriptDo struct{ gen.DO }

type IDetailTranscriptDo interface {
	gen.SubQuery
	Debug() IDetailTranscriptDo
	WithContext(ctx context.Context) IDetailTranscriptDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDetailTranscriptDo
	WriteDB() IDetailTranscriptDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDetailTranscriptDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDetailTranscriptDo
	Not(conds ...gen.Condition) IDetailTranscriptDo
	Or(conds ...gen.Condition) IDetailTranscriptDo
	Select(conds ...field.Expr) IDetailTranscriptDo
	Where(conds ...gen.Condition) IDetailTranscriptDo
	Order(conds ...field.Expr) IDetailTranscriptDo
	Distinct(cols ...field.Expr) IDetailTranscriptDo
	Omit(cols ...field.Expr) IDetailTranscriptDo
	Join(table schema.Tabler, on ...field.Expr) IDetailTranscriptDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDetailTranscriptDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDetailTranscriptDo
	Group(cols ...field.Expr) IDetailTranscriptDo
	Having(conds ...gen.Condition) IDetailTranscriptDo
	Limit(limit int) IDetailTranscriptDo
	Offset(offset int) IDetailTranscriptDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDetailTranscriptDo
	Unscoped() IDetailTranscriptDo
	Create(values ...*model.DetailTranscript) error
	CreateInBatches(values []*model.DetailTranscript, batchSize int) error
	Save(values ...*model.DetailTranscript) error
	First() (*model.DetailTranscript, error)
	Take() (*model.DetailTranscript, error)
	Last() (*model.DetailTranscript, error)
	Find() ([]*model.DetailTranscript, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DetailTranscript, err error)
	FindInBatches(result *[]*model.DetailTranscript, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DetailTranscript) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDetailTranscriptDo
	Assign(attrs ...field.AssignExpr) IDetailTranscriptDo
	Joins(fields ...field.RelationField) IDetailTranscriptDo
	Preload(fields ...field.RelationField) IDetailTranscriptDo
	FirstOrInit() (*model.DetailTranscript, error)
	FirstOrCreate() (*model.DetailTranscript, error)
	FindByPage(offset int, limit int) (result []*model.DetailTranscript, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDetailTranscriptDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d detailTranscriptDo) Debug() IDetailTranscriptDo {
	return d.withDO(d.DO.Debug())
}

func (d detailTranscriptDo) WithContext(ctx context.Context) IDetailTranscriptDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d detailTranscriptDo) ReadDB() IDetailTranscriptDo {
	return d.Clauses(dbresolver.Read)
}

func (d detailTranscriptDo) WriteDB() IDetailTranscriptDo {
	return d.Clauses(dbresolver.Write)
}

func (d detailTranscriptDo) Session(config *gorm.Session) IDetailTranscriptDo {
	return d.withDO(d.DO.Session(config))
}

func (d detailTranscriptDo) Clauses(conds ...clause.Expression) IDetailTranscriptDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d detailTranscriptDo) Returning(value interface{}, columns ...string) IDetailTranscriptDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d detailTranscriptDo) Not(conds ...gen.Condition) IDetailTranscriptDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d detailTranscriptDo) Or(conds ...gen.Condition) IDetailTranscriptDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d detailTranscriptDo) Select(conds ...field.Expr) IDetailTranscriptDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d detailTranscriptDo) Where(conds ...gen.Condition) IDetailTranscriptDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d detailTranscriptDo) Order(conds ...field.Expr) IDetailTranscriptDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d detailTranscriptDo) Distinct(cols ...field.Expr) IDetailTranscriptDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d detailTranscriptDo) Omit(cols ...field.Expr) IDetailTranscriptDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d detailTranscriptDo) Join(table schema.Tabler, on ...field.Expr) IDetailTranscriptDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d detailTranscriptDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDetailTranscriptDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d detailTranscriptDo) RightJoin(table schema.Tabler, on ...field.Expr) IDetailTranscriptDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d detailTranscriptDo) Group(cols ...field.Expr) IDetailTranscriptDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d detailTranscriptDo) Having(conds ...gen.Condition) IDetailTranscriptDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d detailTranscriptDo) Limit(limit int) IDetailTranscriptDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d detailTranscriptDo) Offset(offset int) IDetailTranscriptDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d detailTranscriptDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDetailTranscriptDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d detailTranscriptDo) Unscoped() IDetailTranscriptDo {
	return d.withDO(d.DO.Unscoped())
}

func (d detailTranscriptDo) Create(values ...*model.DetailTranscript) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d detailTranscriptDo) CreateInBatches(values []*model.DetailTranscript, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d detailTranscriptDo) Save(values ...*model.DetailTranscript) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d detailTranscriptDo) First() (*model.DetailTranscript, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetailTranscript), nil
	}
}

func (d detailTranscriptDo) Take() (*model.DetailTranscript, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetailTranscript), nil
	}
}

func (d detailTranscriptDo) Last() (*model.DetailTranscript, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetailTranscript), nil
	}
}

func (d detailTranscriptDo) Find() ([]*model.DetailTranscript, error) {
	result, err := d.DO.Find()
	return result.([]*model.DetailTranscript), err
}

func (d detailTranscriptDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DetailTranscript, err error) {
	buf := make([]*model.DetailTranscript, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d detailTranscriptDo) FindInBatches(result *[]*model.DetailTranscript, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d detailTranscriptDo) Attrs(attrs ...field.AssignExpr) IDetailTranscriptDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d detailTranscriptDo) Assign(attrs ...field.AssignExpr) IDetailTranscriptDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d detailTranscriptDo) Joins(fields ...field.RelationField) IDetailTranscriptDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d detailTranscriptDo) Preload(fields ...field.RelationField) IDetailTranscriptDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d detailTranscriptDo) FirstOrInit() (*model.DetailTranscript, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetailTranscript), nil
	}
}

func (d detailTranscriptDo) FirstOrCreate() (*model.DetailTranscript, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetailTranscript), nil
	}
}

func (d detailTranscriptDo) FindByPage(offset int, limit int) (result []*model.DetailTranscript, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d detailTranscriptDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d detailTranscriptDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d detailTranscriptDo) Delete(models ...*model.DetailTranscript) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *detailTranscriptDo) withDO(do gen.Dao) *detailTranscriptDo {
	d.DO = *do.(*gen.DO)
	return d
}
