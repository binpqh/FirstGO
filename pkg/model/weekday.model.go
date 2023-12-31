package model

const TableNameWeekday = "Weekday"

// Weekday mapped from table <Weekday>
type Weekday struct {
	WeekdayID      int64  `gorm:"column:WeekdayId;primaryKey" json:"WeekdayId"`
	Name           string `gorm:"column:Name" json:"Name"`
	MaxLessonInDay int64  `gorm:"column:MaxLessonInDay" json:"MaxLessonInDay"`
}

// TableName Weekday's table name
func (*Weekday) TableName() string {
	return TableNameWeekday
}
