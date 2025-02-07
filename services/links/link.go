package links

import (
	"example/graphql/graph/model"
	"example/graphql/ports"
)

type HackerNewsLinks struct {
	db ports.HackerDB
}

func NewLink(db ports.HackerDB) ports.LinkSrv {
	return &HackerNewsLinks{
		db: db,
	}
}

func (u *HackerNewsLinks) All(userid int) ([]*model.Link, error) {
	return u.db.AllLinks(userid)
}

func (u *HackerNewsLinks) Create(link *model.Link) (*model.Link, error) {
	return u.db.CreateLink(link)
}
