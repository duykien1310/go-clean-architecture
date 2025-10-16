package payload

type FollowingCreate struct {
	UserId       int `json:"userId"`
	FollowUserId int `json:"followUserId"`
}
