package todo

type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Pasword  string `json:"pasword"`
}
