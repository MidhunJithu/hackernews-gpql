package graph

import (
	"example/graphql/ports"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Usr  ports.UserSrv
	Link ports.LinkSrv
}
