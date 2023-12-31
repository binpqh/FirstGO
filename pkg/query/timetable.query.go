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

func newTimeTable(db *gorm.DB, opts ...gen.DOOption) timeTable {
	_timeTable := timeTable{}

	_timeTable.timeTableDo.UseDB(db, opts...)
	_timeTable.timeTableDo.UseModel(&model.TimeTable{})

	tableName := _timeTable.timeTableDo.TableName()
	_timeTable.ALL = field.NewAsterisk(tableName)
	_timeTable.TimeTableID = field.NewInt64(tableName, "TimeTableId")
	_timeTable.ClassID = field.NewInt64(tableName, "ClassId")
	_timeTable.SubjectID = field.NewInt64(tableName, "SubjectId")
	_timeTable.WeekdayID = field.NewInt64(tableName, "WeekdayId")
	_timeTable.LessonIndex = field.NewInt64(tableName, "LessonIndex")
	_timeTable.LessonNumber = field.NewInt64(tableName, "LessonNumber")

	_timeTable.fillFieldMap()

	return _timeTable
}

type timeTable struct {
	timeTableDo

	ALL          field.Asterisk
	TimeTableID  field.Int64
	ClassID      field.Int64
	SubjectID    field.Int64
	WeekdayID    field.Int64
	LessonIndex  field.Int64
	LessonNumber field.Int64

	fieldMap map[string]field.Expr
}

func (t timeTable) Table(newTableName string) *timeTable {
	t.timeTableDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t timeTable) As(alias string) *timeTable {
	t.timeTableDo.DO = *(t.timeTableDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *timeTable) updateTableName(table string) *timeTable {
	t.ALL = field.NewAsterisk(table)
	t.TimeTableID = field.NewInt64(table, "TimeTableId")
	t.ClassID = field.NewInt64(table, "ClassId")
	t.SubjectID = field.NewInt64(table, "SubjectId")
	t.WeekdayID = field.NewInt64(table, "WeekdayId")
	t.LessonIndex = field.NewInt64(table, "LessonIndex")
	t.LessonNumber = field.NewInt64(table, "LessonNumber")

	t.fillFieldMap()

	return t
}

func (t *timeTable) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *timeTable) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["TimeTableId"] = t.TimeTableID
	t.fieldMap["ClassId"] = t.ClassID
	t.fieldMap["SubjectId"] = t.SubjectID
	t.fieldMap["WeekdayId"] = t.WeekdayID
	t.fieldMap["LessonIndex"] = t.LessonIndex
	t.fieldMap["LessonNumber"] = t.LessonNumber
}

func (t timeTable) clone(db *gorm.DB) timeTable {
	t.timeTableDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t timeTable) replaceDB(db *gorm.DB) timeTable {
	t.timeTableDo.ReplaceDB(db)
	return t
}

type timeTableDo struct{ gen.DO }

type ITimeTableDo interface {
	gen.SubQuery
	Debug() ITimeTableDo
	WithContext(ctx context.Context) ITimeTableDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITimeTableDo
	WriteDB() ITimeTableDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITimeTableDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITimeTableDo
	Not(conds ...gen.Condition) ITimeTableDo
	Or(conds ...gen.Condition) ITimeTableDo
	Select(conds ...field.Expr) ITimeTableDo
	Where(conds ...gen.Condition) ITimeTableDo
	Order(conds ...field.Expr) ITimeTableDo
	Distinct(cols ...field.Expr) ITimeTableDo
	Omit(cols ...field.Expr) ITimeTableDo
	Join(table schema.Tabler, on ...field.Expr) ITimeTableDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITimeTableDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITimeTableDo
	Group(cols ...field.Expr) ITimeTableDo
	Having(conds ...gen.Condition) ITimeTableDo
	Limit(limit int) ITimeTableDo
	Offset(offset int) ITimeTableDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITimeTableDo
	Unscoped() ITimeTableDo
	Create(values ...*model.TimeTable) error
	CreateInBatches(values []*model.TimeTable, batchSize int) error
	Save(values ...*model.TimeTable) error
	First() (*model.TimeTable, error)
	Take() (*model.TimeTable, error)
	Last() (*model.TimeTable, error)
	Find() ([]*model.TimeTable, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TimeTable, err error)
	FindInBatches(result *[]*model.TimeTable, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TimeTable) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITimeTableDo
	Assign(attrs ...field.AssignExpr) ITimeTableDo
	Joins(fields ...field.RelationField) ITimeTableDo
	Preload(fields ...field.RelationField) ITimeTableDo
	FirstOrInit() (*model.TimeTable, error)
	FirstOrCreate() (*model.TimeTable, error)
	FindByPage(offset int, limit int) (result []*model.TimeTable, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITimeTableDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t timeTableDo) Debug() ITimeTableDo {
	return t.withDO(t.DO.Debug())
}

func (t timeTableDo) WithContext(ctx context.Context) ITimeTableDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t timeTableDo) ReadDB() ITimeTableDo {
	return t.Clauses(dbresolver.Read)
}

func (t timeTableDo) WriteDB() ITimeTableDo {
	return t.Clauses(dbresolver.Write)
}

func (t timeTableDo) Session(config *gorm.Session) ITimeTableDo {
	return t.withDO(t.DO.Session(config))
}

func (t timeTableDo) Clauses(conds ...clause.Expression) ITimeTableDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t timeTableDo) Returning(value interface{}, columns ...string) ITimeTableDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t timeTableDo) Not(conds ...gen.Condition) ITimeTableDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t timeTableDo) Or(conds ...gen.Condition) ITimeTableDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t timeTableDo) Select(conds ...field.Expr) ITimeTableDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t timeTableDo) Where(conds ...gen.Condition) ITimeTableDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t timeTableDo) Order(conds ...field.Expr) ITimeTableDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t timeTableDo) Distinct(cols ...field.Expr) ITimeTableDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t timeTableDo) Omit(cols ...field.Expr) ITimeTableDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t timeTableDo) Join(table schema.Tabler, on ...field.Expr) ITimeTableDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t timeTableDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITimeTableDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t timeTableDo) RightJoin(table schema.Tabler, on ...field.Expr) ITimeTableDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t timeTableDo) Group(cols ...field.Expr) ITimeTableDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t timeTableDo) Having(conds ...gen.Condition) ITimeTableDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t timeTableDo) Limit(limit int) ITimeTableDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t timeTableDo) Offset(offset int) ITimeTableDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t timeTableDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITimeTableDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t timeTableDo) Unscoped() ITimeTableDo {
	return t.withDO(t.DO.Unscoped())
}

func (t timeTableDo) Create(values ...*model.TimeTable) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t timeTableDo) CreateInBatches(values []*model.TimeTable, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t timeTableDo) Save(values ...*model.TimeTable) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t timeTableDo) First() (*model.TimeTable, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeTable), nil
	}
}

func (t timeTableDo) Take() (*model.TimeTable, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeTable), nil
	}
}

func (t timeTableDo) Last() (*model.TimeTable, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeTable), nil
	}
}

func (t timeTableDo) Find() ([]*model.TimeTable, error) {
	result, err := t.DO.Find()
	return result.([]*model.TimeTable), err
}

func (t timeTableDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TimeTable, err error) {
	buf := make([]*model.TimeTable, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t timeTableDo) FindInBatches(result *[]*model.TimeTable, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t timeTableDo) Attrs(attrs ...field.AssignExpr) ITimeTableDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t timeTableDo) Assign(attrs ...field.AssignExpr) ITimeTableDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t timeTableDo) Joins(fields ...field.RelationField) ITimeTableDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t timeTableDo) Preload(fields ...field.RelationField) ITimeTableDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t timeTableDo) FirstOrInit() (*model.TimeTable, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeTable), nil
	}
}

func (t timeTableDo) FirstOrCreate() (*model.TimeTable, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeTable), nil
	}
}

func (t timeTableDo) FindByPage(offset int, limit int) (result []*model.TimeTable, count int64, err error) {
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

func (t timeTableDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t timeTableDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t timeTableDo) Delete(models ...*model.TimeTable) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *timeTableDo) withDO(do gen.Dao) *timeTableDo {
	t.DO = *do.(*gen.DO)
	return t
}
