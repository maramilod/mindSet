package form

type StudentRegister struct {
	Name     string `form:"name" json:"name"  binding:"required"`
	Email    string `form:"email" json:"email"`
	Phone    string `form:"phone" json:"phone"  binding:"required"`
	Gender   string `gorm:"column:gender;type:varchar(190);default:'male'" json:"gender"`
	PassWord string `form:"password" json:"password"  binding:"required,min=6"`
}

func StudentRegisterForm() *StudentRegister {
	return &StudentRegister{}
}

type Login struct {
	Phone    string `form:"phone" json:"phone"  binding:"required"`
	PassWord string `form:"password" json:"password"  binding:"required,min=6"`
}

func StudentLoginForm() *Login {
	return &Login{}
}

type Forget struct {
	Phone            string `form:"phone" json:"phone"  binding:"required"`
	VerificationCode string `form:"verification_code" json:"verification_code"  binding:"required"`
	PassWord         string `form:"password" json:"password"  binding:"required,min=6"`
}

func StudentForgetForm() *Forget {
	return &Forget{}
}

type Verification struct {
	Phone string `form:"phone" json:"phone"  binding:"required"`
}

func StudentVerificationForm() *Verification {
	return &Verification{}
}

type UpdateInfo struct {
	Name        string `form:"name" json:"name"  binding:"required"`
	Description string `form:"description" json:"description"  binding:"required"`
	Gender      int    `form:"gender" json:"gender"  binding:"required"`
}

func StudentUpdateInfoForm() *UpdateInfo {
	return &UpdateInfo{}
}

type FcmToken struct {
	Fcmtoken string `form:"fcmtoken" json:"fcmtoken"  binding:"required"`
}

func FcmtokenForm() *FcmToken {
	return &FcmToken{}
}

type UpdatePassword struct {
	Password   string `form:"password" json:"password"  binding:"required,min=6"`     //
	RePassword string `form:"repassword" json:"repassword"  binding:"required,min=6"` //
}

func StudentUpdatePasswordForm() *UpdatePassword {
	return &UpdatePassword{}
}
