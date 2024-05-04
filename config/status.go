package config

type CustomStatus string

const (
	WRONG_USERNAME       CustomStatus = "ER001"
	WRONG_PASSWORD       CustomStatus = "ER002"
	DUPLICATE_PASSWORD   CustomStatus = "ER003"
	INVALID_OTP          CustomStatus = "ER004"
	INVALID_INFORMATION  CustomStatus = "ER005"
	MUST_CHANGE_PASSWORD CustomStatus = "ER006"
	MUST_VERIFY          CustomStatus = "ER007"
)
