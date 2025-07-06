package gormmodel

import (
	"time"

	"github.com/wang900115/Perry/internal/domain/entity"
	"github.com/wang900115/utils/convert"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string
	Password    string
	FullName    string
	NickName    string
	AvatarURL   string
	Phone       string
	Email       string
	Location    Location
	Description string
}

type UserStatus struct {
	User
	Device     string
	LastIP     string
	LastLogin  time.Time
	LastLogout time.Time
}

type Location struct {
	Country   string
	City      string
	Latitude  float64
	Longitude float64
}

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
		User:       *u.User.ToDomain(),
		Device:     u.Device,
		LastIP:     u.LastIP,
		LastLogin:  convert.FromTimeTimeToTimestamp(u.LastLogin),
		LastLogout: convert.FromTimeTimeToTimestamp(u.LastLogout),
	}
}
