package models

type User struct {
	ID        uint64 `gorm:"column:id;primaryKey"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"password" json:"password"`
	OTP       string `json:"-"`
	OTPHash   string `gorm:"column:otp_hash;size:64"`
	OTPExpire int64  `json:"-"`
}
