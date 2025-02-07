package user

import (
	"example/graphql/graph/model"
	"example/graphql/ports"
)

type HackerNewsUser struct {
	db ports.HackerDB
}

func NewUser(db ports.HackerDB) ports.UserSrv {
	return &HackerNewsUser{
		db: db,
	}
}

func (u *HackerNewsUser) GenerateToken(username string) (string, error) {
	return "", nil
}

func (u *HackerNewsUser) ParseToken(token string) (string, error) {
	return "", nil
}
func (u *HackerNewsUser) Create(usr model.NewUser) (*model.User, error) {
	return u.db.CreateUser(usr)
}

func (u *HackerNewsUser) UserByName(name string) (int, error) {
	return u.db.UserByName(name)
}

func (u *HackerNewsUser) AuthenticateUser(username string, password string) bool {
	return u.db.AuthenticateUser(username, password)
}
