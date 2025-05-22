package model

type Verification struct {
	BaseModel
	Phone string `gorm:"column:phone" json:"phone"`
	Code  string `gorm:"column:code" json:"code"`
	Ip    string `gorm:"column:ip" json:"ip"`
}

func NewVerification() *Verification {
	return &Verification{}
}

// TableName
func (u *Verification) TableName() string {
	return "verifications"
}
