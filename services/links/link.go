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

func (u *HackerNewsLinks) VoteLink(vote model.VoteInput, userid int) (int32, error) {
	score, err := u.db.VoteLink(vote, userid)
	return int32(score), err
}
