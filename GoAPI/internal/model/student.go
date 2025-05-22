package model

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(190);default:''" json:"name"`
	//LastName  string `gorm:"column:last_name;type:varchar(190);default:''" json:"last_name"`
	Email     string `gorm:"column:email;type:varchar(120);unique;default:''" json:"email"`
	Phone     string `gorm:"column:phone;type:varchar(120);unique;default:''" json:"phone"`
	PassWord  string `gorm:"column:password;type:varchar(255);default:''" json:"-"`
	Gender    string `gorm:"column:gender;type:varchar(190);default:'male'" json:"gender"`
	Level     string `gorm:"column:level;type:varchar(500);default:'1'" json:"address"`
	Lesson    string `gorm:"column:lesson;type:varchar(500);default:'0'" json:"lesson"`
	Score     int    `gorm:"column:score;type:int(4);default:10" json:"score"`
	Star      int    `gorm:"column:stars;type:int(6);default:30" json:"star"`
	Fcmtoken  string `gorm:"column:fcmtoken;type:varchar(190);default:''" json:"fcmtoken"`
	Token     string `gorm:"column:token;type:varchar(190);unique:true;default:''" json:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func NewStudent() *Student {
	return &Student{}
}

func (c *Student) TableName() string {
	return "students"
}

// GetUserById for auth file
func (c *Student) GetStudentById(id int) *Student {
	if err := c.DB().First(c, id).Error; err != nil {
		return nil
	}
	return c
}
