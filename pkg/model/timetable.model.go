package model

const TableNameTimeTable = "TimeTable"

// TimeTable mapped from table <TimeTable>
type TimeTable struct {
	TimeTableID  int64 `gorm:"column:TimeTableId;primaryKey" json:"TimeTableId"`
	ClassID      int64 `gorm:"column:ClassId;not null" json:"ClassId"`
	SubjectID    int64 `gorm:"column:SubjectId;not null" json:"SubjectId"`
	WeekdayID    int64 `gorm:"column:WeekdayId" json:"WeekdayId"`
	LessonIndex  int64 `gorm:"column:LessonIndex" json:"LessonIndex"`
	LessonNumber int64 `gorm:"column:LessonNumber" json:"LessonNumber"`
}

// TableName TimeTable's table name
func (*TimeTable) TableName() string {
	return TableNameTimeTable
}
