package model

const TableNameUser = "User"

// User mapped from table <User>
type User struct {
	UserID   int64  `gorm:"column:UserId;primaryKey" json:"UserId"`
	Password string `gorm:"column:Password" json:"Password"`
	IsAdmin  bool   `gorm:"column:IsAdmin" json:"IsAdmin"`
	Status   int64  `gorm:"column:Status" json:"Status"`
	Username string `gorm:"column:Username" json:"Username"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
