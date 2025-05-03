package pojo

type User struct {
	UserId       int    `json:"UserId"`
	UserName     string `json:"UserName"`
	UserPassWord string `json:"UserPassWord"`
	UserEmail    string `json:"UserEmail"`
}
