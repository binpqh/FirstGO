package model

const TableNameTranscript = "Transcript"

// Transcript mapped from table <Transcript>
type Transcript struct {
	TranscriptID int64  `gorm:"column:TranscriptId;primaryKey" json:"TranscriptId"`
	StudentID    int64  `gorm:"column:StudentId" json:"StudentId"`
	ClassID      int64  `gorm:"column:ClassId" json:"ClassId"`
	Rank         int64  `gorm:"column:Rank" json:"Rank"`
	Description  string `gorm:"column:Description" json:"Description"`
}

// TableName Transcript's table name
func (*Transcript) TableName() string {
	return TableNameTranscript
}
