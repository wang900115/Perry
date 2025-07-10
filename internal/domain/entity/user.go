package entity

type User struct {
	UserId       uint
	Username     string
	PasswordHash string
	FullName     string
	NickName     string
	AvatarURL    string
	Phone        string
	Email        string
	Country      string
	City         string
}

type UserStatus struct {
	UserId     uint
	Device     string
	LastIP     string
	LastLogin  int64
	LastLogout int64
}
