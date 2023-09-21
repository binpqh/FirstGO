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

func newClass(db *gorm.DB, opts ...gen.DOOption) class {
	_class := class{}

	_class.classDo.UseDB(db, opts...)
	_class.classDo.UseModel(&model.Class{})

	tableName := _class.classDo.TableName()
	_class.ALL = field.NewAsterisk(tableName)
	_class.ClassID = field.NewInt64(tableName, "ClassId")
	_class.Name = field.NewString(tableName, "Name")

	_class.fillFieldMap()

	return _class
}

type class struct {
	classDo

	ALL        field.Asterisk
	ClassID    field.Int64
	Name       field.String
	SchoolYear field.Int64
	Grade      field.Int64
	fieldMap   map[string]field.Expr
}

func (c class) Table(newTableName string) *class {
	c.classDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c class) As(alias string) *class {
	c.classDo.DO = *(c.classDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *class) updateTableName(table string) *class {
	c.ALL = field.NewAsterisk(table)
	c.ClassID = field.NewInt64(table, "ClassId")
	c.Name = field.NewString(table, "Name")
	c.SchoolYear = field.NewInt64(table, "SchoolYear")
	c.Grade = field.NewInt64(table, "Grade")
	c.fillFieldMap()

	return c
}

func (c *class) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *class) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 3)
	c.fieldMap["ClassId"] = c.ClassID
	c.fieldMap["Name"] = c.Name
}

func (c class) clone(db *gorm.DB) class {
	c.classDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c class) replaceDB(db *gorm.DB) class {
	c.classDo.ReplaceDB(db)
	return c
}

type classDo struct{ gen.DO }

type IClassDo interface {
	gen.SubQuery
	Debug() IClassDo
	WithContext(ctx context.Context) IClassDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IClassDo
	WriteDB() IClassDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IClassDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IClassDo
	Not(conds ...gen.Condition) IClassDo
	Or(conds ...gen.Condition) IClassDo
	Select(conds ...field.Expr) IClassDo
	Where(conds ...gen.Condition) IClassDo
	Order(conds ...field.Expr) IClassDo
	Distinct(cols ...field.Expr) IClassDo
	Omit(cols ...field.Expr) IClassDo
	Join(table schema.Tabler, on ...field.Expr) IClassDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IClassDo
	RightJoin(table schema.Tabler, on ...field.Expr) IClassDo
	Group(cols ...field.Expr) IClassDo
	Having(conds ...gen.Condition) IClassDo
	Limit(limit int) IClassDo
	Offset(offset int) IClassDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IClassDo
	Unscoped() IClassDo
	Create(values ...*model.Class) error
	CreateInBatches(values []*model.Class, batchSize int) error
	Save(values ...*model.Class) error
	First() (*model.Class, error)
	Take() (*model.Class, error)
	Last() (*model.Class, error)
	Find() ([]*model.Class, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Class, err error)
	FindInBatches(result *[]*model.Class, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Class) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IClassDo
	Assign(attrs ...field.AssignExpr) IClassDo
	Joins(fields ...field.RelationField) IClassDo
	Preload(fields ...field.RelationField) IClassDo
	FirstOrInit() (*model.Class, error)
	FirstOrCreate() (*model.Class, error)
	FindByPage(offset int, limit int) (result []*model.Class, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IClassDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c classDo) Debug() IClassDo {
	return c.withDO(c.DO.Debug())
}

func (c classDo) WithContext(ctx context.Context) IClassDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c classDo) ReadDB() IClassDo {
	return c.Clauses(dbresolver.Read)
}

func (c classDo) WriteDB() IClassDo {
	return c.Clauses(dbresolver.Write)
}

func (c classDo) Session(config *gorm.Session) IClassDo {
	return c.withDO(c.DO.Session(config))
}

func (c classDo) Clauses(conds ...clause.Expression) IClassDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c classDo) Returning(value interface{}, columns ...string) IClassDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c classDo) Not(conds ...gen.Condition) IClassDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c classDo) Or(conds ...gen.Condition) IClassDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c classDo) Select(conds ...field.Expr) IClassDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c classDo) Where(conds ...gen.Condition) IClassDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c classDo) Order(conds ...field.Expr) IClassDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c classDo) Distinct(cols ...field.Expr) IClassDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c classDo) Omit(cols ...field.Expr) IClassDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c classDo) Join(table schema.Tabler, on ...field.Expr) IClassDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c classDo) LeftJoin(table schema.Tabler, on ...field.Expr) IClassDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c classDo) RightJoin(table schema.Tabler, on ...field.Expr) IClassDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c classDo) Group(cols ...field.Expr) IClassDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c classDo) Having(conds ...gen.Condition) IClassDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c classDo) Limit(limit int) IClassDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c classDo) Offset(offset int) IClassDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c classDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IClassDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c classDo) Unscoped() IClassDo {
	return c.withDO(c.DO.Unscoped())
}

func (c classDo) Create(values ...*model.Class) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c classDo) CreateInBatches(values []*model.Class, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c classDo) Save(values ...*model.Class) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c classDo) First() (*model.Class, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Class), nil
	}
}

func (c classDo) Take() (*model.Class, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Class), nil
	}
}

func (c classDo) Last() (*model.Class, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Class), nil
	}
}

func (c classDo) Find() ([]*model.Class, error) {
	result, err := c.DO.Find()
	return result.([]*model.Class), err
}

func (c classDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Class, err error) {
	buf := make([]*model.Class, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c classDo) FindInBatches(result *[]*model.Class, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c classDo) Attrs(attrs ...field.AssignExpr) IClassDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c classDo) Assign(attrs ...field.AssignExpr) IClassDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c classDo) Joins(fields ...field.RelationField) IClassDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c classDo) Preload(fields ...field.RelationField) IClassDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c classDo) FirstOrInit() (*model.Class, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Class), nil
	}
}

func (c classDo) FirstOrCreate() (*model.Class, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Class), nil
	}
}

func (c classDo) FindByPage(offset int, limit int) (result []*model.Class, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c classDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c classDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c classDo) Delete(models ...*model.Class) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *classDo) withDO(do gen.Dao) *classDo {
	c.DO = *do.(*gen.DO)
	return c
}
