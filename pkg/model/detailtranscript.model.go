package model

const TableNameDetailTranscript = "DetailTranscript"

// DetailTranscript mapped from table <DetailTranscript>
type DetailTranscript struct {
	TranscriptID int64   `gorm:"column:TranscriptId;not null" json:"TranscriptId"`
	SubjectID    int64   `gorm:"column:SubjectId;not null" json:"SubjectId"`
	Score        float64 `gorm:"column:Score" json:"Score"`
}

// TableName DetailTranscript's table name
func (*DetailTranscript) TableName() string {
	return TableNameDetailTranscript
}
