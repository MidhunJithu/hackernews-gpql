package ports

import "example/graphql/graph/model"

type LinkSrv interface {
	Create(link *model.Link) (*model.Link, error)
	All(int) ([]*model.Link, error)
}
