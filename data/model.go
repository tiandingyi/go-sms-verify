package data

type OTPData struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
}

// omitempty 标记，表示当这个字段的值为空值时，不会在 JSON 中输出这个字段

type VerifyData struct {
	User *OTPData `json:"user,omitempty" validate:"required"`
	Code string   `json:"code,omitempty" validate:"required"`
}
