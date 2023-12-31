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

func newWeekday(db *gorm.DB, opts ...gen.DOOption) weekday {
	_weekday := weekday{}

	_weekday.weekdayDo.UseDB(db, opts...)
	_weekday.weekdayDo.UseModel(&model.Weekday{})

	tableName := _weekday.weekdayDo.TableName()
	_weekday.ALL = field.NewAsterisk(tableName)
	_weekday.WeekdayID = field.NewInt64(tableName, "WeekdayId")
	_weekday.Name = field.NewString(tableName, "Name")
	_weekday.MaxLessonInDay = field.NewInt64(tableName, "MaxLessonInDay")

	_weekday.fillFieldMap()

	return _weekday
}

type weekday struct {
	weekdayDo

	ALL            field.Asterisk
	WeekdayID      field.Int64
	Name           field.String
	MaxLessonInDay field.Int64

	fieldMap map[string]field.Expr
}

func (w weekday) Table(newTableName string) *weekday {
	w.weekdayDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w weekday) As(alias string) *weekday {
	w.weekdayDo.DO = *(w.weekdayDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *weekday) updateTableName(table string) *weekday {
	w.ALL = field.NewAsterisk(table)
	w.WeekdayID = field.NewInt64(table, "WeekdayId")
	w.Name = field.NewString(table, "Name")
	w.MaxLessonInDay = field.NewInt64(table, "MaxLessonInDay")

	w.fillFieldMap()

	return w
}

func (w *weekday) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *weekday) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 3)
	w.fieldMap["WeekdayId"] = w.WeekdayID
	w.fieldMap["Name"] = w.Name
	w.fieldMap["MaxLessonInDay"] = w.MaxLessonInDay
}

func (w weekday) clone(db *gorm.DB) weekday {
	w.weekdayDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w weekday) replaceDB(db *gorm.DB) weekday {
	w.weekdayDo.ReplaceDB(db)
	return w
}

type weekdayDo struct{ gen.DO }

type IWeekdayDo interface {
	gen.SubQuery
	Debug() IWeekdayDo
	WithContext(ctx context.Context) IWeekdayDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWeekdayDo
	WriteDB() IWeekdayDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWeekdayDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWeekdayDo
	Not(conds ...gen.Condition) IWeekdayDo
	Or(conds ...gen.Condition) IWeekdayDo
	Select(conds ...field.Expr) IWeekdayDo
	Where(conds ...gen.Condition) IWeekdayDo
	Order(conds ...field.Expr) IWeekdayDo
	Distinct(cols ...field.Expr) IWeekdayDo
	Omit(cols ...field.Expr) IWeekdayDo
	Join(table schema.Tabler, on ...field.Expr) IWeekdayDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWeekdayDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWeekdayDo
	Group(cols ...field.Expr) IWeekdayDo
	Having(conds ...gen.Condition) IWeekdayDo
	Limit(limit int) IWeekdayDo
	Offset(offset int) IWeekdayDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWeekdayDo
	Unscoped() IWeekdayDo
	Create(values ...*model.Weekday) error
	CreateInBatches(values []*model.Weekday, batchSize int) error
	Save(values ...*model.Weekday) error
	First() (*model.Weekday, error)
	Take() (*model.Weekday, error)
	Last() (*model.Weekday, error)
	Find() ([]*model.Weekday, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Weekday, err error)
	FindInBatches(result *[]*model.Weekday, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Weekday) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWeekdayDo
	Assign(attrs ...field.AssignExpr) IWeekdayDo
	Joins(fields ...field.RelationField) IWeekdayDo
	Preload(fields ...field.RelationField) IWeekdayDo
	FirstOrInit() (*model.Weekday, error)
	FirstOrCreate() (*model.Weekday, error)
	FindByPage(offset int, limit int) (result []*model.Weekday, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWeekdayDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w weekdayDo) Debug() IWeekdayDo {
	return w.withDO(w.DO.Debug())
}

func (w weekdayDo) WithContext(ctx context.Context) IWeekdayDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w weekdayDo) ReadDB() IWeekdayDo {
	return w.Clauses(dbresolver.Read)
}

func (w weekdayDo) WriteDB() IWeekdayDo {
	return w.Clauses(dbresolver.Write)
}

func (w weekdayDo) Session(config *gorm.Session) IWeekdayDo {
	return w.withDO(w.DO.Session(config))
}

func (w weekdayDo) Clauses(conds ...clause.Expression) IWeekdayDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w weekdayDo) Returning(value interface{}, columns ...string) IWeekdayDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w weekdayDo) Not(conds ...gen.Condition) IWeekdayDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w weekdayDo) Or(conds ...gen.Condition) IWeekdayDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w weekdayDo) Select(conds ...field.Expr) IWeekdayDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w weekdayDo) Where(conds ...gen.Condition) IWeekdayDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w weekdayDo) Order(conds ...field.Expr) IWeekdayDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w weekdayDo) Distinct(cols ...field.Expr) IWeekdayDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w weekdayDo) Omit(cols ...field.Expr) IWeekdayDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w weekdayDo) Join(table schema.Tabler, on ...field.Expr) IWeekdayDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w weekdayDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWeekdayDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w weekdayDo) RightJoin(table schema.Tabler, on ...field.Expr) IWeekdayDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w weekdayDo) Group(cols ...field.Expr) IWeekdayDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w weekdayDo) Having(conds ...gen.Condition) IWeekdayDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w weekdayDo) Limit(limit int) IWeekdayDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w weekdayDo) Offset(offset int) IWeekdayDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w weekdayDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWeekdayDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w weekdayDo) Unscoped() IWeekdayDo {
	return w.withDO(w.DO.Unscoped())
}

func (w weekdayDo) Create(values ...*model.Weekday) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w weekdayDo) CreateInBatches(values []*model.Weekday, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w weekdayDo) Save(values ...*model.Weekday) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w weekdayDo) First() (*model.Weekday, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Weekday), nil
	}
}

func (w weekdayDo) Take() (*model.Weekday, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Weekday), nil
	}
}

func (w weekdayDo) Last() (*model.Weekday, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Weekday), nil
	}
}

func (w weekdayDo) Find() ([]*model.Weekday, error) {
	result, err := w.DO.Find()
	return result.([]*model.Weekday), err
}

func (w weekdayDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Weekday, err error) {
	buf := make([]*model.Weekday, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w weekdayDo) FindInBatches(result *[]*model.Weekday, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w weekdayDo) Attrs(attrs ...field.AssignExpr) IWeekdayDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w weekdayDo) Assign(attrs ...field.AssignExpr) IWeekdayDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w weekdayDo) Joins(fields ...field.RelationField) IWeekdayDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w weekdayDo) Preload(fields ...field.RelationField) IWeekdayDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w weekdayDo) FirstOrInit() (*model.Weekday, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Weekday), nil
	}
}

func (w weekdayDo) FirstOrCreate() (*model.Weekday, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Weekday), nil
	}
}

func (w weekdayDo) FindByPage(offset int, limit int) (result []*model.Weekday, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w weekdayDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w weekdayDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w weekdayDo) Delete(models ...*model.Weekday) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *weekdayDo) withDO(do gen.Dao) *weekdayDo {
	w.DO = *do.(*gen.DO)
	return w
}
