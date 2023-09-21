package model

const TableNameClass = "Class"

// Class mapped from table <Class>
type Class struct {
	ClassID int64  `gorm:"column:ClassId;primaryKey" json:"ClassId"`
	Name    string `gorm:"column:Name" json:"Name"`
}

// TableName Class's table name
func (*Class) TableName() string {
	return TableNameClass
}
