package user

import (
	"log"
	"os"

	"github.com/antonlindstrom/pgstore"
	"github.com/labstack/echo"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{repository: repository}
}

var SESSION_ID = "id"

// func newCookieStore() *sessions.CookieStore {
// 	authKey := []byte("my-auth-key-very-secret")
// 	encryptionKey := []byte("my-encryption-key-very-secret123")

// 	store := sessions.NewCookieStore(authKey, encryptionKey)
// 	store.Options.Path = "/"
// 	store.Options.MaxAge = 86400 * 7
// 	store.Options.HttpOnly = true

// 	return store
// }

func newPostgresStore() *pgstore.PGStore {
	url := "postgres://postgresuser:postgrespassword@postgres:5432/postgres?sslmode=disable"
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
	if err != nil {
		log.Println("ERROR", err.Error())
		os.Exit(0)
	}

	return store
}

var store = newPostgresStore()

func (service *Service) Register(id int, username string, firstName string, lastName string, password string, c echo.Context) (err error) {
	var (
		request User
	)

	request.ID = id
	request.Username = username
	request.Firstname = firstName
	request.Lastname = lastName
	request.Password = password

	err = service.repository.Register(request)
	if err != nil {
		return
	}

	return
}

func (service *Service) Login(username, password string, c echo.Context) (err error) {
	var (
		getUser User
		request User
	)

	request.Username = username
	request.Password = password

	getUser, err = service.repository.Login(request)
	if err != nil {
		return
	}

	log.Println("user", getUser)

	session, err := store.Get(c.Request(), SESSION_ID)
	if err != nil {
		return
	}

	session.Values["username"] = getUser.Username
	session.Save(c.Request(), c.Response())

	log.Println("hit sini")

	return
}

func (service *Service) Logout(c echo.Context) (err error) {
	session, _ := store.Get(c.Request(), SESSION_ID)
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())
	return nil
}
