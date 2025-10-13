package payload

type Register struct {
	UserName  string `json:"userName" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
