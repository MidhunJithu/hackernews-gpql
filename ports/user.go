package ports

import "example/graphql/graph/model"

type UserSrv interface {
	GenerateToken(username string) (string, error)
	ParseToken(token string) (string, error)
	Create(model.NewUser) (*model.User, error)
	UserByName(string) (int, error)
	AuthenticateUser(string, string) bool
}
