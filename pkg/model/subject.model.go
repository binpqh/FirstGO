package model

const TableNameSubject = "Subject"

// Subject mapped from table <Subject>
type Subject struct {
	SubjectID int64  `gorm:"column:SubjectId;primaryKey" json:"SubjectId"`
	Name      string `gorm:"column:Name" json:"Name"`
}

// TableName Subject's table name
func (*Subject) TableName() string {
	return TableNameSubject
}
