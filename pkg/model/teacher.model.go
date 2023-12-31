package model

import (
	"time"
)

const TableNameTeacher = "Teacher"

// Teacher mapped from table <Teacher>
type Teacher struct {
	TeacherID int64     `gorm:"column:TeacherId;primaryKey" json:"TeacherId"`
	Name      string    `gorm:"column:Name" json:"Name"`
	Birthday  time.Time `gorm:"column:Birthday" json:"Birthday"`
	UserID    int64     `gorm:"column:UserId" json:"UserId"`
}

// TableName Teacher's table name
func (*Teacher) TableName() string {
	return TableNameTeacher
}
