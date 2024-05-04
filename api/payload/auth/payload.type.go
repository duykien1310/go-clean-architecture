package payload

type Register struct {
	Email       string `json:"email" binding:"required"`
	UserName    string `json:"userName" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	OtpCode     string `json:"otpCode" binding:"required"`
}

type SendMailRegister struct {
	Email string `json:"email" binding:"required"`
}

type VerifyUsernamePayload struct {
	UserName string `json:"userName" binding:"required"`
}
