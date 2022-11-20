package main

import (
	"io"
	"session-9/database"
	"session-9/user"
	"text/template"

	"github.com/labstack/echo"
)

func main() {
	r := echo.New()

	// Initializes database
	db := database.ConnectPGLocal()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.NewUserController(userService)

	r.Renderer = NewRenderer("template/*.html", true)
	r.Any("/register", userHandler.RegisterPage)
	r.Any("/login", userHandler.LoginPage)
	r.Any("/home", userHandler.Homepage)
	r.POST("/register/user", userHandler.Register)
	r.POST("/login/user", userHandler.Login)
	r.POST("/logout", userHandler.Logout)

	r.Start(":9000")
}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}
