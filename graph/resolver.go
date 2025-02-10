package graph

import (
	"example/graphql/graph/model"
	"example/graphql/ports"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Usr  ports.UserSrv
	Link ports.LinkSrv
}

type subs struct {
	mu      sync.Mutex
	clients []chan *model.Link
}

var linkSubs subs

func publishLink(link *model.Link) {
	linkSubs.mu.Lock()
	defer linkSubs.mu.Unlock()
	for _, v := range linkSubs.clients {
		v <- link
	}
}
