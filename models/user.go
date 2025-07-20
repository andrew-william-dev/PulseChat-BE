package models

type User struct {
	ID        int   `json:"id"`
	Email    string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	Bio       string `json:"bio"`
	Location  string `json:"location"`
	Website   string `json:"website"`
	Interests   string `json:"interests"`
	Theme     string `json:"theme_color"`
}
