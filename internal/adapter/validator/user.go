package validator

type RegisterRequest struct {
	Username    string           `json:"username" validate:"required"`
	Password    string           `json:"password" validate:"required"`
	FullName    string           `json:"fullName" validate:"required"`
	NickName    string           `json:"nickName" validate:"required"`
	AvatarURL   string           `json:"avatarURL" validate:"omitempty"`
	Phone       string           `json:"phone" validate:"omitempty,e164"`
	Email       string           `json:"email" validate:"required,email"`
	Location    RegisterLocation `json:"location"`
	Description string           `json:"description" validate:"omitempty,max=300"`
}

type RegisterLocation struct {
	Country   string  `json:"country" validate:"required"`
	City      string  `json:"city" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"omitempty"`
	Longitude float64 `json:"longitude" validate:"omitempty"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateSettingsRequest struct {
	Username    string                 `json:"username" validate:"omitempty"`
	FullName    string                 `json:"fullName" validate:"omitempty"`
	NickName    string                 `json:"nickName" validate:"omitempty"`
	AvatarURL   string                 `json:"avatarURL" validate:"omitempty"`
	Phone       string                 `json:"phone" validate:"omitempty"`
	Email       string                 `json:"email" validate:"omitempty"`
	Location    UpdateSettingsLocation `json:"location"`
	Description string                 `json:"description"`
}

type UpdateSettingsLocation struct {
	Country   string  `json:"country" validate:"omitempty"`
	City      string  `json:"city" validate:"omitempty"`
	Latitude  float64 `json:"latitude" validate:"omitempty"`
	Longitude float64 `json:"longitude" validate:"omitempty"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" validate:"required"`
}
