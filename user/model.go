package user

import "html/template"

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}
