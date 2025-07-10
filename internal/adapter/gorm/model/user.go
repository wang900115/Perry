package gormmodel

import (
	"time"

	"github.com/wang900115/Perry/internal/domain/entity"
	"github.com/wang900115/utils/convert"
	"gorm.io/gorm"
)

// #region 表格欄位
type User struct {
	gorm.Model
	Username    string   `gorm:"column:username"`
	Password    string   `gorm:"column:password"`
	FullName    string   `gorm:"column:full_name"`
	NickName    string   `gorm:"column:nick_name"`
	AvatarURL   string   `gorm:"column:avatar_url"`
	Phone       string   `gorm:"column:phone"`
	Email       string   `gorm:"column:email"`
	Location    Location `gorm:"embedded"`
	Description string   `gorm:"column:description"`
}

type UserStatus struct {
	UserID uint `gorm:"column:user_id"`
	User   User `gorm:"foreignKey:UserID"`

	Device     string    `gorm:"column:device"`
	LastIP     string    `gorm:"column:last_ip"`
	LastLogin  time.Time `gorm:"column:last_login"`
	LastLogout time.Time `gorm:"column:last_logout"`
}

// #endregion

// #region 嵌入式欄位
type Location struct {
	Country   string  `gorm:"column:country"`
	City      string  `gorm:"column:city"`
	Latitude  float64 `gorm:"column:latitude"`
	Longitude float64 `gorm:"column:longitude"`
}

// #endregion

// #region 表格名稱
func (User) TableName() string {
	return "user"
}

func (UserStatus) TableName() string {
	return "user_status"
}

// #endregion

// #region -> Domain
func (u *User) ToDomain() *entity.User {
	return &entity.User{
		UserId:    u.ID,
		Username:  u.Username,
		FullName:  u.FullName,
		NickName:  u.NickName,
		AvatarURL: u.AvatarURL,
		Phone:     u.Phone,
		Email:     u.Email,
		Country:   u.Location.City,
		City:      u.Location.Country,
	}
}

func (u *UserStatus) ToDomain() *entity.UserStatus {
	return &entity.UserStatus{
		UserId:     u.UserID,
		Device:     u.Device,
		LastIP:     u.LastIP,
		LastLogin:  convert.FromTimeTimeToTimestamp(u.LastLogin),
		LastLogout: convert.FromTimeTimeToTimestamp(u.LastLogout),
	}
}

// #endregion
