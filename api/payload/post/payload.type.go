package payload

type PostCreate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"userId"`
}
