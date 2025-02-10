package ports

import "example/graphql/graph/model"

type HackerDB interface {
	Close()
	Migrate() error
	CreateLink(link *model.Link) (*model.Link, error)
	AllLinks(int) ([]*model.Link, error)
	CreateUser(model.NewUser) (*model.User, error)
	UserByName(string) (int, error)
	AuthenticateUser(string, string) bool
	VoteLink(model.VoteInput, int) (int, error)
}
