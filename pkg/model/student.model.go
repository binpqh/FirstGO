package model

import (
	"time"
)

const TableNameStudent = "Student"

// Student mapped from table <Student>
type Student struct {
	StudentID int64     `gorm:"column:StudentId;primaryKey" json:"StudentId"`
	Name      string    `gorm:"column:Name" json:"Name"`
	Birthday  time.Time `gorm:"column:Birthday" json:"Birthday"`
	ClassID   int64     `gorm:"column:ClassId" json:"ClassId"`
	UserID    int64     `gorm:"column:UserId" json:"UserId"`
}

// TableName Student's table name
func (*Student) TableName() string {
	return TableNameStudent
}
